package service

type Service struct {
	Service      string            `json:"service"`
	Docs         string            `json:"docs"`
	Commands     map[string]string `json:"commands"`
	Enabled      bool              `json:"enabled"`
	DisableUser  []string          `json:"disable_user"`
	DisableGroup []string          `json:"disable_group"`
}
