package framework

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lijw30/signaling/src/framework/xrpc"
	"strings"
	"time"
)

var xrpcClients map[string]*xrpc.Client = make(map[string]*xrpc.Client)

func loadXrpc() error {
	xrpcConfig := GetXrpcConfig()
	arrServer := strings.Split(xrpcConfig.Server, ",")
	for i, server := range arrServer {
		arrServer[i] = strings.TrimSpace(server)
	}

	client := xrpc.NewClient(arrServer)
	client.ConnectTimeout = time.Duration(xrpcConfig.ConnectTimeout) * time.Millisecond
	client.ReadTimeout = time.Duration(xrpcConfig.ReadTimeout) * time.Millisecond
	client.WriteTimeout = time.Duration(xrpcConfig.WriteTimeout) * time.Millisecond

	fmt.Println(arrServer, client)
	xrpcClients["xrtc"] = client //todo 后边可能有问题，需要改

	return nil
}

func Call(serviceName string, request interface{}, response interface{}, logId uint32) error {
	fmt.Println("call:", serviceName)

	client, ok := xrpcClients[serviceName]
	if !ok {
		return fmt.Errorf("[%s] service not found", serviceName)
	}

	content, _ := json.Marshal(request)
	req := xrpc.NewRequest(bytes.NewReader(content), logId)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
