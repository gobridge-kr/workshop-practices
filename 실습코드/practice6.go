package main

import "fmt"

func main(){
	/*
	실습6)
	손님은, 2000원 이하인 음료만 주문하고 싶어합니다. 
	또한 총 금액이 7000원 이상 되면 더 이상 주문하지 않는다고 합니다.
	주문한 메뉴와, 최종 금액을 출력해주세요. 
	*/
  drinkNames := []string{"Americano", "Latte", "CafeMocha", "GreenTeaLatte",
    "HotChoco", "ChamomileTea", "StrawberrySmoothie", "HotMilk"}
  drinkPrices := []int{1500, 1700, 1900, 2300, 2000, 1800, 3500, 1500}

  totalPrice := 0 //총 금액
  orderDrinkes :="" //주문한 음료

  for /*이 부분을 채워 넣으세요 */ { //음료 배열의 갯수 만큼 반복
    if /*이 부분을 채워 넣으세요 */{ //음료 가격이 2000원 초과이면 주문 건너뜀.
      /*이 부분을 채워 넣으세요 */
    }
    if /*이 부분을 채워 넣으세요 */{ //총 금액이 7000원 이상이면 주문 종료.
      /*이 부분을 채워 넣으세요 */
    }
    totalPrice += /*이 부분을 채워 넣으세요 */
    orderDrinkes += drinkNames[i] + ", "
  }
  fmt.Println("주문하신 음료 ", orderDrinkes, "나왔습니다.")
  fmt.Println("총금액은", totalPrice, "원 입니다.")
}
