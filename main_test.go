// Based on errata-ai/packages
// https://github.com/errata-ai/packages/blob/master/pkg/main_test.go
//
// MIT License
//
// Copyright (c) 2019 errata.ai

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmdtest"
)

func Test(t *testing.T) {
	testdataDir, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}

	for _, dir := range testdataDir {
		if !dir.IsDir() {
			continue
		}

		path := filepath.Join("testdata/", dir.Name())
		fmt.Printf("Running test: %s", path)

		doTest(path, t)
	}
}

func doTest(name string, t *testing.T) {
	ts, err := cmdtest.Read(name)
	if err != nil {
		t.Fatal(err)
	}

	ts.Setup = func(_ string) error {
		_, testFileName, _, ok := runtime.Caller(0)
		if !ok {
			return fmt.Errorf("failed get real working directory from caller")
		}

		projectRootDir := filepath.Join(filepath.Dir(testFileName), name)
		if err := os.Setenv("ROOTDIR", projectRootDir); err != nil {
			return fmt.Errorf("failed change 'ROOTDIR' to caller working directory: %v", err)
		}

		return nil
	}

	path, err := exec.LookPath("vale")
	if err != nil {
		t.Fatal(err)
	}

	ts.Commands["vale"] = cmdtest.Program(path)
	ts.Commands["cdf"] = cmdtest.InProcessProgram("cdf", cdf)

	ts.Run(t, false)
}