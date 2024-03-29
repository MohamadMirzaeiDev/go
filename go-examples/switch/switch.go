package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write " , i , " as ")
	switch i {
		case 1 :
			print("One")
		case 2 : 
			print("Two")
		case 3 : 
			print("Three")
	} 
	print("\n")


	switch time.Now().Weekday() {
		case time.Saturday , time.Sunday :
			fmt.Println("It's the weekend")
		default : 
			fmt.Println("It's the weekday")
	}

	t := time.Now()

	switch {
		case t.Hour() < 12 :
			fmt.Println("It's the before noon")
		default : 
			fmt.Println("It's the after noon")		
	}

	whatAmI := func (i interface{})  {
		switch t := i.(type) {
			case bool :
				fmt.Println("I'm a bool")
			case int : 
				fmt.Println("I'm a int")
			default : 
				fmt.Printf("This type not supported %T\n" , t)
		}
	}

	whatAmI(150)
	whatAmI(false)
	whatAmI("hey")

	var age int 

	fmt.Println("Please enter age : ")

	fmt.Scanf("%d" , &age)

	createMessage(age)
}


func createMessage(age int) {
	switch age {
		case 15:
			fmt.Println("hey thanks for use this pelase click " , age)
		case 18 : 
			fmt.Println("hey thanks for use this pelase click " , age)
		default : 
			fmt.Println("your age is" , age)
	}
}