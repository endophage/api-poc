package v125

import "fmt"

type Client struct{}

func (Client) CreateContainer(name string, privileged bool, net string) bool {
	fmt.Printf("v125.Client::CreateContainer(%s, %t, %s)\n", name, privileged, net)
	return true
}

func (Client) ListContainers() []string {
	fmt.Printf("v125.Client::ListContainers()\n")
	return nil
}

func (Client) SwarmMode() bool {
	fmt.Printf("v125.Client::SwarmMode()\n")
	return true
}
