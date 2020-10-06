// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

// +build linux,service

package main

import (
	"context"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeebo/errs"
)

func cmdRestart(cmd *cobra.Command, args []string) error {
	return nil
}

func restartService(ctx context.Context, service, binaryLocation, newVersionPath, backupPath string) error {
	if err := os.Rename(binaryLocation, backupPath); err != nil {
		return errs.Wrap(err)
	}

	if err := os.Rename(newVersionPath, binaryLocation); err != nil {
		return errs.Combine(err, os.Rename(backupPath, binaryLocation), os.Remove(newVersionPath))
	}

	if service == updaterServiceName {
		os.Exit(1)
	}

	if err := stopProcess(service); err != nil {
		err = errs.New("error stopping %s service: %v", service, err)
		return errs.Combine(err, os.Rename(backupPath, binaryLocation))
	}

	return nil
}

func stopProcess(service string) (err error) {
	pid, err := getServicePID(service)
	if err != nil {
		return err
	}
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return p.Signal(os.Interrupt)

}

func getServicePID(service string) (int, error) {
	args := []string{
		"show",
		"--property=MainPID",
		service,
	}

	cmd := exec.Command("systemctl", args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}

	trimmed := strings.TrimPrefix(string(out), "MainPID=")
	trimmed = strings.TrimSuffix(trimmed, "\n")

	pid, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, err
	}

	return pid, nil
}
