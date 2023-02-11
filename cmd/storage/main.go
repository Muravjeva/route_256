package main

import (
	"fmt"
	"github.com/Muravjeva/route_256/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("Storage works", st)
}
