package configs

type ConfigModel struct {
	WebsocketURL   string              `yaml:"WebsocketURL"`
	Debug          bool                `yaml:"Debug"`
	SuperUsers     []int64             `yaml:"SuperUsers"`
	Nickname       []string            `yaml:"Nickname"`
	CommandPrefix  string              `yaml:"CommandPrefix"`
	AccessToken    string              `yaml:"AccessToken"`
	GoCQHTTP       ConfigGoCQHTTPModel `yaml:"GoCQHTTP"`
	SetuIsUseProxy bool                `yaml:"SetuIsUseProxy"`
	SauceNaoKey    string              `yaml:"SauceNaoKey"`
}

type ConfigGoCQHTTPModel struct {
	Enabled         bool   `yaml:"Enabled"`
	Account         string `yaml:"Account"`
	Password        string `yaml:"Password"`
	Protocol        string `yaml:"Protocol"`
	DownloadVersion string `yaml:"DownloadVersion"`
}
