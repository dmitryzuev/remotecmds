package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"rpc"
)

func handleClient(client *rpc.RemoteCmdClient) (err error) {
	utc, _ := client.Utc()
	fmt.Println("Current time:", utc)

	return err
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	return handleClient(rpc.NewRemoteCmdClientFactory(transport, protocolFactory))
}
