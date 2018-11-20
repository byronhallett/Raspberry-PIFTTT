package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func callProgram(name string) {
	cmd := exec.Command(name)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Command failed to run")
		return
	}
	fmt.Println("Finished running", name)
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error in parsing")
		return
	}
	scriptName := string(body)
	callProgram(scriptName)
}

func main() {
	port, found := os.LookupEnv("IFTTT_PORT")
	if !found {
		log.Fatal("Port not found")
	}
	fmt.Println("Serving on port", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
