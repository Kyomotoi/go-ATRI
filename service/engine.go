package service

import zero "github.com/wdvxdr1123/ZeroBot"

// 对 ZeroBot/engine.go 再封装，以适用于 service

var defEngine = zero.New()


func On(typ string, rules ...Rule) *Matcher {
	return defEngine.On(typ, rules...)
}

func (e *Engine) On(typ string, rules ...Rule) *Matcher {
	matcher := &Matcher{
		Type:   zero.Type(typ),
		Rules:  rules,
		Engine: (*zero.Engine)(e),
	}
	return zero.StoreMatcher(matcher)
}

func OnMessage(rules ...Rule) *Matcher {
	return zero.On("message", rules...)
}

func (e *Engine) OnMessage(rules ...Rule) *Matcher {
	return e.On("message", rules...)
}

func OnNotice(rules ...Rule) *Matcher {
	return On("notice", rules...)
}

func OnRequest(rules ...Rule) *Matcher {
	return On("request", rules...)
}

func OnCommand(commands string, rules ...Rule) *Matcher {
	return defEngine.OnCommand(commands, rules...)
}
