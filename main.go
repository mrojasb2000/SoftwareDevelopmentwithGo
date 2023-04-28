package main

import (
	"log"
	"syscall"
)

func main() {
	log.Println("Getpid : ", syscall.Getpid())   // Obtains the process id of the current running app
	log.Println("Getpgrp : ", syscall.Getpgrp()) // Obtains the group process id of the current running app
	log.Println("Getppid : ", syscall.Getppid()) // Obtains the parent process id of the current running app
	log.Println("Gettid : ", syscall.Getppid())  // Obtains the caller's thread it

	wd, err := syscall.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(wd))
}
