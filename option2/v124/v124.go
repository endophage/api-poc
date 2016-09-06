package v124

import (
	"fmt"

	"github.com/vieux/api-poc/option2/v123"
)

type Client interface {
	CreateContainer(name string, privileged bool) bool
	ListContainers() []string
}

type v124ClientImplem struct{}

func (v124ClientImplem) CreateContainer(name string, privileged bool) bool {
	fmt.Printf("v124ClientImplem::CreateContainer(%s, %t)\n", name, privileged)
	return true
}

func (v124ClientImplem) ListContainers() []string {
	fmt.Printf("v124ClientImplem::ListContainers()\n")
	return nil
}

type v124to123Adapter struct {
	impl v123.Client
}

func (a v124to123Adapter) CreateContainer(name string, privileged bool) bool {
	fmt.Printf("v124to123Adapter::CreateContainer(%s, %t)\n", name, privileged)
	return a.impl.CreateContainer(name)
}

func (a v124to123Adapter) ListContainers() []string {
	fmt.Printf("v124to123Adapter::ListContainers()\n")
	return a.impl.ListContainers()
}

func NewClient(backendVersion string) (Client, error) {
	if backendVersion == "1.24" {
		return v124ClientImplem{}, nil
	}

	// Should be: if older than "1.24"
	fallback, err := v123.NewClient(backendVersion)
	if err != nil {
		return nil, err
	}
	return v124to123Adapter{fallback}, nil
}
