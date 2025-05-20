package main

import (
    "fmt"
    "log"
    "github.com/Manavk2004/gator/internal/config"
)

func main(){
	cfg, err := config.Read()
	if err != nil{
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("Manav")
	if err != nil{
		log.Fatalf("Couldnt set current user: %v", err)
	}

	cfg, err := config.Read()
	if err != nil{
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}