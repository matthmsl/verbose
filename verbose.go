package verbose

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/getsentry/sentry-go"
	"os"
	"time"
)

var debug = false
var file = false
var filename string
var sentryReporting = false

func SetLogFile(name string) {
	file = true
	filename = name
	_, err := os.Create(filename)
	if err != nil {
		Err(err)
	}
}

func SetDebugMode() {
	debug = true
	if sentryReporting {
		Warn("[VERBOSE] Sentry was initialized before debug mode! Enable debug beforehand to debug sentry")
	}
}

func EnableSentry(dsn string, release string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:     dsn,
		Release: release,
		Debug:   debug,
	})
	if err != nil {
		Err(err)
	} else {
		sentryReporting = true
	}
}

func Info(info ...interface{}) {
	msg := "[INFO] " + time.Now().String() + " " + fmt.Sprint(info...)
	if sentryReporting {
		sentryProcess(msg, sentry.LevelInfo)
	}
	color.Green(msg)
	log(msg)
}

func Warn(warning ...interface{}) {
	msg := "[WARN] " + time.Now().String() + " " + fmt.Sprint(warning...)
	if sentryReporting {
		sentryProcess(msg, sentry.LevelWarning)
	}
	color.Yellow(msg)
	log(msg)
}

func Debug(debugInfo ...interface{}) {
	if debug {
		msg := "[DEBUG] " + time.Now().String() + " " + fmt.Sprint(debugInfo...)
		if sentryReporting {
			sentryProcess(msg, sentry.LevelDebug)
		}
		color.Blue(msg)
		log(msg)
	}
}

func Err(error ...interface{}) {
	msg := "[ERROR] " + time.Now().String() + " " + fmt.Sprint(error...)
	if sentryReporting {
		sentryProcess(msg, sentry.LevelError)
	}
	color.Red(msg)
	log(msg)
}

func log(msg string) {
	if file {
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		_, err = f.WriteString(msg + "\n")
		if err != nil {
			panic(err)
		}
	}
}

func sentryProcess(msg string, severity sentry.Level) {
	evt := sentry.NewEvent()
	evt.Level = severity
	evt.Message = msg
	evt.Exception = append(evt.Exception, sentry.Exception{
		Stacktrace: sentry.NewStacktrace(),
		Type:       msg,
	})
	sentry.CaptureEvent(evt)
	sentry.Flush(time.Second*5)
}
