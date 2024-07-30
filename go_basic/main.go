package main

import "fmt"

//นำเอา package fmt มาใช้งาน

//bool เป็น Boolean
//int uint เป็น number มีการเก็บจำนวน bit ยิ่งมากเท่าไหร่ยิ่งเก็บค่าได้มากเท่านั้น เช่น int8, int16, int32, int64 และ uint8, uint16, uint32, uint64, unintptr uint คือ number ที่เก็บค่าจำนวนเต็มที่ไม่มีเครืองหมาย
//string เป็น string
//float เป็น float 32 และ float 64

func main() {
	// fmt.Println("Hello, World!")

	//นิยามตัวเเปร--------------------
	// var name string = "John"
	// name = "Chin"
	// age := 27 //var age int = 27
	// var score float32 = 3.5
	// var isPass bool = true

	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(score)
	// fmt.Println(isPass)

	// //นิยามค่าคงที่
	// const pi float32 = 3.14
	// // pi = 3.41 มันจะ error คือ ไม่สามารถเปลี่ยนค่าได้
	// fmt.Println(pi)

	// //แสดงชนิดข้อมูล ใช้ Printf และ ใช้ \n ในการขึ้นบรรทัดใหม่---------------------------------
	// fmt.Printf("Type of name = %T \n", name) //%T แสดงชนิดข้อมูล
	// fmt.Printf("My name is %v \n", name) //%v แสดงค่า

	//การดำเนินการทางคณิตศาสตร์-------------------------------------------
	// num1 := 10
	// num2 := 5
	// var num1 , num2 = 10, 5
	// num1, num2 := 10,5
	// fmt.Println(num1, num2)
	// fmt.Println("Plus =", num1 + num2)
	// fmt.Println("Minus =", num1 - num2)
	// fmt.Println("Substract =", num1 * num2)
	// fmt.Println("Divide =", num1 / num2)
	// fmt.Println("Mod =", num1 % num2)

	// //การเปรียบเทียบ-------------------------------------------
	// fmt.Println(num1 ,">", num2, "=", num1 > num2)
	// fmt.Println(num1 ,"<", num2, "=", num1 < num2)
	// fmt.Println(num1 ,"=", num2, "=", num1 == num2)
	// fmt.Println(num1 ,"!=", num2, "=", num1 != num2)
	// fmt.Println(num1 ,">=", num2, "=", num1 >= num2)
	// fmt.Println(num1 ,"<=", num2, "=", num1 <= num2)

	//รับค่าจากคีย์บอร์ด-------------------------------------------
	//fmt.Scanf(string_format,address_list)
	//string_format คือ รูปแบบตัวเเทนชนิดข้อมูล string: %s, int: %d, float: %f
	//address_list คือ ตัวเก็บข้อมูล
	// var score int
	// fmt.Print("Enter your score: ")
	// fmt.Scanf("%d", &score)

	// fmt.Println("Your score is", score)

	// //Condition if else----------------------------------------
	// if score >= 60 {
	// 	fmt.Println("Pass")
	// } else {
	// 	fmt.Println("Fail")
	// }

	// if score % 2 == 0 {
	// 	fmt.Println("Your Score is Even")
	// } else {
	// 	fmt.Println("Your score is Odd")
	// }

	//Array----------------------------------------
	var numbers [5]int
	fmt.Println(numbers)
}