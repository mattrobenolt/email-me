package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/cespare/window"
	"github.com/jeanfric/goembed/countingwriter"
)

const Version = "0.1.0"

var (
	toFlag      = flag.String("to", "", "email address to send output to")
	subjectFlag = flag.String("s", "", "subject of email (optional)")
	maxFlag     = flag.Int("max", 10000, "max bytes to capture for stdout/stderr")
)

func usageAndExit(s string) {
	fmt.Printf("!! %s\n", s)
	flag.Usage()
	fmt.Println()
	fmt.Printf("%s version: %s (%s on %s/%s; %s)\n", os.Args[0], Version, runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.Compiler)
	os.Exit(1)
}

func init() {
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "usage: email-me [flags] [command]\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	if *toFlag == "" {
		usageAndExit("missing -to=[address]")
	}

	args := flag.Args()
	if len(args) == 0 {
		usageAndExit("missing [command]")
	}

	truncStdout := window.NewWriter(*maxFlag)
	truncStderr := window.NewWriter(*maxFlag)
	countedStdout := countingwriter.New(truncStdout)
	countedStderr := countingwriter.New(truncStderr)
	child := exec.Command(args[0], args[1:]...)
	child.Stdout = io.MultiWriter(os.Stdout, countedStdout)
	child.Stderr = io.MultiWriter(os.Stderr, countedStderr)

	start := time.Now()
	err := child.Run()
	end := time.Now()

	stdout := truncStdout.Bytes()
	stderr := truncStderr.Bytes()

	r := &Result{
		Cmd:         child,
		Error:       err,
		Start:       start,
		End:         end,
		Duration:    end.Sub(start),
		Stdout:      string(stdout),
		Stderr:      string(stderr),
		StdoutExtra: countedStdout.BytesWritten() - len(stdout),
		StderrExtra: countedStderr.BytesWritten() - len(stderr),
	}

	me := identity()
	subject := *subjectFlag
	if subject == "" {
		subject = fmt.Sprintf("%s", child.Args)
	}

	m := &Message{
		To:      *toFlag,
		From:    me,
		Subject: subject,
		Result:  r,
	}

	if err := findMailer().Send([]string{*toFlag}, me, m.Bytes()); err != nil {
		log.Fatal(err)
	}

	if !child.ProcessState.Success() {
		os.Exit(1)
	}
}
