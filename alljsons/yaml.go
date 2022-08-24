package alljsons

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// when using YAML always add captial of all the letters in struct 
type User struct {
	Name string      // N will be capital  even if in yaml file it is small.
	Occupation string    // O will be capital  even if in yaml file it is small.
}

func Yaml(){
	//read the file 

	source,err := ioutil.ReadFile("items.yaml")
	if err != nil{
		log.Println(err)
	}

	//declaring map

	data := make(map[string]User)
	// unmarshall the data
    err2 := yaml.Unmarshal(source, &data) 
	if err2 != nil {
		log.Println(err)
	}

	for k,v := range data {
		log.Printf("%s : %s",k,v)
	}

}