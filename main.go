package main

import (
	"fmt"

	"github.com/chonlatee/simplecache/cache"
)

func main() {

	c := cache.New(4)

	c.Set("k1", 1)
	c.Set("k2", 2)
	c.Set("k3", 3)
	c.Set("k4", 4)

	v := c.Get("k2")
	fmt.Printf("key: k2, value: %v\n", v)
	c.Print()
	c.Printlist()

	fmt.Printf("-----\n\n\n")

	c.Set("k5", 5)
	c.Print()
	c.Printlist()

	fmt.Printf("-----\n\n\n")

	v = c.Get("k3")
	fmt.Printf("key: k3, value: %v\n", v)
	c.Print()
	c.Printlist()

	fmt.Printf("-----\n\n\n")
	c.Set("k2", 777)
	c.Print()
	c.Printlist()
}
