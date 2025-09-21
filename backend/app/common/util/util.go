// Date:   Sat Jun 28 20:33:40 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package util

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"os"

	consul "github.com/hashicorp/consul/api"
)

func GetLocalIpAddr() (ip string) {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ip = ipnet.IP.String()
			break
		}
	}
	return ip
}

func GetEnv(key, defaultVal string) string {
  if val, ok := os.LookupEnv(key); ok {
    return val
  }
  return defaultVal
}

func ConsulRegisterService(id string, name string, addr string) {
	config := consul.DefaultConfig()
	config.Address = fmt.Sprintf("%s:8500", GetEnv("CONSULADDR", "localhost"))

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
		log.Fatalf("register service at (%s) failed: %s\n", addr, err.Error())
	}
	
	fmt.Printf("Service %s registered successfully!\n", id)
}
