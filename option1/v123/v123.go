package v123

import "fmt"

type Client struct{}

func (Client) CreateContainer(name string, _ bool, _ string) bool {
	fmt.Printf("v123.Client::CreateContainer(%s)\n", name)
	return true
}

func (Client) ListContainers() []string {
	fmt.Printf("v123.Client::ListContainers()\n")
	return nil
}

func (Client) SwarmMode() bool {
	fmt.Printf("v123.Client::NotImplemented\n")
	return true
}
