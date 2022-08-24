package alljsons

import (
	"log"

	"github.com/spf13/viper"
)


func All(){

	vi := viper.New()
	vi.SetConfigFile("items.yaml")
	vi.ReadInConfig()
	 a := vi.GetString("user1.name")
	 b := vi.GetString("user2.name")

	 log.Println(a,b)

}