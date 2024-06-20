package main

import "fmt"

//ประกาศตัวแปล array
// var productName [4]string
// var price [4]float32

func main() {

	//ประกาศตัวแปล array พร้อมใช้งาน

	productName := [2]string{"Macbook", "iPhone"}
	price := [4]float32{20000, 30000, 20000, 2000}

	fmt.Println(productName)
	fmt.Println(price)
}
