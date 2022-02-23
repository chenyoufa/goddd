package main

import (
	"fmt"
	_ "gdemo1/docs"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	defer func() {
		defer func() {
			fmt.Println("q", recover())
		}()

	}()
	fmt.Println("a")
	panic(1)

	// defer func() {
	// 	fmt.Println("c")
	// 	fmt.Println(recover())
	// }()

	// defer func() {
	// 	defer func() {
	// 		fmt.Println("q", recover())
	// 	}()
	// 	fmt.Println("b")
	// 	panic(1)
	// }()

	// defer recover()
	// fmt.Println("a")
	// panic(2)

}
