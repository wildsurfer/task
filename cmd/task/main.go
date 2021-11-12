package main

import (
	"log"
	"main/internal/task"
)

func main() {
	log.Println("I'm ready!")
	log.Fatal(task.ListenAndServe().Error())
}
