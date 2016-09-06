package client

type Client interface {
	CreateContainer(name string, privileged bool, net string) bool
	ListContainers() []string
	SwarmMode() bool
}
