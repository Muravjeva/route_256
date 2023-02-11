package main

import (
	"fmt"
	"github.com/Muravjeva/route_256/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
		return
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it restored", restoredFile)
}
