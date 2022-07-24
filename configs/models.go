package configs

type ConfigModel struct {
	WebsocketURL   string   `yaml:"WebsocketURL"`
	Debug          bool     `yaml:"Debug"`
	SuperUsers     []int64  `yaml:"SuperUsers"`
	Nickname       []string `yaml:"Nickname"`
	CommandPrefix  string   `yaml:"CommandPrefix"`
	AccessToken    string   `yaml:"AccessToken"`
	SetuIsUseProxy bool     `yaml:"SetuIsUseProxy"`
	SauceNaoKey    string   `yaml:"SauceNaoKey"`
}
