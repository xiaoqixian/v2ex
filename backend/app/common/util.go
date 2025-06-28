// Date:   Sat Jun 28 20:33:40 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package util

import (
	"fmt"
	"log"

	consul "github.com/hashicorp/consul/api"
)

func RegisterService(id string, name string, port int) {
	config := consul.DefaultConfig()
	config.Address = "localhost:8500"
	
	client, err := consul.NewClient(config)
	if err != nil {
		log.Fatalf("consul create new client failed: %s", err.Error())
	}
	
	registration := &consul.AgentServiceRegistration {
		ID: id,
		Name: name,
		Address: "localhost",
		Port: port,
		Check: &consul.AgentServiceCheck {
			HTTP: fmt.Sprintf("http://localhost:%d/health", port),
			Interval: "10s",
			Timeout: "2s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("register service failed: %s\n", err.Error())
	}
	
	fmt.Println("Service registered successfully!")
}
