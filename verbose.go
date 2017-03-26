package verbose

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

var debug = false

func SetDebugMode() {
	debug = true
}

func Info(info ...interface{}) {
	color.Green("[INFO] " + time.Now().String() + " " + fmt.Sprint(info))
}

func Warn(warning ...interface{}) {
	color.Yellow("[WARN] " + time.Now().String() + " " + fmt.Sprint(warning))
}

func Debug(debugInfo ...interface{}) {
	if debug {
		color.Blue("[DEBUG] " + time.Now().String() + " " + fmt.Sprint(debugInfo))
	}
}

func Err(error ...interface{}) {
	color.Red("[ERROR] " + time.Now().String() + " " + fmt.Sprint(error))
}
