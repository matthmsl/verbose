package verbose

import (
	"fmt"
	"github.com/fatih/color"
	"time"
	"os"
)

var debug = false
var file = false
var filename string

func SetLogFile(name string){
	file=true
	filename=name
	os.Create(filename)
}

func SetDebugMode() {
	debug = true
}

func Info(info ...interface{}) {
	msg:="[INFO] " + time.Now().String() + " " + fmt.Sprint(info)
	color.Green(msg)
	log(msg)
}

func Warn(warning ...interface{}) {
	msg:="[WARN] " + time.Now().String() + " " + fmt.Sprint(warning)
	color.Yellow(msg)
	log(msg)
}

func Debug(debugInfo ...interface{}) {
	if debug {
		msg:="[DEBUG] " + time.Now().String() + " " + fmt.Sprint(debugInfo)
		color.Blue(msg)
		log(msg)
	}
}

func Err(error ...interface{}) {
	msg:="[ERROR] " + time.Now().String() + " " + fmt.Sprint(error)
	color.Red(msg)
	log(msg)
}

func log(msg string){
	if file{
		f,err:=os.OpenFile(filename, os.O_APPEND|os.O_WRONLY,0600)
		defer f.Close()
		if err!=nil{ panic(err)}
		_,err=f.WriteString(msg)
		if err!=nil{ panic(err)}
	}

}
