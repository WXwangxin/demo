package config

import (
	"github.com/kataras/iris/v12/middleware/accesslog"
	"io"
	"os"
)

// Read the example and its comments carefully.
func MakeAccessLog() *accesslog.AccessLog {
	// Initialize a new access log middleware.
	//ac := accesslog.File("./access.log")

	logFile, _ := os.OpenFile("./access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	ac := accesslog.New(io.MultiWriter(os.Stdout, logFile))

	// The default configuration:
	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	// Default line format if formatter is missing:
	// Time|Latency|Code|Method|Path|IP|Path Params Query Fields|Bytes Received|Bytes Sent|Request|Response|
	//
	// Set Custom Formatter:
	/*ac.SetFormatter(&accesslog.JSON{
		Indent:    "  ",
		HumanTime: true,
	})*/
	// ac.SetFormatter(&accesslog.CSV{})
	// ac.SetFormatter(&accesslog.Template{Text: "{{.Code}}"})

	return ac
}
