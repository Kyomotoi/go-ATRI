package configs

type Config struct {
	Bot    Bot    `yaml:"bot"`
	Driver Driver `yaml:"driver"`
	Plugin Plugin `yaml:"plugin"`
}

type Bot struct {
	Host          string   `yaml:"host"`
	Port          int      `yaml:"port"`
	Debug         bool     `yaml:"debug"`
	Superusers    []int64  `yaml:"superusers"`
	Nickname      []string `yaml:"nickname"`
	CommandPrefix string   `yaml:"command_prefix"`
	AccessToken   string   `yaml:"access_token"`
}

type Driver struct {
	Gocqhttp Gocqhttp `yaml:"gocqhttp"`
}

type Gocqhttp struct {
	Enabled         bool   `yaml:"enabled"`
	Account         int64  `yaml:"account"`
	Password        string `yaml:"password"`
	Protocol        int    `yaml:"protocol"`
	DownloadVersion string `yaml:"download_version"`
}

type Plugin struct {
	SetuIsUseProxy bool   `yaml:"setu_is_use_proxy"`
	SaucenaoKey    string `yaml:"saucenao_key"`
}
