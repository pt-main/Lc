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
	defer func() {
		if r := recover(); r != nil {
			err = core.Err(errors.ParsingError, "Panic recovered during parsing: %v", r).
				WithMeta(core.EMK(0, "string"), fmt.Sprintf("%v", r))
		}
		if err != nil {
			// Обогащаем ошибку контекстом, если это ещё не core.Error или не содержит нужных метаданных
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
			// Проверяем, является ли err уже core.Error, если да, добавляем метаданные, если их нет
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
				// Обёртываем в core.Error, если это не core.Error
				err = core.Wrap(errors.ParsingError, err, "Parsing error at cmd=%v, raw=%v, idx=%d", cmdStr, rawStr, idxVal).
					WithMeta(core.EMK(0, "string"), cmdStr).
					WithMeta(core.EMK(1, "string"), rawStr).
					WithMeta(core.EMK(2, "int"), idxVal)
			}
		}
	}()
	log := func(text string) {
		if len(opts) > 0 {
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
			// shift error — оборачиваем с метаданными позиции
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
			if arglen == 0 {
				err = core.Err(errors.ParsingError, "Cannot form args with 0 argument length for argument #%d", argNum).
					WithMeta(core.EMK(0, "int"), argNum)
				return
			}
			log(fmt.Sprintf("arglen %v", arglen))
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
