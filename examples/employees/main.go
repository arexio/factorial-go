package main

import (
	"fmt"
	"log"

	factorial "github.com/charly3pins/factorial-go"
)

func main() {
	cli, err := factorial.New("ACCESS_TOKEN")
	if err != nil {
		log.Println(err)
		return
	}

	employees, err := cli.ListEmployees()
	if err != nil {
		log.Println(err)
		return
	}

	for _, e := range employees {
		fmt.Printf("%+v\n", e)
	}
}
