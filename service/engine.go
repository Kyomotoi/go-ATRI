package service

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

// NewService 注册一个服务
func NewService(name string, docs string, onlyAdmin bool, mainCommand string, rules ...zero.Rule) *Service {
	return &Service{
		Service:     name,
		Docs:        docs,
		OnlyAdmin:   onlyAdmin,
		MainCommand: mainCommand,
		Rule:        rules,
	}
}

func (s *Service) OnMessage(name string, docs string, rules ...zero.Rule) *zero.Matcher {
	cmd_list := LoadCommandList(s.Service)
	cmd_list[name+"-onmsg"] = CommandInfo{
		Type: "message",
		Docs: docs,
	}

	StoneCommandList(s.Service, cmd_list)

	return zero.On("message", rules...)
}

func (s *Service) OnNotice(name string, docs string, rules ...zero.Rule) *zero.Matcher {
	cmd_list := LoadCommandList(s.Service)
	cmd_list[name+"-onntc"] = CommandInfo{
		Type: "notice",
		Docs: docs,
	}

	StoneCommandList(s.Service, cmd_list)

	return zero.On("notice", rules...)
}

func (s *Service) OnRequest(name string, docs string, rules ...zero.Rule) *zero.Matcher {
	cmd_list := LoadCommandList(s.Service)
	cmd_list[name+"-onreq"] = CommandInfo{
		Type: "request",
		Docs: docs,
	}

	StoneCommandList(s.Service, cmd_list)

	return zero.On("request", rules...)
}

func (s *Service) OnMetaEvent(name string, docs string, rules ...zero.Rule) *zero.Matcher {
	cmd_list := LoadCommandList(s.Service)
	cmd_list[name+"-onmeta"] = CommandInfo{
		Type: "meta",
		Docs: docs,
	}

	StoneCommandList(s.Service, cmd_list)

	return zero.On("meta_event", rules...)
}

func (s *Service) OnCommand(command string, docs string, aliases []string, rules ...zero.Rule) *zero.Matcher {
	cmd_list := LoadCommandList(s.Service)
	cmd_list[command] = CommandInfo{
		Type:    "command",
		Docs:    "",
		Aliases: aliases,
	}

	StoneCommandList(s.Service, cmd_list)

	commands := append(aliases, command)
	matcher := &zero.Matcher{
		Type:  zero.Type("message"),
		Rules: append([]zero.Rule{zero.CommandRule(commands...)}, rules...),
	}
	return zero.StoreMatcher(matcher)
}

func (s *Service) OnRegex(pattern string, docs string, rules ...zero.Rule) *zero.Matcher {
	cmd_list := LoadCommandList(s.Service)
	cmd_list[pattern] = CommandInfo{
		Type: "regex",
		Docs: docs,
	}

	StoneCommandList(s.Service, cmd_list)

	matcher := &zero.Matcher{
		Type:  zero.Type("message"),
		Rules: append([]zero.Rule{zero.RegexRule(pattern)}, rules...),
	}
	return zero.StoreMatcher(matcher)
}
