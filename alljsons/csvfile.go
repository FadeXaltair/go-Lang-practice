package alljsons

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Data struct{
	Name string
	Country string
	Age int
}


func Csv(){
source,err := os.Open("data.csv")
if err != nil{
	log.Println(err)
}

csvReader := csv.NewReader(source)

for{
	rec,err :=csvReader.Read()
	if err ==io.EOF{
		break
	}
	if err != nil{
		log.Println(err)
	}

	log.Printf("%v\n", rec)
}
}