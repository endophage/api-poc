package main

import (
	"fmt"

	"github.com/vieux/api-poc/option3/client"
)

func main() {
	for _, v := range []string{"1.25", "1.24", "1.23", "1.22"} {
		fmt.Printf("[+] Testing against version %q\n", v)
		c := client.NewClient(v)
		if _, err := c.CreateContainer(client.CreateContainerOptions{"foo", true, "host"}); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
		if _, err := c.ListContainers(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
		if _, err := c.SwarmMode(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}
}
