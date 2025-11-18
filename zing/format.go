package zing

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	out     = os.Stdout
	errOut  = os.Stderr
	ok      = color.New(color.FgGreen).SprintFunc()
	info    = color.New(color.FgBlue).SprintFunc()
	warn    = color.New(color.FgYellow).SprintFunc()
	errorCl = color.New(color.FgRed).SprintFunc()
	msgCl   = color.New(color.FgCyan).SprintFunc()
	normal  = color.New(color.FgWhite).SprintFunc()
)

func LogInfo(msgs ...interface{}) {
	w := bufio.NewWriter(out)
	fmt.Fprintln(w, info(msgs...))
	w.Flush()
}

func LogWarn(msgs ...interface{}) {
	w := bufio.NewWriter(out)
	fmt.Fprintln(w, warn(msgs...))
	w.Flush()
}

func LogSuccess(msgs ...interface{}) {
	w := bufio.NewWriter(out)
	fmt.Fprintln(w, ok(msgs...))
	w.Flush()
}

func LogError(msgs ...interface{}) {
	w := bufio.NewWriter(errOut)
	fmt.Fprintln(w, errorCl(msgs...))
	w.Flush()
}

func LogMessage(msgs ...interface{}) {
	w := bufio.NewWriter(errOut)
	fmt.Fprintln(w, msgCl(msgs...))
	w.Flush()
}

func LogNormal(msgs ...interface{}) {
	w := bufio.NewWriter(errOut)
	fmt.Fprint(w, normal(msgs...))
	w.Flush()
}

func LogNormalLn(msgs ...interface{}) {
	w := bufio.NewWriter(errOut)
	fmt.Fprint(w, normal(msgs...))
	w.Flush()
}
