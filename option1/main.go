package main

import (
	"fmt"

	"github.com/vieux/api-poc/option1/client"
)

func main() {
	for _, v := range []string{"1.25", "1.24", "1.23", "1.22"} {
		fmt.Printf("[+] Testing against version %q\n", v)
		c, err := client.NewClient(v)
		if err != nil {
			fmt.Printf("Got error: %v\n", err)
		} else {
			c.CreateContainer("foo", true, "host")
			c.ListContainers()
			c.SwarmMode()
		}
	}
}
