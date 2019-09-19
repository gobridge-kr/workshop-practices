package main

import "fmt"

// 실습 8
type Menu struct {
	name string
	price int
	isCoffee bool
}


// 실습 9
type Customer struct {
	name string
	age int
	phoneNumber string
	isVip bool
	point int
}

var menus	[5]Menu
var membership	[5]Customer

// 실습 10
func main() {
	menus[0] = Menu{"Americano", 1500, true}
	menus[1] = Menu{"Latte", 1700, true}
	menus[2] = Menu{"CafeMocha", 1900, true}
	menus[3] = Menu{"GreenTeaLatte", 2300, false}
	menus[4] = Menu{"HotChoco", 2000, false}

	membership[0] = Customer{"Tom", 23, "010-3212-5578", true, 3200}
	membership[0] = Customer{"Sam", 17, "010-987-5959", true, 2000}
	membership[0] = Customer{"Kim", 42, "010-529-1982", false, 1000}

	// TODO: 여기를 채워 넣으세요
	
}

func printLowPrice() {
	fmt.Println("2000원 미만 메뉴는 다음과 같습니다")
	// TODO: 여기를 채워 넣으세요
	
}

func printCoffee() {
	fmt.Println("커피 메뉴는 다음과 같습니다")
	// TODO: 여기를 채워 넣으세요
	
}

func order(menuNum int, customerNum int) int {
	totalPrice := 0
	// TODO: 메뉴의 가격을 계산, VIP이면 할인률 적용

	// TODO: 포인트 적립 함수 호출
	
	return totalPrice
}

func addPoints(price int, customerNum int) {
	// TODO: 포인트 적립
	
	// TODO: 포인트 적립이 된 후 조건에 맞다면 VIP 승격
	
}