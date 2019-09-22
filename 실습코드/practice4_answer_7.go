package main

import "fmt"

func main() {
  /*
	실습4) 7조 응용
	총 금액이 5000원 초과일 경우 10%를 할인해줍시다.
	총 금액이 5000원 초과이면서, 월요일인 경우엔 추가로 10%를 또 할인해줍니다.
	같은 음료를 4잔 이상 주문시 10% 할인

  7조 응용 코드 추가 (for문 응용)
  - 두 음료 모두 0잔을 주문하면 다시 주문을 받습니다
	*/
  var cafeName string = "GoCafe"
  var isTodayMonday bool = true
  drinkPrices := [2]int{1500, 1700}
  // 사용자에게 입력받을 주문 변수
  var americanoOrder int
  var latteOrder int

  fmt.Println("***어서오세요, " + cafeName + "입니다.***")
  fmt.Println("아메리카노 : ", drinkPrices[0], "원")
	fmt.Println("라떼 : ", drinkPrices[1], "원")
  fmt.Println("주문해주세요")
  // --------------------

  for {
    fmt.Println("아메리카노 몇 잔 드릴까요? ")
    fmt.Scan(&americanoOrder)

    fmt.Println("라떼 몇 잔 드릴까요? ")
    fmt.Scan(&latteOrder)

    if americanoOrder != 0 || latteOrder != 0 {
      break
    }else{
      fmt.Println("주문을 다시해주세요")
    }
  }

  totalPrice := (drinkPrices[0] * americanoOrder) + (drinkPrices[1] * latteOrder)

  if totalPrice > 5000 {
    totalPrice = int(float32(totalPrice) * 0.9)
  }
  if totalPrice > 5000 && isTodayMonday {
    totalPrice = int(float32(totalPrice) * 0.9)
  }
  if americanoOrder >= 4 || latteOrder >= 4 {
    totalPrice = int(float32(totalPrice) * 0.9)
  }

  fmt.Println("총 금액 ", totalPrice)
}
