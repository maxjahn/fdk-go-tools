# fdktools
--
    import "src/public/fdk-go-tools"

Package fdk-go-tools provides some helper functions for writing serverless
functions for fnproject in Go.

## Usage

#### func  InitLogger

```go
func InitLogger(logTag string) (fdklogger *log.Logger)
```
InitLogger will initialize a logger for use with fnproject on Oracle Cloud
(OCI). It is configurable to write either to STDERR (as a default) or a syslog
destination. For using a syslog destination, an URI in the format
"protocol:://logerserver:port" needs to be provided via configuration variable
FN_LOG_DESTINATION. Currently only tcp and udp are supported as protocols. An
example for a valid destination URI is "tcp://logserver.example.com:12345".
