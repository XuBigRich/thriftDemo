package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
	"thriftDemo/gen-go/thriftAPI"
	"time"
)

var defaultCtx = context.Background()

func handleClient(client *thriftAPI.UserInfoServiceClient) (err error) {
	userInfo, _err := client.GetUserByName(defaultCtx, "1")
	if _err==nil{
		fmt.Println("出错了")
	}
	for _,v := range userInfo {
		fmt.Println("%v",v)
	}
	fmt.Println("GetMainPlans()")

	//sum, err := client.GetUserByNameWait(defaultCtx, "1")
	//fmt.Print("1+1=", sum, "\n")
	return err
}
func runClient(transportFactory thrift.TTransportFactory, addr string) (thrift.TTransport, error) {
	var transport thrift.TTransport
	var err error
	//cfg := new(tls.Config)
	//cfg.InsecureSkipVerify = true
	cfg := new(thrift.TConfiguration)
	cfg.ConnectTimeout = time.Second
	cfg.SocketTimeout = time.Second
	//thrift.TConfiguration{
	//	ConnectTimeout: time.Second, // Use 0 for no timeout
	//	SocketTimeout:  time.Second, // Use 0 for no timeout
	//}
	transport = thrift.NewTSocketConf(addr, cfg)
	transport = thrift.NewTFramedTransportConf(transport, cfg)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil, err
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return nil, err
	}
	return transport, nil
}

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(10081)
	transport, err := runClient(transportFactory, "127.0.0.1:10081")
	if err == nil {

	}
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:10081", " ", err)
		os.Exit(1)
	}
	defer transport.Close()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	tMultiplexedProtocol := thrift.NewTMultiplexedProtocol(protocolFactory.GetProtocol(transport), "ThrirftDemo1")
	//client := thriftAPI.NewUserInfoServiceClientFactory(transport, protocolFactory)
	client := thriftAPI.NewUserInfoServiceClientProtocol(transport, tMultiplexedProtocol, tMultiplexedProtocol)
	handleClient(client)
}
