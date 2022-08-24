package alljsons

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"log"
)

type Users struct{
 
	Users [] User1 `json:"users"`

}

type User1 struct{
	Name string    `json:"name"`
	Class int      `json:"class"`
	School string   `json:"school"`
}


func Json() {
	source, err := ioutil.ReadFile("example.json")
	if err != nil {
		log.Println(err)
	}

	var users Users 

     json.Unmarshal(source, &users)

	for i :=0 ; i < len(users.Users) ; i++{
		log.Println("user name :", users.Users[i].Name)
		log.Println("user class :", strconv.Itoa(users.Users[i].Class))
		log.Println("user school :", users.Users[i].School)

	}
  
}