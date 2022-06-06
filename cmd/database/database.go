/*
Copyright © 2022 Parsa <carnoval@protonmail.com>

*/
package database

import (
	"fmt"
	"log"
	"os"
)

func Add(task string) {
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(task + "\n"); err != nil {
		panic(err)
	}
}

func Read() {
	body, err := os.ReadFile("data.txt")
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }
    fmt.Println("\n" + string(body))
}
