package config

type ChatGroup struct {
	IsAnonymous    bool   `yaml:"is_anonymous" json:"is_anonymous"`
	IsShowTime     bool   `yaml:"is_show_time" json:"is_show_time"`
	DefaultLimit   int    `yaml:"default_limit" json:"default_limit"`
	WelcomeTitle   string `yaml:"welcome_title" json:"welcome_title"`
	IsOnlinePeople bool   `yaml:"is_online_people" json:"is_online_people"`
	IsSendImage    bool   `yaml:"is_send_image" json:"is_send_image"`
	IsSendFile     bool   `yaml:"is_send_file" json:"is_send_file"`
	IsMarkdown     bool   `yaml:"is_markdown" json:"is_markdown"`
	IsPaste        bool   `yaml:"is_paste" json:"is_paste"`
	ContentLength  int    `yaml:"content_length" json:"content_length"`
}
