package system

type Engine struct {
	Scope     ScopeType
	Commands  map[string]CommandMeta
	Generator Generator
	Event     Events
}

func (e *Engine) Process(input string) error {
	e.Scope["input_string"] = input
	err1 := e.Event.CallEvents(e, ParseEvent)
	if err1 != nil {
		return err1
	}
	err2 := e.Event.CallEvents(e, CallEvent)
	if err2 != nil {
		return err2
	}
	return nil
}

func (e *Engine) NewCommand(cmd_switch string, handler CommandType, doc string) {
	e.Commands[cmd_switch] = CommandMeta{
		Handler: handler,
		Doc:     doc,
	}
}
