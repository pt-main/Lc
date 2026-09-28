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

// Parser1 decodes a binary stream according to a fixed-length field layout.
// Each bytecode instruction consists of:
//
//	command (CommandBytelen bytes)
//	argscount (ArgscountBytelen bytes)
//	for each argument: arglen (ArglenBytelen bytes) followed by arg data.
//
// Endianness (Little/BigEndian) is used to decode integer fields.
type Parser1 struct {
	Config Parser1Config
}

// Parse reads the byte slice and returns a slice of ParsedBytes.
// Each ParsedBytes contains the raw command bytes, the raw arguments,
// and the original slice of the whole instruction. The Shift utility
// is used internally for safe bounds checking.
//
// Err errors.ParsingError:
//   - On panic recovery. Meta: EMK(0, "string") - the panic value.
//   - On shift error (unexpected end of data). Meta: EMK(0, "int") - attempted length,
//     EMK(1, "int") - current byte index.
//   - On zero argument length. Meta: EMK(0, "int") - argument number.
//   - On any other parsing error. Meta: EMK(0, "int") - command byte index.
//
// The returned error always contains the command bytes, raw bytes, and byte index in metadata.
func (p *Parser1) Parse(code []byte, opts ...*parsing.ParseOption) (result []ParsedBytes, err core.ErrorInterface) {
	lastCmdSwitch := []byte{0}
	var raw []byte
	idxVal := 0
	idx := &idxVal

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
		if err == nil {
			return
		}
		// A failed parse must not hand back a half decoded stream.
		result = nil
		cmdStr := fmt.Sprintf("%v", lastCmdSwitch)
		if len(lastCmdSwitch) == 0 {
			cmdStr = "<unknown>"
		}
		rawStr := fmt.Sprintf("%v", raw)
		if raw == nil {
			rawStr = "<none>"
		}
		if ce, ok := err.(*core.Error); ok {
			if _, ok := ce.Meta[core.EMK(1, "string")]; !ok {
				ce.WithMeta(core.EMK(1, "string"), rawStr)
			}
			if _, ok := ce.Meta[core.EMK(2, "string")]; !ok {
				ce.WithMeta(core.EMK(2, "string"), cmdStr)
			}
			if _, ok := ce.Meta[core.EMK(3, "int")]; !ok {
				ce.WithMeta(core.EMK(3, "int"), *idx)
			}
			return
		}
		err = core.Wrap(errors.ParsingError, err, "Parsing error at cmd=%v, raw=%v, idx=%d", cmdStr, rawStr, *idx).
			WithMeta(core.EMK(0, "string"), cmdStr).
			WithMeta(core.EMK(1, "string"), rawStr).
			WithMeta(core.EMK(2, "int"), *idx)
	}()

	log := func(text string) {
		if len(opts) > 0 && opts[0] != nil && opts[0].UEP != nil {
			logger := opts[0].UEP.Logger
			if logger != nil {
				logger.PrintLog(public.LogParsing, text)
			}
		}
	}

	log("=========== START ===========")
	log(fmt.Sprintf("start parsing code: '%v'", code))
	log(fmt.Sprintf("config: %v", p.Config))

	u := bytecode.Utils{}
	shift := p.Config.Shifter.ShiftError
	p.Config.Shifter.Idx = idx
	p.Config.Shifter.Code = code

	for *idx < len(code) {
		idxStart := *idx
		command, err := shift(p.Config.GConfig.CommandBytelen)
		if err != nil {
			return nil, core.Wrap(errors.ParsingError, err, "Shift error while reading command").
				WithMeta(core.EMK(0, "int"), p.Config.GConfig.CommandBytelen).
				WithMeta(core.EMK(1, "int"), *idx)
		}
		argscountBytes, err := shift(p.Config.GConfig.ArgscountBytelen)
		if err != nil {
			return nil, core.Wrap(errors.ParsingError, err, "Shift error while reading argscount").
				WithMeta(core.EMK(0, "int"), p.Config.GConfig.ArgscountBytelen).
				WithMeta(core.EMK(1, "int"), *idx)
		}
		argscount := u.BytesToInt(argscountBytes, p.Config.GConfig.Endianness)
		if argscount < 0 {
			// BytesToInt sign-extends, so a count with the high bit set decodes
			// negative. It must be rejected instead of yielding a 0-arg node.
			return nil, core.Err(errors.ParsingError, "Negative argument count: %d", argscount).
				WithMeta(core.EMK(0, "int"), argscount)
		}
		args := [][]byte{}
		lastCmdSwitch = command
		log(fmt.Sprintf("cmd %v, argscount %v", command, argscount))

		for argNum := 0; argNum < argscount; argNum++ {
			arglenBytes, err := shift(p.Config.GConfig.ArglenBytelen)
			if err != nil {
				return nil, core.Wrap(errors.ParsingError, err, "Shift error while reading argument length").
					WithMeta(core.EMK(0, "int"), p.Config.GConfig.ArglenBytelen).
					WithMeta(core.EMK(1, "int"), *idx).
					WithMeta(core.EMK(2, "int"), argNum)
			}
			arglen := u.BytesToInt(arglenBytes, p.Config.GConfig.Endianness)
			log(fmt.Sprintf("arglen %v", arglen))
			if arglen == 0 {
				return nil, core.Err(errors.ParsingError, "Zero argument length").
					WithMeta(core.EMK(0, "int"), argNum)
			}
			arg, err := shift(arglen)
			if err != nil {
				return nil, core.Wrap(errors.ParsingError, err, "Shift error while reading argument data").
					WithMeta(core.EMK(0, "int"), arglen).
					WithMeta(core.EMK(1, "int"), *idx).
					WithMeta(core.EMK(2, "int"), argNum)
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
	log(fmt.Sprintf("end parsing code:\n %v", result))
	log("=========== END ===========")
	return result, nil
}

func (p *Parser1) String() string {
	return "lc/parsing/byteParsing/Parser1"
}
