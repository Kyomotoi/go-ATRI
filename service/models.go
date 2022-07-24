package service

import zero "github.com/wdvxdr1123/ZeroBot"

type Service struct {
	Service     string
	Docs        string
	OnlyAdmin   bool
	Rule        []zero.Rule
	MainCommand string
}

type ServiceInfo struct {
	Service      string                 `json:"service"`
	Docs         string                 `json:"docs"`
	CommandList  map[string]CommandInfo `json:"command_list"`
	Enabled      bool                   `json:"enabled"`
	OnlyAdmin    bool                   `json:"only_admin"`
	DisableUser  []string               `json:"disable_user"`
	DisableGroup []string               `json:"disable_group"`
}

type CommandInfo struct {
	Type    string   `json:"type"`
	Docs    string   `json:"docs"`
	Aliases []string `json:"aliases"`
}
