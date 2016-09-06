package client

import (
	"fmt"

	"github.com/vieux/api-poc/option1/v123"
	"github.com/vieux/api-poc/option1/v124"
	"github.com/vieux/api-poc/option1/v125"
)

type Client interface {
	CreateContainer(name string, privileged bool, net string) bool
	ListContainers() []string
	SwarmMode() bool
}

func NewClient(backendVersion string) (Client, error) {
	switch backendVersion {
	case "1.25":
		return v125.Client{}, nil
	case "1.24":
		return v124.Client{}, nil
	case "1.23":
		return v123.Client{}, nil
	}
	return nil, fmt.Errorf("Version %q is too old\n", backendVersion)
}
