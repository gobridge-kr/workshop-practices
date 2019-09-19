package main
import "fmt"
func main(){
	/*
	실습4)
	총 금액이 5000원 초과일 경우 10%를 할인해줍시다.
	총 금액이 5000원 초과이면서, 월요일인 경우엔 추가로 10%를 또 할인해줍니다.
	같은 음료를 4잔 이상 주문시 10% 할인
	*/
	isTodayMonday := true
	drinkNames := [2]string{"Americano", "Latte"}
	drinkPrices := [2]int{1500, 1700}
	var americanoOrder, latteOrder int

	fmt.Println("***어서오세요, GoCafe입니다.***")
	fmt.Println("아메리카노 : ", drinkPrices[0], "원")
	fmt.Println("라떼 : ", drinkPrices[1], "원")
	fmt.Print("아메리카노 몇 잔 드릴까요? ")	
	fmt.Scan(&americanoOrder)//아메리카노 주문량
	fmt.Print("라떼는몇 잔 드릴까요? ")

	fmt.Scan(&latteOrder)//라떼 주문량
	totalPrice := 0	     //총금액

	totalPrice = /*이 부분을 채워 넣으세요*/

	//5000원 초과이면 10% 할인
	if /*이 부분을 채워 넣으세요*/{
		/*이 부분을 채워 넣으세요*/
	}

	//월요일이면 추가할인
	if /*이 부분을 채워 넣으세요*/{
	/*이 부분을 채워 넣으세요*/
	}

	//같은 음료 4잔 이상이면 할인
	if /*이 부분을 채워 넣으세요*/ { 
			/*이 부분을 채워 넣으세요*/
	}
	fmt.Println(drinkNames[0],americanoOrder,"잔, ",drinkNames[1], latteOrder,"주문하셨습니다" )
	fmt.Println("최종 가격은 ", totalPrice, "원 입니다.")
}
