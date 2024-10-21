package main 

import (
	"fmt"
	"github.com/AarizZafar/Nexus_verify.git/router"
)

func main() {
	port := ":8080"
	router := router.Router()

	fmt.Println("\033[97;46m       >>>>>>>>>>>> STARTING SERVER <<<<<<<<<<<<<<<<<<            \033[0m")
    fmt.Printf("\033[97;46m       >>>>>>>>>>>> LISTENING AT PORT %s <<<<<<<<<<            \033[0m\n\n", port)

	router.Run(port)
}