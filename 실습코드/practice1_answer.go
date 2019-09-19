package main

import "fmt"

func main(){
	var cafeName string = "GoCafe"
	var americanoPrice int = 1500
	var lattePrice int = 1700
	var cafePhoneNumber string = "02-428-5544"
	var isOpenToday bool = true
	
	fmt.Println("카페 이름은 " , cafeName , " 입니다.")
	fmt.Println("아메리카노 커피의 가격은 ", americanoPrice, " 입니다.")
	fmt.Println("라떼 커피의 가격은 ", lattePrice, " 입니다.")
	fmt.Println("카페 전화번호는 ", cafePhoneNumber, " 입니다.")
	fmt.Println("오늘 카페 여나요? " ,isOpenToday)
}