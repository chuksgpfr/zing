package zing

import (
	"context"
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
