package main

import (
	"fmt"
	"go-server/utils"
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", utils.FormHandler)
	http.HandleFunc("/hello", utils.HelloHandler)

	fmt.Printf("Port server: 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
} 