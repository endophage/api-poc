package v123

import "fmt"

type Client interface {
	CreateContainer(name string) bool
	ListContainers() []string
}

type v123ClientImplem struct{}

func (v123ClientImplem) CreateContainer(name string) bool {
	fmt.Printf("v123ClientImplem::CreateContainer(%s)\n", name)
	return true
}

func (v123ClientImplem) ListContainers() []string {
	fmt.Printf("v123ClientImplem::ListContainers()\n")
	return nil
}

func NewClient(backendVersion string) (Client, error) {
	if backendVersion == "1.23" {
		return v123ClientImplem{}, nil
	}
	return nil, fmt.Errorf("Version %q is too old\n", backendVersion)
}
