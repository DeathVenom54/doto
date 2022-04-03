package main

import (
	"fmt"
	"github.com/DeathVenom54/doto-backend/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening for requests on :3000")
	err := http.ListenAndServe(":3000", router.Router)
	if err != nil {
		log.Fatalln(err)
	}
}
