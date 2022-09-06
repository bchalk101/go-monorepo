package handler

import (
	"fmt"

	"github.com/bchalk101/go-monorepo/lib"
)

func Handler() {
	help := lib.HelpOut()
	fmt.Printf("Hello %v\n", help)
}
