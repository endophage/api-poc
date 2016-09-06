package v124

import "fmt"

type Client struct{}

func (Client) CreateContainer(name string, privileged bool, _ string) bool {
	fmt.Printf("v124.Client::CreateContainer(%s, %t)\n", name, privileged)
	return true
}

func (Client) ListContainers() []string {
	fmt.Printf("v124.Client::ListContainers()\n")
	return nil
}

func (Client) SwarmMode() bool {
	fmt.Printf("v124.Client::NotImplemented\n")
	return false
}
