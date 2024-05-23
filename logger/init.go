package logger

import (
	"fmt"
)

func InitLogger() {
	// log path
	filePath := fmt.Sprintf("%s%s",
		"./runtime/",
		"logs/",
	)

	// log file name
	fileName := fmt.Sprintf("%s.%s",
		"validator",
		"log",
	)
	// log level
	logLevel := "debug"
	runMode := "debug"
	expireDay := 1
	setup(filePath, fileName, logLevel, runMode, int32(expireDay))
}
