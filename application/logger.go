package application

import (
	"github.com/team-dandelion/go-dandelion/config"
	"github.com/team-dandelion/go-dandelion/logger"
)

func initLogger() {
	if config.Conf.Logger == nil {
		return
	}

	//初始化日志
	if config.Conf.Logger.ConsoleShow ||
		config.Conf.Logger.FileWrite ||
		config.Conf.Logger.MultiFileWrite {
		if config.Conf.Logger.ConsoleShow {
			logger.RegAdapter(logger.AdapterConsole)
			logger.SetLoggerLevel(logger.AdapterConsole, config.Conf.Logger.ConsoleLevel)
		}
		if config.Conf.Logger.FileWrite {
			logger.RegAdapter(logger.AdapterFile)
			logger.SetLoggerLevel(logger.AdapterFile, config.Conf.Logger.FileLevel)
		}
		if config.Conf.Logger.MultiFileWrite {
			logger.RegAdapter(logger.AdapterMultiFile)
			logger.SetLoggerLevel(logger.AdapterMultiFile, config.Conf.Logger.MultiFileLevel)
		}

		logger.Async(1000)
	}
}
