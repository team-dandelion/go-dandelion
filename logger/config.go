package logger

type Config struct {
	ConsoleShow    bool
	ConsoleLevel   int
	FileWrite      bool
	FileLevel      int
	MultiFileWrite bool
	MultiFileLevel int
}

func InitWithDefaultConfig() {
	cfg := Config{}
	cfg.ConsoleShow = true
	cfg.ConsoleLevel = 7
	cfg.FileWrite = false
	cfg.FileLevel = 7
	cfg.MultiFileWrite = true
	cfg.MultiFileLevel = 7

	if cfg.ConsoleShow ||
		cfg.FileWrite ||
		cfg.MultiFileWrite {
		if cfg.ConsoleShow {
			RegAdapter(AdapterConsole)
			SetLoggerLevel(AdapterConsole, cfg.ConsoleLevel)
		}
		if cfg.FileWrite {
			RegAdapter(AdapterFile)
			SetLoggerLevel(AdapterFile, cfg.FileLevel)
		}
		if cfg.MultiFileWrite {
			RegAdapter(AdapterMultiFile)
			SetLoggerLevel(AdapterMultiFile, cfg.MultiFileLevel)
		}
		Async(1000)
	}
}
