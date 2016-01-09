package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/cespare/window"
	"github.com/codegangsta/cli"
	"github.com/jeanfric/goembed/countingwriter"
)

const Version = "0.2.0"

func usageAndExit(s string, c *cli.Context) {
	fmt.Printf("!! %s\n\n", s)
	cli.ShowAppHelp(c)
	os.Exit(1)
}

func init() {
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}

func main() {
	app := cli.NewApp()
	app.Name = "email-me"
	app.Version = Version
	app.Usage = "email me when a thing is done"
	app.Action = main2
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "to",
			Value:  "",
			Usage:  "email address to send output to",
			EnvVar: "EMAIL_ME_TO",
		},
		cli.StringFlag{
			Name:  "subject, s",
			Value: "",
			Usage: "subject of email (optional)",
		},
		cli.IntFlag{
			Name:   "max",
			Value:  10000,
			Usage:  "max bytes to capture for stdout/stderr",
			EnvVar: "EMAIL_ME_MAX",
		},
		cli.BoolFlag{
			Name:  "on-error",
			Usage: "only notify on a non-0 exit code",
		},
	}
	app.Run(os.Args)
}

func main2(c *cli.Context) {
	to := c.String("to")
	max := c.Int("max")
	subject := c.String("subject")
	onError := c.Bool("on-error")

	if to == "" {
		usageAndExit("missing --to=[address]", c)
	}

	args := c.Args()
	if len(args) == 0 {
		usageAndExit("missing [command]", c)
	}

	truncStdout := window.NewWriter(max)
	truncStderr := window.NewWriter(max)
	countedStdout := countingwriter.New(truncStdout)
	countedStderr := countingwriter.New(truncStderr)
	child := exec.Command(args[0], args[1:]...)
	child.Stdout = io.MultiWriter(os.Stdout, countedStdout)
	child.Stderr = io.MultiWriter(os.Stderr, countedStderr)
	child.Stdin = os.Stdin

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
	if subject == "" {
		subject = fmt.Sprintf("%s", child.Args)
	}

	m := &Message{
		To:      to,
		From:    me,
		Subject: subject,
		Result:  r,
	}

	success := child.ProcessState.Success()

	if onError && success {
		os.Exit(0)
	}

	if err := findMailer().Send([]string{to}, me, m.Bytes()); err != nil {
		log.Fatal(err)
	}

	if !success {
		os.Exit(1)
	}
}
