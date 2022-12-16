package config

type Logger struct {
	ConsoleShow    bool `json:"console_show" yaml:"consoleShow"`
	ConsoleLevel   int  `json:"console_level" yaml:"consoleLevel"`
	FileWrite      bool `json:"file_write" yaml:"fileWrite"`
	FileLevel      int  `json:"file_level" yaml:"fileLevel"`
	MultiFileWrite bool `json:"multi_file_write" yaml:"multiFileWrite"`
	MultiFileLevel int  `json:"multi_file_level" yaml:"multiFileLevel"`
}
