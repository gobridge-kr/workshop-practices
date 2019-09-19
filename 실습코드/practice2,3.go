package main

import "fmt"

func main() {
	/*
	실습2) 
	GoCafe에는 메뉴가 아메리카노, 라떼 두 개 있습니다.  
	이 두 메뉴의 이름이 저장된 배열(drinkNames)과 가격이 저장된 배열(drinkPrices)을 정의해주세요. 

	*/
	var drinkNames [2]string
	drinkNames[0] = "Americano"
	drinkNames[1] = "Latte"

	var drinkPrices [2]int
	drinkPrices[0] = 1500 //아메리카노 가격
	drinkPrices[1] = 1700 // 라떼 가격
	fmt.Println("***어서오세요, GoCafe입니다.***")
	fmt.Println("아메리카노 : ", drinkPrices[0], "원")
	fmt.Println("라떼 : ", drinkPrices[1], "원")
	/*
	실습 3) 어떤 손님이 아메리카노 M잔과 라떼 N잔을 주문하였습니다. 총 가격(cost)을 출력해주세요.
	오늘은 개업날이 10% 할인을 해줍시다. 할인된 가격을 출력해주세요.
	(Hint! 정수와 소수를 계산하고 싶을 때에는 명시적형변환을 해주어야 합니다.)
	*/	
	var cost int = 0
	var M, N int
	fmt.Print("아메리카노 몇 잔 드릴까요? ")	
	fmt.Scan(&M)
	fmt.Print("라떼는몇 잔 드릴까요? ")
	fmt.Scan(&N)
	//아메리카노 M잔, 라떼 N잔의 가격구하기
	cost = /*이 부분을 채워 넣으세요*/ 

	fmt.Println("가격은 ", cost, "원 입니다")
	fmt.Println("할인 금액은 ", /*이 부분을 채워 넣으세요*/ , "원 입니다.")
}
