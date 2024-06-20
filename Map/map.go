package main

import "fmt"

var product = make(map[string]float32)

func main() {
	fmt.Println(product)

	//add ข้อมูล
	product["Macbook"] = 40000
	product["Ipad"] = 40000
	fmt.Println(product)

	//ลบข้อมูล
	delete(product, "Ipad")
	fmt.Println(product)

	//update
	product["Ipad"] = 20000
	fmt.Println(product)

	value1 := product["Iphone"]
	fmt.Println(value1)

	courseName := map[string]string{"101": "Java", "102": "python"}
	fmt.Println(courseName)

}
