package main

import (
	"log"
)

//Run program
func main() {
	srv := new(server.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("Error port", err.Error())
	}

}

/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Alex")
	})
	http.ListenAndServe(":8081", nil)

*/
