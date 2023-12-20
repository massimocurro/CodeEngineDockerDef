package main

// This program is a simple HTTP server that will do 2 things:
// 1 - invoke an "init" app (/app/init) at start-up
// 2 - invoke an "app" app (/app/app) upon each HTTP request
// See the README for more info

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var verbose = 1

func Debug(level int, msg string, args ...interface{}) {
	// Note: 0 == always print
	if level > verbose {
		return
	}
	log.Printf(msg, args...)
}

func Run(envs []string, input []byte, cmdStr string, args ...interface{}) (int, string) {
	Debug(1, "Running: %s", fmt.Sprint(cmdStr))
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf(cmdStr, args...))
	if len(envs) != 0 {
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, envs...)
	}

	if input != nil {
		cmd.Stdin = bytes.NewReader(input)
	} else {
		cmd.Stdin = nil
	}

	buf, err := cmd.CombinedOutput()
	if err != nil {
		if len(buf) > 0 {
			buf = append(buf, []byte("\n")...)
		}
		buf = append(buf, []byte(err.Error()+"\n")...)
		Debug(2, "  Error: %s", err)
	}

	Debug(2, "Output:\n%s", string(buf))
	Debug(2, "Exit Code: %d", cmd.ProcessState.ExitCode())

	return cmd.ProcessState.ExitCode(), string(buf)
}

func main() {
	if os.Getenv("DEBUG") != "" {
		verbose = 2
	}

	if _, err := os.Stat("/app/init"); err == nil {
		Run(nil, nil, "/app/init")
	}

	if tmp := os.Getenv("INIT"); tmp != "" {
		Run(nil, nil, tmp)
	}

	Command := "/app/app"
	if tmp := os.Getenv("APP"); tmp != "" {
		Command = tmp
	}

	// Our HTTP handler func
    http.Get("https://function-76.1at6rgz00yjr.eu-de.codeengine.appdomain.cloud")

	http.ListenAndServe(":8080", nil)
}
