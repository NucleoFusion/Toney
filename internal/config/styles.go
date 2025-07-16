package config

type StylesConfig struct {
	Text          string           `mapstructure:"text"`
	Background    string           `mapstructure:"background"`
	Border        string           `mapstructure:"border"`
	FocusedBorder string           `mapstructure:"focused_border"`
	Icons         IconsConfig      `mapstructure:"icons"`
	Renderer      RendererConfig   `mapstructure:"renderer"`
	TaskStyles    TaskStylesConfig `mapstructure:"task_styles"`
}

type IconsConfig struct {
	FolderIcon string    `mapstructure:"folder_icon"`
	FileIcon   string    `mapstructure:"file_icon"`
	TaskIcons  TaskIcons `mapstructure:"task_icons"`
}

type TaskIcons struct {
	CompletedIcon string `mapstructure:"completed_icon"`
	PendingIcon   string `mapstructure:"pending_icon"`
	AbandonedIcon string `mapstructure:"abandoned_icon"`
	StartedIcon   string `mapstructure:"started_icon"`
}

type RendererConfig struct {
	Document       BlockStyle   `mapstructure:"document"`
	BlockQuote     BlockStyle   `mapstructure:"blockquote"`
	Heading        HeadingStyle `mapstructure:"heading"`
	HorizontalRule BlockStyle   `mapstructure:"horizontal_rule"`
	Paragraph      BlockStyle   `mapstructure:"paragraph"`
	List           ListStyle    `mapstructure:"list"`
	Enumeration    ListStyle    `mapstructure:"enumeration"`
	Link           InlineStyle  `mapstructure:"link"`
	Image          InlineStyle  `mapstructure:"image"`
	Emph           InlineStyle  `mapstructure:"emph"`
	Strong         InlineStyle  `mapstructure:"strong"`
	Strikethrough  InlineStyle  `mapstructure:"strikethrough"`
	Code           InlineStyle  `mapstructure:"code"`
	CodeBlock      BlockStyle   `mapstructure:"code_block"`
	Table          TableStyle   `mapstructure:"table"`
}

type BlockStyle struct {
	Margin      [4]int `mapstructure:"margin"`
	Padding     [4]int `mapstructure:"padding"`
	Indent      int    `mapstructure:"indent"`
	Border      string `mapstructure:"border"`
	BorderColor string `mapstructure:"border_color"`
	Color       string `mapstructure:"color"`
	Background  string `mapstructure:"background"`
	Bold        bool   `mapstructure:"bold"`
	Italic      bool   `mapstructure:"italic"`
	Underline   bool   `mapstructure:"underline"`
}

type HeadingStyle struct {
	Level []BlockStyle `mapstructure:"level"`
}

type ListStyle struct {
	Indent int    `mapstructure:"indent"`
	Symbol string `mapstructure:"symbol"`
}

type InlineStyle struct {
	Color         string `mapstructure:"color"`
	Background    string `mapstructure:"background"`
	Bold          bool   `mapstructure:"bold"`
	Italic        bool   `mapstructure:"italic"`
	Underline     bool   `mapstructure:"underline"`
	Strikethrough bool   `mapstructure:"strikethrough"`
}

type TableStyle struct {
	Border bool       `mapstructure:"border"`
	Align  string     `mapstructure:"align"`
	Header BlockStyle `mapstructure:"header"`
	Cell   BlockStyle `mapstructure:"cell"`
}

type TaskStylesConfig struct {
	FocusedBar     string         `mapstructure:"focused_color"`
	UnfocusedBar   string         `mapstructure:"unfocused_color"`
	CompletedStyle TaskStateStyle `mapstructure:"completed_style"`
	PendingStyle   TaskStateStyle `mapstructure:"pending_style"`
	AbandonedStyle TaskStateStyle `mapstructure:"abandoned_style"`
	StartedStyle   TaskStateStyle `mapstructure:"started_style"`
}

type TaskStateStyle struct {
	Color     string `mapstructure:"color"`
	TextColor string `mapstructure:"text_color"`
}
