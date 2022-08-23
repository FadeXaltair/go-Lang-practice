package projects

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func If_Else() {

      godotenv.Load("go.env")

	number := os.Getenv("Number")
   var guess string 
	fmt.Print("enter your guess number : ")
    fmt.Scan(&guess)
	if guess <number{
		fmt.Println("low guess ")
	}
    if guess > number{
		fmt.Println("large number")
	}

	if guess == number {
     fmt.Println("you are right")
	}

	// fmt.Println(guess <= number , guess >=number , guess !=number)
}

// < > == <= >= !=