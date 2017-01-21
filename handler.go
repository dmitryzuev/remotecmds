package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// RPCHandler struct
type RPCHandler struct{}

// NewRPCHandler func
func NewRPCHandler() *RPCHandler {
	return &RPCHandler{}
}

// Utc func
func (p *RPCHandler) Utc() (r string, err error) {
	r = time.Now().UTC().String()
	return
}

// CPU func
func (p *RPCHandler) CPU() (r string, err error) {
	return
}

// RAM func
func (p *RPCHandler) RAM() (r string, err error) {
	return
}

// Download func
func (p *RPCHandler) Download(url string, folder string) (err error) {
	return
}

// Say func
func (p *RPCHandler) Say(phrase string) (err error) {
	switch runtime.GOOS {
	case "darwin":
		cmd := exec.Command("say")
		cmd.Stdin = strings.NewReader(phrase)
		err = cmd.Run()
	}
	return
}

// Screenshot func
func (p *RPCHandler) Screenshot() (r []byte, err error) {
	switch runtime.GOOS {
	case "darwin":
		ch := make(chan time.Duration)
		go func() {
			cmd := exec.Command("screencapture", "-T4", "-x", "screenshot.png")
			err = cmd.Run()
			ch <- 5 * time.Second
		}()
		time.Sleep(<-ch)
		r, err = ioutil.ReadFile("screenshot.png")
		os.Remove("screenshot.png")
	}
	return
}
