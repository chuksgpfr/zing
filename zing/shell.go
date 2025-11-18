package zing

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"
	"time"
)

func RunShellCapture(cmdStr string, timeout time.Duration) (string, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sh", "-lc", cmdStr)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func StreamShell(cmdStr string, timeout time.Duration) (string, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sh", "-lc", cmdStr)

	var buf bytes.Buffer
	stdoutWriter := io.MultiWriter(os.Stdout, &buf)
	stderrWriter := io.MultiWriter(os.Stderr, &buf)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return buf.String(), err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return buf.String(), err
	}

	if err := cmd.Start(); err != nil {
		return buf.String(), err
	}

	// copy stdout and stderr concurrently
	copyStream := func(r io.Reader, w io.Writer) {
		sc := bufio.NewScanner(r)
		for sc.Scan() {
			line := sc.Bytes()
			w.Write(append(line, '\n'))
		}
	}
	go func() { copyStream(stdoutPipe, stdoutWriter) }()
	go func() { copyStream(stderrPipe, stderrWriter) }()

	err = cmd.Wait()
	return buf.String(), err
}
