package gocqhttp

import (
	"time"
)

type GithubReleaseList []struct {
	URL       string `json:"url"`
	AssetsURL string `json:"assets_url"`
	UploadURL string `json:"upload_url"`
	HTMLURL   string `json:"html_url"`
	ID        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeID          string                `json:"node_id"`
	TagName         string                `json:"tag_name"`
	TargetCommitish string                `json:"target_commitish"`
	Name            string                `json:"name"`
	Draft           bool                  `json:"draft"`
	Prerelease      bool                  `json:"prerelease"`
	CreatedAt       time.Time             `json:"created_at"`
	PublishedAt     time.Time             `json:"published_at"`
	Assets          []GithubReleaseAssets `json:"assets"`
	TarballURL      string                `json:"tarball_url"`
	ZipballURL      string                `json:"zipball_url"`
	Body            string                `json:"body"`
	DiscussionURL   string                `json:"discussion_url,omitempty"`
	Reactions       struct {
		URL        string `json:"url"`
		TotalCount int    `json:"total_count"`
		Num1       int    `json:"+1"`
		Num10      int    `json:"-1"`
		Laugh      int    `json:"laugh"`
		Hooray     int    `json:"hooray"`
		Confused   int    `json:"confused"`
		Heart      int    `json:"heart"`
		Rocket     int    `json:"rocket"`
		Eyes       int    `json:"eyes"`
	} `json:"reactions,omitempty"`
	MentionsCount int `json:"mentions_count,omitempty"`
}

type GithubReleaseAssets struct {
	URL      string `json:"url"`
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	Label    string `json:"label"`
	Uploader struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"uploader"`
	ContentType        string    `json:"content_type"`
	State              string    `json:"state"`
	Size               int       `json:"size"`
	DownloadCount      int       `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadURL string    `json:"browser_download_url"`
}

type GocqhttpConfig struct {
	Account            Account            `yaml:"account"`
	Heartbeat          Heartbeat          `yaml:"heartbeat"`
	Message            Message            `yaml:"message"`
	Output             Output             `yaml:"output"`
	DefaultMiddlewares DefaultMiddlewares `yaml:"default-middlewares"`
	Database           Database           `yaml:"database"`
	Servers            []Servers          `yaml:"servers"`
}

type Servers struct {
	Middlewares Middlewares `yaml:"middlewares"`
	Ws          interface{} `yaml:"ws"`
	Host        string      `yaml:"host"`
	Port        int         `yaml:"port"`
}

type Database struct {
	Leveldb Leveldb `yaml:"leveldb"`
	Cache   Cache   `yaml:"cache"`
}

type Heartbeat struct {
	Interval int `yaml:"interval"`
}

type Message struct {
	ForceFragment       bool   `yaml:"force-fragment"`
	FixUrl              bool   `yaml:"fix-url"`
	ExtraReplyData      bool   `yaml:"extra-reply-data"`
	SkipMimeScan        bool   `yaml:"skip-mime-scan"`
	PostFormat          string `yaml:"post-format"`
	IgnoreInvalidCqcode bool   `yaml:"ignore-invalid-cqcode"`
	ProxyRewrite        string `yaml:"proxy-rewrite"`
	ReportSelfMessage   bool   `yaml:"report-self_message"`
	RemoveReplyAt       bool   `yaml:"remove-reply-at"`
}

type Output struct {
	LogForceNew bool   `yaml:"log-force-new"`
	LogColorful bool   `yaml:"log-colorful"`
	Debug       bool   `yaml:"debug"`
	LogLevel    string `yaml:"log-level"`
	LogAging    int    `yaml:"log-aging"`
}

type Middlewares struct {
	AccessToken string               `yaml:"access-token"`
	Filter      string               `yaml:"filter"`
	RateLimit   MiddlewaresRateLimit `yaml:"rate-limit"`
}

type Relogin struct {
	Delay    int `yaml:"delay"`
	Interval int `yaml:"interval"`
	MaxTimes int `yaml:"max-times"`
}

type DefaultMiddlewares struct {
	AccessToken string    `yaml:"access-token"`
	Filter      string    `yaml:"filter"`
	RateLimit   RateLimit `yaml:"rate-limit"`
}

type RateLimit struct {
	Enabled   bool `yaml:"enabled"`
	Frequency int  `yaml:"frequency"`
	Bucket    int  `yaml:"bucket"`
}

type Leveldb struct {
	Enable bool `yaml:"enable"`
}

type Cache struct {
	Image string `yaml:"image"`
	Video string `yaml:"video"`
}

type MiddlewaresRateLimit struct {
	Enabled   bool `yaml:"enabled"`
	Frequency int  `yaml:"frequency"`
	Bucket    int  `yaml:"bucket"`
}

type Account struct {
	Relogin          Relogin `yaml:"relogin"`
	UseSsoAddress    bool    `yaml:"use-sso-address"`
	AllowTempSession bool    `yaml:"allow-temp-session"`
	Uin              int     `yaml:"uin"`
	Password         string  `yaml:"password"`
	Encrypt          bool    `yaml:"encrypt"`
	Status           int     `yaml:"status"`
}
