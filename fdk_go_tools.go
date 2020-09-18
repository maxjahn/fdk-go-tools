// Package fdk-go-tools provides some helper functions for writing serverless functions for
// fnproject in Go.
package fdktools

import (
	"log"
	"log/syslog"
	"os"
	"strings"

)



// InitLogger will initialize a logger for use with fnproject on Oracle Cloud (OCI). It is
// configurable to write either to STDERR (as a default) or a syslog destination.
//For using a syslog destination, an URI in the format "protocol:://logerserver:port" needs to be
//provided via configuration variable FN_LOG_DESTINATION. Currently only tcp and udp are
//supported as protocols. An example for a valid destination URI is
//"tcp://logserver.example.com:12345".
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
