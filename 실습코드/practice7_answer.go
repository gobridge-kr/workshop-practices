package main

import "fmt"

func main() {
	/*
	카페에 전화를 건 손님에게 인사한 후 , 요일에 따라 마감시간을 안내하는 함수를 만들어 봅시다.
	월요일 마감시간 : 10시,화요일 마감시간 : 9시
	수요일 마감시간 : 10시,목요일 마감시간 : 9시
	금요일 마감시간 : 10시,토요일 마감시간 : 8시
	일요일 : 휴무
	*/
	var today string
	fmt.Print("오늘은 무슨요일인가요? ")
	fmt.Scan(&today)
	greeting()
	greeting2(today)
}

func greeting() {
	fmt.Println("전화해주셔서 감사합니다. GoCafe입니다.")
}

func greeting2(input string) {
	if input == "월요일" || input == "수요일" || input == "금요일" {
		fmt.Println("오늘의 마감 시간은 10시 입니다.")
	} else if input == "화요일" || input == "목요일" {
		fmt.Println("오늘의 마감 시간은 9시 입니다.")
	} else if input == "토요일" {
		fmt.Println("오늘의 마감 시간은 8시 입니다.")
	} else if input == "일요일" {
		fmt.Println("오늘은 휴무입니다.")
	}else{
		fmt.Println("*요일의 형태로 다시 입력해주세요.")
	}
}
