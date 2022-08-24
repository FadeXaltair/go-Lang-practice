package alljsons

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(){

 godotenv.Load("go.env")

 number :=  os.Getenv("Number")

 expermimental_value := os.Getenv("Experimental value")



 log.Println(number, expermimental_value)

}