package service

import zero "github.com/wdvxdr1123/ZeroBot"

type Service struct {
	Service      string        `json:"service"`
	Docs         string        `json:"docs"`
	Enabled      bool          `json:"enabled"`
	OnlyAdmin    bool          `json:"only_admin"`
	DisableUser  []string      `json:"disable_user"`
	DisableGroup []string      `json:"disable_group"`
}

type Rule = zero.Rule

type Matcher = zero.Matcher

type Engine zero.Engine
