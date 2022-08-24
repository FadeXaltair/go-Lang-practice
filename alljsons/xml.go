package alljsons

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)


type Userss struct{
   XmlName xml.Name `xml:"users"`
   Userss [] User2   `xml:"user"`
}

type User2 struct{
	Name string   `xml:"name"`
	Class int     `xml:"class"`
	School string  `xml:"school"`
}

func Xml(){
  
	source,err := ioutil.ReadFile("go.xml")
	if err != nil {
		log.Println(err)
	}

	var userss Userss 

	err = xml.Unmarshal(source, &userss)

	if err != nil {
		log.Println(err)
	}

	for i :=0; i <len(userss.Userss); i++{

		log.Println("User name is : ", userss.Userss[i].Name )
		log.Println("User class is : ", userss.Userss[i].Class )
		log.Println("User school is : ", userss.Userss[i].School )
	}

}