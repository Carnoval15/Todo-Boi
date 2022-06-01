package database

import (
	"encoding/json"
	"fmt"
	"os"

	//"github.com/Carnoval15/Todo-Boi/cmd"
)

type Task struct {
	Todo string
}

func Add(task string) {

	input := Task{Todo: task}

	data, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.OpenFile("./data.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	n, err := f.Write(data)
	if n, err = f.WriteString("\n"); err != nil {
		fmt.Println(n, err)
	}

	fmt.Println(string(data))

}
