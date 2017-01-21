package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"rpc"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)
	handler := NewRPCHandler()
	processor := rpc.NewRemoteCmdProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the Remotecommands server... on ", addr)
	return server.Serve()
}
