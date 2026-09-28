#!/usr/bin/env bash
#
# Migrate a Go project from Lc v1 to Lc v2.
#
# Lc v2 lives at github.com/pt-main/lc/v2. Every import of the framework,
# including subpackages, must carry the /v2 suffix:
#
#   github.com/pt-main/lc              -> github.com/pt-main/lc/v2
#   github.com/pt-main/lc/engine/core  -> github.com/pt-main/lc/v2/engine/core
#
# The script rewrites imports in .go files, updates go.mod, runs go mod tidy
# and then build and vet. It never touches your git history and makes no commit.
#
# Usage:
#   ./lc-migrate-v2.sh [project-dir]
#
# Requirements: bash 3.2+, go 1.21+, a clean git worktree is recommended.

set -euo pipefail

OLD_PATH="github.com/pt-main/lc"
NEW_PATH="github.com/pt-main/lc/v2"
VERSION="v2.0.0"

PROJECT_DIR="${1:-.}"

usage() {
    cat <<EOF
Migrate a Go project from Lc v1 to Lc v2.

Usage: $(basename "$0") [project-dir]

  project-dir   Go module to migrate (default: current directory)

Environment:
  LC_V2_VERSION   tag to depend on (default: $VERSION)
EOF
}

if [ "${1:-}" = "-h" ] || [ "${1:-}" = "--help" ]; then
    usage
    exit 0
fi

VERSION="${LC_V2_VERSION:-$VERSION}"

die() {
    printf 'error: %s\n' "$1" >&2
    exit 1
}

info() {
    printf '==> %s\n' "$1"
}

command -v go >/dev/null 2>&1 || die "go is not found in PATH"

cd "$PROJECT_DIR" || die "cannot enter directory: $PROJECT_DIR"

[ -f go.mod ] || die "no go.mod in $(pwd)"

info "Project: $(pwd)"

# Detect a git repository rooted at the project directory, not one that only
# contains it: git would otherwise report the parent repository of a nested
# project and reject the worktree for unrelated changes.
in_git_repo() {
    command -v git >/dev/null 2>&1 || return 1
    [ "$(git rev-parse --show-toplevel 2>/dev/null)" = "$(pwd -P)" ]
}

if in_git_repo; then
    if ! git diff --quiet || ! git diff --cached --quiet; then
        die "worktree has uncommitted changes, commit or stash them first"
    fi
fi

# Collect the .go files that still reference Lc. git ls-files keeps the list
# free of vendor, .git and build cache directories; a plain find is the
# fallback for projects outside git.
list_go_files() {
    if in_git_repo; then
        git ls-files -z '*.go'
    else
        find . -type d \( -name .git -o -name vendor -o -name node_modules \) -prune -o -type f -name '*.go' -print0
    fi
}

info "Searching for imports of $OLD_PATH"

# Detection and rewriting share one perl pattern set. The root path is followed
# by a quote, a subpackage path by a slash. A path that already carries /v2 is
# skipped, so a rerun cannot produce "github.com/pt-main/lc/v2/v2/...". The
# patterns are built from the environment because quoting a double quote inside
# a perl one-liner is unreliable across shells.
changed_files=()
while IFS= read -r -d '' file; do
    if LC_OLD_PATH="$OLD_PATH" perl -ne '
        BEGIN {
            $old = quotemeta $ENV{LC_OLD_PATH};
            $q   = chr 34;
            $found = 0;
        }
        $found = 1 if /$q$old(?![a-zA-Z0-9\/._-])/ || /$q$old\/(?!v2["\/])/;
        END { exit($found ? 0 : 1) }
    ' "$file"; then
        changed_files+=("$file")
    fi
done < <(list_go_files)

if [ "${#changed_files[@]}" -eq 0 ]; then
    info "No imports of Lc v1 found, nothing to rewrite"
else
    printf '    %s\n' "${changed_files[@]}"

    for file in "${changed_files[@]}"; do
        LC_OLD_PATH="$OLD_PATH" LC_NEW_PATH="$NEW_PATH" perl -pi -e '
            BEGIN {
                $old = quotemeta $ENV{LC_OLD_PATH};
                $new = $ENV{LC_NEW_PATH};
                $q   = chr 34;

                $root_re = $q . $old;
                $root_sub = $q . $new;
                $sub_re  = $root_re . "/";
                $sub_sub  = $q . $new . "/";
            }

            s/$root_re(?![a-zA-Z0-9\/._-])/$root_sub/ge;
            s{$sub_re(?!v2["\/])}{$sub_sub}ge;
        ' "$file"
    done

    info "Rewrote imports in ${#changed_files[@]} file(s)"
fi

if [ -f go.mod ]; then
    info "Updating go.mod"

    # Drop the v1 requirement so go mod tidy re-resolves it under the new path.
    go mod edit -droprequire "$OLD_PATH" 2>/dev/null || true
    go mod edit -require="$NEW_PATH@$VERSION"
    info "Set $NEW_PATH@$VERSION"
fi

info "Running go mod tidy"
go mod tidy

info "Building"
go build ./...

info "Running go vet"
go vet ./...

info "Done"
cat <<EOF
Check the result:

  git diff
  go test ./...

Lc v2 also renamed public identifiers and changed some string constants,
see docs/changelog.md in the Lc repository for the full breaking list.
EOF