package usecases

import (
	"fmt"
)

type Core struct {
}

func New() *Core {
	fmt.Println("Initializing UseCases...")

	return &Core{}
}

func (c *Core) Cleanup() {
	fmt.Println("Cleaning UseCases...")
}
