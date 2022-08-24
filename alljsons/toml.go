package alljsons

import (
	"io/ioutil"
	"log"

	"github.com/pelletier/go-toml/v2"
)

func Toml(){

	source,err := ioutil.ReadFile("example.toml")
	if err != nil{
		log.Println(err)
	}

	a := make(map[string]interface{})

	err = toml.Unmarshal(source, &a)
	if err != nil {
    log.Println(err)
	}

	for k,v := range a {
		log.Printf("%s : %v",k,v )
	}
	
}