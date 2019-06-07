package main

import (
	"fmt"

	"github.com/MaxHalford/eaopt"
)

func main() {
	ga, err := eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		panic(err)
	}
	fmt.Println(ga)
}
