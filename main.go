package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {

	server_url := "localhost:22222"

	http.HandleFunc("/", indexAction)

	fmt.Println("Starting server at " + server_url)

	fmt.Println("Trying to open browser..")
	// opening server in browser
	commands := []string{"start", "google-chrome", "firefox"}
	for _, command := range commands {
		cmd := exec.Command(command, "http://"+server_url)
		err := cmd.Run()
		if err == nil {
			break
		}
	}

	err := http.ListenAndServe(server_url, nil)
	if err != nil {
		panic("Error when starting server at " + server_url)
	}

}
