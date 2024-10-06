package main 

import (
	"fmt"
	"github.com/AarizZafar/Nexus_verify.git/router"
)

func main() {
	port := ":8080"
	router := router.Router()

	fmt.Println(">>>>>>>>>>>>>>>>>> Starting server <<<<<<<<<<<<<<<<<<")
	fmt.Printf(">>>>>>>>>>>>>>>>>> Listening at port %s <<<<<<<<<<\n", port)

	router.Run(port)
}