// Date:   Sat Jun 28 20:33:40 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package util

import (
	"fmt"
	"log"
	"net"
	"strconv"

	consul "github.com/hashicorp/consul/api"
)

func RegisterService(id string, name string, addr string) {
	config := consul.DefaultConfig()
	config.Address = "localhost:8500"

	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		log.Fatalf("invalid addr '%s'", addr)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("invalid addr '%s'", addr)
	}
	
	client, err := consul.NewClient(config)
	if err != nil {
		log.Fatalf("consul create new client failed: %s", err.Error())
	}
	
	registration := &consul.AgentServiceRegistration {
		ID: id,
		Name: name,
		Address: host,
		Port: port,
		// Check: &consul.AgentServiceCheck {
		// 	GRPC: addr,
		// 	GRPCUseTLS: false,
		// 	Interval: "10s",
		// 	Timeout: "2s",
		// 	DeregisterCriticalServiceAfter: "30s",
		// },
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("register service failed: %s\n", err.Error())
	}
	
	fmt.Printf("Service %s registered successfully!\n", id)
}
