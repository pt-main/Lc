# Files

- [GODOC.md](#godoc-md)
- [LICENSE](#license)
- [README.md](#readme-md)
- [backup/lexer.five](#backup-lexer-five)
- [builder.go](#builder-go)
- [engine/GODOC.md](#engine-godoc-md)
- [engine/byteEngine.go](#engine-byteengine-go)
- [engine/core/GODOC.md](#engine-core-godoc-md)
- [engine/core/errors.go](#engine-core-errors-go)
- [engine/core/events.go](#engine-core-events-go)
- [engine/core/generator.go](#engine-core-generator-go)
- [engine/core/generator_test.go](#engine-core-generator-test-go)
- [engine/core/logger.go](#engine-core-logger-go)
- [engine/core/scope.go](#engine-core-scope-go)
- [engine/core/types.go](#engine-core-types-go)
- [engine/core/universalEngineParams.go](#engine-core-universalengineparams-go)
- [engine/events/GODOC.md](#engine-events-godoc-md)
- [engine/events/byteEngine.go](#engine-events-byteengine-go)
- [engine/events/defaultEvents.go](#engine-events-defaultevents-go)
- [engine/events/stringEngine.go](#engine-events-stringengine-go)
- [engine/interface.go](#engine-interface-go)
- [engine/stringEngine.go](#engine-stringengine-go)
- [engine.go](#engine-go)
- [example/README.md](#example-readme-md)
- [example/langs/calculator/GODOC.md](#example-langs-calculator-godoc-md)
- [example/langs/calculator/main.go](#example-langs-calculator-main-go)
- [example/langs/configLang/GODOC.md](#example-langs-configlang-godoc-md)
- [example/langs/configLang/main.go](#example-langs-configlang-main-go)
- [example/packages/engine/core/events/GODOC.md](#example-packages-engine-core-events-godoc-md)
- [example/packages/engine/core/events/main.go](#example-packages-engine-core-events-main-go)
- [example/packages/engine/core/generator/GODOC.md](#example-packages-engine-core-generator-godoc-md)
- [example/packages/engine/core/generator/main.go](#example-packages-engine-core-generator-main-go)
- [example/packages/engine/core/logger/GODOC.md](#example-packages-engine-core-logger-godoc-md)
- [example/packages/engine/core/logger/main.go](#example-packages-engine-core-logger-main-go)
- [example/packages/engine/core/other/GODOC.md](#example-packages-engine-core-other-godoc-md)
- [example/packages/engine/core/other/main.go](#example-packages-engine-core-other-main-go)
- [example/packages/engine/engines/byte/GODOC.md](#example-packages-engine-engines-byte-godoc-md)
- [example/packages/engine/engines/byte/main.go](#example-packages-engine-engines-byte-main-go)
- [example/packages/engine/engines/string/GODOC.md](#example-packages-engine-engines-string-godoc-md)
- [example/packages/engine/engines/string/main.go](#example-packages-engine-engines-string-main-go)
- [example/readme/byte/GODOC.md](#example-readme-byte-godoc-md)
- [example/readme/byte/main.go](#example-readme-byte-main-go)
- [example/readme/string/GODOC.md](#example-readme-string-godoc-md)
- [example/readme/string/main.go](#example-readme-string-main-go)
- [example/tests/parser3Test/parser3.go](#example-tests-parser3test-parser3-go)
- [example/tests/speedtest/README.md](#example-tests-speedtest-readme-md)
- [example/tests/speedtest/byte/bench/byte_test.go](#example-tests-speedtest-byte-bench-byte-test-go)
- [example/tests/speedtest/byte/tests/main.go](#example-tests-speedtest-byte-tests-main-go)
- [go.mod](#go-mod)
- [go.sum](#go-sum)
- [main.go](#main-go)
- [parsing/GODOC.md](#parsing-godoc-md)
- [parsing/README.md](#parsing-readme-md)
- [parsing/byteParsing/GODOC.md](#parsing-byteparsing-godoc-md)
- [parsing/byteParsing/node.go](#parsing-byteparsing-node-go)
- [parsing/byteParsing/parser1.go](#parsing-byteparsing-parser1-go)
- [parsing/byteParsing/parser1_test.go](#parsing-byteparsing-parser1-test-go)
- [parsing/main.go](#parsing-main-go)
- [parsing/stringParsing/GODOC.md](#parsing-stringparsing-godoc-md)
- [parsing/stringParsing/README.md](#parsing-stringparsing-readme-md)
- [parsing/stringParsing/lexer.go](#parsing-stringparsing-lexer-go)
- [parsing/stringParsing/lexer_test.go](#parsing-stringparsing-lexer-test-go)
- [parsing/stringParsing/node.go](#parsing-stringparsing-node-go)
- [parsing/stringParsing/parser1.go](#parsing-stringparsing-parser1-go)
- [parsing/stringParsing/parser2.go](#parsing-stringparsing-parser2-go)
- [parsing/stringParsing/parser3/GODOC.md](#parsing-stringparsing-parser3-godoc-md)
- [parsing/stringParsing/parser3/engineAdapter.go](#parsing-stringparsing-parser3-engineadapter-go)
- [parsing/stringParsing/parser3/errors.go](#parsing-stringparsing-parser3-errors-go)
- [parsing/stringParsing/parser3/formatter.go](#parsing-stringparsing-parser3-formatter-go)
- [parsing/stringParsing/parser3/grammar.go](#parsing-stringparsing-parser3-grammar-go)
- [parsing/stringParsing/parser3/parser.go](#parsing-stringparsing-parser3-parser-go)
- [parsing/stringParsing/utils.go](#parsing-stringparsing-utils-go)
- [public/GODOC.md](#public-godoc-md)
- [public/errors/GODOC.md](#public-errors-godoc-md)
- [public/errors/engines.go](#public-errors-engines-go)
- [public/errors/events.go](#public-errors-events-go)
- [public/errors/generator.go](#public-errors-generator-go)
- [public/errors/main.go](#public-errors-main-go)
- [public/errors/others.go](#public-errors-others-go)
- [public/events.go](#public-events-go)
- [public/logging.go](#public-logging-go)
- [public/scope.go](#public-scope-go)
- [public/types.go](#public-types-go)
- [scan_results.txt](#scan-results-txt)
- [tooling/README.md](#tooling-readme-md)
- [tooling/astools/GODOC.md](#tooling-astools-godoc-md)
- [tooling/astools/main.go](#tooling-astools-main-go)
- [tooling/bytecode/GODOC.md](#tooling-bytecode-godoc-md)
- [tooling/bytecode/instruction.go](#tooling-bytecode-instruction-go)
- [tooling/bytecode/utils.go](#tooling-bytecode-utils-go)
- [tooling/debugging/extensiblePlugin/GODOC.md](#tooling-debugging-extensibleplugin-godoc-md)
- [tooling/debugging/extensiblePlugin/config.go](#tooling-debugging-extensibleplugin-config-go)
- [tooling/debugging/extensiblePlugin/events.go](#tooling-debugging-extensibleplugin-events-go)
- [tooling/debugging/extensiblePlugin/plugin.go](#tooling-debugging-extensibleplugin-plugin-go)
- [tooling/debugging/extensiblePlugin/structs.go](#tooling-debugging-extensibleplugin-structs-go)
- [tooling/debugging/profiler/GODOC.md](#tooling-debugging-profiler-godoc-md)
- [tooling/debugging/profiler/main.go](#tooling-debugging-profiler-main-go)
- [tooling/plugin/GODOC.md](#tooling-plugin-godoc-md)
- [tooling/plugin/interface.go](#tooling-plugin-interface-go)
- [tooling/plugin/manager.go](#tooling-plugin-manager-go)
- [tooling/plugin/realization.go](#tooling-plugin-realization-go)
- [tooling/plugin/tools.go](#tooling-plugin-tools-go)

---

> Note: all ``` ` ``` symbols was replaced to '

# GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# lc

'''go
import "github.com/pt-main/lc"
'''

## Index

- [Constants](<#constants>)
- [func NewByteEngine\(generator\_res\_type public.ResType, pipeline \[\]string, add\_default\_events bool, parser byteParser, endianess public.EndianType, colorEnable bool, context context.Context\) \*engine.ByteEngine](<#NewByteEngine>)
- [func NewStringEngine\(generator\_res\_type public.ResType, pipeline \[\]string, add\_default\_events bool, parser stringParser, colorEnable bool, context context.Context\) \*engine.StringEngine](<#NewStringEngine>)
- [type EngineBuilder](<#EngineBuilder>)
  - [func NewEngineBuilder\(engineType public.EngineType, resType public.ResType\) \*EngineBuilder](<#NewEngineBuilder>)
  - [func \(b \*EngineBuilder\) Build\(\) \(\*EngineUniversal, error\)](<#EngineBuilder.Build>)
  - [func \(b \*EngineBuilder\) WithByteParser\(parser byteParser\) \*EngineBuilder](<#EngineBuilder.WithByteParser>)
  - [func \(b \*EngineBuilder\) WithColors\(\) \*EngineBuilder](<#EngineBuilder.WithColors>)
  - [func \(b \*EngineBuilder\) WithContext\(ctx context.Context\) \*EngineBuilder](<#EngineBuilder.WithContext>)
  - [func \(b \*EngineBuilder\) WithDefaultEvents\(add bool\) \*EngineBuilder](<#EngineBuilder.WithDefaultEvents>)
  - [func \(b \*EngineBuilder\) WithEndianess\(endianess public.EndianType\) \*EngineBuilder](<#EngineBuilder.WithEndianess>)
  - [func \(b \*EngineBuilder\) WithLogger\(logger \*core.Logger\) \*EngineBuilder](<#EngineBuilder.WithLogger>)
  - [func \(b \*EngineBuilder\) WithPipeline\(pipeline \[\]string\) \*EngineBuilder](<#EngineBuilder.WithPipeline>)
  - [func \(b \*EngineBuilder\) WithPlugins\(plugins ...plugin.PluginInterface\) \*EngineBuilder](<#EngineBuilder.WithPlugins>)
  - [func \(b \*EngineBuilder\) WithScope\(scope core.ScopeType\) \*EngineBuilder](<#EngineBuilder.WithScope>)
  - [func \(b \*EngineBuilder\) WithStringParser\(parser stringParser\) \*EngineBuilder](<#EngineBuilder.WithStringParser>)
- [type EngineUniversal](<#EngineUniversal>)
  - [func \(e \*EngineUniversal\) CheckEnded\(\) \(err core.ErrorInterface\)](<#EngineUniversal.CheckEnded>)
  - [func \(e \*EngineUniversal\) End\(\) \(err error\)](<#EngineUniversal.End>)
  - [func \(e \*EngineUniversal\) GetUEP\(\) \(\*core.UniversalEngineParams, error\)](<#EngineUniversal.GetUEP>)
  - [func \(e \*EngineUniversal\) NewCommandByte\(opcode int, handler core.CommandType\[engine.ByteEngineInterface, byteParsing.ParsedBytes\], name string, autoByecodeIdxShift bool\) error](<#EngineUniversal.NewCommandByte>)
  - [func \(e \*EngineUniversal\) NewCommandString\(cmdSwitch string, handler core.CommandType\[engine.StringEngineInterface, stringParsing.ParsedNode\], doc string\) error](<#EngineUniversal.NewCommandString>)
  - [func \(e \*EngineUniversal\) ProcessBytes\(input \[\]byte\) core.ErrorInterface](<#EngineUniversal.ProcessBytes>)
  - [func \(e \*EngineUniversal\) ProcessBytesWithCtx\(input \[\]byte, ctx context.Context\) core.ErrorInterface](<#EngineUniversal.ProcessBytesWithCtx>)
  - [func \(e \*EngineUniversal\) ProcessString\(input string\) core.ErrorInterface](<#EngineUniversal.ProcessString>)
  - [func \(e \*EngineUniversal\) ProcessStringWithCtx\(input string, ctx context.Context\) core.ErrorInterface](<#EngineUniversal.ProcessStringWithCtx>)


## Constants

<a name="Version"></a>

'''go
const Version = "1.5.7"
'''

<a name="NewByteEngine"></a>
## func [NewByteEngine](<https://github.com/pt-main/Lc/blob/main/main.go#L58-L66>)

'''go
func NewByteEngine(generator_res_type public.ResType, pipeline []string, add_default_events bool, parser byteParser, endianess public.EndianType, colorEnable bool, context context.Context) *engine.ByteEngine
'''

NewByteEngine creates a byte\-oriented engine for binary formats or bytecode.

The endianess parameter \(e.g., bytecode.LittleEndian\) is stored in scope.

It registers default events when add\_default\_events is true.

The parser must implement paraing.ParserInterface.

<a name="NewStringEngine"></a>
## func [NewStringEngine](<https://github.com/pt-main/Lc/blob/main/main.go#L26-L33>)

'''go
func NewStringEngine(generator_res_type public.ResType, pipeline []string, add_default_events bool, parser stringParser, colorEnable bool, context context.Context) *engine.StringEngine
'''

NewStringEngine creates a ready\-to\-use string\-based engine. Parameters:

'''
generator_res_type – core.StringResType (usually) for text generation.
pipeline – ordered list of generation points (e.g., []string{"pre","main"}).
add_default_events – if true, registers standard parsing and call events.
parser – an implementation parser.ParserInterface.
'''

Returns a StringEngine with empty command map and initialized UEP.

<a name="EngineBuilder"></a>
## type [EngineBuilder](<https://github.com/pt-main/Lc/blob/main/builder.go#L22-L37>)

EngineBuilder is a fluent builder for constructing universal engines. It allows to configure pipeline stages, event handling, logging, custom parsers, scope variables, and byte order before calling Build\(\). Use NewEngineBuilder to create a builder instance.

'''go
type EngineBuilder struct {
    // contains filtered or unexported fields
}
'''

<a name="NewEngineBuilder"></a>
### func [NewEngineBuilder](<https://github.com/pt-main/Lc/blob/main/builder.go#L48>)

'''go
func NewEngineBuilder(engineType public.EngineType, resType public.ResType) *EngineBuilder
'''

NewEngineBuilder creates a new EngineBuilder for the given engine type. engineType must be either ByteEngineType or StringEngineType. Defaults: pipeline = \[\]string\{"main"\}, default events enabled, endianess = bytecode.LittleEndian, empty scope. Example:

'''
builder := lc.NewEngineBuilder(lc.StringEngineType).
            WithPipeline([]string{"pre","main"}).
            WithStringParser(myParser)
'''

<a name="EngineBuilder.Build"></a>
### func \(\*EngineBuilder\) [Build](<https://github.com/pt-main/Lc/blob/main/builder.go#L117>)

'''go
func (b *EngineBuilder) Build() (*EngineUniversal, error)
'''

Build constructs and returns an EngineUniversal or an error if required components are missing \(e.g., a string parser for a StringEngine\). The returned engineUniversal can process strings or bytes depending on its type and provides methods to register commands.

<a name="EngineBuilder.WithByteParser"></a>
### func \(\*EngineBuilder\) [WithByteParser](<https://github.com/pt-main/Lc/blob/main/builder.go#L97>)

'''go
func (b *EngineBuilder) WithByteParser(parser byteParser) *EngineBuilder
'''



<a name="EngineBuilder.WithColors"></a>
### func \(\*EngineBuilder\) [WithColors](<https://github.com/pt-main/Lc/blob/main/builder.go#L87>)

'''go
func (b *EngineBuilder) WithColors() *EngineBuilder
'''



<a name="EngineBuilder.WithContext"></a>
### func \(\*EngineBuilder\) [WithContext](<https://github.com/pt-main/Lc/blob/main/builder.go#L65>)

'''go
func (b *EngineBuilder) WithContext(ctx context.Context) *EngineBuilder
'''



<a name="EngineBuilder.WithDefaultEvents"></a>
### func \(\*EngineBuilder\) [WithDefaultEvents](<https://github.com/pt-main/Lc/blob/main/builder.go#L70>)

'''go
func (b *EngineBuilder) WithDefaultEvents(add bool) *EngineBuilder
'''



<a name="EngineBuilder.WithEndianess"></a>
### func \(\*EngineBuilder\) [WithEndianess](<https://github.com/pt-main/Lc/blob/main/builder.go#L102>)

'''go
func (b *EngineBuilder) WithEndianess(endianess public.EndianType) *EngineBuilder
'''



<a name="EngineBuilder.WithLogger"></a>
### func \(\*EngineBuilder\) [WithLogger](<https://github.com/pt-main/Lc/blob/main/builder.go#L75>)

'''go
func (b *EngineBuilder) WithLogger(logger *core.Logger) *EngineBuilder
'''



<a name="EngineBuilder.WithPipeline"></a>
### func \(\*EngineBuilder\) [WithPipeline](<https://github.com/pt-main/Lc/blob/main/builder.go#L60>)

'''go
func (b *EngineBuilder) WithPipeline(pipeline []string) *EngineBuilder
'''



<a name="EngineBuilder.WithPlugins"></a>
### func \(\*EngineBuilder\) [WithPlugins](<https://github.com/pt-main/Lc/blob/main/builder.go#L107>)

'''go
func (b *EngineBuilder) WithPlugins(plugins ...plugin.PluginInterface) *EngineBuilder
'''



<a name="EngineBuilder.WithScope"></a>
### func \(\*EngineBuilder\) [WithScope](<https://github.com/pt-main/Lc/blob/main/builder.go#L80>)

'''go
func (b *EngineBuilder) WithScope(scope core.ScopeType) *EngineBuilder
'''



<a name="EngineBuilder.WithStringParser"></a>
### func \(\*EngineBuilder\) [WithStringParser](<https://github.com/pt-main/Lc/blob/main/builder.go#L92>)

'''go
func (b *EngineBuilder) WithStringParser(parser stringParser) *EngineBuilder
'''



<a name="EngineUniversal"></a>
## type [EngineUniversal](<https://github.com/pt-main/Lc/blob/main/engine.go#L17-L26>)



'''go
type EngineUniversal struct {
    Plugins      *lcplugin.PluginManager
    Type         public.EngineType
    StringEngine engine.EngineInterface[string, string, stringParsing.ParsedNode]
    ByteEngine   engine.EngineInterface[int, []byte, byteParsing.ParsedBytes]

    Context        context.Context
    CtxCancelCause context.CancelCauseFunc
    // contains filtered or unexported fields
}
'''

<a name="EngineUniversal.CheckEnded"></a>
### func \(\*EngineUniversal\) [CheckEnded](<https://github.com/pt-main/Lc/blob/main/engine.go#L141>)

'''go
func (e *EngineUniversal) CheckEnded() (err core.ErrorInterface)
'''



<a name="EngineUniversal.End"></a>
### func \(\*EngineUniversal\) [End](<https://github.com/pt-main/Lc/blob/main/engine.go#L123>)

'''go
func (e *EngineUniversal) End() (err error)
'''

End \- function for stop engines lifecycle.

<a name="EngineUniversal.GetUEP"></a>
### func \(\*EngineUniversal\) [GetUEP](<https://github.com/pt-main/Lc/blob/main/engine.go#L65>)

'''go
func (e *EngineUniversal) GetUEP() (*core.UniversalEngineParams, error)
'''



<a name="EngineUniversal.NewCommandByte"></a>
### func \(\*EngineUniversal\) [NewCommandByte](<https://github.com/pt-main/Lc/blob/main/engine.go#L78-L81>)

'''go
func (e *EngineUniversal) NewCommandByte(opcode int, handler core.CommandType[engine.ByteEngineInterface, byteParsing.ParsedBytes], name string, autoByecodeIdxShift bool) error
'''

NewCommandByte registers a bytecode command identified by an opcode. If opcode == \-1, the engine automatically assigns the next available opcode. handler receives \(\*ByteEngine, ParsedBytes\).

<a name="EngineUniversal.NewCommandString"></a>
### func \(\*EngineUniversal\) [NewCommandString](<https://github.com/pt-main/Lc/blob/main/engine.go#L107-L109>)

'''go
func (e *EngineUniversal) NewCommandString(cmdSwitch string, handler core.CommandType[engine.StringEngineInterface, stringParsing.ParsedNode], doc string) error
'''

NewCommandString registers a text\-based command in a StringEngine. cmdSwitch is the command name \(e.g., "print"\). handler must have signature func\(\[\]interface\{\}\) error where arguments are \(\*StringEngine, ParsedNode\). doc is an optional documentation string.

<a name="EngineUniversal.ProcessBytes"></a>
### func \(\*EngineUniversal\) [ProcessBytes](<https://github.com/pt-main/Lc/blob/main/engine.go#L61>)

'''go
func (e *EngineUniversal) ProcessBytes(input []byte) core.ErrorInterface
'''

ProcessBytes feeds a byte slice into the engine \(ByteEngineType only\). The input is passed via scope under key "input\_\[\]byte", then parsed and processed.

<a name="EngineUniversal.ProcessBytesWithCtx"></a>
### func \(\*EngineUniversal\) [ProcessBytesWithCtx](<https://github.com/pt-main/Lc/blob/main/engine.go#L40>)

'''go
func (e *EngineUniversal) ProcessBytesWithCtx(input []byte, ctx context.Context) core.ErrorInterface
'''



<a name="EngineUniversal.ProcessString"></a>
### func \(\*EngineUniversal\) [ProcessString](<https://github.com/pt-main/Lc/blob/main/engine.go#L55>)

'''go
func (e *EngineUniversal) ProcessString(input string) core.ErrorInterface
'''

ProcessString feeds a string input into the engine. It works only for engines of type StringEngineType; otherwise returns an core.ErrorInterface. Internally triggers the parse and call events, executing registered handlers.

<a name="EngineUniversal.ProcessStringWithCtx"></a>
### func \(\*EngineUniversal\) [ProcessStringWithCtx](<https://github.com/pt-main/Lc/blob/main/engine.go#L28>)

'''go
func (e *EngineUniversal) ProcessStringWithCtx(input string, ctx context.Context) core.ErrorInterface
'''



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# LICENSE

```
                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

   TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

   1. Definitions.

      "License" shall mean the terms and conditions for use, reproduction,
      and distribution as defined by Sections 1 through 9 of this document.

      "Licensor" shall mean the copyright owner or entity authorized by
      the copyright owner that is granting the License.

      "Legal Entity" shall mean the union of the acting entity and all
      other entities that control, are controlled by, or are under common
      control with that entity. For the purposes of this definition,
      "control" means (i) the power, direct or indirect, to cause the
      direction or management of such entity, whether by contract or
      otherwise, or (ii) ownership of fifty percent (50%) or more of the
      outstanding shares, or (iii) beneficial ownership of such entity.

      "You" (or "Your") shall mean an individual or Legal Entity
      exercising permissions granted by this License.

      "Source" form shall mean the preferred form for making modifications,
      including but not limited to software source code, documentation
      source, and configuration files.

      "Object" form shall mean any form resulting from mechanical
      transformation or translation of a Source form, including but
      not limited to compiled object code, generated documentation,
      and conversions to other media types.

      "Work" shall mean the work of authorship, whether in Source or
      Object form, made available under the License, as indicated by a
      copyright notice that is included in or attached to the work
      (an example is provided in the Appendix below).

      "Derivative Works" shall mean any work, whether in Source or Object
      form, that is based on (or derived from) the Work and for which the
      editorial revisions, annotations, elaborations, or other modifications
      represent, as a whole, an original work of authorship. For the purposes
      of this License, Derivative Works shall not include works that remain
      separable from, or merely link (or bind by name) to the interfaces of,
      the Work and Derivative Works thereof.

      "Contribution" shall mean any work of authorship, including
      the original version of the Work and any modifications or additions
      to that Work or Derivative Works thereof, that is intentionally
      submitted to Licensor for inclusion in the Work by the copyright owner
      or by an individual or Legal Entity authorized to submit on behalf of
      the copyright owner. For the purposes of this definition, "submitted"
      means any form of electronic, verbal, or written communication sent
      to the Licensor or its representatives, including but not limited to
      communication on electronic mailing lists, source code control systems,
      and issue tracking systems that are managed by, or on behalf of, the
      Licensor for the purpose of discussing and improving the Work, but
      excluding communication that is conspicuously marked or otherwise
      designated in writing by the copyright owner as "Not a Contribution."

      "Contributor" shall mean Licensor and any individual or Legal Entity
      on behalf of whom a Contribution has been received by Licensor and
      subsequently incorporated within the Work.

   2. Grant of Copyright License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      copyright license to reproduce, prepare Derivative Works of,
      publicly display, publicly perform, sublicense, and distribute the
      Work and such Derivative Works in Source or Object form.

   3. Grant of Patent License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      (except as stated in this section) patent license to make, have made,
      use, offer to sell, sell, import, and otherwise transfer the Work,
      where such license applies only to those patent claims licensable
      by such Contributor that are necessarily infringed by their
      Contribution(s) alone or by combination of their Contribution(s)
      with the Work to which such Contribution(s) was submitted. If You
      institute patent litigation against any entity (including a
      cross-claim or counterclaim in a lawsuit) alleging that the Work
      or a Contribution incorporated within the Work constitutes direct
      or contributory patent infringement, then any patent licenses
      granted to You under this License for that Work shall terminate
      as of the date such litigation is filed.

   4. Redistribution. You may reproduce and distribute copies of the
      Work or Derivative Works thereof in any medium, with or without
      modifications, and in Source or Object form, provided that You
      meet the following conditions:

      (a) You must give any other recipients of the Work or
          Derivative Works a copy of this License; and

      (b) You must cause any modified files to carry prominent notices
          stating that You changed the files; and

      (c) You must retain, in the Source form of any Derivative Works
          that You distribute, all copyright, patent, trademark, and
          attribution notices from the Source form of the Work,
          excluding those notices that do not pertain to any part of
          the Derivative Works; and

      (d) If the Work includes a "NOTICE" text file as part of its
          distribution, then any Derivative Works that You distribute must
          include a readable copy of the attribution notices contained
          within such NOTICE file, excluding those notices that do not
          pertain to any part of the Derivative Works, in at least one
          of the following places: within a NOTICE text file distributed
          as part of the Derivative Works; within the Source form or
          documentation, if provided along with the Derivative Works; or,
          within a display generated by the Derivative Works, if and
          wherever such third-party notices normally appear. The contents
          of the NOTICE file are for informational purposes only and
          do not modify the License. You may add Your own attribution
          notices within Derivative Works that You distribute, alongside
          or as an addendum to the NOTICE text from the Work, provided
          that such additional attribution notices cannot be construed
          as modifying the License.

      You may add Your own copyright statement to Your modifications and
      may provide additional or different license terms and conditions
      for use, reproduction, or distribution of Your modifications, or
      for any such Derivative Works as a whole, provided Your use,
      reproduction, and distribution of the Work otherwise complies with
      the conditions stated in this License.

   5. Submission of Contributions. Unless You explicitly state otherwise,
      any Contribution intentionally submitted for inclusion in the Work
      by You to the Licensor shall be under the terms and conditions of
      this License, without any additional terms or conditions.
      Notwithstanding the above, nothing herein shall supersede or modify
      the terms of any separate license agreement you may have executed
      with Licensor regarding such Contributions.

   6. Trademarks. This License does not grant permission to use the trade
      names, trademarks, service marks, or product names of the Licensor,
      except as required for reasonable and customary use in describing the
      origin of the Work and reproducing the content of the NOTICE file.

   7. Disclaimer of Warranty. Unless required by applicable law or
      agreed to in writing, Licensor provides the Work (and each
      Contributor provides its Contributions) on an "AS IS" BASIS,
      WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
      implied, including, without limitation, any warranties or conditions
      of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
      PARTICULAR PURPOSE. You are solely responsible for determining the
      appropriateness of using or redistributing the Work and assume any
      risks associated with Your exercise of permissions under this License.

   8. Limitation of Liability. In no event and under no legal theory,
      whether in tort (including negligence), contract, or otherwise,
      unless required by applicable law (such as deliberate and grossly
      negligent acts) or agreed to in writing, shall any Contributor be
      liable to You for damages, including any direct, indirect, special,
      incidental, or consequential damages of any character arising as a
      result of this License or out of the use or inability to use the
      Work (including but not limited to damages for loss of goodwill,
      work stoppage, computer failure or malfunction, or any and all
      other commercial damages or losses), even if such Contributor
      has been advised of the possibility of such damages.

   9. Accepting Warranty or Additional Liability. While redistributing
      the Work or Derivative Works thereof, You may choose to offer,
      and charge a fee for, acceptance of support, warranty, indemnity,
      or other liability obligations and/or rights consistent with this
      License. However, in accepting such obligations, You may act only
      on Your own behalf and on Your sole responsibility, not on behalf
      of any other Contributor, and only if You agree to indemnify,
      defend, and hold each Contributor harmless for any liability
      incurred by, or claims asserted against, such Contributor by reason
      of your accepting any such warranty or additional liability.

   END OF TERMS AND CONDITIONS

   APPENDIX: How to apply the Apache License to your work.

      To apply the Apache License to your work, attach the following
      boilerplate notice, with the fields enclosed by brackets "[]"
      replaced with your own identifying information. (Don't include
      the brackets!)  The text should be enclosed in the appropriate
      comment syntax for the file format. We also recommend that a
      file or class name and description of purpose be included on the
      same "printed page" as the copyright notice for easier
      identification within third-party archives.

   Copyright 2026 Pt

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```

---

# README.md

```md
<h1 align="center">Lc — Language Creator & Devkit</h1>
<p align="center">
  <img alt="banner-low-50p-8b" src="https://github.com/user-attachments/assets/8fa74598-5cee-403e-a9dc-417e86d22dcd" />
</p>
<p align="center">
  <a href="https://pkg.go.dev/github.com/pt-main/lc"><img src="https://img.shields.io/badge/Go-Reference-007d9c?logo=go&logoColor=white"></a>
  <a href="https://github.com/pt-main/lc/releases"><img src="https://img.shields.io/github/v/release/pt-main/lc?color=blue"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-Apache_2.0-yellow"></a>
  <a href="https://github.com/pt-main/Lc/wiki"><img src="https://img.shields.io/badge/Project-Wiki-red"></a>
</p>


> '''bash
> go get github.com/pt-main/lc
> '''
> Lc is a production-oriented framework and toolkit for building language runtimes, compiler-like execution pipelines, command interpreters, and bytecode-driven processors in Go.

Lc contains - 
- Byte & String Engine with Universal Engine abstraction
- Parser (byte & string) - simple parsers, peg parser, lexers, etc.
- Plugin system (works with Universal Engine) 
- Tooling (bytecode, ast, profilers, etc.)

It is intentionally straightforward to adopt, while preserving industrial runtime properties:
- explicit execution lifecycle (just use 'EngineUniversal.End()'),
- deterministic output assembly,
- context-aware cancellation,
- thread-safe core primitives,
- clear extension contracts for parsers and command handlers, plugins.

Lc does not enforce one grammar style or one VM model.  
Instead, it gives you one runtime surface with two engine backends:
- **String Engine** for text-first processing.
- **Byte Engine** for binary instruction execution.
- **Universal Engine** - abstraction for work with string/byte engine with plugins, context (with cancelation), and simple building.

## Table of Contents
- [Quick start](#quick-start)
- [Engine model](#engine-model)
- [Tools and features](#tools-and-features)
- [Parsers](#parsers)
- [Plugin system](#plugin-system)
- [Context support](#context-support)
- [License](#license)

# Quick start
<details> <summary>- StringEngine Example</summary>

'''go
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
)

func main() {
	parser := &stringParsing.Parser2{} // parsing format: 'command arg1, arg2...'

	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(parser).
		WithDefaultEvents(true).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandString("log", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		args, _ := node.Metadata["args"].(string)
		return se.GetUep().Generator.AddString(fmt.Sprintf("Log [%v]: %v",
			time.Now().Format(time.Stamp), args), "main")
	}, "append log with timestamp")
	if err != nil {
		panic(err)
	}

	err = engine.ProcessString(strings.Join([]string{
		"log service_start",
		"log service_ready",
	}, "\n"))
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
'''

'''bash
$ go run ./example/readme/byte
abc
'''

</details> <details> <summary>- ByteEngine Example</summary>

'''go
package main

import (
	"fmt"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/bytecode"
)

func main() {
	// Parsing format:
	// instruction {
	//     [bytes : cmd] [bytes : argscount] [bytes : arglen]  [bytes arglen : arg],
	//                                       [bytes : arglen2] [bytes arglen2 : arg2]...
	// }
	parser := &byteParsing.Parser1{
		Config: byteParsing.Parser1Config{
			GConfig: bytecode.GenerationConfig{
				CommandBytelen:   1,
				ArgscountBytelen: 1,
				ArglenBytelen:    2,
				Endianess:        public.LittleEndian,
			},
			Shifter: bytecode.Shift{},
		},
	}

	engine, err := lc.NewEngineBuilder(public.ByteEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithByteParser(parser).
		WithDefaultEvents(true).
		WithColors().
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandByte(1, func(be enginepkg.ByteEngineInterface, node *byteParsing.ParsedBytes) core.ErrorInterface {
		for _, arg := range node.Args {
			if err := be.GetUep().Generator.AddString(string(arg), "main"); err != nil {
				return err
			}
		}
		return nil
	}, "add to output instruction", true)
	if err != nil {
		panic(err)
	}

	code := []byte{
		0x01,       		// opcode=1
		0x01,       		// argsCount=1
		0x03, 0x00, 		// arglen=3 (little endian, 2 bytes)
		0x61, 0x62, 0x63, 	// args="abc" (3 bytes)
	}

	err = engine.ProcessBytes(code)
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", out)
}
'''

'''bash
$ go run ./example/readme/string
Log [Aug  7 18:44:40]: service_start
Log [Aug  7 18:44:40]: service_ready
'''

</details>

You can find more examples at 'examles/'


# Engine model
## String Engine
Input string (code) and process that - edit, execute, generate code, etc.

Default lifecycle:
1. store input in scope;
2. parse input to '[]ParsedNode';
3. dispatch handlers by 'ParsedNode.Switch';
4. emit output through 'UEP.Generator' (if need).

## Byte Engine
Input bytecode and process that. Very fast hotloop (raw speed - ~200m ops/s on 'i7-4770HQ').

Default lifecycle:
1. store input in scope;
2. parse input to '[]ParsedBytes';
3. convert 'ParsedBytes' to 'ByteCallAttr' - small structure for hotloop;
4. dispatch opcode handler;
5. advance instruction pointer automatically or manually.


# Tools and features

## Powerful core and UEP (Universal Engine Params)
Engines core contains all necessary tools for runtime work. UEP contains then.

You can use it like:

'''go
engine, _ := lc.NewStringEngine(...)
engine.UEP.Generator.AddString(...)
engine.UEP...
'''

Or:

'''go
engine, _ := lc.NewEngineBuilder(...).
	[...].
	Build()
uep, _ := engine.GetUEP()
uep.Generator.AddString(...)
uep...
'''

### 'Events'
Engine arch is event-driven. Events can communicate with 'Events.Scope', work with context ('Events.Context'), call by pipeline. 

Event handlers input '*Events, *EventInput'.

You can override Events by implementing 'core.EventsInterface'.

#### Example
'''go
events := core.NewEvents(context.Background()) // new manager
events.NewEvent("event1", handler1) // create main handler in "event1" event
events.NewEvent("event1", handler2) // append handler to end of "event1"
events.NewEventBefore("event1", handler3) // append handler to start of "event1"
// "event1" - [handler3, handler1, handler2]
'''

### 'Generator'
Powerful tool for codegen.

Work with points pipeline for storing code in independent points. Can generate bytes or string.

#### Example

'''go
pipeline := []string{"pre", "main"}
generator := core.NewGenerator([result-type], pipeline)
generator.AddStrings([]string{ // add strings to main
	"string1 ",
	"string2.",
}, "main")
generator.AddStrings([]string{ // add strings to pre
	"string3 ",
	"string4. ",
}, "pre")
res := core.GetStringRes(generator, "") // get code
// res = string3 string4. string1 string2.
'''

### 'Scope'
The Scope is a thread-safe 'map[string]interface{}' shared across all event handlers, parsers, and commands. It serves as a runtime context for passing data between pipeline stages.

**Important:** Do not overwrite keys from 'public/' package in your custom handlers unless you know exactly what you're doing — they are used by default events.

#### Custom scope usage
'''go
engine, _ := lc.NewEngineBuilder(...).
	WithScope(core.ScopeType{
		"tenant_id": "prod-001",
		"env":       "production",
	}).
	Build()

// later, in your command handler:
func myHandler(se *engine.StringEngine, node stringParsing.ParsedNode) error {
	tenant, _ := core.ScopeGet[string](se.UEP.Scope, "tenant_id")
	fmt.Println("Running for tenant:", tenant)
	return nil
}
'''

### 'Logger'
Structured logger built into UEP. Supports status-based formatting and log level filtering.

#### Example
'''go
logger := core.NewLogger("") // uses default format: "[?BE]%s[?RT] [?CN][%v][?RT] [?GN][%s][?RT]\n"
logger.Logging["debug"] = true  // enable debug output
// other logging will be disabled

// in your engine builder:
engine, _ := lc.NewEngineBuilder(...).
	WithLogger(logger).
	Build()

// in your handlers:
func myHandler(se *engine.StringEngine, node stringParsing.ParsedNode) error {
	se.UEP.Logger.PrintLog("debug", "Processing node: "+node.Switch)
	se.UEP.Logger.PrintLog("error", "Error: "+...) // disabled
	...
}
'''

#### Custom status format
'''go
logger := core.NewLogger("")
logger.Statuses["warn"] = "[?YW]WARN[?RT] [%v] [?RD]%s[?RT]\n" // pt-main/tap color format
logger.PrintLog("warn", "This is a warning")
'''

## Plugin System
Lc has a built‑in plugin manager that allows dynamic registration and execution of external logic. Plugins has their own events and scope, and not isolated (have access to engine and plugin manager).

### Creating a plugin
'''go
import "github.com/pt-main/lc/tooling/plugin"

myPlugin := plugin.NewPlugin(
	"my_plugin",          // name
	"init_event",         // event called on init
	"main_event",         // event called on Run()
	"close_event",        // event called on Close()
	"scope_return",       // plugin.Run (or plugin method) event can put output here
)

// Add handlers to plugin events
myPlugin.Events.NewEvent("init_event", func(ev *core.Events, i *EventInput) error {
	ev.Scope["plugin_ready"] = true
	return nil
})

myPlugin.Events.NewEvent("main_event", func(ev *core.Events, i *EventInput) error {
	// i.Input is whatever was passed to plugin.Run()
	return nil
})
'''

### Registering and using a plugin
'''go
engne, _ := lc.NewEngineBuilder(...).
	WithPlugins(myPlugin). // call "init_event"
	Build()

// Later, call plugin methods:
result, err := engine.Plugins.RunPlugin("my_plugin", "some input") // call "main_event"
'''

## Parsers — ready‑to‑use implementations
Lc ships with several parsers for different use cases:

### StringParsing parsers

| Parser | Description | Best for |
|--------|-------------|----------|
| **Lexer** | Token-based lexer with regexp2 rules, supports bracket balancing and prev/next links | Tokenization |
| **Parser1** | Regex-based grammar with line continuation and bracket balancing | DSLs with line-oriented syntax |
| **Parser2** | Simple 'command args' line parser | Quick prototyping, shell-like languages |
| **Parser3** | PEG-inspired parser with combinators (Sequence, Choice, Repeat, Optional, Named) | Complex grammars, AST generation |
| **Adapter** | 'Parser3' adapter for string engine. |

#### Example: Parser2 (simplest)
'''go
parser := &stringParsing.Parser2{}
// Input: "print hello world"
// Output: ParsedNode{Switch: "print", Metadata: {args: "hello world"}}
'''

### ByteParsing parsers
| Parser | Description |
|--------|-------------|
| **Parser1** | Binary instruction decoder with configurable field lengths and endianness |

'''go
parser := &byteParsing.Parser1{
	Config: byteParsing.Parser1Config{
		GConfig: bytecode.GenerationConfig{
			CommandBytelen:   1,
			ArgscountBytelen: 1,
			ArglenBytelen:    1,
			Endianess:        public.LittleEndian,
		},
		Shifter: bytecode.Shift{},
	},
}
'''

## Context support
All 'Process*' methods have 'WithCtx' variants that accept 'context.Context'. This allows:
- Timeout-based cancellation
- Graceful shutdown
- Request-scoped values

'''go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

err := engine.ProcessStringWithCtx(input, ctx)
if errors.Is(err, context.DeadlineExceeded) {
	fmt.Println("Execution timed out")
}
'''

## License
Apache 2.0 - see 'LICENSE'.

By Pt.
```

---

# backup/lexer.five

```five
Can't read file: 'utf-8' codec can't decode byte 0xf6 in position 80: invalid start byte
```

---

# builder.go

```go
package lc

import (
	"context"
	"errors"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/plugin"
)

type stringParser parsing.ParserInterface[string, stringParsing.ParsedNode]
type byteParser parsing.ParserInterface[[]byte, byteParsing.ParsedBytes]

// EngineBuilder is a fluent builder for constructing universal engines.
// It allows to configure pipeline stages, event handling, logging,
// custom parsers, scope variables, and byte order before calling Build().
// Use NewEngineBuilder to create a builder instance.
type EngineBuilder struct {
	engineType       public.EngineType
	pipeline         []string
	addDefaultEvents bool
	logger           *core.Logger
	scope            core.ScopeType
	stringParser     stringParser
	byteParser       byteParser
	endianess        public.EndianType
	colorEnabled     bool
	context          context.Context
	pm               bool
	plugins          []plugin.PluginInterface
	cancel           context.CancelCauseFunc
	resType          public.ResType
}

// NewEngineBuilder creates a new EngineBuilder for the given engine type.
// engineType must be either ByteEngineType or StringEngineType.
// Defaults: pipeline = []string{"main"}, default events enabled,
// endianess = bytecode.LittleEndian, empty scope.
// Example:
//
//	builder := lc.NewEngineBuilder(lc.StringEngineType).
//	            WithPipeline([]string{"pre","main"}).
//	            WithStringParser(myParser)
func NewEngineBuilder(engineType public.EngineType, resType public.ResType) *EngineBuilder {
	return &EngineBuilder{
		engineType:       engineType,
		resType:          resType,
		pipeline:         []string{"main"},
		addDefaultEvents: true,
		endianess:        public.LittleEndian,
		scope:            make(core.ScopeType),
		context:          context.Background(),
	}
}

func (b *EngineBuilder) WithPipeline(pipeline []string) *EngineBuilder {
	b.pipeline = pipeline
	return b
}

func (b *EngineBuilder) WithContext(ctx context.Context) *EngineBuilder {
	b.context, b.cancel = context.WithCancelCause(ctx)
	return b
}

func (b *EngineBuilder) WithDefaultEvents(add bool) *EngineBuilder {
	b.addDefaultEvents = add
	return b
}

func (b *EngineBuilder) WithLogger(logger *core.Logger) *EngineBuilder {
	b.logger = logger
	return b
}

func (b *EngineBuilder) WithScope(scope core.ScopeType) *EngineBuilder {
	for k, v := range scope {
		b.scope[k] = v
	}
	return b
}

func (b *EngineBuilder) WithColors() *EngineBuilder {
	b.colorEnabled = true
	return b
}

func (b *EngineBuilder) WithStringParser(parser stringParser) *EngineBuilder {
	b.stringParser = parser
	return b
}

func (b *EngineBuilder) WithByteParser(parser byteParser) *EngineBuilder {
	b.byteParser = parser
	return b
}

func (b *EngineBuilder) WithEndianess(endianess public.EndianType) *EngineBuilder {
	b.endianess = endianess
	return b
}

func (b *EngineBuilder) WithPlugins(plugins ...plugin.PluginInterface) *EngineBuilder {
	b.pm = true
	b.plugins = append(b.plugins, plugins...)
	return b
}

// Build constructs and returns an EngineUniversal or an error if
// required components are missing (e.g., a string parser for a StringEngine).
// The returned engineUniversal can process strings or bytes depending
// on its type and provides methods to register commands.
func (b *EngineBuilder) Build() (*EngineUniversal, error) {
	var eu *EngineUniversal
	switch b.engineType {
	case public.StringEngineType:
		if b.stringParser == nil {
			return nil, errors.New("string parser is required for StringEngine")
		}
		strEngine := NewStringEngine(
			b.resType,
			b.pipeline,
			b.addDefaultEvents,
			b.stringParser,
			b.colorEnabled,
			b.context,
		)
		if b.logger != nil {
			strEngine.UEP.Logger = b.logger
		}
		for k, v := range b.scope {
			strEngine.UEP.Scope[k] = v
		}
		eu = &EngineUniversal{
			Plugins:        &plugin.PluginManager{},
			Type:           b.engineType,
			StringEngine:   strEngine,
			opcode_counter: 0,
		}

	case public.ByteEngineType:
		if b.byteParser == nil {
			return nil, errors.New("byte parser is required for ByteEngine")
		}
		byteEngine := NewByteEngine(
			b.resType,
			b.pipeline,
			b.addDefaultEvents,
			b.byteParser,
			b.endianess,
			b.colorEnabled,
			b.context,
		)
		if b.logger != nil {
			byteEngine.UEP.Logger = b.logger
		}
		for k, v := range b.scope {
			byteEngine.UEP.Scope[k] = v
		}
		eu = &EngineUniversal{
			Plugins:        &plugin.PluginManager{},
			Type:           b.engineType,
			ByteEngine:     byteEngine,
			opcode_counter: 0,
		}

	default:
		return nil, errors.New("EngineBuilder.Build: unknown engine type")
	}
	pm := &plugin.PluginManager{
		Plugins: make(map[string]plugin.PluginInterface),
		Scope:   core.ScopeType{public.PluginsScopeEuPtr: eu},
	}
	uep, _ := eu.GetUEP()
	if b.pm {
		if b.plugins != nil {
			for _, plugin := range b.plugins {
				err := pm.AddPlugin(plugin)
				if err != nil {
					return nil, errors.New("EngineBuilder.Build: " + err.Error())
				}
			}
			for k, v := range uep.Scope {
				pm.Scope[k] = v
			}
		}
	}
	eu.CtxCancelCause = b.cancel
	eu.ended = false
	eu.Plugins = pm
	uep.Scope[public.EuScopePmPtr] = pm
	return eu, nil
}
```

---

# engine/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# engine

'''go
import "github.com/pt-main/lc/engine"
'''

## Index

- [Constants](<#constants>)
- [type ByteEngine](<#ByteEngine>)
  - [func \(e \*ByteEngine\) AddToBytecodeIdx\(n int\)](<#ByteEngine.AddToBytecodeIdx>)
  - [func \(e \*ByteEngine\) GetBytecodeIdx\(\) \(\*int, error\)](<#ByteEngine.GetBytecodeIdx>)
  - [func \(e \*ByteEngine\) GetCommands\(\) map\[int\]byteCommandMeta](<#ByteEngine.GetCommands>)
  - [func \(e \*ByteEngine\) GetParser\(\) byteParser](<#ByteEngine.GetParser>)
  - [func \(e \*ByteEngine\) GetUep\(\) \*core.UniversalEngineParams](<#ByteEngine.GetUep>)
  - [func \(e \*ByteEngine\) NewCommand\(cmd\_switch int, handler byteCmdType, o \*core.SimpleInput\) error](<#ByteEngine.NewCommand>)
  - [func \(e \*ByteEngine\) NewCommandFull\(cmd\_switch int, handler core.CommandType\[ByteEngineInterface, byteParsing.ParsedBytes\], name string, autoBytecodeIndexShift bool\)](<#ByteEngine.NewCommandFull>)
  - [func \(e \*ByteEngine\) Process\(input \[\]byte\) core.ErrorInterface](<#ByteEngine.Process>)
  - [func \(e \*ByteEngine\) SetBytecodeIdx\(n int\)](<#ByteEngine.SetBytecodeIdx>)
- [type ByteEngineInterface](<#ByteEngineInterface>)
- [type EngineInterface](<#EngineInterface>)
- [type StringEngine](<#StringEngine>)
  - [func \(e \*StringEngine\) GetCommands\(\) map\[string\]stringCommandMeta](<#StringEngine.GetCommands>)
  - [func \(e \*StringEngine\) GetParser\(\) stringParser](<#StringEngine.GetParser>)
  - [func \(e \*StringEngine\) GetUep\(\) \*core.UniversalEngineParams](<#StringEngine.GetUep>)
  - [func \(e \*StringEngine\) NewCommand\(cmd\_switch string, handler core.CommandType\[StringEngineInterface, stringParsing.ParsedNode\], o \*core.SimpleInput\) error](<#StringEngine.NewCommand>)
  - [func \(e \*StringEngine\) NewCommandFull\(cmd\_switch string, handler core.CommandType\[StringEngineInterface, stringParsing.ParsedNode\], doc string\)](<#StringEngine.NewCommandFull>)
  - [func \(e \*StringEngine\) Process\(input string\) core.ErrorInterface](<#StringEngine.Process>)
- [type StringEngineInterface](<#StringEngineInterface>)


## Constants

<a name="AutoshiftNewCommandFlag"></a>

'''go
const AutoshiftNewCommandFlag = "autoShift"
'''

<a name="ByteEngine"></a>
## type [ByteEngine](<https://github.com/pt-main/Lc/blob/main/engine/byteEngine.go#L22-L28>)

ByteEngine handles binary inputs. Commands are indexed by integer opcodes. It uses a byte parser to decode raw bytes into ParsedBytes structures. The Process method triggers ByteParseEvent and ByteCallEvent in order.

'''go
type ByteEngine struct {
    Commands               map[int]byteCommandMeta
    Parser                 byteParser
    AutoBytecodeIndexShift map[int]bool
    UEP                    *core.UniversalEngineParams
    // contains filtered or unexported fields
}
'''

<a name="ByteEngine.AddToBytecodeIdx"></a>
### func \(\*ByteEngine\) [AddToBytecodeIdx](<https://github.com/pt-main/Lc/blob/main/engine/byteEngine.go#L109>)

'''go
func (e *ByteEngine) AddToBytecodeIdx(n int)
'''



<a name="ByteEngine.GetBytecodeIdx"></a>
### func \(\*ByteEngine\) [GetBytecodeIdx](<https://github.com/pt-main/Lc/blob/main/engine/byteEngine.go#L121>)

'''go
func (e *ByteEngine) GetBytecodeIdx() (*int, error)
'''



<a name="ByteEngine.GetCommands"></a>
### func \(\*ByteEngine\) [GetCommands](<https://github.com/pt-main/Lc/blob/main/engine/byteEngine.go#L96>)

'''go
func (e *ByteEngine) GetCommands() map[int]byteCommandMeta
'''

For interface

<a name="ByteEngine.GetParser"></a>
### func \(\*ByteEngine\) [GetParser](<https://github.com/pt-main/Lc/blob/main/engine/byteEngine.go#L105>)

'''go
func (e *ByteEngine) GetParser() byteParser
'''



<a name="ByteEngine.GetUep"></a>
### func \(\*ByteEngine\) [GetUep](<https://github.com/pt-main/Lc/blob/main/engine/byteEngine.go#L101>)

'''go
func (e *ByteEngine) GetUep() *core.UniversalEngineParams
'''

For interface

<a name="ByteEngine.NewCommand"></a>
### func \(\*ByteEngine\) [NewCommand](<https://github.com/pt-main/Lc/blob/main/engine/byteEngine.go#L77-L79>)

'''go
func (e *ByteEngine) NewCommand(cmd_switch int, handler byteCmdType, o *core.SimpleInput) error
'''

For interface. o.Option.Flags\[AutoshiftNewCommandFlag\] = autoBytecodeIndexShift, o.Input string = name

Err errors.CorePackageSystemError.

<a name="ByteEngine.NewCommandFull"></a>
### func \(\*ByteEngine\) [NewCommandFull](<https://github.com/pt-main/Lc/blob/main/engine/byteEngine.go#L62-L64>)

'''go
func (e *ByteEngine) NewCommandFull(cmd_switch int, handler core.CommandType[ByteEngineInterface, byteParsing.ParsedBytes], name string, autoBytecodeIndexShift bool)
'''

Your handler MUST shift bytecode index if autoBytecodeIndexShift false\!

Usually it's like a:

'''
AddToBytecodeIdx(1) // next instruction
'''

Or:

'''
SetBytecodeIdx(10) // jump
AddToBytecodeIdx(-1) // prev instruction
'''

<a name="ByteEngine.Process"></a>
### func \(\*ByteEngine\) [Process](<https://github.com/pt-main/Lc/blob/main/engine/byteEngine.go#L35>)

'''go
func (e *ByteEngine) Process(input []byte) core.ErrorInterface
'''

Process transforms a byte slice by parsing it and invoking the registered bytecode handlers.

Err errors.ByteEngineProcessError1 | errors.ByteEngineProcessError2. \(cause from 'CallEvents'\)

<a name="ByteEngine.SetBytecodeIdx"></a>
### func \(\*ByteEngine\) [SetBytecodeIdx](<https://github.com/pt-main/Lc/blob/main/engine/byteEngine.go#L115>)

'''go
func (e *ByteEngine) SetBytecodeIdx(n int)
'''



<a name="ByteEngineInterface"></a>
## type [ByteEngineInterface](<https://github.com/pt-main/Lc/blob/main/engine/interface.go#L22>)



'''go
type ByteEngineInterface = EngineInterface[int, []byte, byteParsing.ParsedBytes]
'''

<a name="EngineInterface"></a>
## type [EngineInterface](<https://github.com/pt-main/Lc/blob/main/engine/interface.go#L10-L19>)



'''go
type EngineInterface[CmdT int | string | byte | float32 | float64,
    ParserInput any, ParserOutput any] interface {
    Process(ParserInput) core.ErrorInterface
    NewCommand(CmdT, core.CommandType[EngineInterface[
        CmdT, ParserInput, ParserOutput], ParserOutput], *core.SimpleInput) error
    GetUep() *core.UniversalEngineParams
    GetParser() parsing.ParserInterface[ParserInput, ParserOutput]
    GetCommands() map[CmdT]core.CommandMeta[EngineInterface[
        CmdT, ParserInput, ParserOutput], ParserOutput]
}
'''

<a name="StringEngine"></a>
## type [StringEngine](<https://github.com/pt-main/Lc/blob/main/engine/stringEngine.go#L19-L24>)

StringEngine is the core for text‑based languages. It holds command definitions, a parser, and universal engine parameters \(UEP\) that include generator, events, scope, and logger. The Process method drives compilation.

'''go
type StringEngine struct {
    Commands map[string]stringCommandMeta
    Parser   stringParser
    UEP      *core.UniversalEngineParams
    // contains filtered or unexported fields
}
'''

<a name="StringEngine.GetCommands"></a>
### func \(\*StringEngine\) [GetCommands](<https://github.com/pt-main/Lc/blob/main/engine/stringEngine.go#L82>)

'''go
func (e *StringEngine) GetCommands() map[string]stringCommandMeta
'''

For interface

<a name="StringEngine.GetParser"></a>
### func \(\*StringEngine\) [GetParser](<https://github.com/pt-main/Lc/blob/main/engine/stringEngine.go#L60>)

'''go
func (e *StringEngine) GetParser() stringParser
'''



<a name="StringEngine.GetUep"></a>
### func \(\*StringEngine\) [GetUep](<https://github.com/pt-main/Lc/blob/main/engine/stringEngine.go#L87>)

'''go
func (e *StringEngine) GetUep() *core.UniversalEngineParams
'''

For interface

<a name="StringEngine.NewCommand"></a>
### func \(\*StringEngine\) [NewCommand](<https://github.com/pt-main/Lc/blob/main/engine/stringEngine.go#L65-L67>)

'''go
func (e *StringEngine) NewCommand(cmd_switch string, handler core.CommandType[StringEngineInterface, stringParsing.ParsedNode], o *core.SimpleInput) error
'''

For interface. o.Input string = doc

<a name="StringEngine.NewCommandFull"></a>
### func \(\*StringEngine\) [NewCommandFull](<https://github.com/pt-main/Lc/blob/main/engine/stringEngine.go#L50-L51>)

'''go
func (e *StringEngine) NewCommandFull(cmd_switch string, handler core.CommandType[StringEngineInterface, stringParsing.ParsedNode], doc string)
'''



<a name="StringEngine.Process"></a>
### func \(\*StringEngine\) [Process](<https://github.com/pt-main/Lc/blob/main/engine/stringEngine.go#L33>)

'''go
func (e *StringEngine) Process(input string) core.ErrorInterface
'''

Process executes the compilation pipeline for a string input. It stores the input in scope\["input\_string"\], then calls the StringParseEvent \(to parse into \[\]ParsedNode\) and StringCallEvent \(to dispatch commands\). Any error stops execution.

Err errors.StringEngineProcessError1 | errors.StringEngineProcessError2. \(cause from 'CallEvents'\)

<a name="StringEngineInterface"></a>
## type [StringEngineInterface](<https://github.com/pt-main/Lc/blob/main/engine/interface.go#L21>)



'''go
type StringEngineInterface = EngineInterface[string, string, stringParsing.ParsedNode]
'''

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# engine/byteEngine.go

```go
package engine

import (
	"sync"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

const AutoshiftNewCommandFlag = "autoShift"

type byteParser = parsing.ParserInterface[[]byte, byteParsing.ParsedBytes]
type byteCmdType = core.CommandType[ByteEngineInterface, byteParsing.ParsedBytes]
type byteCommandMeta = core.CommandMeta[ByteEngineInterface, byteParsing.ParsedBytes]

// ByteEngine handles binary inputs. Commands are indexed by integer opcodes.
// It uses a byte parser to decode raw bytes into ParsedBytes structures.
// The Process method triggers ByteParseEvent and ByteCallEvent in order.
type ByteEngine struct {
	Commands               map[int]byteCommandMeta
	Parser                 byteParser
	AutoBytecodeIndexShift map[int]bool
	UEP                    *core.UniversalEngineParams
	mu                     sync.RWMutex
}

// Process transforms a byte slice by parsing it and invoking the registered
// bytecode handlers.
//
// Err errors.ByteEngineProcessError1 | errors.ByteEngineProcessError2.
// (cause from 'CallEvents')
func (e *ByteEngine) Process(input []byte) core.ErrorInterface {
	e.UEP.Scope[public.ByteEngineScopeInput] = input
	err1 := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.ByteParseEvent, false)
	if err1 != nil {
		return core.Wrap(errors.ByteEngineProcessError1, err1, core.GetRealError(err1))
	}
	err2 := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.ByteCallEvent, false)
	if err2 != nil {
		return core.Wrap(errors.ByteEngineProcessError2, err2, core.GetRealError(err2))
	}
	return nil
}

// Your handler MUST shift bytecode index if autoBytecodeIndexShift false!
//
// Usually it's like a:
//
//	AddToBytecodeIdx(1) // next instruction
//
// Or:
//
//	SetBytecodeIdx(10) // jump
//	AddToBytecodeIdx(-1) // prev instruction
func (e *ByteEngine) NewCommandFull(
	cmd_switch int, handler core.CommandType[ByteEngineInterface, byteParsing.ParsedBytes],
	name string, autoBytecodeIndexShift bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Commands[cmd_switch] = core.CommandMeta[ByteEngineInterface, byteParsing.ParsedBytes]{
		Handler: handler,
		Doc:     name,
	}
	e.AutoBytecodeIndexShift[cmd_switch] = autoBytecodeIndexShift
}

// For interface. o.Option.Flags[AutoshiftNewCommandFlag] = autoBytecodeIndexShift, o.Input string = name
//
// Err errors.CorePackageSystemError.
func (e *ByteEngine) NewCommand(
	cmd_switch int, handler byteCmdType,
	o *core.SimpleInput) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	name, ok := o.Input.(string)
	if !ok {
		return core.Err(errors.CorePackageSystemError, "Invalid input: 'o.Input' must be string")
	}
	autoBytecodeIndexShift := o.Option.HasFlag(AutoshiftNewCommandFlag)
	e.Commands[cmd_switch] = core.CommandMeta[ByteEngineInterface, byteParsing.ParsedBytes]{
		Handler: handler,
		Doc:     name,
	}
	e.AutoBytecodeIndexShift[cmd_switch] = autoBytecodeIndexShift
	return nil
}

// For interface
func (e *ByteEngine) GetCommands() map[int]byteCommandMeta {
	return e.Commands
}

// For interface
func (e *ByteEngine) GetUep() *core.UniversalEngineParams {
	return e.UEP
}

func (e *ByteEngine) GetParser() byteParser {
	return e.Parser
}

func (e *ByteEngine) AddToBytecodeIdx(n int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	*e.UEP.Scope[public.ByteEngineScopeBytecodeIdx].(*int) += n
}

func (e *ByteEngine) SetBytecodeIdx(n int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	*e.UEP.Scope[public.ByteEngineScopeBytecodeIdx].(*int) = n
}

func (e *ByteEngine) GetBytecodeIdx() (*int, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	_idx, ok := e.UEP.Scope[public.ByteEngineScopeBytecodeIdx]
	if !ok {
		return nil, core.Err(errors.CorePackageSystemError, "Can't get bytecode index: invalid scope")
	}
	idx, ok := _idx.(*int)
	if !ok {
		return nil, core.Err(errors.CorePackageSystemError, "Can't get bytecode index: invalid interface in scope")
	}
	return idx, nil
}
```

---

# engine/core/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# core

'''go
import "github.com/pt-main/lc/engine/core"
'''

## Index

- [Variables](<#variables>)
- [func EMK\(n int, valType string\) errors.ErrorMetaType](<#EMK>)
- [func GetMetaValue\[T any\]\(in error, n int, valType string\) \(res T, err error\)](<#GetMetaValue>)
- [func GetRealError\(err error\) string](<#GetRealError>)
- [func GetRealErrorReverse\(err error\) string](<#GetRealErrorReverse>)
- [func NewUniversalEngineParams\(generator \*Generator, events \*Events, scope ScopeType, logger \*Logger, context context.Context\) \(\*UniversalEngineParams, ErrorInterface\)](<#NewUniversalEngineParams>)
- [type CommandMeta](<#CommandMeta>)
- [type CommandType](<#CommandType>)
- [type Error](<#Error>)
  - [func Err\(code errors.ErrorCodeType, format string, args ...interface\{\}\) \*Error](<#Err>)
  - [func Wrap\(code errors.ErrorCodeType, cause error, format string, args ...interface\{\}\) \*Error](<#Wrap>)
  - [func \(e \*Error\) Error\(\) string](<#Error.Error>)
  - [func \(e \*Error\) Format\(\) string](<#Error.Format>)
  - [func \(e \*Error\) GetCode\(\) string](<#Error.GetCode>)
  - [func \(e \*Error\) GetMeta\(\) map\[errors.ErrorMetaType\]interface\{\}](<#Error.GetMeta>)
  - [func \(e \*Error\) GetMsg\(\) string](<#Error.GetMsg>)
  - [func \(e \*Error\) Unwrap\(\) error](<#Error.Unwrap>)
  - [func \(e \*Error\) WithMeta\(key errors.ErrorMetaType, value interface\{\}\) \*Error](<#Error.WithMeta>)
- [type ErrorInterface](<#ErrorInterface>)
  - [func GetErr\(ei ErrorInterface\) \(res ErrorInterface\)](<#GetErr>)
  - [func GetStringRes\(g \*Generator, sep string\) \(string, ErrorInterface\)](<#GetStringRes>)
  - [func ScopeGet\[T any\]\(st ScopeType, what string\) \(T, ErrorInterface\)](<#ScopeGet>)
- [type EventInput](<#EventInput>)
- [type EventType](<#EventType>)
- [type Events](<#Events>)
  - [func NewEvents\(context context.Context\) \*Events](<#NewEvents>)
  - [func \(e \*Events\) CallEvents\(input \*EventInput, name string, canWorkWithoutHandler bool\) ErrorInterface](<#Events.CallEvents>)
  - [func \(e \*Events\) CoreEvents\(\) map\[string\]int](<#Events.CoreEvents>)
  - [func \(e \*Events\) GetCoreEventIdx\(name string\) \(int, ErrorInterface\)](<#Events.GetCoreEventIdx>)
  - [func \(e \*Events\) GetEvents\(name string\) \(\[\]EventType, ErrorInterface\)](<#Events.GetEvents>)
  - [func \(e \*Events\) NewEvent\(name string, event EventType\)](<#Events.NewEvent>)
  - [func \(e \*Events\) NewEventBefore\(name string, event EventType\) ErrorInterface](<#Events.NewEventBefore>)
  - [func \(e \*Events\) ReplaceEvent\(name string\)](<#Events.ReplaceEvent>)
  - [func \(e \*Events\) Scope\(\) ScopeType](<#Events.Scope>)
  - [func \(e \*Events\) SetEvents\(name string, events \[\]EventType, coreEvent int\)](<#Events.SetEvents>)
  - [func \(e \*Events\) SetProperty\(name string, value interface\{\}\) ErrorInterface](<#Events.SetProperty>)
- [type EventsInterface](<#EventsInterface>)
- [type EventsTools](<#EventsTools>)
  - [func \(et \*EventsTools\) ChangeCoreEvent\(name string, event EventType\) ErrorInterface](<#EventsTools.ChangeCoreEvent>)
  - [func \(et \*EventsTools\) GetCoreEvent\(name string\) \(EventType, ErrorInterface\)](<#EventsTools.GetCoreEvent>)
- [type Generator](<#Generator>)
  - [func NewGenerator\(res\_type public.ResType, pipeline \[\]string\) \*Generator](<#NewGenerator>)
  - [func \(g \*Generator\) AddBytes\(code \[\]byte, point string\) ErrorInterface](<#Generator.AddBytes>)
  - [func \(g \*Generator\) AddString\(code string, point string\) ErrorInterface](<#Generator.AddString>)
  - [func \(g \*Generator\) AddStrings\(code \[\]string, point string\) ErrorInterface](<#Generator.AddStrings>)
  - [func \(g \*Generator\) GetBytesRes\(\) \(\[\]byte, ErrorInterface\)](<#Generator.GetBytesRes>)
  - [func \(g \*Generator\) GetStringArrRes\(\) \(\[\]string, ErrorInterface\)](<#Generator.GetStringArrRes>)
- [type Logger](<#Logger>)
  - [func NewLogger\(defaultStatusForm string\) \*Logger](<#NewLogger>)
  - [func \(l \*Logger\) GetLog\(\) string](<#Logger.GetLog>)
  - [func \(l \*Logger\) GetStatusForm\(status string\) string](<#Logger.GetStatusForm>)
  - [func \(l \*Logger\) PrintLog\(status string, message string\)](<#Logger.PrintLog>)
- [type Option](<#Option>)
  - [func \(o \*Option\) HasFlag\(name string\) bool](<#Option.HasFlag>)
- [type ScopeType](<#ScopeType>)
- [type SimpleInput](<#SimpleInput>)
- [type UniversalEngineParams](<#UniversalEngineParams>)
  - [func \(p \*UniversalEngineParams\) GetContext\(\) context.Context](<#UniversalEngineParams.GetContext>)


## Variables

<a name="ErrExit"></a>

'''go
var ErrExit = Err(errors.ErrExit, "")
'''

<a name="EMK"></a>
## func [EMK](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L110>)

'''go
func EMK(n int, valType string) errors.ErrorMetaType
'''

Error Meta Key n

<a name="GetMetaValue"></a>
## func [GetMetaValue](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L114>)

'''go
func GetMetaValue[T any](in error, n int, valType string) (res T, err error)
'''



<a name="GetRealError"></a>
## func [GetRealError](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L134>)

'''go
func GetRealError(err error) string
'''



<a name="GetRealErrorReverse"></a>
## func [GetRealErrorReverse](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L161>)

'''go
func GetRealErrorReverse(err error) string
'''



<a name="NewUniversalEngineParams"></a>
## func [NewUniversalEngineParams](<https://github.com/pt-main/Lc/blob/main/engine/core/universalEngineParams.go#L46-L52>)

'''go
func NewUniversalEngineParams(generator *Generator, events *Events, scope ScopeType, logger *Logger, context context.Context) (*UniversalEngineParams, ErrorInterface)
'''

NewUniversalEngineParams constructs an initialized UniversalEngineParams. It automatically injects two event handlers into the Events system:

- CallEventsStartEvent \- logs the start of any event call \(debug level\)
- CallEventsEndEvent \- logs the end, including any error

Parameters: generator, events, scope, logger. All must be non‑nil. Returns a filled struct or an error if event registration fails.

Err errors.CorePackageSystemError

<a name="CommandMeta"></a>
## type [CommandMeta](<https://github.com/pt-main/Lc/blob/main/engine/core/types.go#L9-L12>)



'''go
type CommandMeta[EI, N any] struct {
    Handler CommandType[EI, N]
    Doc     string
}
'''

<a name="CommandType"></a>
## type [CommandType](<https://github.com/pt-main/Lc/blob/main/engine/core/types.go#L7>)



'''go
type CommandType[EI, N any] func(EI, *N) ErrorInterface
'''

<a name="Error"></a>
## type [Error](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L12-L17>)



'''go
type Error struct {
    Code  errors.ErrorCodeType
    Msg   string
    Meta  map[errors.ErrorMetaType]interface{}
    Cause error
}
'''

<a name="Err"></a>
### func [Err](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L93>)

'''go
func Err(code errors.ErrorCodeType, format string, args ...interface{}) *Error
'''



<a name="Wrap"></a>
### func [Wrap](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L97>)

'''go
func Wrap(code errors.ErrorCodeType, cause error, format string, args ...interface{}) *Error
'''



<a name="Error.Error"></a>
### func \(\*Error\) [Error](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L28>)

'''go
func (e *Error) Error() string
'''



<a name="Error.Format"></a>
### func \(\*Error\) [Format](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L52>)

'''go
func (e *Error) Format() string
'''



<a name="Error.GetCode"></a>
### func \(\*Error\) [GetCode](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L36>)

'''go
func (e *Error) GetCode() string
'''



<a name="Error.GetMeta"></a>
### func \(\*Error\) [GetMeta](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L48>)

'''go
func (e *Error) GetMeta() map[errors.ErrorMetaType]interface{}
'''



<a name="Error.GetMsg"></a>
### func \(\*Error\) [GetMsg](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L44>)

'''go
func (e *Error) GetMsg() string
'''



<a name="Error.Unwrap"></a>
### func \(\*Error\) [Unwrap](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L40>)

'''go
func (e *Error) Unwrap() error
'''



<a name="Error.WithMeta"></a>
### func \(\*Error\) [WithMeta](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L101>)

'''go
func (e *Error) WithMeta(key errors.ErrorMetaType, value interface{}) *Error
'''



<a name="ErrorInterface"></a>
## type [ErrorInterface](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L19-L26>)



'''go
type ErrorInterface interface {
    Error() string
    Format() string
    GetCode() string
    GetMsg() string
    GetMeta() map[errors.ErrorMetaType]interface{}
    Unwrap() error
}
'''

<a name="GetErr"></a>
### func [GetErr](<https://github.com/pt-main/Lc/blob/main/engine/core/errors.go#L145>)

'''go
func GetErr(ei ErrorInterface) (res ErrorInterface)
'''



<a name="GetStringRes"></a>
### func [GetStringRes](<https://github.com/pt-main/Lc/blob/main/engine/core/generator.go#L110>)

'''go
func GetStringRes(g *Generator, sep string) (string, ErrorInterface)
'''

Err errors.GeneratorGenerationTypeError.

<a name="ScopeGet"></a>
### func [ScopeGet](<https://github.com/pt-main/Lc/blob/main/engine/core/scope.go#L9>)

'''go
func ScopeGet[T any](st ScopeType, what string) (T, ErrorInterface)
'''

Err errors.ScopeGetError. With meta: EMK\(0, "string"\) \- key

<a name="EventInput"></a>
## type [EventInput](<https://github.com/pt-main/Lc/blob/main/engine/core/types.go#L30>)



'''go
type EventInput = SimpleInput
'''

<a name="EventType"></a>
## type [EventType](<https://github.com/pt-main/Lc/blob/main/engine/core/types.go#L14>)



'''go
type EventType func(*Events, *EventInput) ErrorInterface
'''

<a name="Events"></a>
## type [Events](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L28-L35>)

Events manages an ordered collection of event handlers. Each event has a name \(string\) and a list of EventType functions. The CallEvents method invokes all handlers of an event in registration order. Events can also automatically wrap calls with start/end events for logging.

'''go
type Events struct {
    Context context.Context
    // contains filtered or unexported fields
}
'''

<a name="NewEvents"></a>
### func [NewEvents](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L174>)

'''go
func NewEvents(context context.Context) *Events
'''

NewEvents creates an empty Events instance with an ordered map. The Scope map is initially empty but can be used to pass data between event handlers.

<a name="Events.CallEvents"></a>
### func \(\*Events\) [CallEvents](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L122-L123>)

'''go
func (e *Events) CallEvents(input *EventInput, name string, canWorkWithoutHandler bool) ErrorInterface
'''

Err errors.EventsEventError \(from 'callEvents'\)

<a name="Events.CoreEvents"></a>
### func \(\*Events\) [CoreEvents](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L69>)

'''go
func (e *Events) CoreEvents() map[string]int
'''



<a name="Events.GetCoreEventIdx"></a>
### func \(\*Events\) [GetCoreEventIdx](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L59>)

'''go
func (e *Events) GetCoreEventIdx(name string) (int, ErrorInterface)
'''

Err errors.EventsEventIsNotFound, Msg=name

<a name="Events.GetEvents"></a>
### func \(\*Events\) [GetEvents](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L38>)

'''go
func (e *Events) GetEvents(name string) ([]EventType, ErrorInterface)
'''

Err errors.EventsEventIsNotFound, Msg=name

<a name="Events.NewEvent"></a>
### func \(\*Events\) [NewEvent](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L73>)

'''go
func (e *Events) NewEvent(name string, event EventType)
'''



<a name="Events.NewEventBefore"></a>
### func \(\*Events\) [NewEventBefore](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L88>)

'''go
func (e *Events) NewEventBefore(name string, event EventType) ErrorInterface
'''

Err errors.EventsSystemError. With meta: EMK\(0, "string"\) \- event name, EMK\(1, "error"\) \(from 'GetEvents'\)

<a name="Events.ReplaceEvent"></a>
### func \(\*Events\) [ReplaceEvent](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L147>)

'''go
func (e *Events) ReplaceEvent(name string)
'''



<a name="Events.Scope"></a>
### func \(\*Events\) [Scope](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L167>)

'''go
func (e *Events) Scope() ScopeType
'''



<a name="Events.SetEvents"></a>
### func \(\*Events\) [SetEvents](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L48>)

'''go
func (e *Events) SetEvents(name string, events []EventType, coreEvent int)
'''



<a name="Events.SetProperty"></a>
### func \(\*Events\) [SetProperty](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L153>)

'''go
func (e *Events) SetProperty(name string, value interface{}) ErrorInterface
'''

Err errors.EventsSystemError

<a name="EventsInterface"></a>
## type [EventsInterface](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L11-L22>)



'''go
type EventsInterface interface {
    GetEvents(name string) ([]EventType, ErrorInterface)
    GetCoreEventIdx(name string) (int, ErrorInterface)
    SetEvents(name string, events []EventType, idx int)
    NewEvent(name string, event EventType)
    NewEventBefore(name string, event EventType) ErrorInterface
    CallEvents(input *EventInput, name string, canWorkWithoutHandler bool) ErrorInterface
    Scope() ScopeType
    CoreEvents() map[string]int
    ReplaceEvent(name string)
    SetProperty(name string, value interface{}) ErrorInterface
}
'''

<a name="EventsTools"></a>
## type [EventsTools](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L183-L185>)



'''go
type EventsTools struct {
    Events EventsInterface
}
'''

<a name="EventsTools.ChangeCoreEvent"></a>
### func \(\*EventsTools\) [ChangeCoreEvent](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L187>)

'''go
func (et *EventsTools) ChangeCoreEvent(name string, event EventType) ErrorInterface
'''



<a name="EventsTools.GetCoreEvent"></a>
### func \(\*EventsTools\) [GetCoreEvent](<https://github.com/pt-main/Lc/blob/main/engine/core/events.go#L207>)

'''go
func (et *EventsTools) GetCoreEvent(name string) (EventType, ErrorInterface)
'''



<a name="Generator"></a>
## type [Generator](<https://github.com/pt-main/Lc/blob/main/engine/core/generator.go#L16-L21>)

Generator accumulates code fragments \(strings or bytes\) into named points \(e.g., "pre", "main"\). The Pipeline defines the order in which points are emitted. It supports both text and binary generation modes via res\_type. Thread‑safe due to internal mutex.

'''go
type Generator struct {
    Pipeline []string
    // contains filtered or unexported fields
}
'''

<a name="NewGenerator"></a>
### func [NewGenerator](<https://github.com/pt-main/Lc/blob/main/engine/core/generator.go#L128>)

'''go
func NewGenerator(res_type public.ResType, pipeline []string) *Generator
'''

NewGenerator initializes a Generator with a given resource type \(StringResType or ByteResType\) and a pipeline slice. Empty code slices are pre‑created for each pipeline point.

<a name="Generator.AddBytes"></a>
### func \(\*Generator\) [AddBytes](<https://github.com/pt-main/Lc/blob/main/engine/core/generator.go#L51>)

'''go
func (g *Generator) AddBytes(code []byte, point string) ErrorInterface
'''

Err errors.GeneratorGenerationTypeError.

<a name="Generator.AddString"></a>
### func \(\*Generator\) [AddString](<https://github.com/pt-main/Lc/blob/main/engine/core/generator.go#L42>)

'''go
func (g *Generator) AddString(code string, point string) ErrorInterface
'''

Err errors.GeneratorGenerationTypeError.

<a name="Generator.AddStrings"></a>
### func \(\*Generator\) [AddStrings](<https://github.com/pt-main/Lc/blob/main/engine/core/generator.go#L33>)

'''go
func (g *Generator) AddStrings(code []string, point string) ErrorInterface
'''

Err errors.GeneratorGenerationTypeError.

<a name="Generator.GetBytesRes"></a>
### func \(\*Generator\) [GetBytesRes](<https://github.com/pt-main/Lc/blob/main/engine/core/generator.go#L61>)

'''go
func (g *Generator) GetBytesRes() ([]byte, ErrorInterface)
'''

Err errors.GeneratorGenerationTypeError. Error 'Can't find ...' with EMK\(0, "string"\) \- point

<a name="Generator.GetStringArrRes"></a>
### func \(\*Generator\) [GetStringArrRes](<https://github.com/pt-main/Lc/blob/main/engine/core/generator.go#L85>)

'''go
func (g *Generator) GetStringArrRes() ([]string, ErrorInterface)
'''

Err errors.GeneratorGenerationTypeError. Error 'Can't find ...' with EMK\(0, "string"\) \- point

<a name="Logger"></a>
## type [Logger](<https://github.com/pt-main/Lc/blob/main/engine/core/logger.go#L16-L23>)

Logger is a thread\-safe structured logger for engine diagnostics. It stores a list of log lines, supports custom status formats, and allows different formatting per status \(e.g., "error", "info", "debug"\). Typical usage: attach to UniversalEngineParams.Logger.

'''go
type Logger struct {
    Logging           map[string]bool
    Log               []string
    MaxLogLength      int
    Statuses          map[string]string
    DefaultStatusForm string
    // contains filtered or unexported fields
}
'''

<a name="NewLogger"></a>
### func [NewLogger](<https://github.com/pt-main/Lc/blob/main/engine/core/logger.go#L65>)

'''go
func NewLogger(defaultStatusForm string) *Logger
'''

NewLogger creates a new Logger with an optional defaultStatusForm. The format uses three placeholders: %s for status, %v for timestamp, and %s for the message. Example default: "%s \[%v\] \[%s\]\\n" If empty string is passed, the default format is used.

<a name="Logger.GetLog"></a>
### func \(\*Logger\) [GetLog](<https://github.com/pt-main/Lc/blob/main/engine/core/logger.go#L57>)

'''go
func (l *Logger) GetLog() string
'''

GetLog returns the entire log as a single string with newline separators. It is useful for saving logs to a file or showing them after execution.

<a name="Logger.GetStatusForm"></a>
### func \(\*Logger\) [GetStatusForm](<https://github.com/pt-main/Lc/blob/main/engine/core/logger.go#L28>)

'''go
func (l *Logger) GetStatusForm(status string) string
'''

GetStatusForm returns the format string associated with the given status, falling back to DefaultStatusForm if no custom format is set. Custom formats can be registered by directly assigning to l.Statuses map.

<a name="Logger.PrintLog"></a>
### func \(\*Logger\) [PrintLog](<https://github.com/pt-main/Lc/blob/main/engine/core/logger.go#L41>)

'''go
func (l *Logger) PrintLog(status string, message string)
'''

PrintLog writes a log entry to stdout and appends it to the internal slice. The status string determines the format via GetStatusForm \(if a custom format for that status exists\). The message is inserted into the format. Example: logger.PrintLog\("error", "failed to parse token"\)

<a name="Option"></a>
## type [Option](<https://github.com/pt-main/Lc/blob/main/engine/core/types.go#L16-L19>)



'''go
type Option struct {
    Flags []string
    Scope ScopeType
}
'''

<a name="Option.HasFlag"></a>
### func \(\*Option\) [HasFlag](<https://github.com/pt-main/Lc/blob/main/engine/core/types.go#L21>)

'''go
func (o *Option) HasFlag(name string) bool
'''



<a name="ScopeType"></a>
## type [ScopeType](<https://github.com/pt-main/Lc/blob/main/engine/core/scope.go#L5>)



'''go
type ScopeType map[string]interface{}
'''

<a name="SimpleInput"></a>
## type [SimpleInput](<https://github.com/pt-main/Lc/blob/main/engine/core/types.go#L25-L28>)



'''go
type SimpleInput struct {
    Option *Option
    Input  interface{}
}
'''

<a name="UniversalEngineParams"></a>
## type [UniversalEngineParams](<https://github.com/pt-main/Lc/blob/main/engine/core/universalEngineParams.go#L15-L28>)

UniversalEngineParams is a container shared by both StringEngine and ByteEngine. It holds the Generator \(for output accumulation\), Events \(for hooks\), Scope \(for passing arbitrary data between stages\), and Logger \(for diagnostics\). This struct is embedded \(not composed by pointer\) in the engine types, promoting its fields and methods to the engine itself.

'''go
type UniversalEngineParams struct {
    // Generator *Generator - controls code/output generation across pipeline points.
    Generator *Generator
    // Event *Events - allows hooking into parsing and command dispatch.
    Event EventsInterface
    // Scope ScopeType - a map[string]interface{} that can be used to share
    //   variables between events, parsers, and command handlers.
    Scope ScopeType
    // Logger *Logger - if set, logs internal steps (event calls, errors).
    Logger *Logger

    Context context.Context
}
'''

<a name="UniversalEngineParams.GetContext"></a>
### func \(\*UniversalEngineParams\) [GetContext](<https://github.com/pt-main/Lc/blob/main/engine/core/universalEngineParams.go#L30>)

'''go
func (p *UniversalEngineParams) GetContext() context.Context
'''



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# engine/core/errors.go

```go
package core

import (
	goerr "errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pt-main/lc/public/errors"
)

type Error struct {
	Code  errors.ErrorCodeType
	Msg   string
	Meta  map[errors.ErrorMetaType]interface{}
	Cause error
}

type ErrorInterface interface {
	Error() string
	Format() string
	GetCode() string
	GetMsg() string
	GetMeta() map[errors.ErrorMetaType]interface{}
	Unwrap() error
}

func (e *Error) Error() string {
	var b strings.Builder
	b.WriteString(string(e.Code))
	b.WriteString(": ")
	b.WriteString(e.Msg)
	return b.String()
}

func (e *Error) GetCode() string {
	return string(e.Code)
}

func (e *Error) Unwrap() error {
	return e.Cause
}

func (e *Error) GetMsg() string {
	return e.Msg
}

func (e *Error) GetMeta() map[errors.ErrorMetaType]interface{} {
	return e.Meta
}

func (e *Error) Format() string {
	var b strings.Builder
	e.writeFull(&b, "")
	return b.String()
}

func (e *Error) writeFull(b *strings.Builder, indent string) {
	b.WriteString(indent)
	b.WriteString(e.Error())
	b.WriteString("\n")

	tab := "    |"

	if len(e.Meta) > 0 {
		b.WriteString(indent)
		b.WriteString("  Meta:\n")
		for k, v := range e.Meta {
			b.WriteString(indent)
			b.WriteString(tab)
			b.WriteString(string(k))
			b.WriteString(": ")
			b.WriteString(fmt.Sprintf("%v", v))
			b.WriteString("\n")
		}
	}

	if e.Cause != nil {
		b.WriteString(indent)
		b.WriteString("  Caused by:\n")
		if ce, ok := e.Cause.(*Error); ok {
			ce.writeFull(b, indent+tab)
		} else {
			causeText := strings.ReplaceAll(e.Cause.Error(), "\n", "\n"+indent+tab)
			b.WriteString(indent)
			b.WriteString(tab)
			b.WriteString(causeText)
			b.WriteString("\n")
		}
	}
}

func Err(code errors.ErrorCodeType, format string, args ...interface{}) *Error {
	return &Error{Code: code, Msg: fmt.Sprintf(format, args...)}
}

func Wrap(code errors.ErrorCodeType, cause error, format string, args ...interface{}) *Error {
	return &Error{Code: code, Msg: fmt.Sprintf(format, args...), Cause: cause}
}

func (e *Error) WithMeta(key errors.ErrorMetaType, value interface{}) *Error {
	if e.Meta == nil {
		e.Meta = make(map[errors.ErrorMetaType]interface{})
	}
	e.Meta[key] = value
	return e
}

// Error Meta Key n
func EMK(n int, valType string) errors.ErrorMetaType {
	return errors.ErrorMetaType(strconv.Itoa(n) + "_META:" + valType)
}

func GetMetaValue[T any](in error, n int, valType string) (res T, err error) {
	newE, ok := in.(*Error)
	if !ok {
		err = Err(errors.CorePackageSystemError, "Invalid input: error is not lc *Error")
		return
	}
	key := EMK(n, valType)
	val, ok := newE.Meta[key]
	if !ok {
		err = Err(errors.CorePackageSystemError, "Key not found: %v", key)
		return
	}
	res, ok = val.(T)
	if !ok {
		err = Err(errors.CorePackageSystemError, "Invalid type: meta type and generic type is different")
		return
	}
	return
}

func GetRealError(err error) string {
	if err != nil {
		errText := err.Error()
		if ce, ok := err.(ErrorInterface); ok {
			errText = ce.Format()
		}
		return errText
	}
	return ""
}

func GetErr(ei ErrorInterface) (res ErrorInterface) {
	inner := ei.Unwrap()
	res, ok := inner.(ErrorInterface)
	if !ok {
		res = &Error{
			Code:  errors.WrappedError,
			Msg:   inner.Error(),
			Meta:  make(map[errors.ErrorMetaType]interface{}),
			Cause: nil,
		}
	}
	return
}

var ErrExit = Err(errors.ErrExit, "")

func GetRealErrorReverse(err error) string {
	if err == nil {
		return ""
	}

	if ce, ok := err.(ErrorInterface); ok {
		return ce.Format()
	}

	var parts []string
	cur := err
	for cur != nil {

		if ce, ok := cur.(ErrorInterface); ok {

			innerFormatted := ce.Format()
			if len(parts) > 0 {

				outerMsg := strings.Join(reverse(parts), ": ")
				return outerMsg + ": " + innerFormatted
			}
			return innerFormatted
		}

		parts = append(parts, cur.Error())

		next := goerr.Unwrap(cur)
		if next == nil {
			break
		}
		cur = next
	}

	if len(parts) > 0 {
		return strings.Join(reverse(parts), ": ")
	}
	return err.Error()
}

func reverse(s []string) []string {
	res := make([]string, len(s))
	for i, v := range s {
		res[len(s)-1-i] = v
	}
	return res
}
```

---

# engine/core/events.go

```go
package core

import (
	"context"
	"sync"

	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

type EventsInterface interface {
	GetEvents(name string) ([]EventType, ErrorInterface)
	GetCoreEventIdx(name string) (int, ErrorInterface)
	SetEvents(name string, events []EventType, idx int)
	NewEvent(name string, event EventType)
	NewEventBefore(name string, event EventType) ErrorInterface
	CallEvents(input *EventInput, name string, canWorkWithoutHandler bool) ErrorInterface
	Scope() ScopeType
	CoreEvents() map[string]int
	ReplaceEvent(name string)
	SetProperty(name string, value interface{}) ErrorInterface
}

// Events manages an ordered collection of event handlers. Each event has a
// name (string) and a list of EventType functions. The CallEvents method
// invokes all handlers of an event in registration order. Events can also
// automatically wrap calls with start/end events for logging.
type Events struct {
	scope      ScopeType
	Context    context.Context
	mu         sync.RWMutex
	debug      bool
	coreEvents map[string]int
	events     map[string][]EventType
}

// Err errors.EventsEventIsNotFound, Msg=name
func (e *Events) GetEvents(name string) ([]EventType, ErrorInterface) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	val, ok := e.events[name]
	if !ok {
		return nil, Err(errors.EventsEventIsNotFound, name)
	}
	return val, nil
}

func (e *Events) SetEvents(name string, events []EventType, coreEvent int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if events == nil {
		events = make([]EventType, 0)
	}
	e.events[name] = events
	e.coreEvents[name] = coreEvent
}

// Err errors.EventsEventIsNotFound, Msg=name
func (e *Events) GetCoreEventIdx(name string) (int, ErrorInterface) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	ce, ok := e.coreEvents[name]
	if !ok {
		return -1, Err(errors.EventsEventIsNotFound, name)
	}
	return ce, nil
}

func (e *Events) CoreEvents() map[string]int {
	return e.coreEvents
}

func (e *Events) NewEvent(name string, event EventType) {
	e.mu.Lock()
	defer e.mu.Unlock()

	val, ok := e.events[name]
	if !ok {
		e.events[name] = []EventType{event}
		e.coreEvents[name] = 0
		return
	}
	e.events[name] = append(val, event)
}

// Err errors.EventsSystemError.
// With meta: EMK(0, "string") - event name, EMK(1, "error") (from 'GetEvents')
func (e *Events) NewEventBefore(name string, event EventType) ErrorInterface {
	list, err := e.GetEvents(name)
	if err != nil {
		return Wrap(errors.EventsSystemError, err, "Can't put new event before '%s'", name).
			WithMeta(EMK(0, "string"), name)
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	e.events[name] = append([]EventType{event}, list...)
	e.coreEvents[name] += 1
	return nil
}

// Err errors.EventsEventError.
// Cause: errors.EventsEventIsNotFound (from 'GetEvents') or event error, msg = event ErrorInterface text
func (e *Events) callEvents(input *EventInput, name string, canWorkWithoutHandler bool) ErrorInterface {
	res, err := e.GetEvents(name)
	if err != nil {
		if canWorkWithoutHandler {
			return nil
		} else {
			return Wrap(errors.EventsEventError, err, "Can't found event")
		}
	}
	for _, event := range res {
		err := event(e, input)
		if err != nil {
			return Wrap(errors.EventsEventError, err, "Event handler failed")
		}
	}
	return nil
}

// Err errors.EventsEventError (from 'callEvents')
func (e *Events) CallEvents(input *EventInput, name string,
	canWorkWithoutHandler bool) ErrorInterface {
	e.scope[public.EventsScopeCallName] = name
	var err ErrorInterface
	if e.debug {
		err = e.callEvents(nil, public.CallEventsStartEvent, true)
		if err != nil {
			return err
		}
	}
	err = e.callEvents(input, name, canWorkWithoutHandler)
	e.scope[public.EventsScopeCallError] = err
	if e.debug {
		err1 := e.callEvents(nil, public.CallEventsEndEvent, true)
		if err1 != nil {
			return err1
		}
	}
	if err != nil {
		return err
	}
	return nil

}

func (e *Events) ReplaceEvent(name string) {
	delete(e.events, name)
	delete(e.coreEvents, name)
}

// Err errors.EventsSystemError
func (e *Events) SetProperty(name string, value interface{}) ErrorInterface {
	switch name {
	case "debug":
		var ok bool
		e.debug, ok = value.(bool)
		if !ok {
			return Err(errors.EventsSystemError, "Invalid property value (must be bool): %v", value)
		}
		return nil
	default:
		return Err(errors.EventsSystemError, "Invalid property name: %v", name)
	}
}

func (e *Events) Scope() ScopeType {
	return e.scope
}

// NewEvents creates an empty Events instance with an ordered map.
// The Scope map is initially empty but can be used to pass data between
// event handlers.
func NewEvents(context context.Context) *Events {
	return &Events{
		scope:      make(ScopeType),
		events:     make(map[string][]EventType),
		coreEvents: make(map[string]int),
		Context:    context,
	}
}

type EventsTools struct {
	Events EventsInterface
}

func (et *EventsTools) ChangeCoreEvent(name string, event EventType) ErrorInterface {
	e := et.Events
	idx, err := e.GetCoreEventIdx(name)
	if err != nil {
		return err
	}
	events, err := e.GetEvents(name)
	if err != nil {
		return err
	}
	if idx < 0 {
		return Err(errors.EventsSystemError, "Can't change core event: %v", name)
	}
	pre := events[:idx]
	post := events[idx+1:]
	done := append(append(pre, event), post...)
	e.SetEvents(name, done, idx)
	return nil
}

func (et *EventsTools) GetCoreEvent(name string) (EventType, ErrorInterface) {
	var etn EventType
	e := et.Events
	idx, err := e.GetCoreEventIdx(name)
	if err != nil {
		return etn, err
	}
	ev, err := e.GetEvents(name)
	if err != nil {
		return etn, err
	}
	if idx < 0 || idx > (len(ev)-1) {
		return etn, Err(errors.EventsSystemError, "Invalid core event idx")
	}
	return ev[idx], nil
}
```

---

# engine/core/generator.go

```go
package core

import (
	"sync"

	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

type codetype any

// Generator accumulates code fragments (strings or bytes) into named points
// (e.g., "pre", "main"). The Pipeline defines the order in which points are
// emitted. It supports both text and binary generation modes via res_type.
// Thread‑safe due to internal mutex.
type Generator struct {
	mu       sync.RWMutex
	code     map[string][]codetype
	Pipeline []string
	res_type public.ResType
}

func (g *Generator) addToCode(code codetype, point string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.code[point]; !ok {
		g.code[point] = []codetype{}
	}
	g.code[point] = append(g.code[point], code)
}

// Err errors.GeneratorGenerationTypeError.
func (g *Generator) AddStrings(code []string, point string) ErrorInterface {
	if g.res_type != public.StringResType {
		return Err(errors.GeneratorAddingTypeError, "Can't add strings to bytes")
	}
	g.addToCode(code, point)
	return nil
}

// Err errors.GeneratorGenerationTypeError.
func (g *Generator) AddString(code string, point string) ErrorInterface {
	if g.res_type != public.StringResType {
		return Err(errors.GeneratorAddingTypeError, "Can't add strings to bytes")
	}
	g.addToCode([]string{code}, point)
	return nil
}

// Err errors.GeneratorGenerationTypeError.
func (g *Generator) AddBytes(code []byte, point string) ErrorInterface {
	if g.res_type != public.ByteResType {
		return Err(errors.GeneratorAddingTypeError, "Can't add bytes to strings")
	}
	g.addToCode(code, point)
	return nil
}

// Err errors.GeneratorGenerationTypeError.
// Error 'Can't find ...' with EMK(0, "string") - point
func (g *Generator) GetBytesRes() ([]byte, ErrorInterface) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	if g.res_type != public.ByteResType {
		return nil, Err(errors.GeneratorGenerationTypeError, "Can't get bytes from generated string code")
	}
	res := []byte{}
	for _, point := range g.Pipeline {
		point_code, ok := g.code[point]
		if !ok {
			return nil, Err(errors.GeneratorGenerationTypeError, "Can't find '%v' pipeline point", point).
				WithMeta(EMK(0, "string"), point)
		}
		for _, code := range point_code {
			if bytes, ok := code.([]byte); ok {
				res = append(res, bytes...)
			}
		}
	}
	return res, nil
}

// Err errors.GeneratorGenerationTypeError.
// Error 'Can't find ...' with EMK(0, "string") - point
func (g *Generator) GetStringArrRes() ([]string, ErrorInterface) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	if g.res_type != public.StringResType {
		return nil, Err(errors.GeneratorGenerationTypeError, "Can't get string from generated bytes")
	}
	res := []string{}
	for _, point := range g.Pipeline {
		point_code, ok := g.code[point]
		if !ok {
			return nil, Err(errors.GeneratorGenerationTypeError, "Can't find '%v' pipeline point", point).
				WithMeta(EMK(0, "string"), point)
		}
		for _, code := range point_code {
			if str, ok := code.([]string); ok {
				res = append(res, str...)
			} else {
				return nil, Err(errors.GeneratorGenerationTypeError, "Unexpected type in Generator")
			}
		}
	}
	return res, nil
}

// Err errors.GeneratorGenerationTypeError.
func GetStringRes(g *Generator, sep string) (string, ErrorInterface) {
	res := ""
	arr, err := g.GetStringArrRes()
	if err != nil {
		return "", err
	}
	for idx, val := range arr {
		res += val
		if idx != (len(arr) - 1) {
			res += sep
		}
	}
	return res, nil
}

// NewGenerator initializes a Generator with a given resource type
// (StringResType or ByteResType) and a pipeline slice. Empty code slices are
// pre‑created for each pipeline point.
func NewGenerator(res_type public.ResType, pipeline []string) *Generator {
	g := &Generator{
		code:     map[string][]codetype{},
		Pipeline: pipeline,
		res_type: res_type,
	}
	for _, point := range pipeline {
		g.code[point] = []codetype{}
	}
	return g
}
```

---

# engine/core/generator_test.go

```go
package core

import (
	"reflect"
	"testing"

	"github.com/pt-main/lc/public"
)

func TestGenerator_AddAndGetString(t *testing.T) {
	gen := NewGenerator(public.StringResType, []string{"pre", "main"})

	err := gen.AddString("hello", "pre")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = gen.AddStrings([]string{" ", "world"}, "main")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	res, err := GetStringRes(gen, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "hello world"
	if res != expected {
		t.Errorf("got %q, want %q", res, expected)
	}
}

func TestGenerator_AddBytes(t *testing.T) {
	gen := NewGenerator(public.ByteResType, []string{"main"})
	err := gen.AddBytes([]byte{0x01, 0x02}, "main")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	res, err := gen.GetBytesRes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []byte{0x01, 0x02}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("got %v, want %v", res, expected)
	}
}

func TestGenerator_WrongType(t *testing.T) {
	gen := NewGenerator(public.StringResType, []string{"main"})
	err := gen.AddBytes([]byte{0x01}, "main")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
```

---

# engine/core/logger.go

```go
package core

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/pt-main/tap/color"
)

// Logger is a thread-safe structured logger for engine diagnostics.
// It stores a list of log lines, supports custom status formats, and
// allows different formatting per status (e.g., "error", "info", "debug").
// Typical usage: attach to UniversalEngineParams.Logger.
type Logger struct {
	mu                sync.RWMutex
	Logging           map[string]bool
	Log               []string
	MaxLogLength      int
	Statuses          map[string]string
	DefaultStatusForm string
}

// GetStatusForm returns the format string associated with the given status,
// falling back to DefaultStatusForm if no custom format is set.
// Custom formats can be registered by directly assigning to l.Statuses map.
func (l *Logger) GetStatusForm(status string) string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if s, ok := l.Statuses[status]; ok {
		return s
	}
	return l.DefaultStatusForm
}

// PrintLog writes a log entry to stdout and appends it to the internal slice.
// The status string determines the format via GetStatusForm (if a custom
// format for that status exists). The message is inserted into the format.
// Example: logger.PrintLog("error", "failed to parse token")
func (l *Logger) PrintLog(status string, message string) {
	format := l.GetStatusForm(status)
	line := fmt.Sprintf(format, status, time.Now().UTC(), message)
	if ok, val := l.Logging[status]; ok && val {
		fmt.Println(color.Set(line))
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Log = append(l.Log, color.ReplaceColors(line))
	if len(l.Log) > l.MaxLogLength && l.MaxLogLength > 0 {
		l.Log = l.Log[1:]
	}
}

// GetLog returns the entire log as a single string with newline separators.
// It is useful for saving logs to a file or showing them after execution.
func (l *Logger) GetLog() string {
	return strings.Join(l.Log, "\n")
}

// NewLogger creates a new Logger with an optional defaultStatusForm.
// The format uses three placeholders: %s for status, %v for timestamp,
// and %s for the message. Example default: "%s [%v] [%s]\n"
// If empty string is passed, the default format is used.
func NewLogger(defaultStatusForm string) *Logger {
	if defaultStatusForm == "" {
		defaultStatusForm = "[?BE]%s[?RT] [?CN][%v][?RT] [?GN][%s][?RT]\n"
	}
	return &Logger{
		Log:               make([]string, 0),
		Statuses:          make(map[string]string),
		DefaultStatusForm: defaultStatusForm,
		Logging:           make(map[string]bool),
		MaxLogLength:      -1,
	}
}
```

---

# engine/core/scope.go

```go
package core

import "github.com/pt-main/lc/public/errors"

type ScopeType map[string]interface{}

// Err errors.ScopeGetError.
// With meta: EMK(0, "string") - key
func ScopeGet[T any](st ScopeType, what string) (T, ErrorInterface) {
	var nul T
	val, ok := st[what]
	if !ok {
		return nul, Err(errors.ScopeGetError, "Invalid key: %v", what).
			WithMeta(EMK(0, "string"), what)
	}
	res, ok := val.(T)
	if !ok {
		return nul, Err(errors.ScopeGetError, "Invalid type for key: %v", what).
			WithMeta(EMK(0, "string"), what)
	}
	return res, nil
}
```

---

# engine/core/types.go

```go
package core

import (
	"slices"
)

type CommandType[EI, N any] func(EI, *N) ErrorInterface

type CommandMeta[EI, N any] struct {
	Handler CommandType[EI, N]
	Doc     string
}

type EventType func(*Events, *EventInput) ErrorInterface

type Option struct {
	Flags []string
	Scope ScopeType
}

func (o *Option) HasFlag(name string) bool {
	return slices.Contains(o.Flags, name)
}

type SimpleInput struct {
	Option *Option
	Input  interface{}
}

type EventInput = SimpleInput
```

---

# engine/core/universalEngineParams.go

```go
package core

import (
	"context"

	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

// UniversalEngineParams is a container shared by both StringEngine and ByteEngine.
// It holds the Generator (for output accumulation), Events (for hooks),
// Scope (for passing arbitrary data between stages), and Logger (for diagnostics).
// This struct is embedded (not composed by pointer) in the engine types,
// promoting its fields and methods to the engine itself.
type UniversalEngineParams struct {
	// Generator *Generator - controls code/output generation across pipeline points.
	Generator *Generator
	// Event *Events - allows hooking into parsing and command dispatch.
	Event EventsInterface
	// Scope ScopeType - a map[string]interface{} that can be used to share
	//   variables between events, parsers, and command handlers.
	Scope ScopeType
	// Logger *Logger - if set, logs internal steps (event calls, errors).
	Logger *Logger
	// If true - enable colored messages in enine (like errors, etc.).

	Context context.Context
}

func (p *UniversalEngineParams) GetContext() context.Context {
	if p.Context == nil {
		return context.Background()
	}
	return p.Context
}

// NewUniversalEngineParams constructs an initialized UniversalEngineParams.
// It automatically injects two event handlers into the Events system:
//   - CallEventsStartEvent - logs the start of any event call (debug level)
//   - CallEventsEndEvent   - logs the end, including any error
//
// Parameters: generator, events, scope, logger. All must be non‑nil.
// Returns a filled struct or an error if event registration fails.
//
// Err errors.CorePackageSystemError
func NewUniversalEngineParams(
	generator *Generator,
	events *Events,
	scope ScopeType,
	logger *Logger,
	context context.Context,
) (*UniversalEngineParams, ErrorInterface) {
	if generator == nil || events == nil || logger == nil {
		return nil, Err(errors.CorePackageSystemError, "Invalid input: nil refs")
	}
	logS := func(e *Events, _ *EventInput) ErrorInterface {
		name, err := ScopeGet[string](e.scope, public.EventsScopeCallName)
		if err != nil {
			return Wrap(errors.CorePackageSystemError, err, "LogEvent Start failed")
		}
		logger.PrintLog("event", "Start call '"+name+"' event")
		return nil
	}
	logE := func(e *Events, _ *EventInput) ErrorInterface {
		name, err := ScopeGet[string](e.scope, public.EventsScopeCallName)
		if err != nil {
			return Wrap(errors.CorePackageSystemError, err, "LogEvent End failed")
		}
		text := "End call '" + name + "' event"
		logger.PrintLog("event", text)
		return nil
	}
	events.NewEvent(public.CallEventsStartEvent, logS)
	events.NewEvent(public.CallEventsEndEvent, logE)
	return &UniversalEngineParams{
		Generator: generator,
		Event:     events,
		Scope:     scope,
		Logger:    logger,
		Context:   context,
	}, nil
}
```

---

# engine/events/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# events

'''go
import "github.com/pt-main/lc/engine/events"
'''

## Index

- [type ByteCLDType](<#ByteCLDType>)
- [type ByteCallAttr](<#ByteCallAttr>)
- [type CallLoopData](<#CallLoopData>)
- [type DefaultEvents](<#DefaultEvents>)
  - [func \(de \*DefaultEvents\) ByteCallEvent\(events \*core.Events, i \*core.EventInput\) \(err core.ErrorInterface\)](<#DefaultEvents.ByteCallEvent>)
  - [func \(de \*DefaultEvents\) ByteCallEventIteration\(idx \*int, parsed \*ByteCallAttr, e engine.ByteEngineInterface\) core.ErrorInterface](<#DefaultEvents.ByteCallEventIteration>)
  - [func \(de \*DefaultEvents\) ByteCallHotLoopEvent\(events \*core.Events, i \*core.EventInput\) \(err core.ErrorInterface\)](<#DefaultEvents.ByteCallHotLoopEvent>)
  - [func \(de \*DefaultEvents\) ByteCallPreprocess\(parsed \[\]byteParsing.ParsedBytes, endianess public.EndianType, u bytecode.Utils, abis map\[int\]bool, cmds map\[int\]core.CommandMeta\[engine.ByteEngineInterface, byteParsing.ParsedBytes\]\) \(\[\]ByteCallAttr, core.ErrorInterface\)](<#DefaultEvents.ByteCallPreprocess>)
  - [func \(de \*DefaultEvents\) ByteParsingEvent\(events \*core.Events, i \*core.EventInput\) core.ErrorInterface](<#DefaultEvents.ByteParsingEvent>)
  - [func \(de \*DefaultEvents\) StringCallEvent\(events \*core.Events, i \*core.EventInput\) \(err core.ErrorInterface\)](<#DefaultEvents.StringCallEvent>)
  - [func \(de \*DefaultEvents\) StringCallEventIteration\(parsed \[\]stringParsing.ParsedNode, idx \*int, events \*core.Events, ctx context.Context, e engine.StringEngineInterface\) \(err core.ErrorInterface\)](<#DefaultEvents.StringCallEventIteration>)
  - [func \(de \*DefaultEvents\) StringCallLoopEvent\(events \*core.Events, i \*core.EventInput\) \(err core.ErrorInterface\)](<#DefaultEvents.StringCallLoopEvent>)
  - [func \(de \*DefaultEvents\) StringParsingEvent\(events \*core.Events, i \*core.EventInput\) core.ErrorInterface](<#DefaultEvents.StringParsingEvent>)
- [type StringCLDType](<#StringCLDType>)


<a name="ByteCLDType"></a>
## type [ByteCLDType](<https://github.com/pt-main/Lc/blob/main/engine/events/byteEngine.go#L15>)



'''go
type ByteCLDType CallLoopData[ByteCallAttr, engine.ByteEngineInterface]
'''

<a name="ByteCallAttr"></a>
## type [ByteCallAttr](<https://github.com/pt-main/Lc/blob/main/engine/events/byteEngine.go#L38-L42>)



'''go
type ByteCallAttr struct {
    RawNode *byteParsing.ParsedBytes
    Abis    bool
    Handler core.CommandType[engine.ByteEngineInterface, byteParsing.ParsedBytes]
}
'''

<a name="CallLoopData"></a>
## type [CallLoopData](<https://github.com/pt-main/Lc/blob/main/engine/events/defaultEvents.go#L9-L14>)



'''go
type CallLoopData[P, E any] struct {
    Ctx    context.Context
    Parsed []P
    Engine E
    Idx    *int
}
'''

<a name="DefaultEvents"></a>
## type [DefaultEvents](<https://github.com/pt-main/Lc/blob/main/engine/events/defaultEvents.go#L7>)



'''go
type DefaultEvents struct{}
'''

<a name="DefaultEvents.ByteCallEvent"></a>
### func \(\*DefaultEvents\) [ByteCallEvent](<https://github.com/pt-main/Lc/blob/main/engine/events/byteEngine.go#L77>)

'''go
func (de *DefaultEvents) ByteCallEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface)
'''

Err errors.DefaultEventsSystemError. Err errors.DefaultEventsPanicError. Err errors.DefaultEventsCallErrorContexted. With meta: EMK\(0, "int"\) \- cmd, EMK\(1, "int"\) \- bcIdx, EMK\(2, "string"\) \- pb.

<a name="DefaultEvents.ByteCallEventIteration"></a>
### func \(\*DefaultEvents\) [ByteCallEventIteration](<https://github.com/pt-main/Lc/blob/main/engine/events/byteEngine.go#L177-L180>)

'''go
func (de *DefaultEvents) ByteCallEventIteration(idx *int, parsed *ByteCallAttr, e engine.ByteEngineInterface) core.ErrorInterface
'''

Err errors.DefaultEventsCallErrorHandler.

<a name="DefaultEvents.ByteCallHotLoopEvent"></a>
### func \(\*DefaultEvents\) [ByteCallHotLoopEvent](<https://github.com/pt-main/Lc/blob/main/engine/events/byteEngine.go#L144>)

'''go
func (de *DefaultEvents) ByteCallHotLoopEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface)
'''

Err errors.DefaultEventsCallErrorContex. Err errors.DefaultEventsCallErrorContexted.

<a name="DefaultEvents.ByteCallPreprocess"></a>
### func \(\*DefaultEvents\) [ByteCallPreprocess](<https://github.com/pt-main/Lc/blob/main/engine/events/byteEngine.go#L46-L50>)

'''go
func (de *DefaultEvents) ByteCallPreprocess(parsed []byteParsing.ParsedBytes, endianess public.EndianType, u bytecode.Utils, abis map[int]bool, cmds map[int]core.CommandMeta[engine.ByteEngineInterface, byteParsing.ParsedBytes]) ([]ByteCallAttr, core.ErrorInterface)
'''

Err errors.DefaultEventsCallErrorCmdNotFound. With meta: EMK\(0, "int"\) \- opcode.

<a name="DefaultEvents.ByteParsingEvent"></a>
### func \(\*DefaultEvents\) [ByteParsingEvent](<https://github.com/pt-main/Lc/blob/main/engine/events/byteEngine.go#L20>)

'''go
func (de *DefaultEvents) ByteParsingEvent(events *core.Events, i *core.EventInput) core.ErrorInterface
'''

Err errors.DefaultEventsSystemError. With meta: EMK\(0, "string"\) \- expected type. Cause from core.ScopeGet, e.Parser.Parse.

<a name="DefaultEvents.StringCallEvent"></a>
### func \(\*DefaultEvents\) [StringCallEvent](<https://github.com/pt-main/Lc/blob/main/engine/events/stringEngine.go#L37>)

'''go
func (de *DefaultEvents) StringCallEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface)
'''

Err errors.DefaultEventsSystemError. Err errors.DefaultEventsPanicError. Err errors.DefaultEventsCallErrorContexted. With meta: EMK\(0, "string"\) \- raw line.

<a name="DefaultEvents.StringCallEventIteration"></a>
### func \(\*DefaultEvents\) [StringCallEventIteration](<https://github.com/pt-main/Lc/blob/main/engine/events/stringEngine.go#L103-L109>)

'''go
func (de *DefaultEvents) StringCallEventIteration(parsed []stringParsing.ParsedNode, idx *int, events *core.Events, ctx context.Context, e engine.StringEngineInterface) (err core.ErrorInterface)
'''

Err errors.DefaultEventsCallErrorContex. Err errors.DefaultEventsCallErrorUnknown. Err errors.DefaultEventsCallErrorHandler.

<a name="DefaultEvents.StringCallLoopEvent"></a>
### func \(\*DefaultEvents\) [StringCallLoopEvent](<https://github.com/pt-main/Lc/blob/main/engine/events/stringEngine.go#L80>)

'''go
func (de *DefaultEvents) StringCallLoopEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface)
'''

Err errors.DefaultEventsCallErrorContex. Err errors.DefaultEventsCallErrorContexted.

<a name="DefaultEvents.StringParsingEvent"></a>
### func \(\*DefaultEvents\) [StringParsingEvent](<https://github.com/pt-main/Lc/blob/main/engine/events/stringEngine.go#L17>)

'''go
func (de *DefaultEvents) StringParsingEvent(events *core.Events, i *core.EventInput) core.ErrorInterface
'''

Err errors.DefaultEventsSystemError. Cause from core.ScopeGet, e.Parser.Parse.

<a name="StringCLDType"></a>
## type [StringCLDType](<https://github.com/pt-main/Lc/blob/main/engine/events/stringEngine.go#L14>)



'''go
type StringCLDType CallLoopData[stringParsing.ParsedNode, engine.StringEngineInterface]
'''

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# engine/events/byteEngine.go

```go
package events

import (
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
	"github.com/pt-main/lc/tooling/bytecode"
)

type ByteCLDType CallLoopData[ByteCallAttr, engine.ByteEngineInterface]

// Err errors.DefaultEventsSystemError.
// With meta: EMK(0, "string") - expected type.
// Cause from core.ScopeGet, e.Parser.Parse.
func (de *DefaultEvents) ByteParsingEvent(events *core.Events, i *core.EventInput) core.ErrorInterface {
	e, ok := i.Input.(*engine.ByteEngine)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid input: expected *engine.ByteEngine").
			WithMeta(core.EMK(0, "string"), "*engine.ByteEngine")
	}
	input, err := core.ScopeGet[[]byte](e.UEP.Scope, public.ByteEngineScopeInput)
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Cannot get input from scope")
	}
	nodes, err := e.Parser.Parse(input, &parsing.ParseOption{UEP: e.UEP})
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Parser failed")
	}
	e.UEP.Scope[public.ByteEngineScopeParsed] = nodes
	return nil
}

type ByteCallAttr struct {
	RawNode *byteParsing.ParsedBytes
	Abis    bool
	Handler core.CommandType[engine.ByteEngineInterface, byteParsing.ParsedBytes]
}

// Err errors.DefaultEventsCallErrorCmdNotFound.
// With meta: EMK(0, "int") - opcode.
func (de *DefaultEvents) ByteCallPreprocess(
	parsed []byteParsing.ParsedBytes, endianess public.EndianType,
	u bytecode.Utils, abis map[int]bool,
	cmds map[int]core.CommandMeta[engine.ByteEngineInterface, byteParsing.ParsedBytes],
) ([]ByteCallAttr, core.ErrorInterface) {
	res := make([]ByteCallAttr, 0, len(parsed))
	for _, node := range parsed {
		cmdSwitch := u.BytesToInt(node.Switch, endianess)
		handler, ok := cmds[cmdSwitch]
		if !ok {
			return nil, core.Err(errors.DefaultEventsCallErrorCmdNotFound, "Opcode %d not registered", cmdSwitch).
				WithMeta(core.EMK(0, "int"), cmdSwitch)
		}
		autoshift, ok := abis[cmdSwitch]
		if !ok {
			return nil, core.Err(errors.DefaultEventsCallErrorCmdNotFound, "Autoshift config missing for opcode %d", cmdSwitch).
				WithMeta(core.EMK(0, "int"), cmdSwitch)
		}
		res = append(res, ByteCallAttr{
			RawNode: &node,
			Handler: handler.Handler,
			Abis:    autoshift,
		})
	}
	return res, nil
}

// Err errors.DefaultEventsSystemError.
// Err errors.DefaultEventsPanicError.
// Err errors.DefaultEventsCallErrorContexted.
// With meta: EMK(0, "int") - cmd, EMK(1, "int") - bcIdx, EMK(2, "string") - pb.
func (de *DefaultEvents) ByteCallEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	var idx *int
	var parsed2 []ByteCallAttr
	var lastCmd int
	defer func() {
		if r := recover(); r != nil {
			err = core.Err(errors.DefaultEventsPanicError, "Panic recovered: %v", r)
		}
		if err != nil {
			if err.Error() == core.ErrExit.Error() {
				if idx != nil {
					*idx = -1
				}
				return
			}
			idxV := 0
			if idx != nil {
				idxV = *idx
			}
			err = core.Wrap(errors.DefaultEventsCallErrorContexted, err,
				"Error at cmd=%v, bcIdx=%v", lastCmd, idxV).
				WithMeta(core.EMK(0, "int"), lastCmd).
				WithMeta(core.EMK(1, "int"), idxV)
		}
	}()
	e, ok := i.Input.(*engine.ByteEngine)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid input: expected *engine.ByteEngine").
			WithMeta(core.EMK(0, "string"), "*engine.ByteEngine")
	}
	_parsed, ok := e.GetUep().Scope[public.ByteEngineScopeParsed]
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Parsed data not found in scope")
	}
	parsed, ok := _parsed.([]byteParsing.ParsedBytes)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Parsed data has wrong type")
	}
	u := bytecode.Utils{}
	endianess, ok := e.GetUep().Scope[public.ByteEngineScopeEndianess].(public.EndianType)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid endianess in scope")
	}
	idx, err = core.ScopeGet[*int](e.GetUep().Scope, public.ByteEngineScopeBytecodeIdx)
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Cannot get bytecode index")
	}
	ctx := e.GetUep().GetContext()
	cmds := e.Commands
	abis := e.AutoBytecodeIndexShift

	parsed2, err = de.ByteCallPreprocess(parsed, endianess, u, abis, cmds)
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Preprocessing failed")
	}

	if err = events.CallEvents(&core.EventInput{Input: ByteCLDType{
		Ctx: ctx, Parsed: parsed2, Engine: e, Idx: idx, Other: &parsed,
	}}, public.ByteCallHotloopEvent, false); err != nil {
		return core.Wrap(errors.DefaultEventsCallErrorContexted, err, "Hot-loop event failed")
	}
	return nil
}

// Err errors.DefaultEventsCallErrorContex.
// Err errors.DefaultEventsCallErrorContexted.
func (de *DefaultEvents) ByteCallHotLoopEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	hld, ok := i.Input.(ByteCLDType)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid event input: expected ByteCLDType").
			WithMeta(core.EMK(0, "string"), "ByteCLDType")
	}
	idx := hld.Idx
	ctx := hld.Ctx
	parsed := hld.Parsed
	p2len := len(parsed)
	e := hld.Engine
	iter := 0
	var checkInterval int
	checkInterval, err = core.ScopeGet[int](e.GetUep().Scope, public.ByteEngineScopeBytecodeIdx)
	if err != nil {
		checkInterval = 255 // 2^8-1
	}
	for {
		iter++
		if iter&int(checkInterval) == 0 {
			if ctx.Err() != nil {
				return core.Wrap(errors.DefaultEventsCallErrorContex, ctx.Err(), "Context cancelled (at %v iter)", iter)
			}
		}
		idxN := *idx
		if uint(idxN) >= uint(p2len) {
			break
		}
		node := &parsed[idxN]
		//go:inline
		if err = de.ByteCallEventIteration(idx, node, e); err != nil {
			if node.Abis == true {
				(*idx)--
			}
			return core.Wrap(errors.DefaultEventsCallErrorHandler, err, "Handler failed")
		}
	}
	return nil
}

// Err errors.DefaultEventsCallErrorHandler.
func (de *DefaultEvents) ByteCallEventIteration(
	idx *int,
	parsed *ByteCallAttr, e engine.ByteEngineInterface,
) core.ErrorInterface {
	if parsed.Abis {
		*idx++
	}
	//go:inline
	return parsed.Handler(e, parsed.RawNode)
}
```

---

# engine/events/defaultEvents.go

```go
package events

import (
	"context"
)

type DefaultEvents struct{}

type CallLoopData[P, E any] struct {
	Ctx    context.Context
	Parsed []P
	Other  any
	Engine E
	Idx    *int
}
```

---

# engine/events/stringEngine.go

```go
package events

import (
	"context"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

type StringCLDType CallLoopData[stringParsing.ParsedNode, engine.StringEngineInterface]

// Err errors.DefaultEventsSystemError. Cause from core.ScopeGet, e.Parser.Parse.
func (de *DefaultEvents) StringParsingEvent(events *core.Events, i *core.EventInput) core.ErrorInterface {
	e, ok := i.Input.(*engine.StringEngine)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid input: expected *engine.StringEngine").
			WithMeta(core.EMK(0, "string"), "*engine.StringEngine")
	}
	input, ok := e.UEP.Scope[public.StringEngineScopeInput].(string)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Input not found or invalid type in scope")
	}
	nodes, err := e.Parser.Parse(input, &parsing.ParseOption{UEP: e.UEP})
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Parser failed")
	}
	e.UEP.Scope[public.StringEngineScopeParsed] = nodes
	return nil
}

// Err errors.DefaultEventsSystemError. Err errors.DefaultEventsPanicError.
// Err errors.DefaultEventsCallErrorContexted. With meta: EMK(0, "string") - raw line.
func (de *DefaultEvents) StringCallEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	events.Scope()[public.EventsScopeDERawLine] = "[NIL]"
	defer func() {
		if r := recover(); r != nil {
			err = core.Err(errors.DefaultEventsPanicError, "Panic recovered: %v", r)
		}
		if err != nil {
			if err.Error() == core.ErrExit.Error() {
				return
			}
			raw, _ := core.ScopeGet[string](events.Scope(), public.EventsScopeDERawLine)
			if raw == "" {
				raw = "[NIL]"
			}
			err = core.Wrap(errors.DefaultEventsCallErrorContexted, err,
				"Error at line: %q", raw).
				WithMeta(core.EMK(0, "string"), raw)
		}
	}()
	e, ok := i.Input.(*engine.StringEngine)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid input: expected *engine.StringEngine").
			WithMeta(core.EMK(0, "string"), "*engine.StringEngine")
	}
	parsed, err := core.ScopeGet[[]stringParsing.ParsedNode](e.UEP.Scope, public.StringEngineScopeParsed)
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Cannot get parsed nodes")
	}
	idx := 0
	cld := StringCLDType{
		Ctx:    e.UEP.GetContext(),
		Parsed: parsed,
		Engine: e,
		Idx:    &idx,
	}
	e.UEP.Scope[public.StringEngineScopeInstrIdx] = cld.Idx
	if err = events.CallEvents(&core.EventInput{Input: cld}, public.StringCallCalloopEvent, false); err != nil {
		return core.Wrap(errors.DefaultEventsCallErrorContexted, err, "Call loop failed")
	}
	return nil
}

// Err errors.DefaultEventsCallErrorContex. Err errors.DefaultEventsCallErrorContexted.
func (de *DefaultEvents) StringCallLoopEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	cld := i.Input.(StringCLDType)
	idx := cld.Idx
	parsed := cld.Parsed
	pLen := len(parsed)
	ctx := cld.Ctx
	e := cld.Engine
	for *idx < pLen {
		if ctx.Err() != nil {
			return core.Wrap(errors.DefaultEventsCallErrorContex, ctx.Err(), "Context cancelled")
		}
		if err = de.StringCallEventIteration(parsed, idx, events, ctx, e); err != nil {
			return core.Wrap(errors.DefaultEventsCallErrorContexted, err, "Iteration failed")
		}
		if *idx < 0 {
			return nil
		}
	}
	return nil
}

// Err errors.DefaultEventsCallErrorContex. Err errors.DefaultEventsCallErrorUnknown.
// Err errors.DefaultEventsCallErrorHandler.
func (de *DefaultEvents) StringCallEventIteration(
	parsed []stringParsing.ParsedNode,
	idx *int,
	events *core.Events,
	ctx context.Context,
	e engine.StringEngineInterface,
) (err core.ErrorInterface) {
	node := parsed[*idx]
	canBeUnknown, err := core.ScopeGet[bool](e.GetUep().Scope, public.StringEngineScopeCanBeUnknown)
	if err != nil {
		canBeUnknown = true
	}
	events.Scope()[public.EventsScopeDERawLine] = node.Raw
	if ctx.Err() != nil {
		return core.Wrap(errors.DefaultEventsCallErrorContex, ctx.Err(), "Context cancelled")
	}
	handler, ok := e.GetCommands()[node.Switch]
	if ok {
		err = handler.Handler(e, &node)
	} else if !canBeUnknown {
		return core.Err(errors.DefaultEventsCallErrorUnknown, "Unknown command: %s", node.Switch).
			WithMeta(core.EMK(0, "string"), node.Switch)
	}
	if err != nil {
		if err.Error() == core.ErrExit.Error() {
			*idx = -1
			return nil
		}
		return core.Wrap(errors.DefaultEventsCallErrorHandler, err, "Handler failed for %s", node.Switch).
			WithMeta(core.EMK(0, "string"), node.Switch)
	}
	(*idx)++
	return nil
}
```

---

# engine/interface.go

```go
package engine

import (
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

type EngineInterface[CmdT int | string | byte | float32 | float64,
	ParserInput any, ParserOutput any] interface {
	Process(ParserInput) core.ErrorInterface
	NewCommand(CmdT, core.CommandType[EngineInterface[
		CmdT, ParserInput, ParserOutput], ParserOutput], *core.SimpleInput) error
	GetUep() *core.UniversalEngineParams
	GetParser() parsing.ParserInterface[ParserInput, ParserOutput]
	GetCommands() map[CmdT]core.CommandMeta[EngineInterface[
		CmdT, ParserInput, ParserOutput], ParserOutput]
}

type StringEngineInterface = EngineInterface[string, string, stringParsing.ParsedNode]
type ByteEngineInterface = EngineInterface[int, []byte, byteParsing.ParsedBytes]
```

---

# engine/stringEngine.go

```go
package engine

import (
	"sync"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

type stringParser = parsing.ParserInterface[string, stringParsing.ParsedNode]
type stringCommandMeta = core.CommandMeta[StringEngineInterface, stringParsing.ParsedNode]

// StringEngine is the core for text‑based languages. It holds command
// definitions, a parser, and universal engine parameters (UEP) that include
// generator, events, scope, and logger. The Process method drives compilation.
type StringEngine struct {
	Commands map[string]stringCommandMeta
	Parser   stringParser
	UEP      *core.UniversalEngineParams
	mu       sync.RWMutex
}

// Process executes the compilation pipeline for a string input.
// It stores the input in scope["input_string"], then calls the
// StringParseEvent (to parse into []ParsedNode) and StringCallEvent
// (to dispatch commands). Any error stops execution.
//
// Err errors.StringEngineProcessError1 | errors.StringEngineProcessError2.
// (cause from 'CallEvents')
func (e *StringEngine) Process(input string) core.ErrorInterface {
	e.UEP.Scope[public.StringEngineScopeInput] = input
	err1 := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.StringParseEvent, false)
	if err1 != nil {
		return core.Wrap(errors.StringEngineProcessError1, err1, core.GetRealError(err1))
	}
	err2 := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.StringCallEvent, false)
	if err2 != nil {
		return core.Wrap(errors.StringEngineProcessError2, err2, core.GetRealError(err2))
	}
	return nil
}

func (e *StringEngine) NewCommandFull(cmd_switch string,
	handler core.CommandType[StringEngineInterface, stringParsing.ParsedNode], doc string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Commands[cmd_switch] = core.CommandMeta[StringEngineInterface, stringParsing.ParsedNode]{
		Handler: handler,
		Doc:     doc,
	}
}

func (e *StringEngine) GetParser() stringParser {
	return e.Parser
}

// For interface. o.Input string = doc
func (e *StringEngine) NewCommand(cmd_switch string,
	handler core.CommandType[StringEngineInterface, stringParsing.ParsedNode],
	o *core.SimpleInput) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	doc, ok := o.Input.(string)
	if !ok {
		return core.Err(errors.CorePackageSystemError, "Invalid input: 'o.Input' must be string")
	}
	e.Commands[cmd_switch] = core.CommandMeta[StringEngineInterface, stringParsing.ParsedNode]{
		Handler: handler,
		Doc:     doc,
	}
	return nil
}

// For interface
func (e *StringEngine) GetCommands() map[string]stringCommandMeta {
	return e.Commands
}

// For interface
func (e *StringEngine) GetUep() *core.UniversalEngineParams {
	return e.UEP
}
```

---

# engine.go

```go
package lc

import (
	"context"
	goerr "errors"
	"fmt"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
	lcplugin "github.com/pt-main/lc/tooling/plugin"
)

type EngineUniversal struct {
	Plugins        *lcplugin.PluginManager
	Type           public.EngineType
	StringEngine   engine.EngineInterface[string, string, stringParsing.ParsedNode]
	ByteEngine     engine.EngineInterface[int, []byte, byteParsing.ParsedBytes]
	opcode_counter int
	Context        context.Context
	CtxCancelCause context.CancelCauseFunc
	ended          bool
}

func (e *EngineUniversal) ProcessStringWithCtx(input string, ctx context.Context) core.ErrorInterface {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.StringEngineType {
		return core.Err(errors.CorePackageLcError, "Can't process string in byte engine")
	}
	uep, _ := e.GetUEP()
	uep.Context = ctx
	return e.StringEngine.Process(input)
}

func (e *EngineUniversal) ProcessBytesWithCtx(input []byte, ctx context.Context) core.ErrorInterface {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.ByteEngineType {
		return core.Err(errors.CorePackageLcError, "Can't process bytes in string engine")
	}
	uep, _ := e.GetUEP()
	uep.Context = ctx
	return e.ByteEngine.Process(input)
}

// ProcessString feeds a string input into the engine.
// It works only for engines of type StringEngineType; otherwise returns an core.ErrorInterface.
// Internally triggers the parse and call events, executing registered handlers.
func (e *EngineUniversal) ProcessString(input string) core.ErrorInterface {
	return e.ProcessStringWithCtx(input, context.Background())
}

// ProcessBytes feeds a byte slice into the engine (ByteEngineType only).
// The input is passed via scope under key "input_[]byte", then parsed and processed.
func (e *EngineUniversal) ProcessBytes(input []byte) core.ErrorInterface {
	return e.ProcessBytesWithCtx(input, context.Background())
}

func (e *EngineUniversal) GetUEP() (*core.UniversalEngineParams, error) {
	if err := e.CheckEnded(); err != nil {
		return nil, err
	}
	if e.Type == public.StringEngineType {
		return e.StringEngine.GetUep(), nil
	}
	return e.ByteEngine.GetUep(), nil
}

// NewCommandByte registers a bytecode command identified by an opcode.
// If opcode == -1, the engine automatically assigns the next available opcode.
// handler receives (*ByteEngine, ParsedBytes).
func (e *EngineUniversal) NewCommandByte(
	opcode int, handler core.CommandType[engine.ByteEngineInterface, byteParsing.ParsedBytes], name string,
	autoByecodeIdxShift bool,
) error {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.ByteEngineType {
		return goerr.New("Can't add byte command to string engine")
	}
	finalOpcode := opcode
	if opcode == -1 {
		finalOpcode = e.opcode_counter
		e.opcode_counter++
	} else {
		e.opcode_counter = max(opcode, e.opcode_counter)
	}

	e.ByteEngine.NewCommand(finalOpcode, handler, &core.SimpleInput{
		Input:  name,
		Option: &core.Option{Flags: []string{engine.AutoshiftNewCommandFlag}},
	})
	return nil
}

// NewCommandString registers a text-based command in a StringEngine.
// cmdSwitch is the command name (e.g., "print"). handler must have signature
// func([]interface{}) error where arguments are (*StringEngine, ParsedNode).
// doc is an optional documentation string.
func (e *EngineUniversal) NewCommandString(
	cmdSwitch string, handler core.CommandType[engine.StringEngineInterface, stringParsing.ParsedNode], doc string,
) error {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.StringEngineType {
		return goerr.New("Can't add string command to byte engine")
	}
	e.StringEngine.NewCommand(cmdSwitch, handler, &core.SimpleInput{
		Input: doc,
	})
	return nil
}

// End - function for stop engines lifecycle.
func (e *EngineUniversal) End() (err error) {
	err = e.CheckEnded()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[%v] Panic recovered: %v. ", err, r)
		}
	}()
	if e.Plugins != nil {
		err = fmt.Errorf("[%v] Plugin error: %v. ", err, e.Plugins.End())
	}
	if e.Context.Err() != nil {
		e.CtxCancelCause(fmt.Errorf("EngineUniversal: lifecycle end."))
	}
	e.ByteEngine = nil
	e.StringEngine = nil
	return
}

func (e *EngineUniversal) CheckEnded() (err core.ErrorInterface) {
	if e.ended {
		err = core.Err(errors.CorePackageLcLifecycleError, "EngineUniversal: lifecycle ended.")
	}
	return err
}
```

---

# example/README.md

```md
# Examples

- 'langs/' full complex examples
    - 'langs/calculator' - test expression calculator works in console
    - 'langs/configLang' - test ini-like config -> json tranlator (with profiler plugin)
- 'readme/' - examples from readme
- 'tests/'
    - 'tests/speedtest/' - test of engines speed
    - 'tests/parser3Test/' - parser3 test
- 'packages/' - examples for core concepts in framework
```

---

# example/langs/calculator/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# calculator

'''go
import "github.com/pt-main/lc/example/langs/calculator"
'''

## Index



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# example/langs/calculator/main.go

```go
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/parsing/stringParsing/parser3"
	"github.com/pt-main/lc/tooling/astools"
)

func createLexer() *stringParsing.Lexer {
	rules := []stringParsing.LexerRule{
		{Type: "NUMBER", Pattern: regexp2.MustCompile('\d+(\.\d+)?', 0)},
		{Type: "PLUS", Pattern: regexp2.MustCompile('\+', 0)},
		{Type: "MINUS", Pattern: regexp2.MustCompile('-', 0)},
		{Type: "POW", Pattern: regexp2.MustCompile('\*\*', 0)},
		{Type: "MUL", Pattern: regexp2.MustCompile('\*', 0)},
		{Type: "DIV", Pattern: regexp2.MustCompile('/', 0)},
		{Type: "LPAREN", Pattern: regexp2.MustCompile('\(', 0)},
		{Type: "RPAREN", Pattern: regexp2.MustCompile('\)', 0)},
		{Type: "WHITESPACE", Pattern: regexp2.MustCompile('\s+', 0)},
	}
	config := &stringParsing.LexerConfig{
		UseBracketBalance: false,
	}
	return stringParsing.NewLexer(rules, config)
}

func createGrammar() parser3.Grammar {
	return parser3.Grammar{
		"expr": {
			Name: "expr",
			Expr: parser3.NodeExpr{
				NodeType: "expr",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.NamedExpr{RuleName: "term"},
						parser3.RepeatExpr{
							Expr: parser3.SequenceExpr{
								Exprs: []parser3.Expr{
									parser3.ChoiceExpr{
										Alternatives: []parser3.Expr{
											parser3.TokenExpr{TokenType: "PLUS"},
											parser3.TokenExpr{TokenType: "MINUS"},
										},
									},
									parser3.NamedExpr{RuleName: "term"},
								},
							},
							Min: 0,
						},
					},
				},
			},
		},
		"term": {
			Name: "term",
			Expr: parser3.NodeExpr{
				NodeType: "term",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.NamedExpr{RuleName: "factor"},
						parser3.RepeatExpr{
							Expr: parser3.SequenceExpr{
								Exprs: []parser3.Expr{
									parser3.ChoiceExpr{
										Alternatives: []parser3.Expr{
											parser3.TokenExpr{TokenType: "MUL"},
											parser3.TokenExpr{TokenType: "DIV"},
											parser3.TokenExpr{TokenType: "POW"},
										},
									},
									parser3.NamedExpr{RuleName: "factor"},
								},
							},
							Min: 0,
						},
					},
				},
			},
		},
		"factor": {
			Name: "factor",
			Expr: parser3.NodeExpr{
				NodeType: "factor",
				Expr: parser3.ChoiceExpr{
					Alternatives: []parser3.Expr{
						parser3.TokenExpr{TokenType: "NUMBER"},
						parser3.SequenceExpr{
							Exprs: []parser3.Expr{
								parser3.TokenExpr{TokenType: "LPAREN"},
								parser3.NamedExpr{RuleName: "expr"},
								parser3.TokenExpr{TokenType: "RPAREN"},
							},
						},
					},
				},
			},
		},
	}
}

func evalExpr(node *stringParsing.ParsedNode) (float64, error) {
	switch node.Switch {
	case "expr":
		children := astools.GetChildren(node)
		if len(children) == 0 {
			return 0, errors.New("empty expr")
		}
		val, err := evalExpr(&children[0])
		if err != nil {
			return 0, err
		}
		for i := 1; i < len(children); i += 2 {
			if i+1 >= len(children) {
				break
			}
			opNode := &children[i]
			termNode := &children[i+1]
			termVal, err := evalExpr(termNode)
			if err != nil {
				return 0, err
			}
			switch opNode.Switch {
			case "PLUS":
				val += termVal
			case "MINUS":
				val -= termVal
			}
		}
		return val, nil
	case "term":
		children := astools.GetChildren(node)
		if len(children) == 0 {
			return 0, errors.New("empty term")
		}
		val, err := evalExpr(&children[0])
		if err != nil {
			return 0, err
		}
		for i := 1; i < len(children); i += 2 {
			if i+1 >= len(children) {
				break
			}
			opNode := &children[i]
			factorNode := &children[i+1]
			factorVal, err := evalExpr(factorNode)
			if err != nil {
				return 0, err
			}
			switch opNode.Switch {
			case "MUL":
				val *= factorVal
			case "DIV":
				if factorVal == 0 {
					return 0, errors.New("division by zero")
				}
				val /= factorVal
			case "POW":
				val = math.Pow(val, factorVal)
			}
		}
		return val, nil
	case "factor":
		children := astools.GetChildren(node)
		if len(children) == 0 {
			return 0, errors.New("empty factor")
		}
		child := &children[0]
		if child.Switch == "NUMBER" {
			numStr := child.Raw
			return strconv.ParseFloat(numStr, 64)
		}
		if child.Switch == "LPAREN" && len(children) >= 3 {
			return evalExpr(&children[1])
		}
		return 0, errors.New("unknown factor")
	default:
		return 0, errors.New("unknown node type: " + node.Switch)
	}
}

func main() {
	fmt.Println("Lc version -", lc.Version)

	if len(os.Args) < 2 {
		fmt.Println("Usage: calc <expression>")
		fmt.Println("Example: calc '(2 ** 3) + 4'")
		os.Exit(1)
	}
	expr := os.Args[1]
	lexer := createLexer()
	grammar := createGrammar()
	parser := parser3.NewParser(lexer, grammar, "expr", []string{"WHITESPACE"})
	parsed, err := parser.Parse(expr)
	if err != nil {
		fmt.Println("Parse error:\n", parser3.FormatErrorPretty(err))
		os.Exit(1)
	}
	if len(parsed) == 0 {
		fmt.Println("No nodes parsed")
		os.Exit(1)
	}
	result, err2 := evalExpr(&parsed[0])
	if err2 != nil {
		fmt.Println("Eval error:", err)
		os.Exit(1)
	}
	fmt.Printf("Result: %v\n", result)
}

/*
macbook@MacBook-Pro lc % go run ./example/calculator '(2+3)**4'
Lc version - 1.5.1
Result: 625
macbook@MacBook-Pro lc % go run ./example/calculator '2*3+4'
Lc version - 1.5.1
Result: 10
macbook@MacBook-Pro lc % go run ./example/calculator '2-3*4'
Lc version - 1.5.1
Result: -10
macbook@MacBook-Pro lc % go run ./example/calculator '(2-3*4)+(5*2-1)*2'
Lc version - 1.5.1
Result: 8
*/
```

---

# example/langs/configLang/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# configLang

'''go
import "github.com/pt-main/lc/example/langs/configLang"
'''

## Index

- [func Process\(config string\) \(string, error\)](<#Process>)


<a name="Process"></a>
## func [Process](<https://github.com/pt-main/Lc/blob/main/example/langs/configLang/main.go#L20>)

'''go
func Process(config string) (string, error)
'''



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# example/langs/configLang/main.go

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/debugging/extensiblePlugin"
	"github.com/pt-main/lc/tooling/debugging/profiler"
)

func Process(config string) (string, error) {
	grammar := []stringParsing.GrammarRule{
		{
			Type:    "section",
			Pattern: regexp.MustCompile('^\[(?P<name>[^\]]+)\]$'),
		},
		{
			Type:    "keyval",
			Pattern: regexp.MustCompile('^(?P<key>[^=]+?)\s*=\s*(?P<value>.*)$'),
		},
		{
			Type:    "comment",
			Pattern: regexp.MustCompile('^\s*#[^\n]*$'),
		},
		{
			Type:    "unknown_token",
			Pattern: regexp.MustCompile('.*'),
		},
	}
	cfgp := stringParsing.Parser1Config{
		UseLineContinuation: true,
		SkipEmptyLines:      true,
		UseBracketBalance:   false,
		TrimBlocksSpace:     true,
	}
	parser := stringParsing.NewParser1(grammar, cfgp)
	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(parser).
		WithDefaultEvents(true).
		WithContext(context.Background()).
		WithColors().
		Build()
	if err != nil {
		return "", err
	}
	err = engine.Plugins.AddPlugin(extensiblePlugin.New(engine))
	if err != nil {
		return "", err
	}
	err = engine.Plugins.AddPlugin(profiler.New())
	if err != nil {
		return "", err
	}
	uep, _ := engine.GetUEP()
	uep.Scope[public.StringEngineScopeCanBeUnknown] = false
	uep.Scope["config"] = make(map[string]map[string]interface{})
	uep.Scope["current_section"] = "global"
	engine.NewCommandString("comment", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		return nil
	}, "")
	engine.NewCommandString("section", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		name := node.Metadata["name"].(string)
		if strings.HasPrefix(name, "!") {
			switch name[1:] {
			case "exit":
				return core.ErrExit
			default:
				return core.Err("PROCESS", "Invalid command")
			}
		}
		se.GetUep().Scope["current_section"] = name
		cfg := se.GetUep().Scope["config"].(map[string]map[string]interface{})
		if _, ok := cfg[name]; !ok {
			cfg[name] = make(map[string]interface{})
		}
		return nil
	}, "Set current section")
	engine.NewCommandString("keyval", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		key := strings.TrimSpace(node.Metadata["key"].(string))
		rawVal := strings.TrimSpace(node.Metadata["value"].(string))
		val := parseValue(rawVal)
		sectionName, _ := se.GetUep().Scope["current_section"].(string)
		cfg := se.GetUep().Scope["config"].(map[string]map[string]interface{})
		if _, ok := cfg[sectionName]; !ok {
			cfg[sectionName] = make(map[string]interface{})
		}
		cfg[sectionName][key] = val
		return nil
	}, "Store key-value pair")
	if err := engine.ProcessString(config); err != nil {
		return "", err
	}
	cfg := uep.Scope["config"].(map[string]map[string]interface{})
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	report, err := engine.Plugins.CallPluginMethod("profiler", "report")
	fmt.Println(report, err)
	return string(jsonData), err
}

func parseValue(raw string) interface{} {
	raw = strings.TrimSpace(raw)
	if raw == "true" {
		return true
	}
	if raw == "false" {
		return false
	}
	if strings.HasPrefix(raw, "[") && strings.HasSuffix(raw, "]") {
		inner := strings.TrimSpace(raw[1 : len(raw)-1])
		if inner == "" {
			return []interface{}{}
		}
		parts := strings.Split(inner, ",")
		result := make([]interface{}, 0, len(parts))
		for _, p := range parts {
			result = append(result, parseValue(strings.TrimSpace(p)))
		}
		return result
	}
	if i, err := strconv.Atoi(raw); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(raw, 64); err == nil {
		return f
	}
	if len(raw) >= 2 {
		if (raw[0] == '"' && raw[len(raw)-1] == '"') || (raw[0] == '\'' && raw[len(raw)-1] == '\'') {
			return raw[1 : len(raw)-1]
		}
	}
	return raw
}

func main() {
	fmt.Println("Lc version -", lc.Version)

	fmt.Println(Process('
# Global params
app_name = MyApp
version = 1.2.3

[server]
host = localhost
port = 8080
debug = true
timeout = 30.5

[database]
url = postgres://user:pass@localhost/db
pool_size = 10
features = [caching, logging, metrics]
'))
}

/*
Lc version - 1.5.1
Profiler report (0.00 sec total):

  String commands:
  String calls: 13 (total time: 718.071µs)
    keyval: count=9, total=484.819µs, avg=53.868µs, min=45.808µs, max=68.911µs
    section: count=2, total=92.308µs, avg=46.154µs, min=45.346µs, max=46.962µs

{
  "database": {
    "features": [
      "caching",
      "logging",
      "metrics"
    ],
    "pool_size": 10,
    "url": "postgres://user:pass@localhost/db"
  },
  "global": {
    "app_name": "MyApp",
    "version": "1.2.3"
  },
  "server": {
    "debug": true,
    "host": "localhost",
    "port": 8080,
    "timeout": 30.5
  }
} <nil>
*/
```

---

# example/packages/engine/core/events/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# events

'''go
import "github.com/pt-main/lc/example/packages/engine/core/events"
'''

## Index



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# example/packages/engine/core/events/main.go

```go
// Working with engine/core/Events

package main

import (
	"context"
	"fmt"

	"github.com/pt-main/lc/engine/core"
)

func test1() { // Basic usage
	fmt.Println("=== test1 ===")
	e := core.NewEvents(context.Background())

	// add event 'test'
	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Test event is active")
		return nil
	})

	// call event with empty input and canWorkWithoutHandler=false
	fmt.Println(e.CallEvents(&core.EventInput{}, "test", false)) // "Test event is active\n<nil>"

	// call unregistered event with canWorkWithoutHandler=false
	fmt.Println(e.CallEvents(nil, "unknown", false)) // "Event 'unknown' is not found.""

	// call unregistered event with canWorkWithoutHandler=true
	fmt.Println(e.CallEvents(&core.EventInput{}, "unknown", true)) // <nil>
}

func test2() { // Core events
	fmt.Println("=== test2 ===")
	e := core.NewEvents(context.Background())

	// while event is not registred, it's created as core event
	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Test core event is active")
		return nil
	})

	// after creating core event you can append new handlers after core event
	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Call's after test core event")
		return nil
	})

	// and you can creating events before core event
	e.NewEventBefore("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Call's before test core event")
		return nil
	})

	// you can add more events after/before core event
	e.NewEventBefore("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("(1) Call's before test core event")
		return nil
	})

	fmt.Println(e.CallEvents(nil, "test", false)) // <nil>

	// And you can replace event
	e.ReplaceEvent("test")
	fmt.Println(e.CallEvents(nil, "test", false)) // "Event 'test' is not found."
}

func test3() { // Core events deeper
	fmt.Println("=== test3 ===")
	e := core.NewEvents(context.Background())

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Test core event is active")
		return nil
	})

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Call's after test core event")
		return nil
	})

	e.NewEventBefore("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Call's before test core event")
		return nil
	})

	// Creating EventTools for work with core events
	et := core.EventsTools{
		Events: e,
	}

	// Change only core event
	et.ChangeCoreEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Not test core event")
		return nil
	})

	fmt.Println(e.CallEvents(nil, "test", false))
}

func test4() { // Hard example
	fmt.Println("=== test4 ===")
	e := core.NewEvents(context.Background())

	e.Scope()["numStr"] = "1"

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Print("Test core event is active. ")
		fmt.Println(core.ScopeGet[string](e.Scope(), "numStr"))
		e.Scope()["numStr"] = "2"
		return nil
	})

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Print("Call's after test core event. ")
		fmt.Println(core.ScopeGet[string](e.Scope(), "numStr"))
		e.Scope()["numStr"] = "3"
		return nil
	})

	e.NewEventBefore("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Print("Call's before test core event. ")
		fmt.Println(core.ScopeGet[string](e.Scope(), "numStr"))
		e.Scope()["numStr"] = "4"
		return nil
	})

	numStr, _ := core.ScopeGet[string](e.Scope(), "numStr")
	fmt.Println("Before: " + numStr)
	fmt.Println(e.CallEvents(nil, "test", false))
	numStr, _ = core.ScopeGet[string](e.Scope(), "numStr")
	fmt.Println("After: " + numStr)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}

/*
=== test1 ===
Test event is active
<nil>
Event 'unknown' is not found.
<nil>
=== test2 ===
(1) Call's before test core event
Call's before test core event
Test core event is active
Call's after test core event
<nil>
Event 'test' is not found.
=== test3 ===
Call's before test core event
Not test core event
Call's after test core event
<nil>
=== test4 ===
Before: 1
Call's before test core event. 1 <nil>
Test core event is active. 4 <nil>
Call's after test core event. 2 <nil>
<nil>
After: 3
*/
```

---

# example/packages/engine/core/generator/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# generator

'''go
import "github.com/pt-main/lc/example/packages/engine/core/generator"
'''

## Index



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# example/packages/engine/core/generator/main.go

```go
// Working with engine.core.Generator

package main

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/public"
)

func test1() { // Byte generator
	fmt.Println("=== test1 ===")
	g := core.NewGenerator(public.ByteResType, []string{"main"})
	// can't add string, only bytes
	g.AddBytes([]byte{0, 1, 2, 3, 4}, "main")
	g.AddBytes([]byte{5, 6, 7, 8, 9}, "main")
	fmt.Println(g.GetBytesRes()) // [0 1 2 3 4 5 6 7 8 9]
}

func test2() { // String generator
	fmt.Println("=== test2 ===")
	g := core.NewGenerator(public.StringResType, []string{"main"})
	// can't add bytes, only string
	g.AddString("test", "main")
	g.AddString("test2", "main")
	fmt.Println(g.GetStringArrRes())        // [test test2]
	fmt.Println(core.GetStringRes(g, ", ")) // test, test2
}

func test3() { // non-linear generation
	fmt.Println("=== test3 ===")
	g := core.NewGenerator(public.StringResType, nil)
	g.AddString("test", "1")
	g.AddString("test2", "2")

	g.Pipeline = []string{"1", "2"}
	fmt.Println(core.GetStringRes(g, ", ")) // test, test2

	g.Pipeline = []string{"2", "1"}
	fmt.Println(core.GetStringRes(g, ", ")) // test2, test
}

func main() {
	test1()
	test2()
	test3()
}

/*
=== test1 ===
[0 1 2 3 4 5 6 7 8 9] <nil>
=== test2 ===
[test test2] <nil>
test, test2 <nil>
=== test3 ===
test, test2 <nil>
test2, test <nil>
*/
```

---

# example/packages/engine/core/logger/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# logger

'''go
import "github.com/pt-main/lc/example/packages/engine/core/logger"
'''

## Index



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# example/packages/engine/core/logger/main.go

```go
// Work with engine.core.Logger

package main

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine/core"
)

func test1() { // Simple example
	fmt.Println("=== test1 ===")
	l := core.NewLogger("") // logger with default status form

	// setup logging
	l.Logging["info"] = true

	// log will print if l.Logging["info"] is true
	l.PrintLog("info", "test info log")   // printing
	l.PrintLog("debug", "test debug log") // not printing
}

func test2() { // Hard example
	fmt.Println("=== test2 ===")
	// change default status form
	// 3 values input, colored output
	// using tap.colors for coloring text
	l := core.NewLogger("Status:[?RD]%v [?RT]Time:[?BE][%v] [?RT]Text:[?GN][%v][?RT]")

	l.MaxLogLength = 4 // set the saving logs limit to 4
	l.Logging["info"] = true

	// add 12 logs
	l.PrintLog("info", "first info log")
	for range 10 {
		l.PrintLog("debug", "test debug log")
	}
	l.PrintLog("info", "last info log")

	// print only 4 logs
	fmt.Println("\n[" + strings.Join(l.Log, "]\n[") + "]\n") // print all logs
	// or just
	fmt.Println(l.GetLog())
}

func main() {
	test1()
	test2()
}

/*
=== test1 ===
info [2026-08-04 19:55:38.651721 +0000 UTC] [test info log]

=== test2 ===
Status:info Time:[2026-08-04 19:55:38.651797 +0000 UTC] Text:[first info log]
Status:info Time:[2026-08-04 19:55:38.651977 +0000 UTC] Text:[last info log]

[Status:debug Time:[2026-08-04 19:55:38.651937 +0000 UTC] Text:[test debug log]]
[Status:debug Time:[2026-08-04 19:55:38.65195 +0000 UTC] Text:[test debug log]]
[Status:debug Time:[2026-08-04 19:55:38.651963 +0000 UTC] Text:[test debug log]]
[Status:info Time:[2026-08-04 19:55:38.651977 +0000 UTC] Text:[last info log]]

Status:debug Time:[2026-08-04 19:55:38.651937 +0000 UTC] Text:[test debug log]
Status:debug Time:[2026-08-04 19:55:38.65195 +0000 UTC] Text:[test debug log]
Status:debug Time:[2026-08-04 19:55:38.651963 +0000 UTC] Text:[test debug log]
Status:info Time:[2026-08-04 19:55:38.651977 +0000 UTC] Text:[last info log]
*/
```

---

# example/packages/engine/core/other/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# other

'''go
import "github.com/pt-main/lc/example/packages/engine/core/other"
'''

## Index



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# example/packages/engine/core/other/main.go

```go
// Work with engine.core.Scope

package main

import (
	"fmt"

	"github.com/pt-main/tap/core"
)

func test1() {
	s := core.ScopeType{} // just map[string]any
	s["Key"] = 0
	fmt.Println(core.ScopeGet[int](s, "Key")) // 0
}

func main() {
	test1()
}
```

---

# example/packages/engine/engines/byte/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# byte

'''go
import "github.com/pt-main/lc/example/packages/engine/engines/byte"
'''

## Index



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# example/packages/engine/engines/byte/main.go

```go
// Work with engine.ByteEngine

package main
```

---

# example/packages/engine/engines/string/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# string

'''go
import "github.com/pt-main/lc/example/packages/engine/engines/string"
'''

## Index



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# example/packages/engine/engines/string/main.go

```go
// Work with engine.StringEngine

package main

func test1() {
	// e := lc.NewStringEngine(public.StringResType, []string{"main"})
}
```

---

# example/readme/byte/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# byte

'''go
import "github.com/pt-main/lc/example/readme/byte"
'''

## Index



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# example/readme/byte/main.go

```go
package main

import (
	"fmt"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/bytecode"
)

func main() {
	// Parsing format:
	// instruction {
	//     [bytes : cmd] [bytes : argscount] [bytes : arglen]  [bytes arglen : arg],
	//                                       [bytes : arglen2] [bytes arglen2 : arg2]...
	// }
	parser := &byteParsing.Parser1{
		Config: byteParsing.Parser1Config{
			GConfig: bytecode.GenerationConfig{
				CommandBytelen:   1,
				ArgscountBytelen: 1,
				ArglenBytelen:    2,
				Endianess:        public.LittleEndian,
			},
			Shifter: bytecode.Shift{},
		},
	}

	engine, err := lc.NewEngineBuilder(public.ByteEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithByteParser(parser).
		WithDefaultEvents(true).
		WithColors().
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandByte(1, func(be enginepkg.ByteEngineInterface, node *byteParsing.ParsedBytes) core.ErrorInterface {
		for _, arg := range node.Args {
			if err := be.GetUep().Generator.AddString(string(arg), "main"); err != nil {
				return err
			}
		}
		return nil
	}, "add to output instruction", true)
	if err != nil {
		panic(err)
	}

	code := []byte{
		0x01,       // opcode=1
		0x01,       // argsCount=1
		0x03, 0x00, // arglen=3 (little endian, 2 bytes)
		0x61, 0x62, 0x63, // args="abc" (3 bytes)
	}

	err = engine.ProcessBytes(code)
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", out)
}
```

---

# example/readme/string/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# string

'''go
import "github.com/pt-main/lc/example/readme/string"
'''

## Index



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# example/readme/string/main.go

```go
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
)

func main() {
	parser := &stringParsing.Parser2{} // parsing format: 'command arg1, arg2...'

	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(parser).
		WithDefaultEvents(true).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandString("log", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		args, _ := node.Metadata["args"].(string)
		return se.GetUep().Generator.AddString(fmt.Sprintf("Log [%v]: %v",
			time.Now().Format(time.Stamp), args), "main")
	}, "append log with timestamp")
	if err != nil {
		panic(err)
	}

	err = engine.ProcessString(strings.Join([]string{
		"log service_start",
		"log service_ready",
	}, "\n"))
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
```

---

# example/tests/parser3Test/parser3.go

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/parsing/stringParsing/parser3"
)

func main() {
	fmt.Println("Lc version -", lc.Version)

	rules := []stringParsing.LexerRule{
		{Type: "NUMBER", Pattern: regexp2.MustCompile('\d+', 0)},
		{Type: "PLUS", Pattern: regexp2.MustCompile('\+', 0)},
		{Type: "MINUS", Pattern: regexp2.MustCompile('-', 0)},
		{Type: "STAR", Pattern: regexp2.MustCompile('\*', 0)},
		{Type: "SLASH", Pattern: regexp2.MustCompile('/', 0)},
		{Type: "LPAREN", Pattern: regexp2.MustCompile('\(', 0)},
		{Type: "RPAREN", Pattern: regexp2.MustCompile('\)', 0)},
		{Type: "WHITESPACE", Pattern: regexp2.MustCompile('\s+', 0)},
	}
	lexer := stringParsing.NewLexer(rules, &stringParsing.LexerConfig{
		UseBracketBalance: false,
	})

	grammar := parser3.Grammar{
		"expr": {
			Name: "expr",
			Expr: parser3.NodeExpr{
				NodeType: "expr",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.NamedExpr{RuleName: "term"},
						parser3.RepeatExpr{
							Expr: parser3.SequenceExpr{
								Exprs: []parser3.Expr{
									parser3.ChoiceExpr{
										Alternatives: []parser3.Expr{
											parser3.TokenExpr{TokenType: "PLUS"},
											parser3.TokenExpr{TokenType: "MINUS"},
										},
									},
									parser3.NamedExpr{RuleName: "term"},
								},
							},
							Min: 0,
						},
					},
				},
			},
		},
		"term": {
			Name: "term",
			Expr: parser3.NodeExpr{
				NodeType: "term",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.NamedExpr{RuleName: "factor"},
						parser3.RepeatExpr{
							Expr: parser3.SequenceExpr{
								Exprs: []parser3.Expr{
									parser3.ChoiceExpr{
										Alternatives: []parser3.Expr{
											parser3.TokenExpr{TokenType: "STAR"},
											parser3.TokenExpr{TokenType: "SLASH"},
										},
									},
									parser3.NamedExpr{RuleName: "factor"},
								},
							},
							Min: 0,
						},
					},
				},
			},
		},
		"factor": {
			Name: "factor",
			Expr: parser3.NodeExpr{
				NodeType: "factor",
				Expr: parser3.ChoiceExpr{
					Alternatives: []parser3.Expr{
						parser3.TokenExpr{TokenType: "NUMBER"},
						parser3.SequenceExpr{
							Exprs: []parser3.Expr{
								parser3.TokenExpr{TokenType: "LPAREN"},
								parser3.NamedExpr{RuleName: "expr"},
								parser3.TokenExpr{TokenType: "RPAREN"},
							},
						},
					},
				},
			},
		},
	}

	parser := parser3.NewParser(lexer, grammar, "expr", []string{"WHITESPACE"})

	expression := "3 + 5 * (2 - 1)"
	nodes, err := parser.Parse(expression)
	if err != nil {
		log.Fatalf("ParsingError: %v", err)
	}
	if len(nodes) == 0 {
		fmt.Println("Has no nodes")
		return
	}
	root := nodes[0]

	fmt.Printf("Raw expr: %q\n", expression)
	fmt.Println("\nAST (as text):")
	printAST(root, 0)

	cleanRoot := cleanNode(root)
	jsonData, err2 := json.MarshalIndent(cleanRoot, "", "  ")
	if err2 != nil {
		log.Fatalf("Json error: %v", err)
	}
	fmt.Println("\nAST (JSON):")
	fmt.Println(string(jsonData))
}

func printAST(node stringParsing.ParsedNode, indent int) {
	prefix := strings.Repeat("  ", indent)
	fmt.Printf("%s[%s] Raw: %q\n", prefix, node.Switch, node.Raw)
	if children, ok := node.Metadata["children"].([]stringParsing.ParsedNode); ok {
		for _, child := range children {
			printAST(child, indent+1)
		}
	}
}

func cleanNode(node stringParsing.ParsedNode) stringParsing.ParsedNode {
	newMeta := make(map[string]interface{})
	for k, v := range node.Metadata {
		if k == "__prev" || k == "__next" {
			continue
		}
		if k == "children" {
			if children, ok := v.([]stringParsing.ParsedNode); ok {
				newChildren := make([]stringParsing.ParsedNode, len(children))
				for i, child := range children {
					newChildren[i] = cleanNode(child)
				}
				newMeta[k] = newChildren
				continue
			}
		}
		newMeta[k] = v
	}
	return stringParsing.ParsedNode{
		Raw:      node.Raw,
		Switch:   node.Switch,
		Metadata: newMeta,
	}
}

/*
Lc version - 1.5.1
Raw expr: "3 + 5 * (2 - 1)"

AST (as text):
[expr] Raw: "3+5*(2-1)"
  [expr] Raw: "3+5*(2-1)"
    [term] Raw: "3"
      [factor] Raw: "3"
        [NUMBER] Raw: "3"
    [PLUS] Raw: "+"
    [term] Raw: "5*(2-1)"
      [factor] Raw: "5"
        [NUMBER] Raw: "5"
      [STAR] Raw: "*"
      [factor] Raw: "(2-1)"
        [LPAREN] Raw: "("
        [expr] Raw: "2-1"
          [term] Raw: "2"
            [factor] Raw: "2"
              [NUMBER] Raw: "2"
          [MINUS] Raw: "-"
          [term] Raw: "1"
            [factor] Raw: "1"
              [NUMBER] Raw: "1"
        [RPAREN] Raw: ")"



AST (JSON):
{
  "Raw": "3+5*(2-1)",
  "Switch": "expr",
  "Metadata": {
    "children": [
      {
        "Raw": "3+5*(2-1)",
        "Switch": "expr",
        "Metadata": {
          "children": [
            {
              "Raw": "3",
              "Switch": "term",
              "Metadata": {
                "children": [
                  {
                    "Raw": "3",
                    "Switch": "factor",
                    "Metadata": {
                      "children": [
                        {
                          "Raw": "3",
                          "Switch": "NUMBER",
                          "Metadata": {
                            "__raw": "3",
                            "__value": "3"
                          }
                        }
                      ]
                    }
                  }
                ]
              }
            },
            {
              "Raw": "+",
              "Switch": "PLUS",
              "Metadata": {
                "__raw": "+",
                "__value": "+"
              }
            },
            {
              "Raw": "5*(2-1)",
              "Switch": "term",
              "Metadata": {
                "children": [
                  {
                    "Raw": "5",
                    "Switch": "factor",
                    "Metadata": {
                      "children": [
                        {
                          "Raw": "5",
                          "Switch": "NUMBER",
                          "Metadata": {
                            "__raw": "5",
                            "__value": "5"
                          }
                        }
                      ]
                    }
                  },
                  {
                    "Raw": "*",
                    "Switch": "STAR",
                    "Metadata": {
                      "__raw": "*",
                      "__value": "*"
                    }
                  },
                  {
                    "Raw": "(2-1)",
                    "Switch": "factor",
                    "Metadata": {
                      "children": [
                        {
                          "Raw": "(",
                          "Switch": "LPAREN",
                          "Metadata": {
                            "__raw": "(",
                            "__value": "("
                          }
                        },
                        {
                          "Raw": "2-1",
                          "Switch": "expr",
                          "Metadata": {
                            "children": [
                              {
                                "Raw": "2",
                                "Switch": "term",
                                "Metadata": {
                                  "children": [
                                    {
                                      "Raw": "2",
                                      "Switch": "factor",
                                      "Metadata": {
                                        "children": [
                                          {
                                            "Raw": "2",
                                            "Switch": "NUMBER",
                                            "Metadata": {
                                              "__raw": "2",
                                              "__value": "2"
                                            }
                                          }
                                        ]
                                      }
                                    }
                                  ]
                                }
                              },
                              {
                                "Raw": "-",
                                "Switch": "MINUS",
                                "Metadata": {
                                  "__raw": "-",
                                  "__value": "-"
                                }
                              },
                              {
                                "Raw": "1",
                                "Switch": "term",
                                "Metadata": {
                                  "children": [
                                    {
                                      "Raw": "1",
                                      "Switch": "factor",
                                      "Metadata": {
                                        "children": [
                                          {
                                            "Raw": "1",
                                            "Switch": "NUMBER",
                                            "Metadata": {
                                              "__raw": "1",
                                              "__value": "1"
                                            }
                                          }
                                        ]
                                      }
                                    }
                                  ]
                                }
                              }
                            ]
                          }
                        },
                        {
                          "Raw": ")",
                          "Switch": "RPAREN",
                          "Metadata": {
                            "__raw": ")",
                            "__value": ")"
                          }
                        }
                      ]
                    }
                  }
                ]
              }
            }
          ]
        }
      }
    ]
  }
}
*/
```

---

# example/tests/speedtest/README.md

```md
# Speed tests

## byte/bench/byte_test.go

Bytecode raw hotloop speed with one instruction without auto bytecode instruction index shift. 

<details> <summary>Result</summary>

'''log
ITERATIONS=1_000_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/byte/bench -count=10
Lc version - 1.5.1
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8              1        5762814944 ns/op               173.5 Mops/s        16344 B/op        176 allocs/op
BenchmarkByteProcessing-8              1        5751776315 ns/op               173.9 Mops/s        16184 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        5897353936 ns/op               169.6 Mops/s        16952 B/op        184 allocs/op
BenchmarkByteProcessing-8              1        6498543824 ns/op               153.9 Mops/s        16376 B/op        176 allocs/op
BenchmarkByteProcessing-8              1        6598746076 ns/op               151.5 Mops/s        14776 B/op        161 allocs/op
BenchmarkByteProcessing-8              1        7007383883 ns/op               142.7 Mops/s        16264 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        6527556070 ns/op               153.2 Mops/s        16280 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        6253077383 ns/op               159.9 Mops/s        16424 B/op        177 allocs/op
BenchmarkByteProcessing-8              1        6086296262 ns/op               164.3 Mops/s        16248 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        6452127563 ns/op               155.0 Mops/s        16248 B/op        175 allocs/op
PASS
ok      github.com/pt-main/lc/example/speedtest/byte/bench   63.261s

ITERATIONS=500_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/byte/bench -count=10
Lc version - 1.5.1
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8              1        2725905999 ns/op               183.4 Mops/s        20904 B/op        174 allocs/op
BenchmarkByteProcessing-8              1        2735958753 ns/op               182.8 Mops/s        15592 B/op        170 allocs/op
BenchmarkByteProcessing-8              1        2721029300 ns/op               183.8 Mops/s        15480 B/op        168 allocs/op
BenchmarkByteProcessing-8              1        2904751570 ns/op               172.1 Mops/s        16408 B/op        177 allocs/op
BenchmarkByteProcessing-8              1        2868240595 ns/op               174.3 Mops/s        16904 B/op        182 allocs/op
BenchmarkByteProcessing-8              1        2987468020 ns/op               167.4 Mops/s        15576 B/op        168 allocs/op
BenchmarkByteProcessing-8              1        2805655377 ns/op               178.2 Mops/s        16344 B/op        176 allocs/op
BenchmarkByteProcessing-8              1        2860468821 ns/op               174.8 Mops/s        15560 B/op        168 allocs/op
BenchmarkByteProcessing-8              1        2881849876 ns/op               173.5 Mops/s        16168 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        3005677711 ns/op               166.4 Mops/s        16952 B/op        182 allocs/op
PASS
ok      github.com/pt-main/lc/example/speedtest/byte/bench   28.963s

ITERATIONS=300_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/byte/bench -count=10
Lc version - 1.5.1
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8              1        1788257526 ns/op               167.8 Mops/s        21896 B/op        183 allocs/op
BenchmarkByteProcessing-8              1        1735185620 ns/op               172.9 Mops/s        17048 B/op        182 allocs/op
BenchmarkByteProcessing-8              1        1743715603 ns/op               172.0 Mops/s        16440 B/op        176 allocs/op
BenchmarkByteProcessing-8              1        1698821182 ns/op               176.6 Mops/s        15512 B/op        168 allocs/op
BenchmarkByteProcessing-8              1        1723587268 ns/op               174.1 Mops/s        16376 B/op        177 allocs/op
BenchmarkByteProcessing-8              1        1753887836 ns/op               171.0 Mops/s        14872 B/op        161 allocs/op
BenchmarkByteProcessing-8              1        1743975178 ns/op               172.0 Mops/s        17064 B/op        182 allocs/op
BenchmarkByteProcessing-8              1        1924000551 ns/op               155.9 Mops/s        15576 B/op        168 allocs/op
BenchmarkByteProcessing-8              1        1807119062 ns/op               166.0 Mops/s        16344 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        1769031176 ns/op               169.6 Mops/s        17000 B/op        182 allocs/op
PASS
ok      github.com/pt-main/lc/example/speedtest/byte/bench   18.120s


ITERATIONS=100_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/byte/bench -count=10
Lc version - 1.5.1
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8            933           1278002 ns/op                83.87 Mops/s       13039 B/op        158 allocs/op
BenchmarkByteProcessing-8            930           1232047 ns/op                87.28 Mops/s       13038 B/op        158 allocs/op
BenchmarkByteProcessing-8            838           1381281 ns/op                86.39 Mops/s       13099 B/op        158 allocs/op
BenchmarkByteProcessing-8            924           1189648 ns/op                90.97 Mops/s       13053 B/op        158 allocs/op
BenchmarkByteProcessing-8            914           1211689 ns/op                90.29 Mops/s       13026 B/op        157 allocs/op
BenchmarkByteProcessing-8            909           1201342 ns/op                91.57 Mops/s       13058 B/op        158 allocs/op
BenchmarkByteProcessing-8            904           1205359 ns/op                91.77 Mops/s       13062 B/op        158 allocs/op
BenchmarkByteProcessing-8            927           1188666 ns/op                90.75 Mops/s       13054 B/op        158 allocs/op
BenchmarkByteProcessing-8            915           1203756 ns/op                90.79 Mops/s       13034 B/op        158 allocs/op
BenchmarkByteProcessing-8            770           1318173 ns/op                98.52 Mops/s       13203 B/op        158 allocs/op
PASS
ok      github.com/pt-main/lc/example/speedtest/byte/bench   92.990s
'''

</details>

## byte/tests/main.go

Bytecode raw hotloop speed with small handler - 

'''go
func(be *engine.ByteEngine, pb *byteParsing.ParsedBytes) error {
    allIters += 1
    if allIters >= (ITERATIONS - 1) {
        timeE = time.Now()
        *idx = -1
    }
    return nil
}
'''

- 'Test1' - test with pre generated instructions
- 'Test2' - test with one instruction without auto bytecode instruction index shift. 

<details> <summary>Result</summary>

'''log
Lc version - 1.5.1
===== Test1 =====
Test1: Iters: 10000000, Time: 83.631101ms, 119047619 mOps/s
Test1: Iters: 10000000, Time: 80.806105ms, 123456790 mOps/s
Test1: Iters: 10000000, Time: 83.966236ms, 119047619 mOps/s
Test1: Iters: 10000000, Time: 84.77732ms, 117647058 mOps/s
Test1: Iters: 10000000, Time: 92.00846ms, 108695652 mOps/s
Test1: Iters: 10000000, Time: 81.311348ms, 123456790 mOps/s
Test1: Iters: 10000000, Time: 81.962906ms, 121951219 mOps/s
Test1: Iters: 10000000, Time: 80.21774ms, 125000000 mOps/s
Test1: Iters: 10000000, Time: 87.345851ms, 114942528 mOps/s
Test1: Iters: 10000000, Time: 82.67824ms, 120481927 mOps/s
Test1: Iters: 100000000, Time: 802.337396ms, 124688279 mOps/s
Test1: Iters: 100000000, Time: 847.627388ms, 117924528 mOps/s
Test1: Iters: 100000000, Time: 798.604082ms, 125156445 mOps/s
Test1: Iters: 100000000, Time: 788.823594ms, 126742712 mOps/s
Test1: Iters: 100000000, Time: 762.808279ms, 131061598 mOps/s
Test1: Iters: 200000000, Time: 1.682078851s, 118906064 mOps/s
Test1: Iters: 200000000, Time: 1.645749681s, 121506682 mOps/s
Test1: Iters: 400000000, Time: 9.133057429s, 43797218 mOps/s
Medium: 116861707, Median: 121506682,
Iters [43797218 108695652 114942528 117647058 117924528 118906064 119047619 119047619 120481927 121506682 121951219 123456790 123456790 124688279 125000000 125156445 126742712 131061598]
===== Test2 =====
Test2: Iters: 10000000, Time: 74.136485ms, 135135135 mOps/s
Test2: Iters: 10000000, Time: 80.690947ms, 123456790 mOps/s
Test2: Iters: 10000000, Time: 74.163965ms, 135135135 mOps/s
Test2: Iters: 10000000, Time: 109.85331ms, 90909090 mOps/s
Test2: Iters: 10000000, Time: 76.635534ms, 129870129 mOps/s
Test2: Iters: 10000000, Time: 66.212814ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 76.931153ms, 129870129 mOps/s
Test2: Iters: 10000000, Time: 66.131764ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 63.494247ms, 158730158 mOps/s
Test2: Iters: 10000000, Time: 66.28764ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 67.429024ms, 149253731 mOps/s
Test2: Iters: 10000000, Time: 67.209301ms, 149253731 mOps/s
Test2: Iters: 10000000, Time: 64.580462ms, 153846153 mOps/s
Test2: Iters: 10000000, Time: 72.29076ms, 138888888 mOps/s
Test2: Iters: 10000000, Time: 65.606012ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 64.321543ms, 156250000 mOps/s
Test2: Iters: 10000000, Time: 64.266137ms, 156250000 mOps/s
Test2: Iters: 10000000, Time: 69.527309ms, 142857142 mOps/s
Test2: Iters: 10000000, Time: 75.758207ms, 131578947 mOps/s
Test2: Iters: 10000000, Time: 65.858495ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 65.835576ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 62.492821ms, 161290322 mOps/s
Test2: Iters: 10000000, Time: 69.924138ms, 142857142 mOps/s
Test2: Iters: 10000000, Time: 64.052846ms, 156250000 mOps/s
Test2: Iters: 10000000, Time: 65.010355ms, 153846153 mOps/s
Test2: Iters: 10000000, Time: 66.219458ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 65.548766ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 70.480818ms, 142857142 mOps/s
Test2: Iters: 10000000, Time: 62.270297ms, 161290322 mOps/s
Test2: Iters: 10000000, Time: 62.410387ms, 161290322 mOps/s
Test2: Iters: 10000000, Time: 61.298797ms, 163934426 mOps/s
Test2: Iters: 10000000, Time: 69.284619ms, 144927536 mOps/s
Test2: Iters: 100000000, Time: 643.836767ms, 155279503 mOps/s
Test2: Iters: 100000000, Time: 685.867418ms, 145772594 mOps/s
Test2: Iters: 100000000, Time: 653.605758ms, 152905198 mOps/s
Test2: Iters: 100000000, Time: 634.22513ms, 157728706 mOps/s
Test2: Iters: 100000000, Time: 629.747093ms, 158730158 mOps/s
Test2: Iters: 100000000, Time: 636.79118ms, 156985871 mOps/s
Test2: Iters: 100000000, Time: 631.620296ms, 158227848 mOps/s
Test2: Iters: 100000000, Time: 637.004552ms, 156985871 mOps/s
Test2: Iters: 100000000, Time: 639.77815ms, 156250000 mOps/s
Test2: Iters: 100000000, Time: 627.149309ms, 159489633 mOps/s
Test2: Iters: 100000000, Time: 630.189662ms, 158730158 mOps/s
Test2: Iters: 100000000, Time: 640.183614ms, 156250000 mOps/s
Test2: Iters: 100000000, Time: 671.145516ms, 149031296 mOps/s
Test2: Iters: 100000000, Time: 673.820605ms, 148367952 mOps/s
Test2: Iters: 100000000, Time: 751.899969ms, 132978723 mOps/s
Test2: Iters: 100000000, Time: 928.532697ms, 107642626 mOps/s
Test2: Iters: 200000000, Time: 1.309047815s, 152788388 mOps/s
Test2: Iters: 200000000, Time: 1.417237526s, 141143260 mOps/s
Test2: Iters: 200000000, Time: 1.3435233s, 148809523 mOps/s
Test2: Iters: 200000000, Time: 1.329056673s, 150489089 mOps/s
Test2: Iters: 200000000, Time: 1.239711748s, 161290322 mOps/s
Test2: Iters: 200000000, Time: 1.234268487s, 162074554 mOps/s
Test2: Iters: 200000000, Time: 1.264810986s, 158102766 mOps/s
Test2: Iters: 200000000, Time: 1.236752963s, 161681487 mOps/s
Test2: Iters: 400000000, Time: 2.558859178s, 156311059 mOps/s
Test2: Iters: 400000000, Time: 2.614027526s, 153022188 mOps/s
Test2: Iters: 400000000, Time: 2.528229248s, 158227848 mOps/s
Test2: Iters: 400000000, Time: 2.580307925s, 155038759 mOps/s
Test2: Iters: 800000000, Time: 5.04125604s, 158698670 mOps/s
Test2: Iters: 800000000, Time: 5.089182014s, 157201807 mOps/s
Medium: 149486864, Median: 152905198,
Iters [90909090 107642626 123456790 129870129 129870129 131578947 132978723 135135135 135135135 138888888 141143260 142857142 142857142 142857142 144927536 145772594 148367952 148809523 149031296 149253731 149253731 150489089 151515151 151515151 151515151 151515151 151515151 151515151 151515151 151515151 152788388 152905198 153022188 153846153 153846153 155038759 155279503 156250000 156250000 156250000 156250000 156250000 156311059 156985871 156985871 157201807 157728706 158102766 158227848 158227848 158698670 158730158 158730158 158730158 159489633 161290322 161290322 161290322 161290322 161681487 162074554 163934426]
'''

</details>
```

---

# example/tests/speedtest/byte/bench/byte_test.go

```go
package speedtest

import (
	"context"
	"testing"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/bytecode"
)

func BenchmarkByteProcessing(b *testing.B) {
	const ITERATIONS = 100_000_000
	_idx := 0
	end := public.BigEndian
	gc := byteParsing.Parser1Config{
		GConfig: bytecode.GenerationConfig{
			CommandBytelen: 1, ArgscountBytelen: 1,
			ArglenBytelen: 1, Endianess: end,
		}, Shifter: *bytecode.NewShift(make([]byte, 0), &_idx),
	}
	e := lc.NewByteEngine(
		0, nil, true, &byteParsing.Parser1{Config: gc},
		end, true, context.Background(),
	)
	var iteration int
	e.NewCommandFull(0, func(bei engine.ByteEngineInterface, pb *byteParsing.ParsedBytes) core.ErrorInterface {
		be := bei.(*engine.ByteEngine)
		iteration += 1
		if iteration+1 == ITERATIONS {
			be.SetBytecodeIdx(-1)
		}
		return nil
	}, "nil", false)

	b.ResetTimer()
	for range b.N {
		iteration = 0
		e.Process([]byte{0, 0})
		b.ReportMetric(float64(ITERATIONS)/b.Elapsed().Seconds()/1e6, "Mops/s")
	}
}

/* ----==== RESULTS ====----

Lc version - 1.5.7

ITERATIONS=1_000_000_000
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/tests/speedtest/byte/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/tests/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8               1        4636165728 ns/op               215.7 Mops/s         6992 B/op        105 allocs/op
BenchmarkByteProcessing-8               1        4678134135 ns/op               213.8 Mops/s         6976 B/op        104 allocs/op
BenchmarkByteProcessing-8               1        4545117089 ns/op               220.0 Mops/s         6576 B/op        102 allocs/op
BenchmarkByteProcessing-8               1        4521750876 ns/op               221.2 Mops/s         6976 B/op        104 allocs/op
BenchmarkByteProcessing-8               1        5308026598 ns/op               188.4 Mops/s         6976 B/op        104 allocs/op
BenchmarkByteProcessing-8               1        4701634606 ns/op               212.7 Mops/s         6976 B/op        104 allocs/op
BenchmarkByteProcessing-8               1        4964960666 ns/op               201.4 Mops/s         6528 B/op        101 allocs/op
BenchmarkByteProcessing-8               1        4707789927 ns/op               212.4 Mops/s         6528 B/op        101 allocs/op
BenchmarkByteProcessing-8               1        4483261343 ns/op               223.1 Mops/s         6528 B/op        101 allocs/op
BenchmarkByteProcessing-8               1        4644296688 ns/op               215.3 Mops/s         6912 B/op        104 allocs/op
PASS
ok      github.com/pt-main/lc/example/tests/speedtest/byte/bench        47.634s

ITERATIONS=500_000_000
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/tests/speedtest/byte/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/tests/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8              1        2416450521 ns/op               206.9 Mops/s         8440 B/op        109 allocs/op
BenchmarkByteProcessing-8              1        2308473392 ns/op               216.6 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2233277662 ns/op               223.9 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2194493841 ns/op               227.8 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2419690868 ns/op               206.6 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2407318933 ns/op               207.7 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        2760941809 ns/op               181.1 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        2616158593 ns/op               191.1 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2629497097 ns/op               190.2 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2482259449 ns/op               201.4 Mops/s         8376 B/op        108 allocs/op
PASS
ok      github.com/pt-main/lc/example/tests/speedtest/byte/bench        24.940s

ITERATIONS=250_000_000
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/tests/speedtest/byte/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/tests/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8              1        1155587732 ns/op               216.3 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        1101299266 ns/op               227.0 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        1200872455 ns/op               208.2 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        1187507477 ns/op               210.5 Mops/s         8392 B/op        109 allocs/op
BenchmarkByteProcessing-8              1        1209859392 ns/op               206.6 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        1182012312 ns/op               211.5 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        1182744295 ns/op               211.4 Mops/s         7784 B/op        104 allocs/op
BenchmarkByteProcessing-8              1        1196835605 ns/op               208.9 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        1188541998 ns/op               210.3 Mops/s         7848 B/op        104 allocs/op
BenchmarkByteProcessing-8              1        1139287512 ns/op               219.4 Mops/s         8376 B/op        108 allocs/op
PASS
ok      github.com/pt-main/lc/example/tests/speedtest/byte/bench        12.220s

ITERATIONS=100_000_000
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/tests/speedtest/byte/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/tests/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8           6278            192517 ns/op                82.74 Mops/s        6591 B/op         98 allocs/op
BenchmarkByteProcessing-8           6279            160522 ns/op                99.21 Mops/s        6591 B/op         98 allocs/op
BenchmarkByteProcessing-8           5919            182307 ns/op                92.67 Mops/s        6496 B/op         98 allocs/op
BenchmarkByteProcessing-8           6784            161409 ns/op                91.32 Mops/s        6550 B/op         98 allocs/op
BenchmarkByteProcessing-8           6258            162812 ns/op                98.15 Mops/s        6592 B/op         98 allocs/op
BenchmarkByteProcessing-8           6234            163855 ns/op                97.90 Mops/s        6594 B/op         98 allocs/op
BenchmarkByteProcessing-8           5823            181810 ns/op                94.46 Mops/s        6503 B/op         98 allocs/op
BenchmarkByteProcessing-8           4544            226392 ns/op                97.21 Mops/s        6499 B/op         98 allocs/op
BenchmarkByteProcessing-8           5347            198100 ns/op                94.41 Mops/s        6543 B/op         98 allocs/op
BenchmarkByteProcessing-8           5827            171721 ns/op                99.94 Mops/s        6503 B/op         98 allocs/op
PASS
ok      github.com/pt-main/lc/example/tests/speedtest/byte/bench        76.599s
*/
```

---

# example/tests/speedtest/byte/tests/main.go

```go
package main

import (
	"context"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/engine/events"
	"github.com/pt-main/lc/parsing/byteParsing"
)

func test1(ITERATIONS int) int {
	fmt.Printf("Test1: ")
	_idx := 0
	idx := &_idx
	e := lc.NewByteEngine(
		0, nil, true, &byteParsing.Parser1{},
		0, true, context.Background(),
	) // pseudo engine

	parsed := make([]events.ByteCallAttr, 0, ITERATIONS)
	rawNode := byteParsing.ParsedBytes{
		Switch: []byte{0},
		Args:   [][]byte{},
	}
	handler := func(be engine.ByteEngineInterface, pb *byteParsing.ParsedBytes) core.ErrorInterface {
		return nil
	}
	bca := events.ByteCallAttr{
		RawNode: &rawNode,
		Handler: handler,
		Abis:    true,
	}
	for i := 0; i < int(ITERATIONS); i++ {
		parsed = append(parsed, bca)
	}

	de := events.DefaultEvents{}
	i := &core.EventInput{
		Input: events.ByteCLDType{
			Ctx:    e.UEP.Context,
			Parsed: parsed,
			Engine: e,
			Idx:    idx,
		},
	}
	ev := e.UEP.Event.(*core.Events)

	// f, err := os.Create("cpu.prof")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// if err := pprof.StartCPUProfile(f); err != nil {
	// 	panic(err)
	// }

	var timeS time.Time = time.Now()
	err := de.ByteCallHotLoopEvent(ev, i)
	timeE := time.Now()
	// pprof.StopCPUProfile()
	time := timeE.Sub(timeS)
	if err != nil {
		fmt.Println(err)
	}
	mops := int(float64(ITERATIONS) / (math.Round(time.Seconds()*1000) / 1000))
	fmt.Printf("Iters: %v, Time: %v, %v mOps/s\n", ITERATIONS, time, mops)
	return mops
}

func test2(ITERATIONS int) int {
	fmt.Printf("Test2: ")
	_idx := 0
	idx := &_idx
	e := lc.NewByteEngine(
		0, nil, true, &byteParsing.Parser1{},
		0, true, context.Background(),
	) // pseudo engine

	allIters := 0

	rawNode := byteParsing.ParsedBytes{
		Switch: []byte{0},
		Args:   [][]byte{},
	}
	var timeE time.Time
	handler := func(bei engine.ByteEngineInterface, pb *byteParsing.ParsedBytes) core.ErrorInterface {
		allIters += 1
		if allIters >= (ITERATIONS - 1) {
			timeE = time.Now()
			*idx = -1
		}
		return nil
	}
	bca := events.ByteCallAttr{
		RawNode: &rawNode,
		Handler: handler,
	}

	de := events.DefaultEvents{}
	i := &core.EventInput{
		Input: events.ByteCLDType{
			Ctx:    e.UEP.Context,
			Parsed: []events.ByteCallAttr{bca},
			Engine: e,
			Idx:    idx,
		},
	}
	ev := e.UEP.Event.(*core.Events)

	// f, err := os.Create("cpu.prof")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// if err := pprof.StartCPUProfile(f); err != nil {
	// 	panic(err)
	// }

	var timeS time.Time = time.Now()
	err := de.ByteCallHotLoopEvent(ev, i)
	// pprof.StopCPUProfile()
	time := timeE.Sub(timeS)
	if err != nil {
		fmt.Println(err)
	}
	mops := int(float64(ITERATIONS) / (math.Round(time.Seconds()*1000) / 1000))
	fmt.Printf("Iters: %v, Time: %v, %v mOps/s\n", ITERATIONS, time, mops)
	return mops
}

func test(testf func(int) int, Start int) {
	iters := []int{}
	for range Start {
		iters = append(iters, testf(5_000_000))
	}
	for range Start / 2 {
		iters = append(iters, testf(10_000_000))
	}
	for range Start / 4 {
		iters = append(iters, testf(20_000_000))
	}
	for range Start / 8 {
		iters = append(iters, testf(40_000_000))
	}
	for range Start / 16 {
		iters = append(iters, testf(80_000_000))
	}
	for range Start / 32 {
		iters = append(iters, testf(160_000_000))
	}
	for range Start / 64 {
		iters = append(iters, testf(320_000_000))
	}
	res := 0
	for _, val := range iters {
		res += val
	}
	res /= len(iters)
	m := int(math.Round(float64(len(iters) / 2)))
	sort.Ints(iters)
	fmt.Printf("Medium: %v, Median: %v, \nIters %v\n", res, iters[m], iters)
}

func main() {
	fmt.Println("Lc version -", lc.Version)
	fmt.Println("===== Test1 =====")
	test(test1, 32)
	fmt.Println("===== Test2 =====")
	test(test2, 32)
}

/*
macbook@MacBook-Pro lc % go run ./example/tests/speedtest/byte/tests -gcflags="-m -m" -ldflags="-s -w"
Lc version - 1.5.7
===== Test1 =====
Test1: Iters: 5000000, Time: 19.331003ms, 263157894 mOps/s
Test1: Iters: 5000000, Time: 18.733908ms, 263157894 mOps/s
Test1: Iters: 5000000, Time: 22.980312ms, 217391304 mOps/s
Test1: Iters: 5000000, Time: 19.558461ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 19.739296ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 21.091147ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 21.742407ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 19.765059ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 21.57979ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 21.565381ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 22.473135ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 22.237547ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 23.38193ms, 217391304 mOps/s
Test1: Iters: 5000000, Time: 22.048882ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 20.128706ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 21.16066ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 19.281225ms, 263157894 mOps/s
Test1: Iters: 5000000, Time: 20.663264ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.035794ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 23.065067ms, 217391304 mOps/s
Test1: Iters: 5000000, Time: 22.874167ms, 217391304 mOps/s
Test1: Iters: 5000000, Time: 20.898845ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.811976ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 21.192832ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 21.092074ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.848486ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.688949ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 19.763027ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 21.939585ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 20.649626ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.758666ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.853927ms, 238095238 mOps/s
Test1: Iters: 10000000, Time: 39.078747ms, 256410256 mOps/s
Test1: Iters: 10000000, Time: 40.295967ms, 250000000 mOps/s
Test1: Iters: 10000000, Time: 38.578624ms, 256410256 mOps/s
Test1: Iters: 10000000, Time: 37.681705ms, 263157894 mOps/s
Test1: Iters: 10000000, Time: 38.255899ms, 263157894 mOps/s
Test1: Iters: 10000000, Time: 40.260515ms, 250000000 mOps/s
Test1: Iters: 10000000, Time: 40.59357ms, 243902439 mOps/s
Test1: Iters: 10000000, Time: 41.570728ms, 238095238 mOps/s
Test1: Iters: 10000000, Time: 38.745396ms, 256410256 mOps/s
Test1: Iters: 10000000, Time: 39.545718ms, 250000000 mOps/s
Test1: Iters: 10000000, Time: 38.340595ms, 263157894 mOps/s
Test1: Iters: 10000000, Time: 37.275479ms, 270270270 mOps/s
Test1: Iters: 10000000, Time: 39.588624ms, 250000000 mOps/s
Test1: Iters: 10000000, Time: 40.150805ms, 250000000 mOps/s
Test1: Iters: 10000000, Time: 37.195498ms, 270270270 mOps/s
Test1: Iters: 10000000, Time: 40.036955ms, 250000000 mOps/s
Test1: Iters: 20000000, Time: 74.349315ms, 270270270 mOps/s
Test1: Iters: 20000000, Time: 74.60037ms, 266666666 mOps/s
Test1: Iters: 20000000, Time: 73.616943ms, 270270270 mOps/s
Test1: Iters: 20000000, Time: 71.842844ms, 277777777 mOps/s
Test1: Iters: 20000000, Time: 76.208526ms, 263157894 mOps/s
Test1: Iters: 20000000, Time: 70.059533ms, 285714285 mOps/s
Test1: Iters: 20000000, Time: 75.241165ms, 266666666 mOps/s
Test1: Iters: 20000000, Time: 71.8311ms, 277777777 mOps/s
Test1: Iters: 40000000, Time: 148.648065ms, 268456375 mOps/s
Test1: Iters: 40000000, Time: 145.840353ms, 273972602 mOps/s
Test1: Iters: 40000000, Time: 146.652388ms, 272108843 mOps/s
Test1: Iters: 40000000, Time: 146.301707ms, 273972602 mOps/s
Test1: Iters: 80000000, Time: 285.775098ms, 279720279 mOps/s
Test1: Iters: 80000000, Time: 308.652823ms, 258899676 mOps/s
Test1: Iters: 160000000, Time: 867.274031ms, 184544405 mOps/s
Medium: 248862061, Median: 250000000,
Iters [184544405 217391304 217391304 217391304 217391304 227272727 227272727 227272727 227272727 227272727 227272727 227272727 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 243902439 250000000 250000000 250000000 250000000 250000000 250000000 250000000 250000000 250000000 250000000 250000000 250000000 256410256 256410256 256410256 258899676 263157894 263157894 263157894 263157894 263157894 263157894 263157894 266666666 266666666 268456375 270270270 270270270 270270270 270270270 272108843 273972602 273972602 277777777 277777777 279720279 285714285]
===== Test2 =====
Test2: Iters: 5000000, Time: 23.743946ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 23.778253ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 24.995652ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.941669ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 24.6286ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 24.779641ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 25.993257ms, 192307692 mOps/s
Test2: Iters: 5000000, Time: 25.318152ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.632493ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 23.744656ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 23.530285ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 22.8118ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 23.455578ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 23.230671ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 24.748188ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.377032ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 24.238652ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 24.465862ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 24.719291ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.440745ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 23.847918ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 24.911544ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.74668ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 22.658085ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 22.688899ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 23.432506ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 22.990193ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 22.863689ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 25.492215ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.930831ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 22.686733ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 23.156547ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.848181ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.559144ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.96151ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.77535ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.637576ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.141271ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 45.407469ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 44.985172ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 44.891841ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 47.209454ms, 212765957 mOps/s
Test2: Iters: 10000000, Time: 45.353728ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 45.751145ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.442977ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 47.030275ms, 212765957 mOps/s
Test2: Iters: 10000000, Time: 45.153691ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 47.250445ms, 212765957 mOps/s
Test2: Iters: 20000000, Time: 90.743824ms, 219780219 mOps/s
Test2: Iters: 20000000, Time: 98.412486ms, 204081632 mOps/s
Test2: Iters: 20000000, Time: 93.030863ms, 215053763 mOps/s
Test2: Iters: 20000000, Time: 90.437058ms, 222222222 mOps/s
Test2: Iters: 20000000, Time: 97.165843ms, 206185567 mOps/s
Test2: Iters: 20000000, Time: 91.614653ms, 217391304 mOps/s
Test2: Iters: 20000000, Time: 93.531025ms, 212765957 mOps/s
Test2: Iters: 20000000, Time: 88.639409ms, 224719101 mOps/s
Test2: Iters: 40000000, Time: 179.760469ms, 222222222 mOps/s
Test2: Iters: 40000000, Time: 185.233176ms, 216216216 mOps/s
Test2: Iters: 40000000, Time: 183.163349ms, 218579234 mOps/s
Test2: Iters: 40000000, Time: 220.352058ms, 181818181 mOps/s
Test2: Iters: 80000000, Time: 373.641796ms, 213903743 mOps/s
Test2: Iters: 80000000, Time: 376.967661ms, 212201591 mOps/s
Test2: Iters: 160000000, Time: 723.861096ms, 220994475 mOps/s
Medium: 212682645, Median: 217391304,
Iters [181818181 192307692 200000000 200000000 200000000 200000000 200000000 200000000 200000000 200000000 204081632 206185567 208333333 208333333 208333333 208333333 208333333 208333333 208333333 208333333 208333333 208333333 208333333 212201591 212765957 212765957 212765957 212765957 213903743 215053763 216216216 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 218579234 219780219 220994475 222222222 222222222 222222222 222222222 222222222 222222222 222222222 222222222 222222222 224719101]
*/
```

---

# go.mod

```mod
module github.com/pt-main/lc

go 1.24.13

require github.com/dlclark/regexp2 v1.12.0

require github.com/pt-main/tap v1.4.7
```

---

# go.sum

```sum
github.com/dlclark/regexp2 v1.12.0 h1:0j4c5qQmnC6XOWNjP3PIXURXN2gWx76rd3KvgdPkCz8=
github.com/dlclark/regexp2 v1.12.0/go.mod h1:DHkYz0B9wPfa6wondMfaivmHpzrQ3v9q8cnmRbL6yW8=
github.com/pt-main/tap v1.4.5 h1:g6W1aIx7qHhVmugicqGRvk5UugpyGWDn/ip57iCIK6E=
github.com/pt-main/tap v1.4.5/go.mod h1:ULQUJ/+8VIji9oq26pr1cmbXv+VUlhjsvq1n/vd4f3I=
github.com/pt-main/tap v1.4.6 h1:Tx+/riiw558qy4Az2QYXlc0QP0UhBConv7D2QZs++a8=
github.com/pt-main/tap v1.4.6/go.mod h1:ULQUJ/+8VIji9oq26pr1cmbXv+VUlhjsvq1n/vd4f3I=
github.com/pt-main/tap v1.4.7 h1:73hoVgdAng1qh/TKLzfR2sREHCTcNFuOpwseQupD+5Y=
github.com/pt-main/tap v1.4.7/go.mod h1:ULQUJ/+8VIji9oq26pr1cmbXv+VUlhjsvq1n/vd4f3I=
```

---

# main.go

```go
package lc

import (
	"context"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/engine/events"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/tap/color"
)

const Version = "1.5.7"

// NewStringEngine creates a ready-to-use string-based engine.
// Parameters:
//
//	generator_res_type – core.StringResType (usually) for text generation.
//	pipeline – ordered list of generation points (e.g., []string{"pre","main"}).
//	add_default_events – if true, registers standard parsing and call events.
//	parser – an implementation parser.ParserInterface.
//
// Returns a StringEngine with empty command map and initialized UEP.
func NewStringEngine(
	generator_res_type public.ResType,
	pipeline []string,
	add_default_events bool,
	parser stringParser,
	colorEnable bool,
	context context.Context,
) *engine.StringEngine {
	color.ColorEnabled = colorEnable
	e := core.NewEvents(context)
	if add_default_events {
		de := events.DefaultEvents{}
		e.NewEvent(public.StringParseEvent, de.StringParsingEvent)
		e.NewEvent(public.StringCallEvent, de.StringCallEvent)
		e.NewEvent(public.StringCallCalloopEvent, de.StringCallLoopEvent)
	}
	uep, _ := core.NewUniversalEngineParams(core.NewGenerator(generator_res_type, pipeline),
		e, make(core.ScopeType), core.NewLogger(""), context)
	return &engine.StringEngine{
		UEP:      uep,
		Commands: make(map[string]core.CommandMeta[engine.StringEngineInterface, stringParsing.ParsedNode]),
		Parser:   parser,
	}
}

// NewByteEngine creates a byte-oriented engine for binary formats or bytecode.
//
// The endianess parameter (e.g., bytecode.LittleEndian) is stored in scope.
//
// It registers default events when add_default_events is true.
//
// The parser must implement paraing.ParserInterface.
func NewByteEngine(
	generator_res_type public.ResType,
	pipeline []string,
	add_default_events bool,
	parser byteParser,
	endianess public.EndianType,
	colorEnable bool,
	context context.Context,
) *engine.ByteEngine {
	color.ColorEnabled = colorEnable
	idx := 0
	e := core.NewEvents(context)
	if add_default_events {
		de := events.DefaultEvents{}
		e.NewEvent(public.ByteParseEvent, de.ByteParsingEvent)
		e.NewEvent(public.ByteCallEvent, de.ByteCallEvent)
		e.NewEvent(public.ByteCallHotloopEvent, de.ByteCallHotLoopEvent)
	}
	uep, _ := core.NewUniversalEngineParams(core.NewGenerator(
		generator_res_type, pipeline,
	), e, core.ScopeType{
		public.ByteEngineScopeEndianess:   endianess,
		public.ByteEngineScopeBytecodeIdx: &idx,
	}, core.NewLogger(""), context)
	return &engine.ByteEngine{
		UEP:                    uep,
		AutoBytecodeIndexShift: make(map[int]bool),
		Commands:               make(map[int]core.CommandMeta[engine.ByteEngineInterface, byteParsing.ParsedBytes]),
		Parser:                 parser,
	}
}
```

---

# parsing/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# parsing

'''go
import "github.com/pt-main/lc/parsing"
'''

## Index

- [type ParseOption](<#ParseOption>)
- [type ParserInterface](<#ParserInterface>)


<a name="ParseOption"></a>
## type [ParseOption](<https://github.com/pt-main/Lc/blob/main/parsing/main.go#L5-L8>)



'''go
type ParseOption struct {
    UEP    *core.UniversalEngineParams
    Option core.Option
}
'''

<a name="ParserInterface"></a>
## type [ParserInterface](<https://github.com/pt-main/Lc/blob/main/parsing/main.go#L10-L13>)



'''go
type ParserInterface[I any, P any] interface {
    Parse(I, ...*ParseOption) ([]P, core.ErrorInterface)
    String() string
}
'''

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# parsing/README.md

```md
# Parsing

Package with default parsers.

- 'stringParsing/'
- 'byteParsing/'
```

---

# parsing/byteParsing/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# byteParsing

'''go
import "github.com/pt-main/lc/parsing/byteParsing"
'''

## Index

- [type ParsedBytes](<#ParsedBytes>)
- [type Parser1](<#Parser1>)
  - [func \(p \*Parser1\) Parse\(code \[\]byte, opts ...\*parsing.ParseOption\) \(result \[\]ParsedBytes, err core.ErrorInterface\)](<#Parser1.Parse>)
  - [func \(p \*Parser1\) String\(\) string](<#Parser1.String>)
- [type Parser1Config](<#Parser1Config>)


<a name="ParsedBytes"></a>
## type [ParsedBytes](<https://github.com/pt-main/Lc/blob/main/parsing/byteParsing/node.go#L7-L17>)

ParsedBytes represents a single parsed instruction or block in binary mode. It holds the raw byte slice of the entire instruction, the command identifier \(Switch\) as a byte slice, the list of argument byte slices \(Args\), and optional Metadata for additional context \(e.g., line numbers, offsets\).

'''go
type ParsedBytes struct {
    // Switch []byte – the command/opcode portion of the instruction.
    Switch []byte
    // Raw []byte – the complete original byte slice that produced this node.
    Raw []byte
    // Args [][]byte – each element is a raw byte slice of an argument.
    Args [][]byte
    // Metadata map[string]interface{} – extensible storage for extra info
    //   (e.g., "offset": 42, "line": 5, "source_file": "main.asm").
    Metadata map[string]interface{}
}
'''

<a name="Parser1"></a>
## type [Parser1](<https://github.com/pt-main/Lc/blob/main/parsing/byteParsing/parser1.go#L26-L28>)

Parser1 decodes a binary stream according to a fixed‑length field layout. Each bytecode instruction consists of:

'''
command (CommandBytelen bytes)
argscount (ArgscountBytelen bytes)
for each argument: arglen (ArglenBytelen bytes) followed by arg data.
'''

Endianess \(Little/BigEndian\) is used to decode integer fields.

'''go
type Parser1 struct {
    Config Parser1Config
}
'''

<a name="Parser1.Parse"></a>
### func \(\*Parser1\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/byteParsing/parser1.go#L43>)

'''go
func (p *Parser1) Parse(code []byte, opts ...*parsing.ParseOption) (result []ParsedBytes, err core.ErrorInterface)
'''

Parse reads the byte slice and returns a slice of ParsedBytes. Each ParsedBytes contains the raw command bytes, the raw arguments, and the original slice of the whole instruction. The ShiftStruct utility is used internally for safe bounds checking.

Err errors.ParsingError:

- On panic recovery. Meta: EMK\(0, "string"\) – the panic value.
- On shift error \(unexpected end of data\). Meta: EMK\(0, "int"\) – attempted length, EMK\(1, "int"\) – current byte index.
- On zero argument length. Meta: EMK\(0, "int"\) – argument number.
- On any other parsing error. Meta: EMK\(0, "int"\) – command byte index.

The returned error always contains the command bytes, raw bytes, and byte index in metadata.

<a name="Parser1.String"></a>
### func \(\*Parser1\) [String](<https://github.com/pt-main/Lc/blob/main/parsing/byteParsing/parser1.go#L187>)

'''go
func (p *Parser1) String() string
'''



<a name="Parser1Config"></a>
## type [Parser1Config](<https://github.com/pt-main/Lc/blob/main/parsing/byteParsing/parser1.go#L13-L16>)



'''go
type Parser1Config struct {
    GConfig bytecode.GenerationConfig
    Shifter bytecode.Shift
}
'''

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# parsing/byteParsing/node.go

```go
package byteParsing

// ParsedBytes represents a single parsed instruction or block in binary mode.
// It holds the raw byte slice of the entire instruction, the command identifier
// (Switch) as a byte slice, the list of argument byte slices (Args), and
// optional Metadata for additional context (e.g., line numbers, offsets).
type ParsedBytes struct {
	// Switch []byte – the command/opcode portion of the instruction.
	Switch []byte
	// Raw []byte – the complete original byte slice that produced this node.
	Raw []byte
	// Args [][]byte – each element is a raw byte slice of an argument.
	Args [][]byte
	// Metadata map[string]interface{} – extensible storage for extra info
	//   (e.g., "offset": 42, "line": 5, "source_file": "main.asm").
	Metadata map[string]interface{}
}
```

---

# parsing/byteParsing/parser1.go

```go
package byteParsing

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
	"github.com/pt-main/lc/tooling/bytecode"
)

type Parser1Config struct {
	GConfig bytecode.GenerationConfig
	Shifter bytecode.Shift
}

// Parser1 decodes a binary stream according to a fixed‑length field layout.
// Each bytecode instruction consists of:
//
//	command (CommandBytelen bytes)
//	argscount (ArgscountBytelen bytes)
//	for each argument: arglen (ArglenBytelen bytes) followed by arg data.
//
// Endianess (Little/BigEndian) is used to decode integer fields.
type Parser1 struct {
	Config Parser1Config
}

// Parse reads the byte slice and returns a slice of ParsedBytes.
// Each ParsedBytes contains the raw command bytes, the raw arguments,
// and the original slice of the whole instruction. The ShiftStruct utility
// is used internally for safe bounds checking.
//
// Err errors.ParsingError:
//   - On panic recovery. Meta: EMK(0, "string") – the panic value.
//   - On shift error (unexpected end of data). Meta: EMK(0, "int") – attempted length,
//     EMK(1, "int") – current byte index.
//   - On zero argument length. Meta: EMK(0, "int") – argument number.
//   - On any other parsing error. Meta: EMK(0, "int") – command byte index.
//
// The returned error always contains the command bytes, raw bytes, and byte index in metadata.
func (p *Parser1) Parse(code []byte, opts ...*parsing.ParseOption) (result []ParsedBytes, err core.ErrorInterface) {
	var lastCmdSwitch []byte = []byte{0}
	var raw []byte
	_idx := 0
	var idx *int = &_idx

	oldIdxPtr := p.Config.Shifter.Idx
	oldIdxVal := 0
	if oldIdxPtr != nil {
		oldIdxVal = *oldIdxPtr
	}

	defer func() {
		p.Config.Shifter.Idx = oldIdxPtr
		if oldIdxPtr != nil {
			*oldIdxPtr = oldIdxVal
		}

		if r := recover(); r != nil {
			err = core.Err(errors.ParsingError, "Panic recovered during parsing: %v", r).
				WithMeta(core.EMK(0, "string"), fmt.Sprintf("%v", r))
		}
		if err != nil {
			var cmdStr string
			var rawStr string
			if len(lastCmdSwitch) > 0 {
				cmdStr = fmt.Sprintf("%v", lastCmdSwitch)
			} else {
				cmdStr = "<unknown>"
			}
			if raw != nil {
				rawStr = fmt.Sprintf("%v", raw)
			} else {
				rawStr = "<none>"
			}
			idxVal := 0
			if idx != nil {
				idxVal = *idx
			}
			if ce, ok := err.(*core.Error); ok {
				if _, ok := ce.Meta[core.EMK(1, "string")]; !ok {
					ce.WithMeta(core.EMK(1, "string"), rawStr)
				}
				if _, ok := ce.Meta[core.EMK(2, "string")]; !ok {
					ce.WithMeta(core.EMK(2, "string"), cmdStr)
				}
				if _, ok := ce.Meta[core.EMK(3, "int")]; !ok {
					ce.WithMeta(core.EMK(3, "int"), idxVal)
				}
				err = ce
			} else {
				err = core.Wrap(errors.ParsingError, err, "Parsing error at cmd=%v, raw=%v, idx=%d", cmdStr, rawStr, idxVal).
					WithMeta(core.EMK(0, "string"), cmdStr).
					WithMeta(core.EMK(1, "string"), rawStr).
					WithMeta(core.EMK(2, "int"), idxVal)
			}
		}
	}()

	log := func(text string) {
		if len(opts) > 0 && opts[0] != nil {
			logger := opts[0].UEP.Logger
			if logger != nil {
				logger.PrintLog(public.LogParsing, text)
			}
		}
	}

	log(fmt.Sprintf("=========== START ==========="))
	log(fmt.Sprintf("start parsing code: '%v'", code))
	log(fmt.Sprintf("config: %v", p.Config))

	u := bytecode.Utils{}
	_Idx := p.Config.Shifter.Idx
	p.Config.Shifter.Idx = idx
	p.Config.Shifter.Code = code
	shift := p.Config.Shifter.ShiftError

	for *idx < len(code) {
		idxStart := *idx
		var command []byte
		command, err = shift(p.Config.GConfig.CommandBytelen)
		if err != nil {
			err = core.Wrap(errors.ParsingError, err, "Shift error while reading command").
				WithMeta(core.EMK(0, "int"), p.Config.GConfig.CommandBytelen).
				WithMeta(core.EMK(1, "int"), *idx)
			return
		}
		var argscountBytes []byte
		argscountBytes, err = shift(p.Config.GConfig.ArgscountBytelen)
		if err != nil {
			err = core.Wrap(errors.ParsingError, err, "Shift error while reading argscount").
				WithMeta(core.EMK(0, "int"), p.Config.GConfig.ArgscountBytelen).
				WithMeta(core.EMK(1, "int"), *idx)
			return
		}
		argscount := u.BytesToInt(argscountBytes, p.Config.GConfig.Endianess)
		args := [][]byte{}
		lastCmdSwitch = command
		log(fmt.Sprintf("cmd %v, argscount %v", command, argscount))

		for argNum := 0; argNum < argscount; argNum++ {
			var arglenBytes []byte
			arglenBytes, err = shift(p.Config.GConfig.ArglenBytelen)
			if err != nil {
				err = core.Wrap(errors.ParsingError, err, "Shift error while reading argument length").
					WithMeta(core.EMK(0, "int"), p.Config.GConfig.ArglenBytelen).
					WithMeta(core.EMK(1, "int"), *idx).
					WithMeta(core.EMK(2, "int"), argNum)
				return
			}
			arglen := u.BytesToInt(arglenBytes, p.Config.GConfig.Endianess)
			log(fmt.Sprintf("arglen %v", arglen))
			if arglen == 0 {
				err = core.Err(errors.ParsingError, "Zero argument length").
					WithMeta(core.EMK(0, "int"), argNum)
				return
			}
			var arg []byte
			arg, err = shift(arglen)
			if err != nil {
				err = core.Wrap(errors.ParsingError, err, "Shift error while reading argument data").
					WithMeta(core.EMK(0, "int"), arglen).
					WithMeta(core.EMK(1, "int"), *idx).
					WithMeta(core.EMK(2, "int"), argNum)
				return
			}
			log(fmt.Sprintf("arglen, args %v; %v", arglen, args))
			args = append(args, arg)
		}
		raw = code[idxStart:*idx]
		result = append(result, ParsedBytes{
			Switch:   command,
			Args:     args,
			Raw:      raw,
			Metadata: make(map[string]interface{}),
		})
	}
	p.Config.Shifter.Idx = _Idx
	log(fmt.Sprintf("end parsing code:\n %v", result))
	log(fmt.Sprintf("=========== END ==========="))
	return result, nil
}

func (p *Parser1) String() string {
	return "lc/parsing/byteParsing/Parser1"
}
```

---

# parsing/byteParsing/parser1_test.go

```go
package byteParsing

import (
	"reflect"
	"testing"

	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/bytecode"
)

func TestParser1_Parse(t *testing.T) {
	config := Parser1Config{
		GConfig: bytecode.GenerationConfig{
			CommandBytelen:   1,
			ArgscountBytelen: 1,
			ArglenBytelen:    1,
			Endianess:        public.LittleEndian,
		},
		Shifter: bytecode.Shift{},
	}
	parser := &Parser1{Config: config}

	code := []byte{0x01, 0x01, 0x03, 0x61, 0x62, 0x63}
	nodes, err := parser.Parse(code)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(nodes))
	}
	node := nodes[0]
	if !reflect.DeepEqual(node.Switch, []byte{0x01}) {
		t.Errorf("switch = %v, want [0x01]", node.Switch)
	}
	if len(node.Args) != 1 {
		t.Fatalf("expected 1 arg, got %d", len(node.Args))
	}
	if !reflect.DeepEqual(node.Args[0], []byte{0x61, 0x62, 0x63}) {
		t.Errorf("arg = %v, want [0x61 0x62 0x63]", node.Args[0])
	}
}

func TestParser1_Parse_Error(t *testing.T) {
	config := Parser1Config{
		GConfig: bytecode.GenerationConfig{
			CommandBytelen:   1,
			ArgscountBytelen: 1,
			ArglenBytelen:    1,
			Endianess:        public.LittleEndian,
		},
		Shifter: bytecode.Shift{},
	}
	parser := &Parser1{Config: config}
	code := []byte{0x01, 0x01, 0x03}
	_, err := parser.Parse(code)
	if err == nil {
		t.Error("expected error, got nil")
	}
}
```

---

# parsing/main.go

```go
package parsing

import "github.com/pt-main/lc/engine/core"

type ParseOption struct {
	UEP    *core.UniversalEngineParams
	Option core.Option
}

type ParserInterface[I any, P any] interface {
	Parse(I, ...*ParseOption) ([]P, core.ErrorInterface)
	String() string
}
```

---

# parsing/stringParsing/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# stringParsing

'''go
import "github.com/pt-main/lc/parsing/stringParsing"
'''

## Index

- [type GrammarRule](<#GrammarRule>)
- [type Lexer](<#Lexer>)
  - [func NewLexer\(rules \[\]LexerRule, config \*LexerConfig\) \*Lexer](<#NewLexer>)
  - [func \(lp \*Lexer\) Parse\(code string, opts ...\*parsing.ParseOption\) \(\[\]ParsedNode, core.ErrorInterface\)](<#Lexer.Parse>)
  - [func \(l \*Lexer\) String\(\) string](<#Lexer.String>)
- [type LexerConfig](<#LexerConfig>)
- [type LexerRule](<#LexerRule>)
- [type ParsedNode](<#ParsedNode>)
- [type Parser1](<#Parser1>)
  - [func NewParser1\(rules \[\]GrammarRule, config Parser1Config\) \*Parser1](<#NewParser1>)
  - [func \(p \*Parser1\) Parse\(code string, opts ...\*parsing.ParseOption\) \(\[\]ParsedNode, core.ErrorInterface\)](<#Parser1.Parse>)
  - [func \(p \*Parser1\) String\(\) string](<#Parser1.String>)
- [type Parser1Config](<#Parser1Config>)
- [type Parser2](<#Parser2>)
  - [func \(p \*Parser2\) Parse\(code string, opts ...\*parsing.ParseOption\) \(\[\]ParsedNode, core.ErrorInterface\)](<#Parser2.Parse>)
  - [func \(p \*Parser2\) String\(\) string](<#Parser2.String>)


<a name="GrammarRule"></a>
## type [GrammarRule](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser1.go#L13-L16>)



'''go
type GrammarRule struct {
    Type    string
    Pattern *regexp.Regexp
}
'''

<a name="Lexer"></a>
## type [Lexer](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/lexer.go#L29-L37>)

Lexer converts a source string into a sequence of ParsedNode objects.

'''go
type Lexer struct {
    // contains filtered or unexported fields
}
'''

<a name="NewLexer"></a>
### func [NewLexer](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/lexer.go#L40>)

'''go
func NewLexer(rules []LexerRule, config *LexerConfig) *Lexer
'''

NewLexer creates a lexer with the given rule set and optional configuration.

<a name="Lexer.Parse"></a>
### func \(\*Lexer\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/lexer.go#L248>)

'''go
func (lp *Lexer) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface)
'''

Parse scans the entire input string and returns a slice of ParsedNode.

Err errors.ParsingError:

- If bracket balancing is enabled and the input has unbalanced brackets. Meta: EMK\(0, "string"\) – the whole input code.
- If a regexp rule fails to match. Meta: EMK\(0, "string"\) – rule type, EMK\(1, "string"\) – substring being matched.
- If no rule matches at the current position. Meta: EMK\(0, "int"\) – line number, EMK\(1, "int"\) – column number, EMK\(2, "string"\) – context snippet.

<a name="Lexer.String"></a>
### func \(\*Lexer\) [String](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/lexer.go#L335>)

'''go
func (l *Lexer) String() string
'''



<a name="LexerConfig"></a>
## type [LexerConfig](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/lexer.go#L23-L26>)

LexerConfig holds configuration options for the lexer.

'''go
type LexerConfig struct {
    UseBracketBalance bool
    Brackets          [][2]string
}
'''

<a name="LexerRule"></a>
## type [LexerRule](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/lexer.go#L17-L20>)

LexerRule defines a single token type and its regular expression pattern.

'''go
type LexerRule struct {
    Type    string
    Pattern *regexp2.Regexp
}
'''

<a name="ParsedNode"></a>
## type [ParsedNode](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/node.go#L9-L18>)

ParsedNode represents a single token or syntactic unit in text mode. It contains the original raw string, the token type or command name \(Switch\), and a metadata map for named groups or additional attributes. After parsing, nodes are automatically enriched with \_\_prev and \_\_next links.

'''go
type ParsedNode struct {
    // Raw string – the exact substring matched (or the full line/block).
    Raw string
    // Switch string – the token type (e.g., "NUMBER", "IDENT") or command name.
    Switch string
    // Metadata core.ScopeType – holds regexp named groups, "__raw" (full
    //   original text, in all basic Parsers), "__value" (matched value, in basic Lexer),
    //   and optionally "__prev"/"__next" which point to neighboring ParsedNode (or nil).
    Metadata core.ScopeType
}
'''

<a name="Parser1"></a>
## type [Parser1](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser1.go#L28-L33>)

Parser1 implements a regex‑based grammar parser with line continuation and bracket balancing support.

'''go
type Parser1 struct {
    // contains filtered or unexported fields
}
'''

<a name="NewParser1"></a>
### func [NewParser1](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser1.go#L36>)

'''go
func NewParser1(rules []GrammarRule, config Parser1Config) *Parser1
'''

NewParser1 constructs a Parser1 with given grammar rules and configuration.

<a name="Parser1.Parse"></a>
### func \(\*Parser1\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser1.go#L60>)

'''go
func (p *Parser1) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface)
'''

Parse splits the input into logical blocks and applies grammar rules.

Err errors.ParsingError:

- If no rule matches a block. Meta: EMK\(0, "string"\) – the block that failed to match.

<a name="Parser1.String"></a>
### func \(\*Parser1\) [String](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser1.go#L173>)

'''go
func (p *Parser1) String() string
'''



<a name="Parser1Config"></a>
## type [Parser1Config](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser1.go#L18-L24>)



'''go
type Parser1Config struct {
    UseLineContinuation bool
    UseBracketBalance   bool
    SkipEmptyLines      bool
    Brackets            []string
    TrimBlocksSpace     bool
}
'''

<a name="Parser2"></a>
## type [Parser2](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser2.go#L12>)

Parser2 is a simple command‑args line parser.

'''go
type Parser2 struct{}
'''

<a name="Parser2.Parse"></a>
### func \(\*Parser2\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser2.go#L19>)

'''go
func (p *Parser2) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface)
'''

Parse converts each non‑empty line into a ParsedNode.

Err errors.ParsingError:

- If no valid lines are found in the input. Meta: EMK\(0, "string"\) – the whole input string.

<a name="Parser2.String"></a>
### func \(\*Parser2\) [String](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser2.go#L57>)

'''go
func (p *Parser2) String() string
'''



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# parsing/stringParsing/README.md

```md
# String Parsing

Package with string parsers, implements ParserInterface for engine: 

- 'Lexer': Simple lexer based on regexp2. Find tokens in full text. Supports multichar bracket balance.
- 'Parser1': Parser based on regexp. Find tokens in lines. Supports simple bracket balance, line continuation (with 'line \ \n...'), and block space trimming.
- 'Parser2': The simplest one parser parsing grammar like 'cmd args ...'.
- 'Parser3': The hardest one peg-like parser. See 'parsing/stringParsing/parser3' for more docs
```

---

# parsing/stringParsing/lexer.go

```go
package stringParsing

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

// LexerRule defines a single token type and its regular expression pattern.
type LexerRule struct {
	Type    string
	Pattern *regexp2.Regexp
}

// LexerConfig holds configuration options for the lexer.
type LexerConfig struct {
	UseBracketBalance bool
	Brackets          [][2]string
}

// Lexer converts a source string into a sequence of ParsedNode objects.
type Lexer struct {
	rules       []LexerRule
	config      LexerConfig
	openToClose map[string]string
	closeToOpen map[string]string
	ruleGroups  [][]string
	openByByte  map[byte][]string
	closeByByte map[byte][]string
}

// NewLexer creates a lexer with the given rule set and optional configuration.
func NewLexer(rules []LexerRule, config *LexerConfig) *Lexer {
	cfg := LexerConfig{}
	if config != nil {
		cfg = *config
	}

	openToClose := make(map[string]string)
	closeToOpen := make(map[string]string)
	openByByte := make(map[byte][]string)
	closeByByte := make(map[byte][]string)

	for _, pair := range cfg.Brackets {
		if len(pair) != 2 {
			continue
		}
		open, close := pair[0], pair[1]
		openToClose[open] = close
		closeToOpen[close] = open
		if len(open) > 0 {
			openByByte[open[0]] = append(openByByte[open[0]], open)
		}
		if len(close) > 0 {
			closeByByte[close[0]] = append(closeByByte[close[0]], close)
		}
	}

	for b := range openByByte {
		sort.Slice(openByByte[b], func(i, j int) bool {
			return len(openByByte[b][i]) > len(openByByte[b][j])
		})
	}
	for b := range closeByByte {
		sort.Slice(closeByByte[b], func(i, j int) bool {
			return len(closeByByte[b][i]) > len(closeByByte[b][j])
		})
	}

	ruleGroups := make([][]string, len(rules))
	for i, rule := range rules {
		ruleGroups[i] = rule.Pattern.GetGroupNames()
	}

	return &Lexer{
		rules:       rules,
		config:      cfg,
		openToClose: openToClose,
		closeToOpen: closeToOpen,
		ruleGroups:  ruleGroups,
		openByByte:  openByByte,
		closeByByte: closeByByte,
	}
}

func snippet(s string, maxRunes int) string {
	runes := []rune(s)
	if len(runes) <= maxRunes {
		return s
	}
	return string(runes[:maxRunes]) + "..."
}

func snippetFromRunes(runes []rune, pos, maxRunes int) string {
	end := pos + maxRunes
	if end >= len(runes) {
		return string(runes[pos:])
	}
	return string(runes[pos:end]) + "..."
}

func posToLineCol(code string, pos int) (line, col int) {
	line, col = 1, 1
	runeIdx := 0
	for _, r := range code {
		if runeIdx == pos {
			return
		}
		if r == '\n' {
			line++
			col = 1
		} else {
			col++
		}
		runeIdx++
	}
	return
}

// isBracketBalanced checks if the accumulated text has balanced brackets.
func (l *Lexer) isBracketBalanced(text string) bool {
	if !l.config.UseBracketBalance || len(l.openToClose) == 0 {
		return true
	}
	stack := make([]string, 0, 8)
	i := 0
	n := len(text)

	for i < n {
		matched := false
		if candidates, ok := l.openByByte[text[i]]; ok {
			for _, open := range candidates {
				if strings.HasPrefix(text[i:], open) {
					stack = append(stack, open)
					i += len(open)
					matched = true
					break
				}
			}
		}
		if matched {
			continue
		}
		if candidates, ok := l.closeByByte[text[i]]; ok {
			for _, close := range candidates {
				if strings.HasPrefix(text[i:], close) {
					if len(stack) == 0 {
						return false
					}
					last := stack[len(stack)-1]
					if open, ok := l.closeToOpen[close]; !ok || last != open {
						return false
					}
					stack = stack[:len(stack)-1]
					i += len(close)
					matched = true
					break
				}
			}
		}
		if matched {
			continue
		}
		_, size := utf8.DecodeRuneInString(text[i:])
		i += size
	}
	return len(stack) == 0
}

// bracketBalanceError returns a detailed error if brackets are unbalanced.
func (l *Lexer) bracketBalanceError(text string) *core.Error {
	if !l.config.UseBracketBalance || len(l.openToClose) == 0 {
		return nil
	}
	stack := make([]string, 0, 8)
	i := 0
	n := len(text)

	for i < n {
		matched := false
		if candidates, ok := l.openByByte[text[i]]; ok {
			for _, open := range candidates {
				if strings.HasPrefix(text[i:], open) {
					stack = append(stack, open)
					i += len(open)
					matched = true
					break
				}
			}
		}
		if matched {
			continue
		}
		if candidates, ok := l.closeByByte[text[i]]; ok {
			for _, close := range candidates {
				if strings.HasPrefix(text[i:], close) {
					if len(stack) == 0 {
						return core.Err(errors.ParsingError, "Unexpected closing bracket %q at byte %d", close, i).
							WithMeta(core.EMK(0, "string"), close).
							WithMeta(core.EMK(1, "int"), i)
					}
					last := stack[len(stack)-1]
					if open, ok := l.closeToOpen[close]; !ok || last != open {
						return core.Err(errors.ParsingError, "Mismatched bracket: expected closing %q for %q, got %q at byte %d",
							l.openToClose[last], last, close, i).
							WithMeta(core.EMK(0, "string"), close).
							WithMeta(core.EMK(1, "string"), last).
							WithMeta(core.EMK(2, "int"), i)
					}
					stack = stack[:len(stack)-1]
					i += len(close)
					matched = true
					break
				}
			}
		}
		if matched {
			continue
		}
		_, size := utf8.DecodeRuneInString(text[i:])
		i += size
	}
	if len(stack) > 0 {
		unclosed := stack[len(stack)-1]
		return core.Err(errors.ParsingError, "Unclosed bracket %q", unclosed).
			WithMeta(core.EMK(0, "string"), unclosed)
	}
	return nil
}

// Parse scans the entire input string and returns a slice of ParsedNode.
//
// Err errors.ParsingError:
//   - If bracket balancing is enabled and the input has unbalanced brackets.
//     Meta: EMK(0, "string") – the whole input code.
//   - If a regexp rule fails to match.
//     Meta: EMK(0, "string") – rule type, EMK(1, "string") – substring being matched.
//   - If no rule matches at the current position.
//     Meta: EMK(0, "int") – line number, EMK(1, "int") – column number,
//     EMK(2, "string") – context snippet.
func (lp *Lexer) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface) {
	var log func(string)
	if len(opts) > 0 && opts[0].UEP != nil && opts[0].UEP.Logger != nil {
		logger := opts[0].UEP.Logger
		log = func(text string) {
			logger.PrintLog(public.LogParsing, "\\n"+text)
		}
	} else {
		log = func(string) {}
	}
	log("start parsing code [" + code + "]")

	if lp.config.UseBracketBalance {
		if err := lp.bracketBalanceError(code); err != nil {
			return nil, core.Wrap(errors.ParsingError, err, "Bracket balance error").
				WithMeta(core.EMK(0, "string"), code)
		}
	}

	var nodes []ParsedNode
	runes := []rune(code)
	pos := 0
	length := len(runes)

	for pos < length {
		log(fmt.Sprintf("pos %v, length %v", pos, length))
		matched := false
		subStr := string(runes[pos:])

		for ruleIdx, rule := range lp.rules {
			m, err := rule.Pattern.FindStringMatch(subStr)
			if err != nil {
				return nil, core.Wrap(errors.ParsingError, err, "Regexp error for rule %q", rule.Type).
					WithMeta(core.EMK(0, "string"), rule.Type).
					WithMeta(core.EMK(1, "string"), subStr)
			}
			log(fmt.Sprintf("rule %v, pos %v, substrLen %v", rule, pos, len(subStr)))
			if m != nil && m.Index == 0 {
				tokenRunes := runes[pos : pos+m.Length]
				tokenValue := string(tokenRunes)
				startPos := pos
				endPos := pos + m.Length

				meta := map[string]interface{}{
					"__raw":   tokenValue,
					"__value": tokenValue,
					"__pos":   startPos,
					"__start": startPos,
					"__end":   endPos,
				}

				groupNames := lp.ruleGroups[ruleIdx]
				for _, name := range groupNames {
					if name != "0" {
						grp := m.GroupByName(name)
						if grp != nil {
							meta[name] = grp.String()
						}
					}
				}

				if lp.config.UseBracketBalance {
					meta["__bracket_balanced"] = lp.isBracketBalanced(tokenValue)
				}

				nodes = append(nodes, ParsedNode{
					Raw:      tokenValue,
					Switch:   rule.Type,
					Metadata: meta,
				})
				pos += m.Length
				matched = true
				break
			}
		}
		if !matched {
			line, col := posToLineCol(code, pos)
			context := snippetFromRunes(runes, pos, 20)
			return nil, core.Err(errors.ParsingError, "Unexpected sequence near %q at line %d, col %d", context, line, col).
				WithMeta(core.EMK(0, "int"), line).
				WithMeta(core.EMK(1, "int"), col).
				WithMeta(core.EMK(2, "string"), context)
		}
	}
	return addPrevNextNodes(nodes), nil
}

func (l *Lexer) String() string {
	return "lc/parsing/stringParsing/Lexer"
}
```

---

# parsing/stringParsing/lexer_test.go

```go
package stringParsing

import (
	"fmt"
	"testing"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
)

func TestLexer_Parse(t *testing.T) {
	rules := []LexerRule{
		{Type: "WHITESPACE", Pattern: regexp2.MustCompile('\s+', 0)},
		{Type: "BLOCK", Pattern: regexp2.MustCompile('(?s)\s*begin\{(.*)?\}end', 0)},
		{Type: "COMMENT", Pattern: regexp2.MustCompile('(?s)/\*\s*@(.+?)@\*/', 0)},
		{Type: "COMMENT", Pattern: regexp2.MustCompile('//@.*', 0)},
		{Type: "IDENT", Pattern: regexp2.MustCompile('[a-zA-Z_][a-zA-Z0-9_]+', 0)},
		{Type: "NUMBER", Pattern: regexp2.MustCompile('[0-9]+(?:\.[0-9]+)?', 0)},
		{Type: "STRING", Pattern: regexp2.MustCompile('"(?:[^"\\]|\\.)*"', 0)},
		{Type: "LBRACE", Pattern: regexp2.MustCompile('\{', 0)},
		{Type: "RBRACE", Pattern: regexp2.MustCompile('\}', 0)},
		{Type: "LPAREN", Pattern: regexp2.MustCompile('\(', 0)},
		{Type: "RPAREN", Pattern: regexp2.MustCompile('\)', 0)},
		{Type: "COMMA", Pattern: regexp2.MustCompile(',', 0)},
		{Type: "EQ", Pattern: regexp2.MustCompile('=', 0)},
	}

	lexer := NewLexer(rules, &LexerConfig{
		UseBracketBalance: true,
		Brackets:          [][2]string{{"begin{", "}end"}},
	})
	nodes, err := lexer.Parse('//@ test
1.0 test_param1 = "..." string "   "
begin{
	test block
}end',
		&parsing.ParseOption{UEP: &core.UniversalEngineParams{Logger: core.NewLogger("")}})
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	types := []string{
		"COMMENT", "WHITESPACE", "NUMBER", "WHITESPACE", "IDENT",
		"WHITESPACE", "EQ", "WHITESPACE", "STRING", "WHITESPACE", "IDENT",
		"WHITESPACE", "STRING", "WHITESPACE", "BLOCK",
	}
	if len(nodes) < len(types) {
		t.Fatalf("Too few tokens: got %d, want at least %d", len(nodes), len(types))
	}
	for idx, expected := range types {
		fmt.Printf("%v '%v' %v : %v\n", idx, nodes[idx].Raw, nodes[idx].Switch, expected)
		if nodes[idx].Switch != expected {
			t.Fatalf("Mismatch at index %d: expected %q, got %q", idx, expected, nodes[idx].Switch)
		}
	}
}
```

---

# parsing/stringParsing/node.go

```go
package stringParsing

import "github.com/pt-main/lc/engine/core"

// ParsedNode represents a single token or syntactic unit in text mode.
// It contains the original raw string, the token type or command name (Switch),
// and a metadata map for named groups or additional attributes.
// After parsing, nodes are automatically enriched with __prev and __next links.
type ParsedNode struct {
	// Raw string – the exact substring matched (or the full line/block).
	Raw string
	// Switch string – the token type (e.g., "NUMBER", "IDENT") or command name.
	Switch string
	// Metadata core.ScopeType – holds regexp named groups, "__raw" (full
	//   original text, in all basic Parsers), "__value" (matched value, in basic Lexer),
	//   and optionally "__prev"/"__next" which point to neighboring ParsedNode (or nil).
	Metadata core.ScopeType
}
```

---

# parsing/stringParsing/parser1.go

```go
package stringParsing

import (
	"regexp"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

type GrammarRule struct {
	Type    string
	Pattern *regexp.Regexp
}

type Parser1Config struct {
	UseLineContinuation bool
	UseBracketBalance   bool
	SkipEmptyLines      bool
	Brackets            []string
	TrimBlocksSpace     bool
}

// Parser1 implements a regex‑based grammar parser with line continuation
// and bracket balancing support.
type Parser1 struct {
	grammar     []GrammarRule
	config      Parser1Config
	openToClose map[rune]rune
	closeToOpen map[rune]rune
}

// NewParser1 constructs a Parser1 with given grammar rules and configuration.
func NewParser1(rules []GrammarRule, config Parser1Config) *Parser1 {
	openToClose := make(map[rune]rune)
	closeToOpen := make(map[rune]rune)
	for _, pair := range config.Brackets {
		runes := []rune(pair)
		if len(runes) == 2 {
			open, close := runes[0], runes[1]
			openToClose[open] = close
			closeToOpen[close] = open
		}
	}
	return &Parser1{
		grammar:     rules,
		config:      config,
		openToClose: openToClose,
		closeToOpen: closeToOpen,
	}
}

// Parse splits the input into logical blocks and applies grammar rules.
//
// Err errors.ParsingError:
//   - If no rule matches a block.
//     Meta: EMK(0, "string") – the block that failed to match.
func (p *Parser1) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface) {
	log := func(text string) {
		text = "\n" + text
		if len(opts) > 0 {
			logger := opts[0].UEP.Logger
			if logger != nil {
				logger.PrintLog(public.LogParsing, text)
			}
		}
	}
	log("start parsing code " + code)
	lines := strings.Split(code, "\n")
	var result []ParsedNode

	var blockLines []string
	bracketStack := []rune{}

	flush := func(block string) core.ErrorInterface {
		node, err := p.matchGrammar(block)
		if err != nil {
			return core.Wrap(errors.ParsingError, err, "Failed to match grammar for block").
				WithMeta(core.EMK(0, "string"), block)
		}
		if node.Raw != "" {
			result = append(result, node)
		}
		return nil
	}

	for _, rawLine := range lines {
		line := strings.TrimRight(rawLine, " \t")

		if line == "" && len(bracketStack) == 0 && !p.config.UseBracketBalance {
			continue
		}

		continues := p.config.UseLineContinuation && strings.HasSuffix(line, "\\")
		if continues {
			line = strings.TrimSuffix(line, "\\")
		}

		blockLines = append(blockLines, line)

		if p.config.UseBracketBalance {
			for _, ch := range line {
				if _, ok := p.openToClose[ch]; ok {
					bracketStack = append(bracketStack, ch)
				} else if closeOpen, ok := p.closeToOpen[ch]; ok {
					if len(bracketStack) > 0 && bracketStack[len(bracketStack)-1] == closeOpen {
						bracketStack = bracketStack[:len(bracketStack)-1]
					}
				}
			}
		}

		blockComplete := true
		if continues {
			blockComplete = false
		}
		if p.config.UseBracketBalance && len(bracketStack) > 0 {
			blockComplete = false
		}

		if blockComplete {
			block := strings.Join(blockLines, "\n")
			if err := flush(block); err != nil {
				return nil, err
			}
			blockLines = nil
			bracketStack = []rune{}
		}
	}

	if len(blockLines) > 0 {
		block := strings.Join(blockLines, "\n")
		if err := flush(block); err != nil {
			return nil, err
		}
	}

	return addPrevNextNodes(result), nil
}

func (p *Parser1) matchGrammar(block string) (ParsedNode, core.ErrorInterface) {
	absolutely_raw := block
	if p.config.TrimBlocksSpace {
		block = strings.TrimSpace(block)
	}
	if block == "" && p.config.SkipEmptyLines {
		return ParsedNode{Raw: ""}, nil
	}
	for _, rule := range p.grammar {
		if rule.Pattern.MatchString(block) {
			meta := make(map[string]interface{})
			matches := rule.Pattern.FindStringSubmatch(block)
			names := rule.Pattern.SubexpNames()
			for i, name := range names {
				if i != 0 && name != "" && i < len(matches) {
					meta[name] = matches[i]
				}
			}
			meta["__raw"] = absolutely_raw
			return ParsedNode{
				Raw:      block,
				Switch:   rule.Type,
				Metadata: meta,
			}, nil
		}
	}
	return ParsedNode{}, core.Err(errors.ParsingError, "Syntax error: no rule matches block: %q", block).
		WithMeta(core.EMK(0, "string"), block)
}

func (p *Parser1) String() string {
	return "lc/parsing/stringParsing/Parser1"
}
```

---

# parsing/stringParsing/parser2.go

```go
package stringParsing

import (
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/public/errors"
)

// Parser2 is a simple command‑args line parser.
type Parser2 struct{}

// Parse converts each non‑empty line into a ParsedNode.
//
// Err errors.ParsingError:
//   - If no valid lines are found in the input.
//     Meta: EMK(0, "string") – the whole input string.
func (p *Parser2) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface) {
	lines := strings.Split(code, "\n")
	result := []ParsedNode{}

	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		command := parts[0]
		args := ""
		if len(parts) > 1 {
			args = parts[1]
		}

		meta := map[string]interface{}{
			"command": command,
			"args":    args,
			"__raw":   rawLine,
		}

		node := ParsedNode{
			Raw:      rawLine,
			Switch:   command,
			Metadata: meta,
		}
		result = append(result, node)
	}

	if len(result) == 0 {
		return nil, core.Err(errors.ParsingError, "No valid lines found in input").
			WithMeta(core.EMK(0, "string"), code)
	}
	return addPrevNextNodes(result), nil
}

func (p *Parser2) String() string {
	return "lc/parsing/stringParsing/Parser2"
}
```

---

# parsing/stringParsing/parser3/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# parser3

'''go
import "github.com/pt-main/lc/parsing/stringParsing/parser3"
'''

## Index

- [Constants](<#constants>)
- [func FormatError\(err error, useColors bool\) string](<#FormatError>)
- [func FormatErrorPretty\(err error\) string](<#FormatErrorPretty>)
- [type ActionExpr](<#ActionExpr>)
  - [func \(a ActionExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#ActionExpr.Parse>)
- [type Adapter](<#Adapter>)
  - [func \(a \*Adapter\) Parse\(code string, o ...\*parsing.ParseOption\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#Adapter.Parse>)
  - [func \(a \*Adapter\) String\(\) string](<#Adapter.String>)
- [type AdapterError](<#AdapterError>)
  - [func \(e \*AdapterError\) Error\(\) string](<#AdapterError.Error>)
  - [func \(e \*AdapterError\) Format\(\) string](<#AdapterError.Format>)
  - [func \(e \*AdapterError\) GetCode\(\) string](<#AdapterError.GetCode>)
  - [func \(e \*AdapterError\) GetMeta\(\) map\[errors.ErrorMetaType\]interface\{\}](<#AdapterError.GetMeta>)
  - [func \(e \*AdapterError\) GetMsg\(\) string](<#AdapterError.GetMsg>)
  - [func \(e \*AdapterError\) Unwrap\(\) error](<#AdapterError.Unwrap>)
- [type AndExpr](<#AndExpr>)
  - [func \(a AndExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#AndExpr.Parse>)
- [type Associativity](<#Associativity>)
- [type ChoiceExpr](<#ChoiceExpr>)
  - [func \(c ChoiceExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#ChoiceExpr.Parse>)
- [type Expr](<#Expr>)
- [type Grammar](<#Grammar>)
- [type GrammarError](<#GrammarError>)
  - [func \(e \*GrammarError\) Error\(\) string](<#GrammarError.Error>)
  - [func \(e \*GrammarError\) Format\(\) string](<#GrammarError.Format>)
  - [func \(e \*GrammarError\) GetCode\(\) string](<#GrammarError.GetCode>)
  - [func \(e \*GrammarError\) GetMeta\(\) map\[errors.ErrorMetaType\]interface\{\}](<#GrammarError.GetMeta>)
  - [func \(e \*GrammarError\) GetMsg\(\) string](<#GrammarError.GetMsg>)
  - [func \(e \*GrammarError\) Unwrap\(\) error](<#GrammarError.Unwrap>)
- [type InfixInfo](<#InfixInfo>)
- [type NamedExpr](<#NamedExpr>)
  - [func \(n NamedExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#NamedExpr.Parse>)
- [type NodeExpr](<#NodeExpr>)
  - [func \(n NodeExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#NodeExpr.Parse>)
- [type NotExpr](<#NotExpr>)
  - [func \(n NotExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#NotExpr.Parse>)
- [type OptionalExpr](<#OptionalExpr>)
  - [func \(o OptionalExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#OptionalExpr.Parse>)
- [type ParseError](<#ParseError>)
  - [func \(e \*ParseError\) Error\(\) string](<#ParseError.Error>)
  - [func \(e \*ParseError\) Format\(\) string](<#ParseError.Format>)
  - [func \(e \*ParseError\) GetCode\(\) string](<#ParseError.GetCode>)
  - [func \(e \*ParseError\) GetMeta\(\) map\[errors.ErrorMetaType\]interface\{\}](<#ParseError.GetMeta>)
  - [func \(e \*ParseError\) GetMsg\(\) string](<#ParseError.GetMsg>)
  - [func \(e \*ParseError\) Unwrap\(\) error](<#ParseError.Unwrap>)
- [type Parser](<#Parser>)
  - [func NewParser\(lexer \*stringParsing.Lexer, grammar Grammar, startRule string, ignoreTypes \[\]string\) \*Parser](<#NewParser>)
  - [func \(p \*Parser\) Expect\(tokenType string\) \(stringParsing.ParsedNode, core.ErrorInterface\)](<#Parser.Expect>)
  - [func \(p \*Parser\) NextToken\(\) \(stringParsing.ParsedNode, error\)](<#Parser.NextToken>)
  - [func \(p \*Parser\) Parse\(code string, opts ...\*parsing.ParseOption\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#Parser.Parse>)
  - [func \(p \*Parser\) Peek\(\) \(stringParsing.ParsedNode, error\)](<#Parser.Peek>)
  - [func \(p \*Parser\) SkipIgnored\(\)](<#Parser.SkipIgnored>)
  - [func \(p \*Parser\) String\(\) string](<#Parser.String>)
- [type PeekExpr](<#PeekExpr>)
  - [func \(p PeekExpr\) Parse\(prs \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#PeekExpr.Parse>)
- [type PrattExpr](<#PrattExpr>)
  - [func \(p \*PrattExpr\) Parse\(prs \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#PrattExpr.Parse>)
- [type RepeatExpr](<#RepeatExpr>)
  - [func \(r RepeatExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#RepeatExpr.Parse>)
- [type Rule](<#Rule>)
- [type SeparatedRepeatExpr](<#SeparatedRepeatExpr>)
  - [func \(s SeparatedRepeatExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#SeparatedRepeatExpr.Parse>)
- [type SequenceExpr](<#SequenceExpr>)
  - [func \(s SequenceExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#SequenceExpr.Parse>)
- [type TokenExpr](<#TokenExpr>)
  - [func \(t TokenExpr\) Parse\(p \*Parser\) \(\[\]stringParsing.ParsedNode, core.ErrorInterface\)](<#TokenExpr.Parse>)


## Constants

<a name="AdapterErrCode"></a>

'''go
const AdapterErrCode = "parser3/adapter"
'''

<a name="GrammarErrCode"></a>

'''go
const GrammarErrCode = "parser3/grammar"
'''

<a name="ParseErrCode"></a>

'''go
const ParseErrCode = "parser3"
'''

<a name="FormatError"></a>
## func [FormatError](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/formatter.go#L13>)

'''go
func FormatError(err error, useColors bool) string
'''

FormatError renders a parser3 error with optional ANSI color codes. Pass useColors=false for plain\-text logs, JSON APIs, or file output.

<a name="FormatErrorPretty"></a>
## func [FormatErrorPretty](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/formatter.go#L40>)

'''go
func FormatErrorPretty(err error) string
'''

FormatErrorPretty is a convenience wrapper for CLI output \(colors enabled\).

<a name="ActionExpr"></a>
## type [ActionExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L172-L175>)



'''go
type ActionExpr struct {
    Expr   Expr
    Action func([]stringParsing.ParsedNode) (stringParsing.ParsedNode, core.ErrorInterface)
}
'''

<a name="ActionExpr.Parse"></a>
### func \(ActionExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L177>)

'''go
func (a ActionExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="Adapter"></a>
## type [Adapter](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/engineAdapter.go#L12-L14>)



'''go
type Adapter struct {
    Parser *Parser
}
'''

<a name="Adapter.Parse"></a>
### func \(\*Adapter\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/engineAdapter.go#L16>)

'''go
func (a *Adapter) Parse(code string, o ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="Adapter.String"></a>
### func \(\*Adapter\) [String](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/engineAdapter.go#L63>)

'''go
func (a *Adapter) String() string
'''



<a name="AdapterError"></a>
## type [AdapterError](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L150-L153>)

AdapterError is raised by the engine adapter when the AST shape is wrong.

'''go
type AdapterError struct {
    Msg   string
    Cause error
}
'''

<a name="AdapterError.Error"></a>
### func \(\*AdapterError\) [Error](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L157>)

'''go
func (e *AdapterError) Error() string
'''



<a name="AdapterError.Format"></a>
### func \(\*AdapterError\) [Format](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L166>)

'''go
func (e *AdapterError) Format() string
'''



<a name="AdapterError.GetCode"></a>
### func \(\*AdapterError\) [GetCode](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L170>)

'''go
func (e *AdapterError) GetCode() string
'''



<a name="AdapterError.GetMeta"></a>
### func \(\*AdapterError\) [GetMeta](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L178>)

'''go
func (e *AdapterError) GetMeta() map[errors.ErrorMetaType]interface{}
'''



<a name="AdapterError.GetMsg"></a>
### func \(\*AdapterError\) [GetMsg](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L174>)

'''go
func (e *AdapterError) GetMsg() string
'''



<a name="AdapterError.Unwrap"></a>
### func \(\*AdapterError\) [Unwrap](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L164>)

'''go
func (e *AdapterError) Unwrap() error
'''



<a name="AndExpr"></a>
## type [AndExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L214-L216>)



'''go
type AndExpr struct {
    Expr Expr
}
'''

<a name="AndExpr.Parse"></a>
### func \(AndExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L218>)

'''go
func (a AndExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="Associativity"></a>
## type [Associativity](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L306>)



'''go
type Associativity int
'''

<a name="LeftAssoc"></a>

'''go
const (
    LeftAssoc Associativity = iota
    RightAssoc
    NonAssoc
)
'''

<a name="ChoiceExpr"></a>
## type [ChoiceExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L58-L60>)



'''go
type ChoiceExpr struct {
    Alternatives []Expr
}
'''

<a name="ChoiceExpr.Parse"></a>
### func \(ChoiceExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L62>)

'''go
func (c ChoiceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="Expr"></a>
## type [Expr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L11-L13>)



'''go
type Expr interface {
    Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
}
'''

<a name="Grammar"></a>
## type [Grammar](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L20>)



'''go
type Grammar map[string]Rule
'''

<a name="GrammarError"></a>
## type [GrammarError](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L107-L111>)

GrammarError is raised when the grammar itself is misconfigured \(undefined rule, missing start rule, etc.\).

'''go
type GrammarError struct {
    Code  string // e.g. "NamedExpr", "ChoiceExpr"
    Msg   string // human-readable description
    Cause error
}
'''

<a name="GrammarError.Error"></a>
### func \(\*GrammarError\) [Error](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L115>)

'''go
func (e *GrammarError) Error() string
'''



<a name="GrammarError.Format"></a>
### func \(\*GrammarError\) [Format](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L133>)

'''go
func (e *GrammarError) Format() string
'''



<a name="GrammarError.GetCode"></a>
### func \(\*GrammarError\) [GetCode](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L145>)

'''go
func (e *GrammarError) GetCode() string
'''



<a name="GrammarError.GetMeta"></a>
### func \(\*GrammarError\) [GetMeta](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L141>)

'''go
func (e *GrammarError) GetMeta() map[errors.ErrorMetaType]interface{}
'''



<a name="GrammarError.GetMsg"></a>
### func \(\*GrammarError\) [GetMsg](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L137>)

'''go
func (e *GrammarError) GetMsg() string
'''



<a name="GrammarError.Unwrap"></a>
### func \(\*GrammarError\) [Unwrap](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L131>)

'''go
func (e *GrammarError) Unwrap() error
'''



<a name="InfixInfo"></a>
## type [InfixInfo](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L314-L317>)



'''go
type InfixInfo struct {
    Precedence int
    Assoc      Associativity
}
'''

<a name="NamedExpr"></a>
## type [NamedExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L125-L127>)



'''go
type NamedExpr struct {
    RuleName string
}
'''

<a name="NamedExpr.Parse"></a>
### func \(NamedExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L129>)

'''go
func (n NamedExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="NodeExpr"></a>
## type [NodeExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L144-L147>)



'''go
type NodeExpr struct {
    NodeType string
    Expr     Expr
}
'''

<a name="NodeExpr.Parse"></a>
### func \(NodeExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L149>)

'''go
func (n NodeExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="NotExpr"></a>
## type [NotExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L197-L199>)



'''go
type NotExpr struct {
    Expr Expr
}
'''

<a name="NotExpr.Parse"></a>
### func \(NotExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L201>)

'''go
func (n NotExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="OptionalExpr"></a>
## type [OptionalExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L111-L113>)



'''go
type OptionalExpr struct {
    Expr Expr
}
'''

<a name="OptionalExpr.Parse"></a>
### func \(OptionalExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L115>)

'''go
func (o OptionalExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="ParseError"></a>
## type [ParseError](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L13-L27>)

ParseError is a structured parsing error. It carries enough context for both plain logging and rich CLI output.

'''go
type ParseError struct {
    // Where it happened.
    TokenIdx int    // index in the token stream
    TokenPos string // human-readable position, e.g. "line 3, col 5" or "start=9-11"

    // What was expected vs what we got.
    Code     string // operation name: "Expect", "Peek", "ChoiceExpr", "PrattExpr", etc.
    Expected string // expected token type / rule / value
    Got      string // actual token type / value
    Raw      string // raw text of the offending token
    Msg      string // free-form message (used when Expected/Got don't fit)

    // Cause holds the underlying error (lexer error, user action error, etc.)
    Cause error
}
'''

<a name="ParseError.Error"></a>
### func \(\*ParseError\) [Error](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L31>)

'''go
func (e *ParseError) Error() string
'''



<a name="ParseError.Format"></a>
### func \(\*ParseError\) [Format](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L82>)

'''go
func (e *ParseError) Format() string
'''



<a name="ParseError.GetCode"></a>
### func \(\*ParseError\) [GetCode](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L86>)

'''go
func (e *ParseError) GetCode() string
'''



<a name="ParseError.GetMeta"></a>
### func \(\*ParseError\) [GetMeta](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L94>)

'''go
func (e *ParseError) GetMeta() map[errors.ErrorMetaType]interface{}
'''



<a name="ParseError.GetMsg"></a>
### func \(\*ParseError\) [GetMsg](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L90>)

'''go
func (e *ParseError) GetMsg() string
'''



<a name="ParseError.Unwrap"></a>
### func \(\*ParseError\) [Unwrap](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/errors.go#L80>)

'''go
func (e *ParseError) Unwrap() error
'''

Unwrap returns the underlying error for errors.Is / errors.As.

<a name="Parser"></a>
## type [Parser](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/parser.go#L28-L36>)



'''go
type Parser struct {
    // contains filtered or unexported fields
}
'''

<a name="NewParser"></a>
### func [NewParser](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/parser.go#L38>)

'''go
func NewParser(lexer *stringParsing.Lexer, grammar Grammar, startRule string, ignoreTypes []string) *Parser
'''



<a name="Parser.Expect"></a>
### func \(\*Parser\) [Expect](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/parser.go#L136>)

'''go
func (p *Parser) Expect(tokenType string) (stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="Parser.NextToken"></a>
### func \(\*Parser\) [NextToken](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/parser.go#L121>)

'''go
func (p *Parser) NextToken() (stringParsing.ParsedNode, error)
'''



<a name="Parser.Parse"></a>
### func \(\*Parser\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/parser.go#L51>)

'''go
func (p *Parser) Parse(code string, opts ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="Parser.Peek"></a>
### func \(\*Parser\) [Peek](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/parser.go#L160>)

'''go
func (p *Parser) Peek() (stringParsing.ParsedNode, error)
'''



<a name="Parser.SkipIgnored"></a>
### func \(\*Parser\) [SkipIgnored](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/parser.go#L115>)

'''go
func (p *Parser) SkipIgnored()
'''



<a name="Parser.String"></a>
### func \(\*Parser\) [String](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/parser.go#L173>)

'''go
func (p *Parser) String() string
'''



<a name="PeekExpr"></a>
## type [PeekExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L232-L234>)



'''go
type PeekExpr struct {
    TokenType string
}
'''

<a name="PeekExpr.Parse"></a>
### func \(PeekExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L236>)

'''go
func (p PeekExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="PrattExpr"></a>
## type [PrattExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L319-L323>)



'''go
type PrattExpr struct {
    Atom     Expr
    Prefixes map[string]Expr
    Infixes  map[string]InfixInfo
}
'''

<a name="PrattExpr.Parse"></a>
### func \(\*PrattExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L327>)

'''go
func (p *PrattExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="RepeatExpr"></a>
## type [RepeatExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L81-L84>)



'''go
type RepeatExpr struct {
    Expr Expr
    Min  int
}
'''

<a name="RepeatExpr.Parse"></a>
### func \(RepeatExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L86>)

'''go
func (r RepeatExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="Rule"></a>
## type [Rule](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L15-L18>)



'''go
type Rule struct {
    Name string
    Expr Expr
}
'''

<a name="SeparatedRepeatExpr"></a>
## type [SeparatedRepeatExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L254-L259>)



'''go
type SeparatedRepeatExpr struct {
    Element Expr
    Sep     string
    Min     int
    Max     int
}
'''

<a name="SeparatedRepeatExpr.Parse"></a>
### func \(SeparatedRepeatExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L261>)

'''go
func (s SeparatedRepeatExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="SequenceExpr"></a>
## type [SequenceExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L38-L40>)



'''go
type SequenceExpr struct {
    Exprs []Expr
}
'''

<a name="SequenceExpr.Parse"></a>
### func \(SequenceExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L42>)

'''go
func (s SequenceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



<a name="TokenExpr"></a>
## type [TokenExpr](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L22-L24>)



'''go
type TokenExpr struct {
    TokenType string
}
'''

<a name="TokenExpr.Parse"></a>
### func \(TokenExpr\) [Parse](<https://github.com/pt-main/Lc/blob/main/parsing/stringParsing/parser3/grammar.go#L26>)

'''go
func (t TokenExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
'''



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# parsing/stringParsing/parser3/engineAdapter.go

```go
package parser3

import (
	"fmt"
	"reflect"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

type Adapter struct {
	Parser *Parser
}

func (a *Adapter) Parse(code string, o ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	nodes, err := a.Parser.Parse(code, o...)
	if err != nil {
		return nil, &AdapterError{
			Msg:   "adapter parse failed",
			Cause: err,
		}
	}
	if len(nodes) != 1 {
		return nil, &AdapterError{
			Msg: fmt.Sprintf("expected exactly 1 root node, got %d", len(nodes)),
		}
	}

	children, ok := nodes[0].Metadata["children"].([]stringParsing.ParsedNode)
	if !ok {
		raw, has := nodes[0].Metadata["children"]
		if !has {
			keys := make([]string, 0, len(nodes[0].Metadata))
			for k := range nodes[0].Metadata {
				keys = append(keys, k)
			}
			return nil, &AdapterError{
				Msg: fmt.Sprintf("root node '%s' has no 'children' metadata (available keys: %v)", nodes[0].Switch, keys),
			}
		}
		rv := reflect.ValueOf(raw)
		if rv.Kind() != reflect.Slice {
			return nil, &AdapterError{
				Msg: fmt.Sprintf("root node '%s' has 'children' metadata of type %T (expected []ParsedNode)", nodes[0].Switch, raw),
			}
		}
		children = make([]stringParsing.ParsedNode, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			elem := rv.Index(i).Interface()
			pn, ok := elem.(stringParsing.ParsedNode)
			if !ok {
				return nil, &AdapterError{
					Msg: fmt.Sprintf("root node '%s' has 'children' metadata with wrong element type %T at index %d (expected ParsedNode)", nodes[0].Switch, elem, i),
				}
			}
			children[i] = pn
		}
	}
	return children, nil
}

func (a *Adapter) String() string {
	return "lc/parsing/stringParsing/parser3/Adapter"
}
```

---

# parsing/stringParsing/parser3/errors.go

```go
package parser3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pt-main/lc/public/errors"
)

// ParseError is a structured parsing error.
// It carries enough context for both plain logging and rich CLI output.
type ParseError struct {
	// Where it happened.
	TokenIdx int    // index in the token stream
	TokenPos string // human-readable position, e.g. "line 3, col 5" or "start=9-11"

	// What was expected vs what we got.
	Code     string // operation name: "Expect", "Peek", "ChoiceExpr", "PrattExpr", etc.
	Expected string // expected token type / rule / value
	Got      string // actual token type / value
	Raw      string // raw text of the offending token
	Msg      string // free-form message (used when Expected/Got don't fit)

	// Cause holds the underlying error (lexer error, user action error, etc.)
	Cause error
}

const ParseErrCode = "parser3"

func (e *ParseError) Error() string {
	var b strings.Builder
	b.WriteString(ParseErrCode)
	if e.Code != "" {
		b.WriteString("/")
		b.WriteString(e.Code)
	}
	b.WriteString(": ")

	parts := 0
	if e.Expected != "" {
		b.WriteString("expected ")
		b.WriteString(strconv.Quote(e.Expected))
		parts++
	}
	if e.Got != "" {
		if parts > 0 {
			b.WriteString(", got ")
		} else {
			b.WriteString("got ")
		}
		b.WriteString(strconv.Quote(e.Got))
		parts++
	}
	if e.Raw != "" && e.Raw != e.Got {
		b.WriteString(fmt.Sprintf(" (raw: %s)", strconv.Quote(e.Raw)))
	}
	if e.Msg != "" {
		if parts > 0 {
			b.WriteString(" — ")
		}
		b.WriteString(e.Msg)
		parts++
	}
	if e.TokenPos != "" {
		b.WriteString(" at ")
		b.WriteString(e.TokenPos)
	}
	if e.Cause != nil {
		b.WriteString(": ")
		b.WriteString(e.Cause.Error())
	}
	if parts == 0 && e.Cause == nil {
		b.WriteString("parse error")
	}
	return b.String()
}

// Unwrap returns the underlying error for errors.Is / errors.As.
func (e *ParseError) Unwrap() error { return e.Cause }

func (e *ParseError) Format() string {
	return FormatError(e, false)
}

func (e *ParseError) GetCode() string {
	return ParseErrCode
}

func (e *ParseError) GetMsg() string {
	return e.Msg
}

func (e *ParseError) GetMeta() map[errors.ErrorMetaType]interface{} {
	return map[errors.ErrorMetaType]interface{}{
		"TokenIdx": e.TokenIdx,
		"TokenPos": e.TokenPos,
		"Code":     e.Code,
		"Expected": e.Expected,
		"Raw":      e.Raw,
		"Got":      e.Got,
	}
}

// GrammarError is raised when the grammar itself is misconfigured
// (undefined rule, missing start rule, etc.).
type GrammarError struct {
	Code  string // e.g. "NamedExpr", "ChoiceExpr"
	Msg   string // human-readable description
	Cause error
}

const GrammarErrCode = "parser3/grammar"

func (e *GrammarError) Error() string {
	var b strings.Builder
	b.WriteString(GrammarErrCode)
	if e.Code != "" {
		b.WriteString("/")
		b.WriteString(e.Code)
	}
	b.WriteString(": ")
	b.WriteString(e.Msg)
	if e.Cause != nil {
		b.WriteString(": ")
		b.WriteString(e.Cause.Error())
	}
	return b.String()
}

func (e *GrammarError) Unwrap() error { return e.Cause }

func (e *GrammarError) Format() string {
	return FormatError(e, false)
}

func (e *GrammarError) GetMsg() string {
	return e.Msg
}

func (e *GrammarError) GetMeta() map[errors.ErrorMetaType]interface{} {
	return nil
}

func (e *GrammarError) GetCode() string {
	return e.Code
}

// AdapterError is raised by the engine adapter when the AST shape is wrong.
type AdapterError struct {
	Msg   string
	Cause error
}

const AdapterErrCode = "parser3/adapter"

func (e *AdapterError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("parser3/adapter: %s: %v", e.Msg, e.Cause)
	}
	return AdapterErrCode + ": " + e.Msg
}

func (e *AdapterError) Unwrap() error { return e.Cause }

func (e *AdapterError) Format() string {
	return FormatError(e, false)
}

func (e *AdapterError) GetCode() string {
	return string(AdapterErrCode)
}

func (e *AdapterError) GetMsg() string {
	return e.Msg
}

func (e *AdapterError) GetMeta() map[errors.ErrorMetaType]interface{} {
	return nil
}
```

---

# parsing/stringParsing/parser3/formatter.go

```go
package parser3

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pt-main/tap/color"
)

// FormatError renders a parser3 error with optional ANSI color codes.
// Pass useColors=false for plain-text logs, JSON APIs, or file output.
func FormatError(err error, useColors bool) string {
	if err == nil {
		return ""
	}

	// Build a chain of structured errors.
	var parts []string
	walk(err, &parts)

	if !useColors {
		return strings.Join(parts, " -> ")
	}

	// Colorize each part.
	colored := make([]string, len(parts))
	for i, p := range parts {
		if i == len(parts)-1 {
			// The root cause — highlight in red.
			colored[i] = color.Set(fmt.Sprintf("[?RD]%s[?RT]", p))
		} else {
			colored[i] = color.Set(fmt.Sprintf("[?YW]%s[?RT]", p))
		}
	}
	return strings.Join(colored, " -> ")
}

// FormatErrorPretty is a convenience wrapper for CLI output (colors enabled).
func FormatErrorPretty(err error) string {
	return FormatError(err, true)
}

// walk unwraps the error chain and collects human-readable messages.
func walk(err error, out *[]string) {
	if err == nil {
		return
	}

	// Try structured types first for rich output.
	var pe *ParseError
	if errors.As(err, &pe) {
		*out = append(*out, formatParseError(pe))
		walk(pe.Cause, out)
		return
	}

	var ge *GrammarError
	if errors.As(err, &ge) {
		*out = append(*out, formatGrammarError(ge))
		walk(ge.Cause, out)
		return
	}

	var ae *AdapterError
	if errors.As(err, &ae) {
		*out = append(*out, formatAdapterError(ae))
		walk(ae.Cause, out)
		return
	}

	// Fallback for plain errors.
	*out = append(*out, err.Error())
}

func formatParseError(e *ParseError) string {
	var b strings.Builder
	b.WriteString("parser3")
	if e.Code != "" {
		b.WriteString("/")
		b.WriteString(e.Code)
	}
	b.WriteString(": ")
	if e.Expected != "" {
		b.WriteString(fmt.Sprintf("expected '%s'", e.Expected))
		if e.Got != "" {
			b.WriteString(fmt.Sprintf(", got '%s'", e.Got))
		}
	} else if e.Got != "" {
		b.WriteString(fmt.Sprintf("got '%s'", e.Got))
	}
	if e.Raw != "" && e.Raw != e.Got {
		b.WriteString(fmt.Sprintf(" (raw: %q)", e.Raw))
	}
	if e.TokenPos != "" {
		b.WriteString(fmt.Sprintf(" at %s", e.TokenPos))
	}
	return b.String()
}

func formatGrammarError(e *GrammarError) string {
	var b strings.Builder
	b.WriteString("grammar")
	if e.Code != "" {
		b.WriteString("/")
		b.WriteString(e.Code)
	}
	b.WriteString(": ")
	b.WriteString(e.Msg)
	return b.String()
}

func formatAdapterError(e *AdapterError) string {
	return "adapter: " + e.Msg
}
```

---

# parsing/stringParsing/parser3/grammar.go

```go
package parser3

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
)

type Expr interface {
	Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
}

type Rule struct {
	Name string
	Expr Expr
}

type Grammar map[string]Rule

type TokenExpr struct {
	TokenType string
}

func (t TokenExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	tok, err := p.Expect(t.TokenType)
	if err != nil {
		return nil, &GrammarError{
			Code:  "TokenExpr",
			Msg:   fmt.Sprintf("expected token '%s'", t.TokenType),
			Cause: err,
		}
	}
	return []stringParsing.ParsedNode{tok}, nil
}

type SequenceExpr struct {
	Exprs []Expr
}

func (s SequenceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	var children []stringParsing.ParsedNode
	for i, e := range s.Exprs {
		nodes, err := e.Parse(p)
		if err != nil {
			return nil, &GrammarError{
				Code:  "SequenceExpr",
				Msg:   fmt.Sprintf("element %d/%d failed", i+1, len(s.Exprs)),
				Cause: err,
			}
		}
		children = append(children, nodes...)
	}
	return children, nil
}

type ChoiceExpr struct {
	Alternatives []Expr
}

func (c ChoiceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	for _, alt := range c.Alternatives {
		nodes, err := alt.Parse(p)
		if err == nil {
			return nodes, nil
		}
		p.pos = savedPos
	}
	altNames := make([]string, len(c.Alternatives))
	for i, alt := range c.Alternatives {
		altNames[i] = fmt.Sprintf("%T", alt)
	}
	return nil, &GrammarError{
		Code: "ChoiceExpr",
		Msg:  fmt.Sprintf("no alternative matched at %s (%d alternatives tried: %v)", tokenPos(p.tokens, p.pos), len(c.Alternatives), altNames),
	}
}

type RepeatExpr struct {
	Expr Expr
	Min  int
}

func (r RepeatExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	var all []stringParsing.ParsedNode
	for {
		savedPos := p.pos
		nodes, err := r.Expr.Parse(p)
		if err != nil {
			p.pos = savedPos
			break
		}
		// Guard against infinite loop on zero-width match.
		if p.pos == savedPos {
			p.pos = savedPos
			break
		}
		all = append(all, nodes...)
	}
	if len(all) < r.Min {
		return nil, &GrammarError{
			Code: "RepeatExpr",
			Msg:  fmt.Sprintf("expected at least %d repetition(s), got %d at %s", r.Min, len(all), tokenPos(p.tokens, p.pos)),
		}
	}
	return all, nil
}

type OptionalExpr struct {
	Expr Expr
}

func (o OptionalExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	nodes, err := o.Expr.Parse(p)
	if err != nil {
		p.pos = savedPos
		return []stringParsing.ParsedNode{}, nil
	}
	return nodes, nil
}

type NamedExpr struct {
	RuleName string
}

func (n NamedExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	rule, ok := p.grammar[n.RuleName]
	if !ok {
		keys := make([]string, 0, len(p.grammar))
		for k := range p.grammar {
			keys = append(keys, k)
		}
		return nil, &GrammarError{
			Code: "NamedExpr",
			Msg:  fmt.Sprintf("undefined rule '%s' (grammar has %d rule(s): %v)", n.RuleName, len(p.grammar), keys),
		}
	}
	return rule.Expr.Parse(p)
}

type NodeExpr struct {
	NodeType string
	Expr     Expr
}

func (n NodeExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	children, err := n.Expr.Parse(p)
	if err != nil {
		return nil, &GrammarError{
			Code:  "NodeExpr",
			Msg:   fmt.Sprintf("building node '%s'", n.NodeType),
			Cause: err,
		}
	}
	var b strings.Builder
	for _, child := range children {
		b.WriteString(child.Raw)
	}
	node := stringParsing.ParsedNode{
		Switch: n.NodeType,
		Raw:    b.String(),
		Metadata: map[string]interface{}{
			"children": children,
		},
	}
	return []stringParsing.ParsedNode{node}, nil
}

type ActionExpr struct {
	Expr   Expr
	Action func([]stringParsing.ParsedNode) (stringParsing.ParsedNode, core.ErrorInterface)
}

func (a ActionExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	children, err := a.Expr.Parse(p)
	if err != nil {
		return nil, &GrammarError{
			Code:  "ActionExpr",
			Msg:   "sub-expression failed",
			Cause: err,
		}
	}
	node, err := a.Action(children)
	if err != nil {
		return nil, &GrammarError{
			Code:  "ActionExpr",
			Msg:   fmt.Sprintf("user action returned error on %d node(s)", len(children)),
			Cause: err,
		}
	}
	return []stringParsing.ParsedNode{node}, nil
}

type NotExpr struct {
	Expr Expr
}

func (n NotExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	_, err := n.Expr.Parse(p)
	p.pos = savedPos
	if err == nil {
		return nil, &GrammarError{
			Code: "NotExpr",
			Msg:  fmt.Sprintf("unexpected match at %s (expression should not match here)", tokenPos(p.tokens, p.pos)),
		}
	}
	return []stringParsing.ParsedNode{}, nil
}

type AndExpr struct {
	Expr Expr
}

func (a AndExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	_, err := a.Expr.Parse(p)
	p.pos = savedPos
	if err != nil {
		return nil, &GrammarError{
			Code:  "AndExpr",
			Msg:   fmt.Sprintf("expression did not match at %s", tokenPos(p.tokens, p.pos)),
			Cause: err,
		}
	}
	return []stringParsing.ParsedNode{}, nil
}

type PeekExpr struct {
	TokenType string
}

func (p PeekExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	tok, err := prs.Peek()
	if err != nil {
		return nil, &GrammarError{
			Code:  "PeekExpr",
			Msg:   fmt.Sprintf("peeking for '%s'", p.TokenType),
			Cause: err,
		}
	}
	if tok.Switch != p.TokenType {
		return nil, &GrammarError{
			Code: "PeekExpr",
			Msg:  fmt.Sprintf("expected '%s', got '%s' (raw: %q) at %s", p.TokenType, tok.Switch, tok.Raw, tokenPos(prs.tokens, prs.pos)),
		}
	}
	return []stringParsing.ParsedNode{}, nil
}

type SeparatedRepeatExpr struct {
	Element Expr
	Sep     string
	Min     int
	Max     int
}

func (s SeparatedRepeatExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	var all []stringParsing.ParsedNode
	count := 0
	firstPos := p.pos
	firstNodes, err := s.Element.Parse(p)
	if err != nil {
		if s.Min == 0 {
			return []stringParsing.ParsedNode{}, nil
		}
		p.pos = firstPos
		return nil, &GrammarError{
			Code:  "SeparatedRepeatExpr",
			Msg:   fmt.Sprintf("expected at least %d element(s), got none at %s (separator: '%s')", s.Min, tokenPos(p.tokens, p.pos), s.Sep),
			Cause: err,
		}
	}
	all = append(all, firstNodes...)
	count++
	for {
		if s.Max > 0 && count >= s.Max {
			break
		}
		savedPos := p.pos
		_, err := p.Expect(s.Sep)
		if err != nil {
			p.pos = savedPos
			break
		}
		nodes, err := s.Element.Parse(p)
		if err != nil {
			p.pos = savedPos
			break
		}
		all = append(all, nodes...)
		count++
	}
	if count < s.Min {
		return nil, &GrammarError{
			Code: "SeparatedRepeatExpr",
			Msg:  fmt.Sprintf("expected at least %d elements, got %d at %s (separator: '%s')", s.Min, count, tokenPos(p.tokens, p.pos), s.Sep),
		}
	}
	return all, nil
}

type Associativity int

const (
	LeftAssoc Associativity = iota
	RightAssoc
	NonAssoc
)

type InfixInfo struct {
	Precedence int
	Assoc      Associativity
}

type PrattExpr struct {
	Atom     Expr
	Prefixes map[string]Expr
	Infixes  map[string]InfixInfo
}

const maxPrefixPrecedence = 100

func (p *PrattExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	node, err := p.parseExpression(prs, 0)
	if err != nil {
		return nil, &GrammarError{
			Code:  "PrattExpr",
			Msg:   "expression parsing failed",
			Cause: err,
		}
	}
	return []stringParsing.ParsedNode{node}, nil
}

func (p *PrattExpr) parseExpression(prs *Parser, minPrec int) (stringParsing.ParsedNode, core.ErrorInterface) {
	var leftNode stringParsing.ParsedNode
	tok, err := prs.Peek()
	if err == nil {
		if _, ok := p.Prefixes[tok.Switch]; ok {
			_, err := prs.Expect(tok.Switch)
			if err != nil {
				return stringParsing.ParsedNode{}, &GrammarError{
					Code:  "PrattExpr",
					Msg:   fmt.Sprintf("prefix operator '%s'", tok.Switch),
					Cause: err,
				}
			}
			rightNode, err := p.parseExpression(prs, maxPrefixPrecedence)
			if err != nil {
				return stringParsing.ParsedNode{}, &GrammarError{
					Code:  "PrattExpr",
					Msg:   fmt.Sprintf("prefix operator '%s' right operand", tok.Switch),
					Cause: err,
				}
			}
			var b strings.Builder
			b.WriteString(tok.Raw)
			b.WriteString(rightNode.Raw)
			leftNode = stringParsing.ParsedNode{
				Switch: "PrefixOp",
				Raw:    b.String(),
				Metadata: map[string]interface{}{
					"operator": tok.Switch,
					"operand":  rightNode,
				},
			}
		} else {
			atomNodes, err := p.Atom.Parse(prs)
			if err != nil {
				return stringParsing.ParsedNode{}, &GrammarError{
					Code:  "PrattExpr",
					Msg:   fmt.Sprintf("atom at %s", tokenPos(prs.tokens, prs.pos)),
					Cause: err,
				}
			}
			if len(atomNodes) == 0 {
				return stringParsing.ParsedNode{}, &GrammarError{
					Code: "PrattExpr",
					Msg:  fmt.Sprintf("atom returned empty at %s", tokenPos(prs.tokens, prs.pos)),
				}
			}
			if len(atomNodes) == 1 {
				leftNode = atomNodes[0]
			} else {
				var b strings.Builder
				for _, n := range atomNodes {
					b.WriteString(n.Raw)
				}
				leftNode = stringParsing.ParsedNode{
					Switch: "Sequence",
					Raw:    b.String(),
					Metadata: map[string]interface{}{
						"children": atomNodes,
					},
				}
			}
		}
	} else {
		return stringParsing.ParsedNode{}, &GrammarError{
			Code: "PrattExpr",
			Msg:  fmt.Sprintf("unexpected end of input at %s (need atom or prefix operator)", tokenPos(prs.tokens, prs.pos)),
		}
	}
	for {
		nextTok, err := prs.Peek()
		if err != nil {
			break
		}
		infix, ok := p.Infixes[nextTok.Switch]
		if !ok {
			break
		}
		if infix.Precedence < minPrec {
			break
		}
		opTok, err := prs.Expect(nextTok.Switch)
		if err != nil {
			return stringParsing.ParsedNode{}, &GrammarError{
				Code:  "PrattExpr",
				Msg:   fmt.Sprintf("infix operator '%s' at %s", nextTok.Switch, tokenPos(prs.tokens, prs.pos)),
				Cause: err,
			}
		}
		nextMinPrec := infix.Precedence
		if infix.Assoc == LeftAssoc {
			nextMinPrec = infix.Precedence + 1
		}
		rightNode, err := p.parseExpression(prs, nextMinPrec)
		if err != nil {
			return stringParsing.ParsedNode{}, &GrammarError{
				Code:  "PrattExpr",
				Msg:   fmt.Sprintf("infix operator '%s' right operand (minPrec=%d)", nextTok.Switch, nextMinPrec),
				Cause: err,
			}
		}
		var b strings.Builder
		b.WriteString(leftNode.Raw)
		b.WriteString(opTok.Raw)
		b.WriteString(rightNode.Raw)
		leftNode = stringParsing.ParsedNode{
			Switch: "BinaryOp",
			Raw:    b.String(),
			Metadata: map[string]interface{}{
				"operator": opTok.Switch,
				"left":     leftNode,
				"right":    rightNode,
			},
		}
	}
	return leftNode, nil
}
```

---

# parsing/stringParsing/parser3/parser.go

```go
package parser3

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

// tokenPos returns a human-readable position string for the token at idx.
// If the lexer injected __start/__end metadata, it includes that.
func tokenPos(tokens []stringParsing.ParsedNode, idx int) string {
	if idx < 0 || idx >= len(tokens) {
		return "EOF"
	}
	tok := tokens[idx]
	if start, ok := tok.Metadata["__start"].(int); ok {
		if end, ok2 := tok.Metadata["__end"].(int); ok2 {
			return fmt.Sprintf("idx=%d start=%d-%d", idx, start, end)
		}
		return fmt.Sprintf("idx=%d start=%d", idx, start)
	}
	return fmt.Sprintf("idx=%d", idx)
}

type Parser struct {
	lexer     *stringParsing.Lexer
	grammar   Grammar
	startRule string
	ignore    map[string]bool

	tokens []stringParsing.ParsedNode
	pos    int
}

func NewParser(lexer *stringParsing.Lexer, grammar Grammar, startRule string, ignoreTypes []string) *Parser {
	ignore := make(map[string]bool, len(ignoreTypes))
	for _, t := range ignoreTypes {
		ignore[t] = true
	}
	return &Parser{
		lexer:     lexer,
		grammar:   grammar,
		startRule: startRule,
		ignore:    ignore,
	}
}

func (p *Parser) Parse(code string, opts ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	tokens, err := p.lexer.Parse(code, opts...)
	if err != nil {
		return nil, &ParseError{
			Code:  "Lexer",
			Cause: err,
		}
	}
	p.tokens = tokens
	p.pos = 0

	startRule, ok := p.grammar[p.startRule]
	if !ok {
		keys := make([]string, 0, len(p.grammar))
		for k := range p.grammar {
			keys = append(keys, k)
		}
		return nil, &GrammarError{
			Code: "StartRule",
			Msg:  fmt.Sprintf("start rule '%s' not found in grammar (available: %v)", p.startRule, keys),
		}
	}

	children, err := startRule.Expr.Parse(p)
	if err != nil {
		return nil, &ParseError{
			Code:     "Parse",
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Cause:    err,
		}
	}

	p.SkipIgnored()
	if p.pos < len(p.tokens) {
		remaining := p.tokens[p.pos:]
		types := make([]string, len(remaining))
		for i, t := range remaining {
			types[i] = t.Switch
		}
		return nil, &ParseError{
			Code:     "UnexpectedToken",
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Got:      p.tokens[p.pos].Switch,
			Raw:      p.tokens[p.pos].Raw,
			Msg:      fmt.Sprintf("parsing complete — %d token(s) unconsumed (%v)", len(remaining), types),
		}
	}

	var b strings.Builder
	for _, child := range children {
		b.WriteString(child.Raw)
	}
	root := stringParsing.ParsedNode{
		Switch: p.startRule,
		Raw:    b.String(),
		Metadata: map[string]interface{}{
			"children": children,
		},
	}
	return []stringParsing.ParsedNode{root}, nil
}

func (p *Parser) SkipIgnored() {
	for p.pos < len(p.tokens) && p.ignore[p.tokens[p.pos].Switch] {
		p.pos++
	}
}

func (p *Parser) NextToken() (stringParsing.ParsedNode, error) {
	p.SkipIgnored()
	if p.pos >= len(p.tokens) {
		return stringParsing.ParsedNode{}, &ParseError{
			Code:     "NextToken",
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Msg:      "unexpected EOF",
		}
	}
	tok := p.tokens[p.pos]
	p.pos++
	return tok, nil
}

func (p *Parser) Expect(tokenType string) (stringParsing.ParsedNode, core.ErrorInterface) {
	tok, err := p.NextToken()
	if err != nil {
		return stringParsing.ParsedNode{}, &ParseError{
			Code:     "Expect",
			Expected: tokenType,
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Cause:    err,
		}
	}
	if tok.Switch != tokenType {
		return stringParsing.ParsedNode{}, &ParseError{
			Code:     "Expect",
			Expected: tokenType,
			Got:      tok.Switch,
			Raw:      tok.Raw,
			TokenIdx: p.pos - 1,
			TokenPos: tokenPos(p.tokens, p.pos-1),
		}
	}
	return tok, nil
}

func (p *Parser) Peek() (stringParsing.ParsedNode, error) {
	p.SkipIgnored()
	if p.pos >= len(p.tokens) {
		return stringParsing.ParsedNode{}, &ParseError{
			Code:     "Peek",
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Msg:      "EOF",
		}
	}
	return p.tokens[p.pos], nil
}

func (p *Parser) String() string {
	return "lc/parsing/stringParsing/parser3/Parser"
}
```

---

# parsing/stringParsing/utils.go

```go
package stringParsing

func addPrevNextNodes(nodes []ParsedNode) []ParsedNode {
	for i := 0; i < len(nodes); i++ {
		if i > 0 {
			nodes[i].Metadata["__prev"] = &nodes[i-1]
		} else {
			nodes[i].Metadata["__prev"] = nil
		}
		if i < len(nodes)-1 {
			nodes[i].Metadata["__next"] = &nodes[i+1]
		} else {
			nodes[i].Metadata["__next"] = nil
		}
	}
	return nodes
}
```

---

# public/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# public

'''go
import "github.com/pt-main/lc/public"
'''

## Index

- [Constants](<#constants>)
- [type EndianType](<#EndianType>)
- [type EngineType](<#EngineType>)
- [type ResType](<#ResType>)


## Constants

<a name="StringParseEvent"></a>

'''go
const (
    StringParseEvent       = "INPUT string->PARSED []ParsedNode"
    StringCallEvent        = "call(PARSED []ParsedNode)"
    StringCallCalloopEvent = "CALLOP call(PARSED []ParsedNode)"
    ByteParseEvent         = "INPUT []byte->PARSED []ParsedBytes"
    ByteCallEvent          = "call(PARSED []ParsedBytes)"
    ByteCallHotloopEvent   = "HOTLOOP call(PARSED []ParsedBytes)"
)
'''

<a name="CallEventsStartEvent"></a>

'''go
const (
    CallEventsStartEvent = "->call(Events.CallEvents)"
    CallEventsEndEvent   = "call(Events.CallEvents)->"
)
'''

<a name="LogEvents"></a>

'''go
const (
    LogEvents  = "SYSTEM:eventLogs"
    LogParsing = "SYSTEM:parsingLogs"
    LogVerbose = "SYSTEM:verboseLogs"
)
'''

<a name="ByteEngineScopeParsed"></a>

'''go
const (
    ByteEngineScopeParsed      = "PARSED []ParsedBytes"
    ByteEngineScopeEndianess   = "ENDIANESS int"
    ByteEngineScopeBytecodeIdx = "BYTECODE_IDX *int"
    ByteEngineScopeInput       = "INPUT []byte"
)
'''

<a name="StringEngineScopeInput"></a>

'''go
const (
    StringEngineScopeInput        = "INPUT string"
    StringEngineScopeParsed       = "PARSED []ParsedNode"
    StringEngineScopeInstrIdx     = "INSTR_IDX *int"
    StringEngineScopeCanBeUnknown = "StringEngineScopeCanBeUnknown bool"
)
'''

<a name="EventsScopeCallName"></a>

'''go
const (
    EventsScopeCallName  = "CALL_NAME string"
    EventsScopeCallError = "CALL_ERROR error"
    EventsScopeDERawLine = "RAW_LINE string"
)
'''

<a name="PluginsScopeEuPtr"></a>

'''go
const (
    PluginsScopeEuPtr = "EuPtr *EngineUniversal"
    EuScopePmPtr      = "PmPtr *PluginManager"
)
'''

<a name="EndianType"></a>
## type [EndianType](<https://github.com/pt-main/Lc/blob/main/public/types.go#L10>)



'''go
type EndianType int
'''

<a name="BigEndian"></a>

'''go
const (
    BigEndian EndianType = iota
    LittleEndian
)
'''

<a name="EngineType"></a>
## type [EngineType](<https://github.com/pt-main/Lc/blob/main/public/types.go#L17>)



'''go
type EngineType int
'''

<a name="ByteEngineType"></a>

'''go
const (
    ByteEngineType EngineType = iota
    StringEngineType
)
'''

<a name="ResType"></a>
## type [ResType](<https://github.com/pt-main/Lc/blob/main/public/types.go#L3>)



'''go
type ResType int
'''

<a name="ByteResType"></a>

'''go
const (
    ByteResType ResType = iota
    StringResType
)
'''

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# public/errors/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# errors

'''go
import "github.com/pt-main/lc/public/errors"
'''

## Index

- [type ErrorCodeType](<#ErrorCodeType>)
- [type ErrorMetaType](<#ErrorMetaType>)


<a name="ErrorCodeType"></a>
## type [ErrorCodeType](<https://github.com/pt-main/Lc/blob/main/public/errors/main.go#L3>)



'''go
type ErrorCodeType string
'''

<a name="ByteEngineProcessError1"></a>

'''go
const (
    ByteEngineProcessError1   ErrorCodeType = "Byte:PROCESS_ERR1"
    ByteEngineProcessError2   ErrorCodeType = "Byte:PROCESS_ERR2"
    StringEngineProcessError1 ErrorCodeType = "String:PROCESS_ERR1"
    StringEngineProcessError2 ErrorCodeType = "String:PROCESS_ERR2"
)
'''

<a name="DefaultEventsSystemError"></a>

'''go
const (
    DefaultEventsSystemError          ErrorCodeType = "SYSTEM@DEDAULT_EVENTS"
    DefaultEventsPanicError           ErrorCodeType = "DEDAULT_EVENTS:PANIC"
    DefaultEventsCallErrorCmdNotFound ErrorCodeType = "DEDAULT_EVENTS:CMD_NOT_FOUND"
    DefaultEventsCallErrorContexted   ErrorCodeType = "DEDAULT_EVENTS:CONTEXTED_ERROR"
    DefaultEventsCallErrorContex      ErrorCodeType = "DEDAULT_EVENTS:CONTEXT_ERROR"
    DefaultEventsCallErrorHandler     ErrorCodeType = "DEDAULT_EVENTS:HANDLER"
    DefaultEventsCallErrorUnknown     ErrorCodeType = "DEDAULT_EVENTS:UNKNOWN"
)
'''

<a name="EventsEventError"></a>

'''go
const (
    EventsEventError      ErrorCodeType = "EVENT_ERROR"
    EventsEventIsNotFound ErrorCodeType = "EVENT_IS_NOT_FOUND"
    EventsSystemError     ErrorCodeType = "SYSTEM"
)
'''

<a name="GeneratorGenerationTypeError"></a>

'''go
const (
    GeneratorGenerationTypeError ErrorCodeType = "GENERATION_TYPE_ERROR"
    GeneratorAddingTypeError     ErrorCodeType = "ADDING_TYPE_ERROR"
)
'''

<a name="CorePackageSystemError"></a>

'''go
const (
    CorePackageSystemError ErrorCodeType = "SYSTEM@CORE"
    CorePackageLcError     ErrorCodeType = "SYSTEM@LC"
    WrappedError           ErrorCodeType = "WrappedError"

    CorePackageLcLifecycleError ErrorCodeType = "SYSTEM@LC:LIFECYCLE"
)
'''

<a name="BytecodeShiftError"></a>

'''go
const (
    BytecodeShiftError ErrorCodeType = "BYTECODE_SHIFT_ERROR"
)
'''

<a name="ExtensiblePluginError"></a>

'''go
const (
    ExtensiblePluginError ErrorCodeType = "EXTENSIBLE_PLUGIN_ERROR"
)
'''

<a name="ParsingError"></a>

'''go
const (
    ParsingError ErrorCodeType = "PARSING_ERROR"
)
'''

<a name="ScopeGetError"></a>

'''go
const (
    ScopeGetError ErrorCodeType = "SCOPE_GET" /* metadata ScopeGetErrorMetakey : string */
)
'''

<a name="ErrExit"></a>

'''go
var ErrExit ErrorCodeType = "EXIT"
'''

<a name="ErrorMetaType"></a>
## type [ErrorMetaType](<https://github.com/pt-main/Lc/blob/main/public/errors/main.go#L4>)



'''go
type ErrorMetaType string
'''

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# public/errors/engines.go

```go
package errors

const (
	ByteEngineProcessError1   ErrorCodeType = "Byte:PROCESS_ERR1"
	ByteEngineProcessError2   ErrorCodeType = "Byte:PROCESS_ERR2"
	StringEngineProcessError1 ErrorCodeType = "String:PROCESS_ERR1"
	StringEngineProcessError2 ErrorCodeType = "String:PROCESS_ERR2"
)

const (
	DefaultEventsSystemError          ErrorCodeType = "SYSTEM@DEDAULT_EVENTS"
	DefaultEventsPanicError           ErrorCodeType = "DEDAULT_EVENTS:PANIC"
	DefaultEventsCallErrorCmdNotFound ErrorCodeType = "DEDAULT_EVENTS:CMD_NOT_FOUND"
	DefaultEventsCallErrorContexted   ErrorCodeType = "DEDAULT_EVENTS:CONTEXTED_ERROR"
	DefaultEventsCallErrorContex      ErrorCodeType = "DEDAULT_EVENTS:CONTEXT_ERROR"
	DefaultEventsCallErrorHandler     ErrorCodeType = "DEDAULT_EVENTS:HANDLER"
	DefaultEventsCallErrorUnknown     ErrorCodeType = "DEDAULT_EVENTS:UNKNOWN"
)
```

---

# public/errors/events.go

```go
package errors

const (
	EventsEventError      ErrorCodeType = "EVENT_ERROR"
	EventsEventIsNotFound ErrorCodeType = "EVENT_IS_NOT_FOUND"
	EventsSystemError     ErrorCodeType = "SYSTEM"
)
```

---

# public/errors/generator.go

```go
package errors

const (
	GeneratorGenerationTypeError ErrorCodeType = "GENERATION_TYPE_ERROR"
	GeneratorAddingTypeError     ErrorCodeType = "ADDING_TYPE_ERROR"
)
```

---

# public/errors/main.go

```go
package errors

type ErrorCodeType string
type ErrorMetaType string
```

---

# public/errors/others.go

```go
package errors

const (
	ScopeGetError ErrorCodeType = "SCOPE_GET" /* metadata ScopeGetErrorMetakey : string */
)

const (
	CorePackageSystemError ErrorCodeType = "SYSTEM@CORE"
	CorePackageLcError     ErrorCodeType = "SYSTEM@LC"
	WrappedError           ErrorCodeType = "WrappedError"

	CorePackageLcLifecycleError ErrorCodeType = "SYSTEM@LC:LIFECYCLE"
)

const (
	ParsingError ErrorCodeType = "PARSING_ERROR"
)

const (
	BytecodeShiftError ErrorCodeType = "BYTECODE_SHIFT_ERROR"
)

const (
	ExtensiblePluginError ErrorCodeType = "EXTENSIBLE_PLUGIN_ERROR"
)

var ErrExit ErrorCodeType = "EXIT"
```

---

# public/events.go

```go
package public

const (
	StringParseEvent       = "INPUT string->PARSED []ParsedNode"
	StringCallEvent        = "call(PARSED []ParsedNode)"
	StringCallCalloopEvent = "CALLOP call(PARSED []ParsedNode)"
	ByteParseEvent         = "INPUT []byte->PARSED []ParsedBytes"
	ByteCallEvent          = "call(PARSED []ParsedBytes)"
	ByteCallHotloopEvent   = "HOTLOOP call(PARSED []ParsedBytes)"
)

const (
	CallEventsStartEvent = "->call(Events.CallEvents)"
	CallEventsEndEvent   = "call(Events.CallEvents)->"
)
```

---

# public/logging.go

```go
package public

const (
	LogEvents  = "SYSTEM:eventLogs"
	LogParsing = "SYSTEM:parsingLogs"
	LogVerbose = "SYSTEM:verboseLogs"
)
```

---

# public/scope.go

```go
package public

const (
	ByteEngineScopeParsed      = "PARSED []ParsedBytes"
	ByteEngineScopeEndianess   = "ENDIANESS int"
	ByteEngineScopeBytecodeIdx = "BYTECODE_IDX *int"
	ByteEngineScopeInput       = "INPUT []byte"
)

const (
	StringEngineScopeInput        = "INPUT string"
	StringEngineScopeParsed       = "PARSED []ParsedNode"
	StringEngineScopeInstrIdx     = "INSTR_IDX *int"
	StringEngineScopeCanBeUnknown = "StringEngineScopeCanBeUnknown bool"
)

const (
	EventsScopeCallName  = "CALL_NAME string"
	EventsScopeCallError = "CALL_ERROR error"
	EventsScopeDERawLine = "RAW_LINE string"
)

const (
	PluginsScopeEuPtr = "EuPtr *EngineUniversal"
	EuScopePmPtr      = "PmPtr *PluginManager"
)
```

---

# public/types.go

```go
package public

type ResType int

const (
	ByteResType ResType = iota
	StringResType
)

type EndianType int

const (
	BigEndian EndianType = iota
	LittleEndian
)

type EngineType int

const (
	ByteEngineType EngineType = iota
	StringEngineType
)
```

---

# scan_results.txt

```txt
Scan Results
Directory: /Users/macbook/Desktop/lc/Lc
Total lines: 23610
Total files: 100

Files:
Lc/merged.md
Lc/README.md
Lc/engine.go
Lc/builder.go
Lc/GODOC.md
Lc/main.go
Lc/tooling/README.md
Lc/tooling/astools/GODOC.md
Lc/tooling/astools/main.go
Lc/tooling/debugging/extensiblePlugin/config.go
Lc/tooling/debugging/extensiblePlugin/events.go
Lc/tooling/debugging/extensiblePlugin/plugin.go
Lc/tooling/debugging/extensiblePlugin/structs.go
Lc/tooling/debugging/extensiblePlugin/GODOC.md
Lc/tooling/debugging/profiler/GODOC.md
Lc/tooling/debugging/profiler/main.go
Lc/tooling/plugin/realization.go
Lc/tooling/plugin/interface.go
Lc/tooling/plugin/tools.go
Lc/tooling/plugin/GODOC.md
Lc/tooling/plugin/manager.go
Lc/tooling/bytecode/utils.go
Lc/tooling/bytecode/GODOC.md
Lc/tooling/bytecode/instruction.go
Lc/example/README.md
Lc/example/tests/speedtest/README.md
Lc/example/tests/speedtest/byte/bench/byte_test.go
Lc/example/tests/speedtest/byte/tests/GODOC.md
Lc/example/tests/speedtest/byte/tests/main.go
Lc/example/tests/parser3Test/parser3.go
Lc/example/tests/parser3Test/GODOC.md
Lc/example/readme/byte/GODOC.md
Lc/example/readme/byte/main.go
Lc/example/readme/string/GODOC.md
Lc/example/readme/string/main.go
Lc/example/langs/calculator/GODOC.md
Lc/example/langs/calculator/main.go
Lc/example/langs/configLang/GODOC.md
Lc/example/langs/configLang/main.go
Lc/example/packages/engine/core/logger/GODOC.md
Lc/example/packages/engine/core/logger/main.go
Lc/example/packages/engine/core/other/GODOC.md
Lc/example/packages/engine/core/other/main.go
Lc/example/packages/engine/core/generator/GODOC.md
Lc/example/packages/engine/core/generator/main.go
Lc/example/packages/engine/core/events/GODOC.md
Lc/example/packages/engine/core/events/main.go
Lc/example/packages/engine/engines/byte/GODOC.md
Lc/example/packages/engine/engines/byte/main.go
Lc/example/packages/engine/engines/string/GODOC.md
Lc/example/packages/engine/engines/string/main.go
Lc/public/logging.go
Lc/public/types.go
Lc/public/events.go
Lc/public/GODOC.md
Lc/public/scope.go
Lc/public/errors/events.go
Lc/public/errors/engines.go
Lc/public/errors/generator.go
Lc/public/errors/GODOC.md
Lc/public/errors/main.go
Lc/public/errors/others.go
Lc/parsing/README.md
Lc/parsing/GODOC.md
Lc/parsing/main.go
Lc/parsing/byteParsing/GODOC.md
Lc/parsing/byteParsing/parser1.go
Lc/parsing/byteParsing/parser1_test.go
Lc/parsing/byteParsing/node.go
Lc/parsing/stringParsing/parser2.go
Lc/parsing/stringParsing/lexer_test.go
Lc/parsing/stringParsing/README.md
Lc/parsing/stringParsing/lexer.go
Lc/parsing/stringParsing/utils.go
Lc/parsing/stringParsing/GODOC.md
Lc/parsing/stringParsing/parser1.go
Lc/parsing/stringParsing/node.go
Lc/parsing/stringParsing/parser3/formatter.go
Lc/parsing/stringParsing/parser3/grammar.go
Lc/parsing/stringParsing/parser3/parser.go
Lc/parsing/stringParsing/parser3/GODOC.md
Lc/parsing/stringParsing/parser3/engineAdapter.go
Lc/parsing/stringParsing/parser3/errors.go
Lc/engine/interface.go
Lc/engine/byteEngine.go
Lc/engine/GODOC.md
Lc/engine/stringEngine.go
Lc/engine/core/types.go
Lc/engine/core/logger.go
Lc/engine/core/events.go
Lc/engine/core/universalEngineParams.go
Lc/engine/core/generator.go
Lc/engine/core/GODOC.md
Lc/engine/core/scope.go
Lc/engine/core/generator_test.go
Lc/engine/core/errors.go
Lc/engine/events/byteEngine.go
Lc/engine/events/GODOC.md
Lc/engine/events/stringEngine.go
Lc/engine/events/defaultEvents.go
```

---

# tooling/README.md

```md
# Tooling

A package of other independent tools:

- 'astools': Simple tools for work with 'parsing/stringParsing/parser3/Parser3 or Adapter' Ast.

- 'bytecode': Package for int/float->bytecode, bytecode->int/float transformations.

- 'debugging': 
    - 'extensiblePlugin': Plugin ('ExtensibleCLPlugin') with hooks in calloop.
    - 'profiler': Plugin for profiling lc execution loop (require 'ExtensibleCLPlugin').

- 'plugin': Plugin manager and PluginInterface realization.
```

---

# tooling/astools/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# astools

'''go
import "github.com/pt-main/lc/tooling/astools"
'''

## Index

- [func FindChild\(node \*stringParsing.ParsedNode, switchName string\) \*stringParsing.ParsedNode](<#FindChild>)
- [func FindChildIndex\(node \*stringParsing.ParsedNode, switchName string\) int](<#FindChildIndex>)
- [func FindChildren\(node \*stringParsing.ParsedNode, switchName string\) \[\]stringParsing.ParsedNode](<#FindChildren>)
- [func GetChildAt\(node \*stringParsing.ParsedNode, index int\) \*stringParsing.ParsedNode](<#GetChildAt>)
- [func GetChildren\(node \*stringParsing.ParsedNode\) \[\]stringParsing.ParsedNode](<#GetChildren>)
- [func GetTokenValue\(node \*stringParsing.ParsedNode\) string](<#GetTokenValue>)
- [func Walk\(node \*stringParsing.ParsedNode, fn func\(\*stringParsing.ParsedNode\) error\) error](<#Walk>)


<a name="FindChild"></a>
## func [FindChild](<https://github.com/pt-main/Lc/blob/main/tooling/astools/main.go#L15>)

'''go
func FindChild(node *stringParsing.ParsedNode, switchName string) *stringParsing.ParsedNode
'''



<a name="FindChildIndex"></a>
## func [FindChildIndex](<https://github.com/pt-main/Lc/blob/main/tooling/astools/main.go#L24>)

'''go
func FindChildIndex(node *stringParsing.ParsedNode, switchName string) int
'''



<a name="FindChildren"></a>
## func [FindChildren](<https://github.com/pt-main/Lc/blob/main/tooling/astools/main.go#L37>)

'''go
func FindChildren(node *stringParsing.ParsedNode, switchName string) []stringParsing.ParsedNode
'''



<a name="GetChildAt"></a>
## func [GetChildAt](<https://github.com/pt-main/Lc/blob/main/tooling/astools/main.go#L60>)

'''go
func GetChildAt(node *stringParsing.ParsedNode, index int) *stringParsing.ParsedNode
'''



<a name="GetChildren"></a>
## func [GetChildren](<https://github.com/pt-main/Lc/blob/main/tooling/astools/main.go#L5>)

'''go
func GetChildren(node *stringParsing.ParsedNode) []stringParsing.ParsedNode
'''



<a name="GetTokenValue"></a>
## func [GetTokenValue](<https://github.com/pt-main/Lc/blob/main/tooling/astools/main.go#L47>)

'''go
func GetTokenValue(node *stringParsing.ParsedNode) string
'''



<a name="Walk"></a>
## func [Walk](<https://github.com/pt-main/Lc/blob/main/tooling/astools/main.go#L68>)

'''go
func Walk(node *stringParsing.ParsedNode, fn func(*stringParsing.ParsedNode) error) error
'''



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# tooling/astools/main.go

```go
package astools

import "github.com/pt-main/lc/parsing/stringParsing"

func GetChildren(node *stringParsing.ParsedNode) []stringParsing.ParsedNode {
	if node == nil {
		return nil
	}
	if children, ok := node.Metadata["children"].([]stringParsing.ParsedNode); ok {
		return children
	}
	return nil
}

func FindChild(node *stringParsing.ParsedNode, switchName string) *stringParsing.ParsedNode {
	for _, child := range GetChildren(node) {
		if child.Switch == switchName {
			return &child
		}
	}
	return nil
}

func FindChildIndex(node *stringParsing.ParsedNode, switchName string) int {
	if node == nil {
		return -1
	}
	children := GetChildren(node)
	for i, child := range children {
		if child.Switch == switchName {
			return i
		}
	}
	return -1
}

func FindChildren(node *stringParsing.ParsedNode, switchName string) []stringParsing.ParsedNode {
	var result []stringParsing.ParsedNode
	for _, child := range GetChildren(node) {
		if child.Switch == switchName {
			result = append(result, child)
		}
	}
	return result
}

func GetTokenValue(node *stringParsing.ParsedNode) string {
	if node == nil {
		return ""
	}
	if node.Raw != "" {
		return node.Raw
	}
	if val, ok := node.Metadata["value"].(string); ok {
		return val
	}
	return ""
}

func GetChildAt(node *stringParsing.ParsedNode, index int) *stringParsing.ParsedNode {
	children := GetChildren(node)
	if index >= 0 && index < len(children) {
		return &children[index]
	}
	return nil
}

func Walk(node *stringParsing.ParsedNode, fn func(*stringParsing.ParsedNode) error) error {
	if node == nil {
		return nil
	}
	if err := fn(node); err != nil {
		return err
	}
	for _, child := range GetChildren(node) {
		if err := Walk(&child, fn); err != nil {
			return err
		}
	}
	return nil
}
```

---

# tooling/bytecode/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# bytecode

'''go
import "github.com/pt-main/lc/tooling/bytecode"
'''

## Index

- [type GenerationConfig](<#GenerationConfig>)
- [type InstructionsGenerator](<#InstructionsGenerator>)
  - [func \(ig \*InstructionsGenerator\) Generate\(opcode int, args \[\]\[\]byte\) \[\]byte](<#InstructionsGenerator.Generate>)
- [type Shift](<#Shift>)
  - [func NewShift\(code \[\]byte, idx \*int\) \*Shift](<#NewShift>)
  - [func \(s \*Shift\) ShiftError\(length int\) \(\[\]byte, core.ErrorInterface\)](<#Shift.ShiftError>)
  - [func \(s \*Shift\) ShiftFloat64Error\(size int, endianess public.EndianType\) \(float64, core.ErrorInterface\)](<#Shift.ShiftFloat64Error>)
  - [func \(s \*Shift\) ShiftFloat64Panic\(size int, endianess public.EndianType\) float64](<#Shift.ShiftFloat64Panic>)
  - [func \(s \*Shift\) ShiftFloat64RangeError\(size int, minVal, maxVal float64, endianess public.EndianType\) \(float64, core.ErrorInterface\)](<#Shift.ShiftFloat64RangeError>)
  - [func \(s \*Shift\) ShiftFloat64RangePanic\(size int, minVal, maxVal float64, endianess public.EndianType\) float64](<#Shift.ShiftFloat64RangePanic>)
  - [func \(s \*Shift\) ShiftPanic\(length int\) \[\]byte](<#Shift.ShiftPanic>)
- [type Utils](<#Utils>)
  - [func \(u \*Utils\) AutoFloat64ToBytes\(value float64, endianess public.EndianType\) \[\]byte](<#Utils.AutoFloat64ToBytes>)
  - [func \(u \*Utils\) AutoIntToBytes\(value int, endianess public.EndianType\) \[\]byte](<#Utils.AutoIntToBytes>)
  - [func \(u \*Utils\) BytesToFloat64\(bytes \[\]byte, endianess public.EndianType\) float64](<#Utils.BytesToFloat64>)
  - [func \(u \*Utils\) BytesToFloat64BigEndian\(bytes \[\]byte\) float64](<#Utils.BytesToFloat64BigEndian>)
  - [func \(u \*Utils\) BytesToFloat64LittleEndian\(bytes \[\]byte\) float64](<#Utils.BytesToFloat64LittleEndian>)
  - [func \(u \*Utils\) BytesToFloat64Range\(bytes \[\]byte, minVal, maxVal float64, endianess public.EndianType\) float64](<#Utils.BytesToFloat64Range>)
  - [func \(u \*Utils\) BytesToInt\(bytes \[\]byte, endianess public.EndianType\) int](<#Utils.BytesToInt>)
  - [func \(u \*Utils\) BytesToIntBigEndian\(bytes \[\]byte\) int](<#Utils.BytesToIntBigEndian>)
  - [func \(u \*Utils\) BytesToIntLittleEndian\(bytes \[\]byte\) int](<#Utils.BytesToIntLittleEndian>)
  - [func \(u \*Utils\) Float64ToBytes\(value float64, size int, endianess public.EndianType\) \[\]byte](<#Utils.Float64ToBytes>)
  - [func \(u \*Utils\) Float64ToBytesBigEndian\(value float64, size int\) \[\]byte](<#Utils.Float64ToBytesBigEndian>)
  - [func \(u \*Utils\) Float64ToBytesLittleEndian\(value float64, size int\) \[\]byte](<#Utils.Float64ToBytesLittleEndian>)
  - [func \(u \*Utils\) Float64ToBytesRange\(value float64, size int, minVal, maxVal float64, endianess public.EndianType\) \[\]byte](<#Utils.Float64ToBytesRange>)
  - [func \(u \*Utils\) IntToBytes\(value int, size int, endianess public.EndianType\) \[\]byte](<#Utils.IntToBytes>)
  - [func \(u \*Utils\) IntToBytesBigEndian\(value int, size int\) \[\]byte](<#Utils.IntToBytesBigEndian>)
  - [func \(u \*Utils\) IntToBytesLittleEndian\(value int, size int\) \[\]byte](<#Utils.IntToBytesLittleEndian>)


<a name="GenerationConfig"></a>
## type [GenerationConfig](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/instruction.go#L5-L10>)



'''go
type GenerationConfig struct {
    CommandBytelen   int
    ArglenBytelen    int
    ArgscountBytelen int
    Endianess        public.EndianType
}
'''

<a name="InstructionsGenerator"></a>
## type [InstructionsGenerator](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/instruction.go#L12-L14>)



'''go
type InstructionsGenerator struct {
    Config GenerationConfig
}
'''

<a name="InstructionsGenerator.Generate"></a>
### func \(\*InstructionsGenerator\) [Generate](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/instruction.go#L20-L22>)

'''go
func (ig *InstructionsGenerator) Generate(opcode int, args [][]byte) []byte
'''

Generate instruction. byteParsing.Parser1 specification

\!\!\! PANIC IF ANY ARGLEN IS ZERO \!\!\!

<a name="Shift"></a>
## type [Shift](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L185-L188>)

Shift provides safe and unsafe byte reading from a buffer with an internal index.

'''go
type Shift struct {
    Code []byte
    Idx  *int
}
'''

<a name="NewShift"></a>
### func [NewShift](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L191>)

'''go
func NewShift(code []byte, idx *int) *Shift
'''

NewShift creates a new Shift instance.

<a name="Shift.ShiftError"></a>
### func \(\*Shift\) [ShiftError](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L206>)

'''go
func (s *Shift) ShiftError(length int) ([]byte, core.ErrorInterface)
'''

ShiftError reads \'length\' bytes from the buffer and returns them. If there is not enough data, returns a core.Error.

Err errors.BytecodeShiftError:

- On unexpected end of data. Meta: EMK\(0, "int"\) – requested length, EMK\(1, "int"\) – current index, EMK\(2, "int"\) – total buffer length.

<a name="Shift.ShiftFloat64Error"></a>
### func \(\*Shift\) [ShiftFloat64Error](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L232>)

'''go
func (s *Shift) ShiftFloat64Error(size int, endianess public.EndianType) (float64, core.ErrorInterface)
'''

ShiftFloat64Error reads \'size\' bytes and interprets them as a float64 in \[\-1,1\]. If there is not enough data, returns a core.Error.

Err errors.BytecodeShiftError \(wrapped from ShiftError\).

<a name="Shift.ShiftFloat64Panic"></a>
### func \(\*Shift\) [ShiftFloat64Panic](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L242>)

'''go
func (s *Shift) ShiftFloat64Panic(size int, endianess public.EndianType) float64
'''

ShiftFloat64Panic reads and converts a float64, panicking on error.

<a name="Shift.ShiftFloat64RangeError"></a>
### func \(\*Shift\) [ShiftFloat64RangeError](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L252>)

'''go
func (s *Shift) ShiftFloat64RangeError(size int, minVal, maxVal float64, endianess public.EndianType) (float64, core.ErrorInterface)
'''

ShiftFloat64RangeError reads \'size\' bytes and interprets them as a float64 in \[minVal, maxVal\]. If there is not enough data, returns a core.Error.

Err errors.BytecodeShiftError \(wrapped from ShiftError\).

<a name="Shift.ShiftFloat64RangePanic"></a>
### func \(\*Shift\) [ShiftFloat64RangePanic](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L262>)

'''go
func (s *Shift) ShiftFloat64RangePanic(size int, minVal, maxVal float64, endianess public.EndianType) float64
'''

ShiftFloat64RangePanic reads and converts a float64 with range, panicking on error.

<a name="Shift.ShiftPanic"></a>
### func \(\*Shift\) [ShiftPanic](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L220>)

'''go
func (s *Shift) ShiftPanic(length int) []byte
'''

ShiftPanic reads \'length\' bytes, panicking if not enough data. Use only in contexts where you are certain the buffer is large enough.

<a name="Utils"></a>
## type [Utils](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L12>)

Utils provides byte\-level conversions and shifting utilities.

'''go
type Utils struct{}
'''

<a name="Utils.AutoFloat64ToBytes"></a>
### func \(\*Utils\) [AutoFloat64ToBytes](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L288>)

'''go
func (u *Utils) AutoFloat64ToBytes(value float64, endianess public.EndianType) []byte
'''

AutoFloat64ToBytes chooses an optimal byte size based on the magnitude of \'value\', then converts it to bytes.

<a name="Utils.AutoIntToBytes"></a>
### func \(\*Utils\) [AutoIntToBytes](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L270>)

'''go
func (u *Utils) AutoIntToBytes(value int, endianess public.EndianType) []byte
'''

AutoIntToBytes chooses the minimal number of bytes needed to represent \'value\' \(excluding sign extension\) and converts it.

<a name="Utils.BytesToFloat64"></a>
### func \(\*Utils\) [BytesToFloat64](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L111>)

'''go
func (u *Utils) BytesToFloat64(bytes []byte, endianess public.EndianType) float64
'''

BytesToFloat64 converts a byte slice to a float64 in range \[\-1,1\].

<a name="Utils.BytesToFloat64BigEndian"></a>
### func \(\*Utils\) [BytesToFloat64BigEndian](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L175>)

'''go
func (u *Utils) BytesToFloat64BigEndian(bytes []byte) float64
'''

BytesToFloat64BigEndian is a convenience wrapper.

<a name="Utils.BytesToFloat64LittleEndian"></a>
### func \(\*Utils\) [BytesToFloat64LittleEndian](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L180>)

'''go
func (u *Utils) BytesToFloat64LittleEndian(bytes []byte) float64
'''

BytesToFloat64LittleEndian is a convenience wrapper.

<a name="Utils.BytesToFloat64Range"></a>
### func \(\*Utils\) [BytesToFloat64Range](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L147>)

'''go
func (u *Utils) BytesToFloat64Range(bytes []byte, minVal, maxVal float64, endianess public.EndianType) float64
'''

BytesToFloat64Range converts a byte slice to a float64 in \[minVal, maxVal\]. Panics: if minVal \>= maxVal.

<a name="Utils.BytesToInt"></a>
### func \(\*Utils\) [BytesToInt](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L76>)

'''go
func (u *Utils) BytesToInt(bytes []byte, endianess public.EndianType) int
'''

BytesToInt converts a byte slice to an int using the specified endianness.

<a name="Utils.BytesToIntBigEndian"></a>
### func \(\*Utils\) [BytesToIntBigEndian](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L49>)

'''go
func (u *Utils) BytesToIntBigEndian(bytes []byte) int
'''

BytesToIntBigEndian converts a big\-endian byte slice to an int. It handles sign extension.

<a name="Utils.BytesToIntLittleEndian"></a>
### func \(\*Utils\) [BytesToIntLittleEndian](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L63>)

'''go
func (u *Utils) BytesToIntLittleEndian(bytes []byte) int
'''

BytesToIntLittleEndian converts a little\-endian byte slice to an int. It handles sign extension.

<a name="Utils.Float64ToBytes"></a>
### func \(\*Utils\) [Float64ToBytes](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L85>)

'''go
func (u *Utils) Float64ToBytes(value float64, size int, endianess public.EndianType) []byte
'''

Float64ToBytes converts a float64 in range \[\-1,1\] to a byte slice of given size. Panics: if size \<= 0 or value exceeds the representable range \(clamped\).

<a name="Utils.Float64ToBytesBigEndian"></a>
### func \(\*Utils\) [Float64ToBytesBigEndian](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L165>)

'''go
func (u *Utils) Float64ToBytesBigEndian(value float64, size int) []byte
'''

Float64ToBytesBigEndian is a convenience wrapper for big\-endian conversion.

<a name="Utils.Float64ToBytesLittleEndian"></a>
### func \(\*Utils\) [Float64ToBytesLittleEndian](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L170>)

'''go
func (u *Utils) Float64ToBytesLittleEndian(value float64, size int) []byte
'''

Float64ToBytesLittleEndian is a convenience wrapper for little\-endian conversion.

<a name="Utils.Float64ToBytesRange"></a>
### func \(\*Utils\) [Float64ToBytesRange](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L126>)

'''go
func (u *Utils) Float64ToBytesRange(value float64, size int, minVal, maxVal float64, endianess public.EndianType) []byte
'''

Float64ToBytesRange converts a float64 in \[minVal, maxVal\] to a byte slice. Panics: if size \<= 0, minVal \>= maxVal, or value out of range \(clamped\).

<a name="Utils.IntToBytes"></a>
### func \(\*Utils\) [IntToBytes](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L40>)

'''go
func (u *Utils) IntToBytes(value int, size int, endianess public.EndianType) []byte
'''

IntToBytes converts an int to a byte slice with the given endianness. Panics: if size \<= 0 \(via called functions\).

<a name="Utils.IntToBytesBigEndian"></a>
### func \(\*Utils\) [IntToBytesBigEndian](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L16>)

'''go
func (u *Utils) IntToBytesBigEndian(value int, size int) []byte
'''

IntToBytesBigEndian converts an int to a big\-endian byte slice of the given size. Panics: if size \<= 0 \(not checked, will panic on slice allocation\).

<a name="Utils.IntToBytesLittleEndian"></a>
### func \(\*Utils\) [IntToBytesLittleEndian](<https://github.com/pt-main/Lc/blob/main/tooling/bytecode/utils.go#L28>)

'''go
func (u *Utils) IntToBytesLittleEndian(value int, size int) []byte
'''

IntToBytesLittleEndian converts an int to a little\-endian byte slice of the given size. Panics: if size \<= 0 \(not checked, will panic on slice allocation\).

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# tooling/bytecode/instruction.go

```go
package bytecode

import "github.com/pt-main/lc/public"

type GenerationConfig struct {
	CommandBytelen   int
	ArglenBytelen    int
	ArgscountBytelen int
	Endianess        public.EndianType
}

type InstructionsGenerator struct {
	Config GenerationConfig
}

// Generate instruction.
// byteParsing.Parser1 specification
//
// !!! PANIC IF ANY ARGLEN IS ZERO !!!
func (ig *InstructionsGenerator) Generate(
	opcode int, args [][]byte,
) []byte {
	u := Utils{}
	res := append(
		append([]byte{}, u.IntToBytes(opcode, ig.Config.CommandBytelen, ig.Config.Endianess)...),
		u.IntToBytes(len(args), ig.Config.ArgscountBytelen, ig.Config.Endianess)...,
	)
	for _, arg := range args {
		if len(arg) == 0 {
			panic("Argument length cannot be zero")
		}
		res = append(res, u.IntToBytes(len(arg), ig.Config.ArglenBytelen, ig.Config.Endianess)...)
		res = append(res, arg...)
	}
	return res
}
```

---

# tooling/bytecode/utils.go

```go
package bytecode

import (
	"math"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

// Utils provides byte-level conversions and shifting utilities.
type Utils struct{}

// IntToBytesBigEndian converts an int to a big-endian byte slice of the given size.
// Panics: if size <= 0 (not checked, will panic on slice allocation).
func (u *Utils) IntToBytesBigEndian(value int, size int) []byte {
	result := make([]byte, size)
	val := uint64(value)
	for i := size - 1; i >= 0; i-- {
		result[i] = byte(val & 0xFF)
		val >>= 8
	}
	return result
}

// IntToBytesLittleEndian converts an int to a little-endian byte slice of the given size.
// Panics: if size <= 0 (not checked, will panic on slice allocation).
func (u *Utils) IntToBytesLittleEndian(value int, size int) []byte {
	result := make([]byte, size)
	val := uint64(value)
	for i := 0; i < size; i++ {
		result[i] = byte(val & 0xFF)
		val >>= 8
	}
	return result
}

// IntToBytes converts an int to a byte slice with the given endianness.
// Panics: if size <= 0 (via called functions).
func (u *Utils) IntToBytes(value int, size int, endianess public.EndianType) []byte {
	if endianess == public.BigEndian {
		return u.IntToBytesBigEndian(value, size)
	}
	return u.IntToBytesLittleEndian(value, size)
}

// BytesToIntBigEndian converts a big-endian byte slice to an int.
// It handles sign extension.
func (u *Utils) BytesToIntBigEndian(bytes []byte) int {
	var val uint64 = 0
	for _, b := range bytes {
		val = (val << 8) | uint64(b)
	}
	bits := uint(len(bytes) * 8)
	if bits > 0 && (val>>(bits-1))&1 == 1 {
		val |= ^uint64(0) << bits
	}
	return int(val)
}

// BytesToIntLittleEndian converts a little-endian byte slice to an int.
// It handles sign extension.
func (u *Utils) BytesToIntLittleEndian(bytes []byte) int {
	var val uint64 = 0
	for i, b := range bytes {
		val |= uint64(b) << (8 * i)
	}
	bits := uint(len(bytes) * 8)
	if bits > 0 && (val>>(bits-1))&1 == 1 {
		val |= ^uint64(0) << bits
	}
	return int(val)
}

// BytesToInt converts a byte slice to an int using the specified endianness.
func (u *Utils) BytesToInt(bytes []byte, endianess public.EndianType) int {
	if endianess == public.BigEndian {
		return u.BytesToIntBigEndian(bytes)
	}
	return u.BytesToIntLittleEndian(bytes)
}

// Float64ToBytes converts a float64 in range [-1,1] to a byte slice of given size.
// Panics: if size <= 0 or value exceeds the representable range (clamped).
func (u *Utils) Float64ToBytes(value float64, size int, endianess public.EndianType) []byte {
	if size <= 0 {
		panic("Float64ToBytes: size must be positive")
	}
	maxValue := uint64(1<<(uint(size)*8) - 1)
	var scaledValue uint64
	if value >= 0 {
		if value > 1.0 {
			scaledValue = maxValue
		} else {
			scaledValue = uint64(value * float64(maxValue))
		}
	} else {
		if value < -1.0 {
			scaledValue = 0
		} else {
			scaledValue = uint64((value + 1.0) * float64(maxValue))
		}
	}
	if scaledValue > maxValue {
		scaledValue = maxValue
	}
	return u.IntToBytes(int(scaledValue), size, endianess)
}

// BytesToFloat64 converts a byte slice to a float64 in range [-1,1].
func (u *Utils) BytesToFloat64(bytes []byte, endianess public.EndianType) float64 {
	size := len(bytes)
	if size == 0 {
		return 0.0
	}
	intValue := uint64(u.BytesToInt(bytes, endianess))
	maxValue := uint64(1<<(uint(size)*8) - 1)
	if maxValue == 0 {
		return 0.0
	}
	return float64(intValue)/float64(maxValue)*2.0 - 1.0
}

// Float64ToBytesRange converts a float64 in [minVal, maxVal] to a byte slice.
// Panics: if size <= 0, minVal >= maxVal, or value out of range (clamped).
func (u *Utils) Float64ToBytesRange(value float64, size int, minVal, maxVal float64, endianess public.EndianType) []byte {
	if size <= 0 {
		panic("Float64ToBytesRange: size must be positive")
	}
	if minVal >= maxVal {
		panic("Float64ToBytesRange: minVal must be less than maxVal")
	}
	if value < minVal {
		value = minVal
	}
	if value > maxVal {
		value = maxVal
	}
	normalized := (value - minVal) / (maxVal - minVal)
	maxValue := uint64(1<<(uint(size)*8) - 1)
	scaledValue := uint64(normalized * float64(maxValue))
	return u.IntToBytes(int(scaledValue), size, endianess)
}

// BytesToFloat64Range converts a byte slice to a float64 in [minVal, maxVal].
// Panics: if minVal >= maxVal.
func (u *Utils) BytesToFloat64Range(bytes []byte, minVal, maxVal float64, endianess public.EndianType) float64 {
	size := len(bytes)
	if size == 0 {
		return minVal
	}
	if minVal >= maxVal {
		panic("BytesToFloat64Range: minVal must be less than maxVal")
	}
	intValue := uint64(u.BytesToInt(bytes, endianess))
	maxValue := uint64(1<<(uint(size)*8) - 1)
	if maxValue == 0 {
		return minVal
	}
	normalized := float64(intValue) / float64(maxValue)
	return minVal + normalized*(maxVal-minVal)
}

// Float64ToBytesBigEndian is a convenience wrapper for big-endian conversion.
func (u *Utils) Float64ToBytesBigEndian(value float64, size int) []byte {
	return u.Float64ToBytes(value, size, public.BigEndian)
}

// Float64ToBytesLittleEndian is a convenience wrapper for little-endian conversion.
func (u *Utils) Float64ToBytesLittleEndian(value float64, size int) []byte {
	return u.Float64ToBytes(value, size, public.LittleEndian)
}

// BytesToFloat64BigEndian is a convenience wrapper.
func (u *Utils) BytesToFloat64BigEndian(bytes []byte) float64 {
	return u.BytesToFloat64(bytes, public.BigEndian)
}

// BytesToFloat64LittleEndian is a convenience wrapper.
func (u *Utils) BytesToFloat64LittleEndian(bytes []byte) float64 {
	return u.BytesToFloat64(bytes, public.LittleEndian)
}

// Shift provides safe and unsafe byte reading from a buffer with an internal index.
type Shift struct {
	Code []byte
	Idx  *int
}

// NewShift creates a new Shift instance.
func NewShift(code []byte, idx *int) *Shift {
	return &Shift{
		Code: code,
		Idx:  idx,
	}
}

// ShiftError reads 'length' bytes from the buffer and returns them.
// If there is not enough data, returns a core.Error.
//
// Err errors.BytecodeShiftError:
//   - On unexpected end of data.
//     Meta: EMK(0, "int") – requested length,
//     EMK(1, "int") – current index,
//     EMK(2, "int") – total buffer length.
func (s *Shift) ShiftError(length int) ([]byte, core.ErrorInterface) {
	if *s.Idx+length > len(s.Code) {
		return nil, core.Err(errors.BytecodeShiftError, "Unexpected end of data").
			WithMeta(core.EMK(0, "int"), length).
			WithMeta(core.EMK(1, "int"), *s.Idx).
			WithMeta(core.EMK(2, "int"), len(s.Code))
	}
	res := s.Code[*s.Idx : *s.Idx+length]
	*s.Idx += length
	return res, nil
}

// ShiftPanic reads 'length' bytes, panicking if not enough data.
// Use only in contexts where you are certain the buffer is large enough.
func (s *Shift) ShiftPanic(length int) []byte {
	bytes, err := s.ShiftError(length)
	if err != nil {
		panic("Can't continue shifting, error: " + err.Error())
	}
	return bytes
}

// ShiftFloat64Error reads 'size' bytes and interprets them as a float64 in [-1,1].
// If there is not enough data, returns a core.Error.
//
// Err errors.BytecodeShiftError (wrapped from ShiftError).
func (s *Shift) ShiftFloat64Error(size int, endianess public.EndianType) (float64, core.ErrorInterface) {
	bytes, err := s.ShiftError(size)
	if err != nil {
		return 0, err
	}
	utils := &Utils{}
	return utils.BytesToFloat64(bytes, endianess), nil
}

// ShiftFloat64Panic reads and converts a float64, panicking on error.
func (s *Shift) ShiftFloat64Panic(size int, endianess public.EndianType) float64 {
	bytes := s.ShiftPanic(size)
	utils := &Utils{}
	return utils.BytesToFloat64(bytes, endianess)
}

// ShiftFloat64RangeError reads 'size' bytes and interprets them as a float64 in [minVal, maxVal].
// If there is not enough data, returns a core.Error.
//
// Err errors.BytecodeShiftError (wrapped from ShiftError).
func (s *Shift) ShiftFloat64RangeError(size int, minVal, maxVal float64, endianess public.EndianType) (float64, core.ErrorInterface) {
	bytes, err := s.ShiftError(size)
	if err != nil {
		return 0, err
	}
	utils := &Utils{}
	return utils.BytesToFloat64Range(bytes, minVal, maxVal, endianess), nil
}

// ShiftFloat64RangePanic reads and converts a float64 with range, panicking on error.
func (s *Shift) ShiftFloat64RangePanic(size int, minVal, maxVal float64, endianess public.EndianType) float64 {
	bytes := s.ShiftPanic(size)
	utils := &Utils{}
	return utils.BytesToFloat64Range(bytes, minVal, maxVal, endianess)
}

// AutoIntToBytes chooses the minimal number of bytes needed to represent 'value'
// (excluding sign extension) and converts it.
func (u *Utils) AutoIntToBytes(value int, endianess public.EndianType) []byte {
	size := 1
	temp := value
	if temp < 0 {
		temp = -temp
	}
	for temp > 0xFF {
		temp >>= 8
		size++
	}
	if value < 0 {
		size = 8
	}
	return u.IntToBytes(value, size, endianess)
}

// AutoFloat64ToBytes chooses an optimal byte size based on the magnitude of 'value',
// then converts it to bytes.
func (u *Utils) AutoFloat64ToBytes(value float64, endianess public.EndianType) []byte {
	size := 1
	absValue := math.Abs(value)
	if absValue > 1.0 {
		if absValue <= 2 {
			size = 2
		} else if absValue <= 4 {
			size = 3
		} else if absValue <= 8 {
			size = 4
		} else {
			size = 8
		}
	} else {
		if absValue == 0 {
			size = 1
		} else if absValue < 0.01 {
			size = 4
		} else if absValue < 0.1 {
			size = 3
		} else if absValue < 0.5 {
			size = 2
		}
	}
	return u.Float64ToBytes(value, size, endianess)
}
```

---

# tooling/debugging/extensiblePlugin/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# extensiblePlugin

'''go
import "github.com/pt-main/lc/tooling/debugging/extensiblePlugin"
'''

## Index

- [Constants](<#constants>)
- [type BCLEData](<#BCLEData>)
- [type CLEData](<#CLEData>)
- [type ExtensibleCLPlugin](<#ExtensibleCLPlugin>)
  - [func New\(eu \*lc.EngineUniversal\) \*ExtensibleCLPlugin](<#New>)
  - [func \(ep \*ExtensibleCLPlugin\) ByteCallHotLoopEvent\(ev \*core.Events, i \*core.EventInput\) \(err core.ErrorInterface\)](<#ExtensibleCLPlugin.ByteCallHotLoopEvent>)
  - [func \(ep \*ExtensibleCLPlugin\) Call\(string, ...core.Option\) \(o any, e error\)](<#ExtensibleCLPlugin.Call>)
  - [func \(ep \*ExtensibleCLPlugin\) Close\(\) error](<#ExtensibleCLPlugin.Close>)
  - [func \(ep \*ExtensibleCLPlugin\) Init\(scope core.ScopeType, pm \*plugin.PluginManager\) error](<#ExtensibleCLPlugin.Init>)
  - [func \(ep \*ExtensibleCLPlugin\) Name\(\) string](<#ExtensibleCLPlugin.Name>)
  - [func \(ep \*ExtensibleCLPlugin\) Run\(input any\) \(o any, e error\)](<#ExtensibleCLPlugin.Run>)
  - [func \(ep \*ExtensibleCLPlugin\) StringCallLoopEvent\(ev \*core.Events, i \*core.EventInput\) \(err core.ErrorInterface\)](<#ExtensibleCLPlugin.StringCallLoopEvent>)
- [type SCLEData](<#SCLEData>)


## Constants

<a name="CLEPreEvent"></a>Сalloop events

'''go
const (
    CLEPreEvent    = "CalloopE PreEvent"
    CLEInPreEvent  = "CalloopE InPreEvent"
    CLEInPostEvent = "CalloopE InPostEvent"
    CLEPostEvent   = "CalloopE PostEvent"
)
'''

<a name="CLEScopeData"></a>

'''go
const (
    CLEScopeData = "ExtensiblePlugin ScopeData CalloopE Data" // Сalloop data (CLEData)
)
'''

<a name="ECLFlag"></a>

'''go
const (
    ECLFlag = "ExtensiblePlugin" // This flag will be added to plugins manager when the plugin is initialized
)
'''

<a name="Name"></a>

'''go
const Name = "extensible call loop"
'''

<a name="BCLEData"></a>
## type [BCLEData](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/structs.go#L26-L30>)



'''go
type BCLEData CLEData[
    events.ByteCLDType,
    events.ByteCallAttr,
    engine.ByteEngineInterface,
]
'''

<a name="CLEData"></a>
## type [CLEData](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/structs.go#L11-L18>)



'''go
type CLEData[I, P, E any] struct {
    Input  I
    Idx    *int
    Parsed []P
    PLen   int
    Ctx    context.Context
    E      E
}
'''

<a name="ExtensibleCLPlugin"></a>
## type [ExtensibleCLPlugin](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/plugin.go#L19-L25>)

### ExtensibleCLPlugin

Replace standart calloops and add calloops with event hooks. System plugin.

'''go
type ExtensibleCLPlugin struct {
    Eu     *lc.EngineUniversal
    Events core.EventsInterface
    ETools core.EventsTools
    WasE   core.EventType
    // contains filtered or unexported fields
}
'''

<a name="New"></a>
### func [New](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/plugin.go#L27>)

'''go
func New(eu *lc.EngineUniversal) *ExtensibleCLPlugin
'''



<a name="ExtensibleCLPlugin.ByteCallHotLoopEvent"></a>
### func \(\*ExtensibleCLPlugin\) [ByteCallHotLoopEvent](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/events.go#L71>)

'''go
func (ep *ExtensibleCLPlugin) ByteCallHotLoopEvent(ev *core.Events, i *core.EventInput) (err core.ErrorInterface)
'''

ByteCallHotLoopEvent is the main loop for bytecode execution. It wraps the standard byte‑call loop with pre‑ and post‑iteration hooks.

Err errors.ExtensiblePluginError:

- If the input is not of type ByteCLDType. Meta: EMK\(0, "string"\) – the actual type \(if available\).
- If core.ScopeGet fails to retrieve BCLEData.
- If ByteCallEventIteration fails.
- Meta: EMK\(0, "int"\) – bytecode index where the error occurred.

<a name="ExtensibleCLPlugin.Call"></a>
### func \(\*ExtensibleCLPlugin\) [Call](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/plugin.go#L108>)

'''go
func (ep *ExtensibleCLPlugin) Call(string, ...core.Option) (o any, e error)
'''



<a name="ExtensibleCLPlugin.Close"></a>
### func \(\*ExtensibleCLPlugin\) [Close](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/plugin.go#L100>)

'''go
func (ep *ExtensibleCLPlugin) Close() error
'''



<a name="ExtensibleCLPlugin.Init"></a>
### func \(\*ExtensibleCLPlugin\) [Init](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/plugin.go#L77>)

'''go
func (ep *ExtensibleCLPlugin) Init(scope core.ScopeType, pm *plugin.PluginManager) error
'''



<a name="ExtensibleCLPlugin.Name"></a>
### func \(\*ExtensibleCLPlugin\) [Name](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/plugin.go#L98>)

'''go
func (ep *ExtensibleCLPlugin) Name() string
'''



<a name="ExtensibleCLPlugin.Run"></a>
### func \(\*ExtensibleCLPlugin\) [Run](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/plugin.go#L112>)

'''go
func (ep *ExtensibleCLPlugin) Run(input any) (o any, e error)
'''



<a name="ExtensibleCLPlugin.StringCallLoopEvent"></a>
### func \(\*ExtensibleCLPlugin\) [StringCallLoopEvent](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/events.go#L16>)

'''go
func (ep *ExtensibleCLPlugin) StringCallLoopEvent(ev *core.Events, i *core.EventInput) (err core.ErrorInterface)
'''

StringCallLoopEvent is the main loop for string\-based command execution. It wraps the standard events with pre‑ and post‑iteration hooks.

Err errors.ExtensiblePluginError:

- If core.ScopeGet fails to retrieve SCLEData.
- If StringCallEventIteration fails.
- Meta: EMK\(0, "string"\) – the raw line that caused the error \(if available\).

<a name="SCLEData"></a>
## type [SCLEData](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/extensiblePlugin/structs.go#L20-L24>)



'''go
type SCLEData CLEData[
    events.StringCLDType,
    stringParsing.ParsedNode,
    engine.StringEngineInterface,
]
'''

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# tooling/debugging/extensiblePlugin/config.go

```go
package extensiblePlugin

// Сalloop events
const (
	CLEPreEvent    = "CalloopE PreEvent"
	CLEInPreEvent  = "CalloopE InPreEvent"
	CLEInPostEvent = "CalloopE InPostEvent"
	CLEPostEvent   = "CalloopE PostEvent"
)

const (
	CLEScopeData = "ExtensiblePlugin ScopeData CalloopE Data" // Сalloop data (CLEData)
)

const (
	ECLFlag = "ExtensiblePlugin" // This flag will be added to plugins manager when the plugin is initialized
)
```

---

# tooling/debugging/extensiblePlugin/events.go

```go
package extensiblePlugin

import (
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/engine/events"
	"github.com/pt-main/lc/public/errors"
)

// StringCallLoopEvent is the main loop for string-based command execution.
// It wraps the standard events with pre‑ and post‑iteration hooks.
//
// Err errors.ExtensiblePluginError:
//   - If core.ScopeGet fails to retrieve SCLEData.
//   - If StringCallEventIteration fails.
//   - Meta: EMK(0, "string") – the raw line that caused the error (if available).
func (ep *ExtensibleCLPlugin) StringCallLoopEvent(ev *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	cld := i.Input.(events.StringCLDType)
	idx := cld.Idx
	parsed := cld.Parsed
	pLen := len(parsed)
	ctx := cld.Ctx
	e := cld.Engine
	ep.Events.Scope()[CLEScopeData] = SCLEData{
		Input:  cld,
		Idx:    idx,
		Parsed: parsed,
		PLen:   pLen,
		Ctx:    ctx,
		E:      e,
	}
	ep.Events.CallEvents(nil, CLEPreEvent, true) // ignore error (canWorkWithoutHandler)
	sd, err := core.ScopeGet[SCLEData](ep.Events.Scope(), CLEScopeData)
	if err != nil {
		return core.Wrap(errors.ExtensiblePluginError, err, "Failed to retrieve SCLEData from scope").
			WithMeta(core.EMK(0, "string"), "SCLEData")
	}
	cld = sd.Input
	idx = sd.Idx
	parsed = sd.Parsed
	pLen = sd.PLen
	ctx = sd.Ctx
	e = sd.E
	for *idx < pLen && *idx >= 0 {
		ep.Events.CallEvents(nil, CLEInPreEvent, true) // ignore error
		err = ep.de.StringCallEventIteration(parsed, idx, ev, ctx, e)
		if err != nil {
			// Обогащаем ошибку, если это core.Error
			if ce, ok := err.(*core.Error); ok {
				ce.WithMeta(core.EMK(0, "string"), "StringCallLoopEvent")
			} else {
				err = core.Wrap(errors.ExtensiblePluginError, err, "StringCallEventIteration failed").
					WithMeta(core.EMK(0, "string"), "StringCallLoopEvent")
			}
			return err
		}
		ep.Events.CallEvents(nil, CLEInPostEvent, true) // ignore error
	}
	ep.Events.CallEvents(nil, CLEPostEvent, true) // ignore error
	return nil
}

// ByteCallHotLoopEvent is the main loop for bytecode execution.
// It wraps the standard byte‑call loop with pre‑ and post‑iteration hooks.
//
// Err errors.ExtensiblePluginError:
//   - If the input is not of type ByteCLDType.
//     Meta: EMK(0, "string") – the actual type (if available).
//   - If core.ScopeGet fails to retrieve BCLEData.
//   - If ByteCallEventIteration fails.
//   - Meta: EMK(0, "int") – bytecode index where the error occurred.
func (ep *ExtensibleCLPlugin) ByteCallHotLoopEvent(ev *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	hld, ok := i.Input.(events.ByteCLDType)
	if !ok {
		return core.Err(errors.ExtensiblePluginError, "Invalid event input: expected ByteCLDType").
			WithMeta(core.EMK(0, "string"), "ByteCLDType")
	}
	idx := hld.Idx
	ctx := hld.Ctx
	parsed := hld.Parsed
	p2len := len(parsed)
	e := hld.Engine
	ep.Events.Scope()[CLEScopeData] = BCLEData{
		Input:  hld,
		Idx:    idx,
		Parsed: parsed,
		PLen:   p2len,
		Ctx:    ctx,
		E:      e,
	}
	ep.Events.CallEvents(nil, CLEPreEvent, true) // ignore error
	sd, err := core.ScopeGet[BCLEData](ep.Events.Scope(), CLEScopeData)
	if err != nil {
		return core.Wrap(errors.ExtensiblePluginError, err, "Failed to retrieve BCLEData from scope").
			WithMeta(core.EMK(0, "string"), "BCLEData")
	}
	hld = sd.Input
	idx = sd.Idx
	parsed = sd.Parsed
	p2len = sd.PLen
	ctx = sd.Ctx
	e = sd.E
	for *idx < p2len && *idx >= 0 {
		ep.Events.CallEvents(nil, CLEInPreEvent, true) // ignore error
		if ctx.Err() != nil {
			err = core.Wrap(errors.ExtensiblePluginError, ctx.Err(), "Context cancelled during loop").
				WithMeta(core.EMK(0, "int"), *idx)
			break
		}
		err = ep.de.ByteCallEventIteration(idx, &parsed[*idx], e)
		if err != nil {
			// Обогащаем ошибку
			if ce, ok := err.(*core.Error); ok {
				ce.WithMeta(core.EMK(0, "int"), *idx)
			} else {
				err = core.Wrap(errors.ExtensiblePluginError, err, "ByteCallEventIteration failed at index %d", *idx).
					WithMeta(core.EMK(0, "int"), *idx)
			}
			break
		}
		ep.Events.CallEvents(nil, CLEInPostEvent, true) // ignore error
	}
	ep.Events.CallEvents(nil, CLEPostEvent, true) // ignore error
	return
}
```

---

# tooling/debugging/extensiblePlugin/plugin.go

```go
package extensiblePlugin

import (
	"fmt"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/engine/events"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/plugin"
)

const Name = "extensible call loop"

// # ExtensibleCLPlugin
//
// Replace standart calloops and add calloops with event hooks.
// System plugin.
type ExtensibleCLPlugin struct {
	de     events.DefaultEvents
	Eu     *lc.EngineUniversal
	Events core.EventsInterface
	ETools core.EventsTools
	WasE   core.EventType
}

func New(eu *lc.EngineUniversal) *ExtensibleCLPlugin {
	uep, err := eu.GetUEP()
	if err != nil {
		panic("Can't add extensible plugin: " + err.Error())
	}
	e := uep.Event
	return &ExtensibleCLPlugin{
		de:     events.DefaultEvents{},
		Eu:     eu,
		Events: e,
		ETools: core.EventsTools{
			Events: e,
		},
		WasE: nil,
	}
}

func (ep *ExtensibleCLPlugin) changeEvents(val bool) (string, error) {
	euType := ep.Eu.Type
	var name string
	var event core.EventType
	switch euType {
	case public.StringEngineType:
		name = public.StringCallCalloopEvent
	case public.ByteEngineType:
		name = public.ByteCallHotloopEvent
	}
	switch val {
	case true:
		var err error
		ep.WasE, err = ep.ETools.GetCoreEvent(name)
		if err != nil {
			return "", err
		}
		switch euType {
		case public.StringEngineType:
			event = ep.StringCallLoopEvent
		case public.ByteEngineType:
			event = ep.ByteCallHotLoopEvent
		}
	default:
		event = ep.WasE
	}
	err := ep.ETools.ChangeCoreEvent(name, event)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (ep *ExtensibleCLPlugin) Init(scope core.ScopeType, pm *plugin.PluginManager) error {
	eu, ok := scope[public.PluginsScopeEuPtr].(*lc.EngineUniversal)
	if !ok {
		return fmt.Errorf("Bad scope: can't find EngineUniversal. Plugins didn't load.")
	}
	ep.Eu = eu
	uep, err := eu.GetUEP()
	if err != nil {
		return err
	}
	_, err = ep.changeEvents(true)
	if err != nil {
		return err
	}
	ep.ETools = core.EventsTools{
		Events: uep.Event,
	}
	(&plugin.Tools{Pm: pm}).SetFlag(ECLFlag)
	return nil
}

func (ep *ExtensibleCLPlugin) Name() string { return Name }

func (ep *ExtensibleCLPlugin) Close() error {
	_, err := ep.changeEvents(false)
	if err != nil {
		return err
	}
	return nil
}

func (ep *ExtensibleCLPlugin) Call(string, ...core.Option) (o any, e error) {
	return
}

func (ep *ExtensibleCLPlugin) Run(input any) (o any, e error) {
	return
}
```

---

# tooling/debugging/extensiblePlugin/structs.go

```go
package extensiblePlugin

import (
	"context"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/events"
	"github.com/pt-main/lc/parsing/stringParsing"
)

type CLEData[I, P, E any] struct {
	Input  I
	Idx    *int
	Parsed []P
	PLen   int
	Ctx    context.Context
	E      E
}

type SCLEData CLEData[
	events.StringCLDType,
	stringParsing.ParsedNode,
	engine.StringEngineInterface,
]

type BCLEData CLEData[
	events.ByteCLDType,
	events.ByteCallAttr,
	engine.ByteEngineInterface,
]
```

---

# tooling/debugging/profiler/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# profiler

'''go
import "github.com/pt-main/lc/tooling/debugging/profiler"
'''

## Index

- [Constants](<#constants>)
- [type Metric](<#Metric>)
  - [func \(m \*Metric\) Avg\(\) time.Duration](<#Metric.Avg>)
- [type ProfilerPlugin](<#ProfilerPlugin>)
  - [func New\(\) \*ProfilerPlugin](<#New>)
  - [func \(p \*ProfilerPlugin\) Call\(name string, opts ...core.Option\) \(any, error\)](<#ProfilerPlugin.Call>)
  - [func \(p \*ProfilerPlugin\) Close\(\) error](<#ProfilerPlugin.Close>)
  - [func \(p \*ProfilerPlugin\) Disable\(\)](<#ProfilerPlugin.Disable>)
  - [func \(p \*ProfilerPlugin\) Enable\(\)](<#ProfilerPlugin.Enable>)
  - [func \(p \*ProfilerPlugin\) Init\(scope core.ScopeType, pm \*plugin.PluginManager\) error](<#ProfilerPlugin.Init>)
  - [func \(p \*ProfilerPlugin\) Name\(\) string](<#ProfilerPlugin.Name>)
  - [func \(p \*ProfilerPlugin\) Report\(\) string](<#ProfilerPlugin.Report>)
  - [func \(p \*ProfilerPlugin\) Reset\(\)](<#ProfilerPlugin.Reset>)
  - [func \(p \*ProfilerPlugin\) Run\(input any\) \(any, error\)](<#ProfilerPlugin.Run>)


## Constants

<a name="Name"></a>

'''go
const Name = "profiler"
'''

<a name="Metric"></a>
## type [Metric](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L20-L25>)



'''go
type Metric struct {
    Count     int64
    TotalTime time.Duration
    MinTime   time.Duration
    MaxTime   time.Duration
}
'''

<a name="Metric.Avg"></a>
### func \(\*Metric\) [Avg](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L27>)

'''go
func (m *Metric) Avg() time.Duration
'''



<a name="ProfilerPlugin"></a>
## type [ProfilerPlugin](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L48-L60>)

### ProfilerPlugin

Connects to EngineUniversal, requires ExtensibleCLPlugin.

### Methods:

'''
Call("report") -> (string, error) // Return report saved from last "reset"

Call("reset") -> ("reset done", error)

Call("enable") -> ("enabled", error)

Call("disable") -> ("disabled", error)
'''

'''go
type ProfilerPlugin struct {
    plugin.Plugin
    // contains filtered or unexported fields
}
'''

<a name="New"></a>
### func [New](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L62>)

'''go
func New() *ProfilerPlugin
'''



<a name="ProfilerPlugin.Call"></a>
### func \(\*ProfilerPlugin\) [Call](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L237>)

'''go
func (p *ProfilerPlugin) Call(name string, opts ...core.Option) (any, error)
'''



<a name="ProfilerPlugin.Close"></a>
### func \(\*ProfilerPlugin\) [Close](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L233>)

'''go
func (p *ProfilerPlugin) Close() error
'''



<a name="ProfilerPlugin.Disable"></a>
### func \(\*ProfilerPlugin\) [Disable](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L231>)

'''go
func (p *ProfilerPlugin) Disable()
'''



<a name="ProfilerPlugin.Enable"></a>
### func \(\*ProfilerPlugin\) [Enable](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L230>)

'''go
func (p *ProfilerPlugin) Enable()
'''



<a name="ProfilerPlugin.Init"></a>
### func \(\*ProfilerPlugin\) [Init](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L75>)

'''go
func (p *ProfilerPlugin) Init(scope core.ScopeType, pm *plugin.PluginManager) error
'''



<a name="ProfilerPlugin.Name"></a>
### func \(\*ProfilerPlugin\) [Name](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L71>)

'''go
func (p *ProfilerPlugin) Name() string
'''



<a name="ProfilerPlugin.Report"></a>
### func \(\*ProfilerPlugin\) [Report](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L180>)

'''go
func (p *ProfilerPlugin) Report() string
'''



<a name="ProfilerPlugin.Reset"></a>
### func \(\*ProfilerPlugin\) [Reset](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L218>)

'''go
func (p *ProfilerPlugin) Reset()
'''



<a name="ProfilerPlugin.Run"></a>
### func \(\*ProfilerPlugin\) [Run](<https://github.com/pt-main/Lc/blob/main/tooling/debugging/profiler/main.go#L255>)

'''go
func (p *ProfilerPlugin) Run(input any) (any, error)
'''



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# tooling/debugging/profiler/main.go

```go
package profiler

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/bytecode"
	"github.com/pt-main/lc/tooling/debugging/extensiblePlugin"
	"github.com/pt-main/lc/tooling/plugin"
)

const Name = "profiler"

type Metric struct {
	Count     int64
	TotalTime time.Duration
	MinTime   time.Duration
	MaxTime   time.Duration
}

func (m *Metric) Avg() time.Duration {
	if m.Count == 0 {
		return 0
	}
	return m.TotalTime / time.Duration(m.Count)
}

// # ProfilerPlugin
//
// Connects to EngineUniversal,
// requires ExtensibleCLPlugin.
//
// # Methods:
//
//	Call("report") -> (string, error) // Return report saved from last "reset"
//
//	Call("reset") -> ("reset done", error)
//
//	Call("enable") -> ("enabled", error)
//
//	Call("disable") -> ("disabled", error)
type ProfilerPlugin struct {
	plugin.Plugin
	mu               sync.Mutex
	byteMetrics      map[int]*Metric
	stringMetrics    map[string]*Metric
	totalByteCalls   int64
	totalStringCalls int64
	totalByteTime    time.Duration
	totalStringTime  time.Duration
	startTime        time.Time
	enabled          bool
	scope            core.ScopeType
}

func New() *ProfilerPlugin {
	return &ProfilerPlugin{
		byteMetrics:   make(map[int]*Metric),
		stringMetrics: make(map[string]*Metric),
		startTime:     time.Now(),
		enabled:       true,
	}
}

func (p *ProfilerPlugin) Name() string {
	return Name
}

func (p *ProfilerPlugin) Init(scope core.ScopeType, pm *plugin.PluginManager) error {
	if !(&plugin.Tools{Pm: pm}).IsPluginInstaled(extensiblePlugin.Name) {
		return fmt.Errorf("Extensible plugin is not installed, can't init profiler")
	}
	p.scope = scope
	eu, ok := scope[public.PluginsScopeEuPtr].(*lc.EngineUniversal)
	if !ok {
		return fmt.Errorf("profiler: can't get EngineUniversal")
	}
	uep, err := eu.GetUEP()
	if err != nil {
		return err
	}
	ep := extensiblePlugin.New(eu)
	err = ep.Init(scope, pm)
	if err != nil {
		return err
	}
	uep.Event.NewEvent(extensiblePlugin.CLEInPreEvent, p.preEvent)
	uep.Event.NewEvent(extensiblePlugin.CLEInPostEvent, p.postEvent)
	return nil
}

func (p *ProfilerPlugin) preEvent(ev *core.Events, i *core.EventInput) core.ErrorInterface {
	if !p.enabled {
		return nil
	}

	ev.Scope()["profiler_start_time"] = time.Now()
	return nil
}

func (p *ProfilerPlugin) postEvent(ev *core.Events, i *core.EventInput) core.ErrorInterface {
	if !p.enabled {
		return nil
	}
	startVal, ok := ev.Scope()["profiler_start_time"]
	if !ok {
		return nil
	}
	start, ok := startVal.(time.Time)
	if !ok {
		return nil
	}
	elapsed := time.Since(start)

	p.mu.Lock()
	defer p.mu.Unlock()

	if data, err := core.ScopeGet[extensiblePlugin.BCLEData](ev.Scope(), extensiblePlugin.CLEScopeData); err == nil {
		p.totalByteCalls++
		p.totalByteTime += elapsed
		idx := data.Idx
		if idx != nil && *idx >= 0 && *idx < len(data.Parsed) {
			attr := data.Parsed[*idx]
			opcode := 0
			if attr.RawNode != nil {
				opcode = int((&bytecode.Utils{}).BytesToInt(attr.RawNode.Switch, public.LittleEndian))
			}
			p.updateMetricByte(opcode, elapsed)
		}
	} else if data, err := core.ScopeGet[extensiblePlugin.SCLEData](ev.Scope(), extensiblePlugin.CLEScopeData); err == nil {
		p.totalStringCalls++
		p.totalStringTime += elapsed
		idx := data.Idx
		if idx != nil && *idx >= 0 && *idx < len(data.Parsed) {
			node := data.Parsed[*idx]
			p.updateMetricString(node.Switch, elapsed)
		}
	}
	return nil
}

func (p *ProfilerPlugin) updateMetricByte(opcode int, elapsed time.Duration) {
	m, ok := p.byteMetrics[opcode]
	if !ok {
		m = &Metric{MinTime: elapsed, MaxTime: elapsed}
		p.byteMetrics[opcode] = m
	}
	m.Count++
	m.TotalTime += elapsed
	if elapsed < m.MinTime {
		m.MinTime = elapsed
	}
	if elapsed > m.MaxTime {
		m.MaxTime = elapsed
	}
}

func (p *ProfilerPlugin) updateMetricString(cmd string, elapsed time.Duration) {
	m, ok := p.stringMetrics[cmd]
	if !ok {
		m = &Metric{MinTime: elapsed, MaxTime: elapsed}
		p.stringMetrics[cmd] = m
	}
	m.Count++
	m.TotalTime += elapsed
	if elapsed < m.MinTime {
		m.MinTime = elapsed
	}
	if elapsed > m.MaxTime {
		m.MaxTime = elapsed
	}
}

func (p *ProfilerPlugin) Report() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	elapsed := time.Since(p.startTime)
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Profiler report (%.2f sec total):\n", elapsed.Seconds()))

	if len(p.byteMetrics) > 0 {
		sb.WriteString("\n  Byte opcodes:\n")
		sb.WriteString(fmt.Sprintf("  Byte calls: %d (total time: %v)\n", p.totalByteCalls, p.totalByteTime))
		var keys []int
		for k := range p.byteMetrics {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, op := range keys {
			m := p.byteMetrics[op]
			sb.WriteString(fmt.Sprintf("    %d: count=%d, total=%v, avg=%v, min=%v, max=%v\n",
				op, m.Count, m.TotalTime, m.Avg(), m.MinTime, m.MaxTime))
		}
	}
	if len(p.stringMetrics) > 0 {
		sb.WriteString("\n  String commands:\n")
		sb.WriteString(fmt.Sprintf("  String calls: %d (total time: %v)\n", p.totalStringCalls, p.totalStringTime))
		var keys []string
		for k := range p.stringMetrics {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, cmd := range keys {
			m := p.stringMetrics[cmd]
			sb.WriteString(fmt.Sprintf("    %s: count=%d, total=%v, avg=%v, min=%v, max=%v\n",
				cmd, m.Count, m.TotalTime, m.Avg(), m.MinTime, m.MaxTime))
		}
	}
	return sb.String()
}

func (p *ProfilerPlugin) Reset() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.byteMetrics = make(map[int]*Metric)
	p.stringMetrics = make(map[string]*Metric)
	p.totalByteCalls = 0
	p.totalStringCalls = 0
	p.totalByteTime = 0
	p.totalStringTime = 0
	p.startTime = time.Now()
}

func (p *ProfilerPlugin) Enable()  { p.enabled = true }
func (p *ProfilerPlugin) Disable() { p.enabled = false }

func (p *ProfilerPlugin) Close() error {
	return nil
}

func (p *ProfilerPlugin) Call(name string, opts ...core.Option) (any, error) {
	switch name {
	case "report":
		return p.Report(), nil
	case "reset":
		p.Reset()
		return "reset done", nil
	case "enable":
		p.Enable()
		return "enabled", nil
	case "disable":
		p.Disable()
		return "disabled", nil
	default:
		return nil, fmt.Errorf("unknown call: %s", name)
	}
}

func (p *ProfilerPlugin) Run(input any) (any, error) {
	return nil, nil
}
```

---

# tooling/plugin/GODOC.md

```md
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# plugin

'''go
import "github.com/pt-main/lc/tooling/plugin"
'''

## Index

- [type Plugin](<#Plugin>)
  - [func NewPlugin\(name, initEvent, mainEvent, closeEvent, scopeRunResultKey, scopeCallResultKey string, context context.Context\) \*Plugin](<#NewPlugin>)
  - [func \(p \*Plugin\) Call\(name string, opts ...core.Option\) \(any, error\)](<#Plugin.Call>)
  - [func \(p \*Plugin\) Close\(\) error](<#Plugin.Close>)
  - [func \(p \*Plugin\) Init\(scope core.ScopeType, pm \*PluginManager\) error](<#Plugin.Init>)
  - [func \(p \*Plugin\) Name\(\) string](<#Plugin.Name>)
  - [func \(p \*Plugin\) Run\(input any\) \(any, error\)](<#Plugin.Run>)
- [type PluginInterface](<#PluginInterface>)
- [type PluginManager](<#PluginManager>)
  - [func NewPluginManager\(scope core.ScopeType\) \*PluginManager](<#NewPluginManager>)
  - [func \(pm \*PluginManager\) AddPlugin\(plugin PluginInterface\) error](<#PluginManager.AddPlugin>)
  - [func \(pm \*PluginManager\) CallPluginMethod\(name, method string, opts ...core.Option\) \(any, error\)](<#PluginManager.CallPluginMethod>)
  - [func \(pm \*PluginManager\) DeletePlugin\(name string\) error](<#PluginManager.DeletePlugin>)
  - [func \(pm \*PluginManager\) End\(\) \(err error\)](<#PluginManager.End>)
  - [func \(pm \*PluginManager\) GetPlugin\(name string\) \(PluginInterface, error\)](<#PluginManager.GetPlugin>)
  - [func \(pm \*PluginManager\) RunPlugin\(name string, input any\) \(any, error\)](<#PluginManager.RunPlugin>)
- [type Tools](<#Tools>)
  - [func \(t \*Tools\) HasFlag\(f string\) bool](<#Tools.HasFlag>)
  - [func \(t \*Tools\) IsPluginInstaled\(p string\) bool](<#Tools.IsPluginInstaled>)
  - [func \(t \*Tools\) SetFlag\(f string\)](<#Tools.SetFlag>)


<a name="Plugin"></a>
## type [Plugin](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/realization.go#L15-L25>)

### Plugin

PluginInterface realization base working on events.

Method calling calls evant named as method name from local Events engine. Name of plugin is constant and immutable.

'''go
type Plugin struct {
    Events             *core.Events
    ScopeRunResultKey  string
    ScopeCallResultKey string

    InitEvent  string
    CloseEvent string
    MainEvent  string
    // contains filtered or unexported fields
}
'''

<a name="NewPlugin"></a>
### func [NewPlugin](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/realization.go#L27-L33>)

'''go
func NewPlugin(name, initEvent, mainEvent, closeEvent, scopeRunResultKey, scopeCallResultKey string, context context.Context) *Plugin
'''



<a name="Plugin.Call"></a>
### func \(\*Plugin\) [Call](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/realization.go#L75>)

'''go
func (p *Plugin) Call(name string, opts ...core.Option) (any, error)
'''



<a name="Plugin.Close"></a>
### func \(\*Plugin\) [Close](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/realization.go#L58>)

'''go
func (p *Plugin) Close() error
'''



<a name="Plugin.Init"></a>
### func \(\*Plugin\) [Init](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/realization.go#L49>)

'''go
func (p *Plugin) Init(scope core.ScopeType, pm *PluginManager) error
'''



<a name="Plugin.Name"></a>
### func \(\*Plugin\) [Name](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/realization.go#L45>)

'''go
func (p *Plugin) Name() string
'''



<a name="Plugin.Run"></a>
### func \(\*Plugin\) [Run](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/realization.go#L64>)

'''go
func (p *Plugin) Run(input any) (any, error)
'''



<a name="PluginInterface"></a>
## type [PluginInterface](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/interface.go#L5-L11>)



'''go
type PluginInterface interface {
    Name() string
    Init(scope core.ScopeType, pm *PluginManager) error
    Close() error
    Call(string, ...core.Option) (any, error)
    Run(input any) (any, error)
}
'''

<a name="PluginManager"></a>
## type [PluginManager](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/manager.go#L26-L30>)

### PluginManager

Plugin manager contains plugins, and scope with flags \(for plugins communicating\)

### Methods:

'''
AddPlugin(PluginInterface) // Add plugin to manager. Name will got from plugin.Name(). If plugin with same name was already registred -> error. Calling plugin.Init(Scope, *PluginManager), and return result.

DeletePlugin(string) // Delete plugin from manager. No return if plugin is not found. Calling plugin.Close(), and return result.

GetPlugin(string) // Get plugin. Return error if not found.

RunPlugin(string, any) // Run plugin. Call plugin.Run(input) and return result. Return error if plugin is not found.

CallPluginMethod(string, string, core.Option...) // Call plugin method. Call plugin.Call(method, opts...) and return result. Return error if plugin not found.

End() // End plugin lifecycle
'''

'''go
type PluginManager struct {
    Plugins map[string]PluginInterface
    Scope   core.ScopeType
    // contains filtered or unexported fields
}
'''

<a name="NewPluginManager"></a>
### func [NewPluginManager](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/manager.go#L37>)

'''go
func NewPluginManager(scope core.ScopeType) *PluginManager
'''

Create new plugin manager.

Args:

\- scope: scope \(of engine, or empty\), or nil

<a name="PluginManager.AddPlugin"></a>
### func \(\*PluginManager\) [AddPlugin](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/manager.go#L49>)

'''go
func (pm *PluginManager) AddPlugin(plugin PluginInterface) error
'''

Add plugin to manager. Name will got from plugin.Name\(\). If plugin with same name was already registred \-\> error. Calling plugin.Init\(Scope, \*PluginManager\), and return result.

<a name="PluginManager.CallPluginMethod"></a>
### func \(\*PluginManager\) [CallPluginMethod](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/manager.go#L90>)

'''go
func (pm *PluginManager) CallPluginMethod(name, method string, opts ...core.Option) (any, error)
'''

Call plugin method. Call plugin.Call\(method, opts...\) and return result. Return error if plugin not found.

<a name="PluginManager.DeletePlugin"></a>
### func \(\*PluginManager\) [DeletePlugin](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/manager.go#L60>)

'''go
func (pm *PluginManager) DeletePlugin(name string) error
'''

Delete plugin from manager. No return if plugin is not found. Calling plugin.Close\(\), and return result.

<a name="PluginManager.End"></a>
### func \(\*PluginManager\) [End](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/manager.go#L99>)

'''go
func (pm *PluginManager) End() (err error)
'''

End plugin lifecycle

<a name="PluginManager.GetPlugin"></a>
### func \(\*PluginManager\) [GetPlugin](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/manager.go#L72>)

'''go
func (pm *PluginManager) GetPlugin(name string) (PluginInterface, error)
'''

Get plugin. Return error if not found.

<a name="PluginManager.RunPlugin"></a>
### func \(\*PluginManager\) [RunPlugin](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/manager.go#L81>)

'''go
func (pm *PluginManager) RunPlugin(name string, input any) (any, error)
'''

Run plugin. Call plugin.Run\(input\) and return result. Return error if plugin is not found.

<a name="Tools"></a>
## type [Tools](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/tools.go#L8-L10>)

### Tools

PluginManager tools for simple plugins using

'''go
type Tools struct {
    Pm *PluginManager
}
'''

<a name="Tools.HasFlag"></a>
### func \(\*Tools\) [HasFlag](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/tools.go#L12>)

'''go
func (t *Tools) HasFlag(f string) bool
'''



<a name="Tools.IsPluginInstaled"></a>
### func \(\*Tools\) [IsPluginInstaled](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/tools.go#L20>)

'''go
func (t *Tools) IsPluginInstaled(p string) bool
'''



<a name="Tools.SetFlag"></a>
### func \(\*Tools\) [SetFlag](<https://github.com/pt-main/Lc/blob/main/tooling/plugin/tools.go#L16>)

'''go
func (t *Tools) SetFlag(f string)
'''



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
```

---

# tooling/plugin/interface.go

```go
package plugin

import "github.com/pt-main/lc/engine/core"

type PluginInterface interface {
	Name() string
	Init(scope core.ScopeType, pm *PluginManager) error
	Close() error
	Call(string, ...core.Option) (any, error)
	Run(input any) (any, error)
}
```

---

# tooling/plugin/manager.go

```go
package plugin

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
)

// # PluginManager
//
// Plugin manager contains plugins, and scope with flags (for plugins communicating)
//
// # Methods:
//
//	AddPlugin(PluginInterface) // Add plugin to manager. Name will got from plugin.Name(). If plugin with same name was already registred -> error. Calling plugin.Init(Scope, *PluginManager), and return result.
//
//	DeletePlugin(string) // Delete plugin from manager. No return if plugin is not found. Calling plugin.Close(), and return result.
//
//	GetPlugin(string) // Get plugin. Return error if not found.
//
//	RunPlugin(string, any) // Run plugin. Call plugin.Run(input) and return result. Return error if plugin is not found.
//
//	CallPluginMethod(string, string, core.Option...) // Call plugin method. Call plugin.Call(method, opts...) and return result. Return error if plugin not found.
//
//	End() // End plugin lifecycle
type PluginManager struct {
	Plugins map[string]PluginInterface
	Scope   core.ScopeType
	flags   []string // You can work with flags with Tools
}

// Create new plugin manager.
//
// Args:
//
// - scope: scope (of engine, or empty), or nil
func NewPluginManager(scope core.ScopeType) *PluginManager {
	if scope == nil {
		scope = make(core.ScopeType)
	}
	return &PluginManager{
		Plugins: make(map[string]PluginInterface),
		Scope:   scope,
	}
}

// Add plugin to manager. Name will got from plugin.Name(). If plugin with same name was
// already registred -> error.  Calling plugin.Init(Scope, *PluginManager), and return result.
func (pm *PluginManager) AddPlugin(plugin PluginInterface) error {
	name := plugin.Name()
	if _, exists := pm.Plugins[name]; exists {
		return fmt.Errorf("Plugin %s already loaded", name)
	}
	pm.Plugins[name] = plugin
	return plugin.Init(pm.Scope, pm)
}

// Delete plugin from manager. No return if plugin is not found. Calling plugin.Close(), and
// return result.
func (pm *PluginManager) DeletePlugin(name string) error {
	if plugin, exists := pm.Plugins[name]; exists {
		err := plugin.Close()
		if err != nil {
			return err
		}
		delete(pm.Plugins, name)
	}
	return nil
}

// Get plugin. Return error if not found.
func (pm *PluginManager) GetPlugin(name string) (PluginInterface, error) {
	plugin, ok := pm.Plugins[name]
	if !ok {
		return nil, fmt.Errorf("Plugin %s not found", name)
	}
	return plugin, nil
}

// Run plugin. Call plugin.Run(input) and return result. Return error if plugin is not found.
func (pm *PluginManager) RunPlugin(name string, input any) (any, error) {
	plugin, err := pm.GetPlugin(name)
	if err != nil {
		return nil, err
	}
	return plugin.Run(input)
}

// Call plugin method. Call plugin.Call(method, opts...) and return result. Return error if plugin not found.
func (pm *PluginManager) CallPluginMethod(name, method string, opts ...core.Option) (any, error) {
	plugin, err := pm.GetPlugin(name)
	if err != nil {
		return nil, err
	}
	return plugin.Call(method, opts...)
}

// End plugin lifecycle
func (pm *PluginManager) End() (err error) {
	for plugin := range pm.Plugins {
		err = pm.DeletePlugin(plugin)
		if err != nil {
			return
		}
	}
	return
}
```

---

# tooling/plugin/realization.go

```go
package plugin

import (
	"context"

	"github.com/pt-main/lc/engine/core"
)

// # Plugin
//
// PluginInterface realization base working on events.
//
// Method calling calls evant named as method name from local Events engine.
// Name of plugin is constant and immutable.
type Plugin struct {
	Events             *core.Events
	ScopeRunResultKey  string
	ScopeCallResultKey string

	InitEvent  string
	CloseEvent string
	MainEvent  string

	name string
}

func NewPlugin(
	name, initEvent,
	mainEvent, closeEvent,
	scopeRunResultKey,
	scopeCallResultKey string,
	context context.Context,
) *Plugin {
	return &Plugin{
		Events:             core.NewEvents(context),
		name:               name,
		InitEvent:          initEvent,
		MainEvent:          mainEvent,
		CloseEvent:         closeEvent,
		ScopeRunResultKey:  scopeCallResultKey,
		ScopeCallResultKey: scopeCallResultKey,
	}
}

func (p *Plugin) Name() string {
	return p.name
}

func (p *Plugin) Init(scope core.ScopeType, pm *PluginManager) error {
	return p.Events.CallEvents(&core.EventInput{
		Input: pm,
		Option: &core.Option{
			Scope: scope,
		},
	}, p.InitEvent, true)
}

func (p *Plugin) Close() error {
	return p.Events.CallEvents(&core.EventInput{
		Input: p,
	}, p.CloseEvent, true)
}

func (p *Plugin) Run(input any) (any, error) {
	err := p.Events.CallEvents(&core.EventInput{
		Input: input,
	}, p.MainEvent, true)
	if err != nil {
		return nil, err
	}
	res, _ := core.ScopeGet[any](p.Events.Scope(), p.ScopeRunResultKey)
	return res, nil
}

func (p *Plugin) Call(name string, opts ...core.Option) (any, error) {
	err := p.Events.CallEvents(&core.EventInput{
		Input: opts,
	}, name, false)
	if err != nil {
		return nil, err
	}
	res, _ := core.ScopeGet[any](p.Events.Scope(), p.ScopeRunResultKey)
	return res, nil
}
```

---

# tooling/plugin/tools.go

```go
package plugin

import "slices"

// # Tools
//
// PluginManager tools for simple plugins using
type Tools struct {
	Pm *PluginManager
}

func (t *Tools) HasFlag(f string) bool {
	return slices.Contains(t.Pm.flags, f)
}

func (t *Tools) SetFlag(f string) {
	t.Pm.flags = append(t.Pm.flags, f)
}

func (t *Tools) IsPluginInstaled(p string) bool {
	_, ok := t.Pm.Plugins[p]
	return ok
}
```

---

