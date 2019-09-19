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

	printLowPrice()
	printCoffee()

	fmt.Println("가격은 ", order(2,0), "원 입니다.")
	fmt.Println(membership[0].name, "님의 잔여 포인트는 ", membership[0].point, "입니다.")
}

func printLowPrice() {
	fmt.Println("2000원 미만 메뉴는 다음과 같습니다")
	for _, menu := range menus {
		if menu.price < 2000 {
			fmt.Println(menu.name)
		}
	}
}

func printCoffee() {
	fmt.Println("커피 메뉴는 다음과 같습니다")
	for _, menu := range menus {
		if menu.isCoffee == true {
			fmt.Println(menu.name)
		}
	}
}

func order(menuNum int, customerNum int) int {
	totalPrice := 0
	if (membership[customerNum].isVip) {
		totalPrice = int(float32(menus[menuNum].price) * 0.9)
	} else {
		totalPrice = menus[menuNum].price
	}

	//포인트 적립
	addPoints(totalPrice, customerNum)
	return totalPrice
}

func addPoints(price int, customerNum int) {
	membership[customerNum].point += int(float32(price) * 0.1)
	if (membership[customerNum].point >= 2000) {
		membership[customerNum].isVip = true
	}
}
