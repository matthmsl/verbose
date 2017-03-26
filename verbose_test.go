package verbose

import "testing"

func TestDebug(t *testing.T) {
	//To enable debug information in console, call SetDebugMode.
	SetDebugMode()
	Debug("This is a debug information")
}

func TestInfo(t *testing.T) {
	Info("This is an information")
}

func TestErr(t *testing.T) {
	Err("This is an error")
}

func TestWarn(t *testing.T) {
	Warn("This is a warning")
}