package v125

import (
	"fmt"

	"github.com/vieux/api-poc/option2/v124"
)

type Client interface {
	CreateContainer(name string, privileged bool, net string) bool
	ListContainers() []string
	SwarmMode() bool
}

type v125ClientImplem struct{}

func (v125ClientImplem) CreateContainer(name string, privileged bool, net string) bool {
	fmt.Printf("v125ClientImplem::CreateContainer(%s, %t, %s)\n", name, privileged, net)
	return true
}

func (v125ClientImplem) ListContainers() []string {
	fmt.Printf("v125ClientImplem::ListContainers()\n")
	return nil
}

func (v125ClientImplem) SwarmMode() bool {
	fmt.Printf("v125ClientImplem::SwarmMode()\n")
	return true
}

type v125to124Adapter struct {
	impl v124.Client
}

func (a v125to124Adapter) CreateContainer(name string, privileged bool, net string) bool {
	fmt.Printf("v125to124Adapter::CreateContainer(%s, %t, %s)\n", name, privileged, net)
	return a.impl.CreateContainer(name, privileged)
}

func (a v125to124Adapter) ListContainers() []string {
	fmt.Printf("v125to124Adapter::ListContainers()\n")
	return a.impl.ListContainers()
}

func (a v125to124Adapter) SwarmMode() bool {
	fmt.Printf("v125to124Adapter::SwarmMode()\n")
	return false
}

func NewClient(backendVersion string) (Client, error) {
	if backendVersion == "1.25" {
		return v125ClientImplem{}, nil
	}

	// Should be: if older than "1.25"
	fallback, err := v124.NewClient(backendVersion)
	if err != nil {
		return nil, err
	}
	return v125to124Adapter{fallback}, nil
}
