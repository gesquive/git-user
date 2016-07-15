package cli

import "os"
import "fmt"
import "github.com/fatih/color"

// Yellow creates a yellow string
var Yellow = color.New(color.FgHiYellow).SprintFunc()

// Green creates a green string
var Green = color.New(color.FgHiGreen).SprintFunc()

// Blue creates a blue string
var Blue = color.New(color.FgHiBlue).SprintFunc()

// Red creates a red string
var Red = color.New(color.FgHiRed).SprintFunc()

// PrintDebug sets if debug messages will be printed or not
var PrintDebug = false

// Infof prints an info line to stdout
func Infof(format string, args ...interface{}) (n int, err error) {
	n, err = fmt.Printf(format, args...)
	if err == nil {
		fmt.Print("\n")
	}
	return
}

// Debugf prints an debug line to stdout if printDebug is set
func Debugf(format string, args ...interface{}) (n int, err error) {
	if PrintDebug {
		n, err = fmt.Printf(format, args...)
		if err == nil {
			fmt.Print("\n")
		}
	}
	return
}

// Errorf prints an error line to stderr
func Errorf(format string, args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(os.Stderr, format, args...)
	if err == nil {
		fmt.Fprintf(os.Stderr, "\n")
	}
	return
}
