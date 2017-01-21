package main

import (
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
	return r, nil
}

// CPUUsage func
func (p *RPCHandler) CPUUsage() (r string, err error) {
	return r, nil
}

// AvailableRAM func
func (p *RPCHandler) AvailableRAM() (r string, err error) {
	return r, nil
}

// CPUUsageLastHour func
func (p *RPCHandler) CPUUsageLastHour() (r string, err error) {
	return r, nil
}

// AvailableRAMLastHour func
func (p *RPCHandler) AvailableRAMLastHour() (r string, err error) {
	return r, nil
}

// DownloadURL func
func (p *RPCHandler) DownloadURL(url string, folder string) (r string, err error) {
	return r, nil
}

// Say func
func (p *RPCHandler) Say(phrase string) (r string, err error) {
	return r, nil
}

// Screenshot func
func (p *RPCHandler) Screenshot() (r []byte, err error) {
	return r, nil
}
