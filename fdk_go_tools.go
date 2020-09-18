package fdktools

import (
	"log"
	"log/syslog"
	"os"
	"strings"

)

func InitLogger(logTag string) (fdklogger *log.Logger) {

	// use STDERR as default logging destination
	fdklogger = log.New(os.Stderr, "", log.LstdFlags)

	// if a syslog destination is configured, try to use it for syslog
	logParts := strings.Split(os.Getenv("FN_LOG_DESTINATION"), "://")

	if len(logParts) == 2 {

		logProtocol := logParts[0]
		logURL := logParts[1]

		sysLog, err := syslog.Dial(logProtocol, logURL, syslog.LOG_WARNING|syslog.LOG_SYSLOG, logTag)

		if err != nil {
			fdklogger.Println("Failed to create syslog connections: ", err)
		} else {
			fdklogger = log.New(sysLog, "", log.LstdFlags)
		}
	}
	return fdklogger
}
