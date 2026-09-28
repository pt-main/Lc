package parser3

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func itoa(n int) string { return strconv.Itoa(n) }

func quote(s string) string { return strconv.Quote(s) }

func typeName(v interface{}) string { return fmt.Sprintf("%T", v) }

func sortStrings(s []string) { sort.Strings(s) }

func joinStrings(s []string) string { return strings.Join(s, ", ") }
