package client

import (
	"fmt"
	"math/big"
)

type Client interface {
	CreateContainer(opts CreateContainerOptions) (bool, error)
	ListContainers() ([]string, error)
	SwarmMode() (bool, error)
}

type CreateContainerOptions struct {
	Name string
	Privileged bool
	Net string
}

type client struct {
	version string
}

func NewClient(version string) Client {
	return &client{version: version}
}


func (c *client) CreateContainer(opts CreateContainerOptions) (bool, error) {
	if err := c.requires(
		Requirement(Boolean(opts.Privileged, "privileged"), GT, "1.24"),
		Requirement(NotEmptyString(opts.Net, "net"), GT, "1.25"),
	); err != nil {
		return false, err
	}

	fmt.Printf("CreateContainer(%s)\n", opts.Name)
	return true, nil
}

func (c *client) ListContainers() ([]string, error) {
	fmt.Printf("ListContainers()\n")
	return nil, nil
}

func (c *client) SwarmMode() (bool, error) {
	if err := c.requires(Requirement(Feature("SwarmMode"), GT, "1.25")); err != nil {
		return false, err
	}

	fmt.Printf("SwarmMode()\n")
	return true, nil
}



type VersionComparator func(version string, required string) error

type versionReq struct {
	Condition Condition
	Comparator VersionComparator
	Requirement string
}

type Condition func() error

func Boolean(value bool, field string) (func() error) {
	return func() error {
		if value {
			return fmt.Errorf(field)
		}
		return nil
	}
}

func NotEmptyString(value string, field string) (func() error) {
	return func() error {
		if value != "" {
			return fmt.Errorf(field)
		}
		return nil
	}
}

func Feature(name string) (func() error) {
	return func() error {
		return fmt.Errorf(name)
	}
}

func Requirement(cond Condition, comp VersionComparator, req string) versionReq {
	return versionReq{Condition: cond, Comparator: comp, Requirement: req}
}

func (c *client) requires(reqs ...versionReq) error {
	for _, req := range reqs {
		condErr := req.Condition()
		if condErr == nil {
			continue
		}
		if err := req.Comparator(c.version, req.Requirement); err != nil {
			return fmt.Errorf(
				"%s %s, and the API version is %s",
				condErr.Error(),
				err.Error(),
				c.version,
			)
		}
	}
	return nil
}

func GT(version string, required string) error {
	if parseVersion(version).Cmp(parseVersion(required)) < 0 {
		return fmt.Errorf("requires at least version %s", required)
	}
	return nil
}

func LT(version string, required string) error {
	if parseVersion(version).Cmp(parseVersion(required)) >= 0 {
		return fmt.Errorf("was removed in version %s", required)
	}
	return nil
}

func parseVersion(version string) *big.Float {
	float := new(big.Float)
	var err error
	float, _, err = float.Parse(version, 10)
	if err != nil {
		panic(err.Error())
	}
	return float
}
