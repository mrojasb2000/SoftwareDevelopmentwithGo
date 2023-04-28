package main

import (
	u "golang.org/x/sys/unix"
	"log"
)

func main() {
	c := make([]byte, 512)

	log.Println("Getpid : ", u.Getpid())   // Obtains the process id of the current running app
	log.Println("Getpgrp : ", u.Getpgrp()) // Obtains the group process id of the current running app
	log.Println("Getppid : ", u.Getppid()) // Obtains the parent process id of the current running app
	log.Println("Gettid : ", u.Getppid())  // Obtains the caller's thread it

	_, err := u.Getcwd(c)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(c))
}
