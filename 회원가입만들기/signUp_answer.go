package main
import "fmt"
/*
* 아이디 중복체크
* 입력한 아이디가 이미 있는 아이디인지 확인합니다.
* 아이디는 ids에 저장되어있습니다.
* 중복이라면 false를, 아니라면 true를 리턴합니다.
*/
func duplicateCheck(value string, ids []string) bool {
   valid := true
   for _, id := range ids {
       if value == id {
           valid = false
           break
       }
   }
   return valid
}
/*
* 패스워드를 확인하는 함수
* 두 번 입력한 패스워드가 일치하는지 검사합니다.
* 일치하면 true를, 아니라면 false를 리턴합니다.
*/
func pwdCheck(pwd1 string, pwd2 string) bool {
   valid := true
   if pwd1 != pwd2 {
       valid = false
   }
   return valid
}
/*
* 생년월일을 확인하는 함수
* 생년월일이 8자리인지 검사합니다.
* 8자리가 맞으면 true를, 아니라면 false를 리턴합니다.
*/
func dateCheck(date string) bool {
   valid := true
   if len(date) != 8 {
       valid = false
   }
   return valid
}
/*
* 정보를 출력해주는 함수
* 아이디, 이름, 생년월일, 연락처를 프린트합니다.
* 개인정보 보호를 위해 마스킹처리 합니다. (심화)
* 아이디는 4글자까지 보여주고 나머지는 ****처리
* 연락처는 010-1234-1234 -> 010-****-1234 로 처리
*/
func printInfo(id, name, birthdate, phone string){
    maskingId := id[0:4]
    masingPhone := phone[0:4] + "****" + phone[8:]
    for i := 0; i < len(id) - 4; i++ {
        maskingId += "*"
    }
    fmt.Println("아이디 : ", maskingId)
    fmt.Println("이름 : ", name)
    fmt.Println("생년월일 : ", birthdate)
    fmt.Println("연락처 : ", masingPhone)
    fmt.Println("-----------------------------------")
}
func main() {
   ids := []string{"parrot0707", "mylove1004", "lkm1154", "test1234"}
   exit := 1
   for exit != 0 {
       var id string        // 아이디를 받는 변수
       var pwd1 string      // 비밀번호를 받는 변수
       var pwd2 string      // 비밀번호 재입력을 받는 변수
       var name string      // 이름을 받는 변수
       var birthdate string // 생년월일을 받는 변수
       var phone string     // 전화번호를 받는 변수
       var valid bool // 유효한 값인지 확인하는 변수
        fmt.Println("---------- GoBridge 회원가입 -------------")
        //1. 아이디
       fmt.Print("아이디: ")
        fmt.Scan(&id)
        //아이디 중복체크
       valid = duplicateCheck(id, ids)
       if valid == false {
           fmt.Println("이미 존재하는 아이디입니다")
           continue
        }
        
        //2. 비밀번호
       fmt.Print("비밀번호를 입력하세요: ")
        fmt.Scan(&pwd1)
       fmt.Print("비밀번호를 재입력하세요: ")
        fmt.Scan(&pwd2)
        //비밀번호 확인 -> 두개가 같은지 확인
       valid = pwdCheck(pwd1, pwd2)
       if valid == false {
           fmt.Println("재입력한 비밀번호가 같지 않습니다")
           continue
        }
        //3. 이름
       fmt.Print("이름: ") // 이름 입력
        fmt.Scan(&name)
       if valid == false {
           fmt.Println("이름을 입력해주세요")
           continue
        }
        
        //4. 생년월일
       fmt.Print("생년월일(8자리로 입력): ") //8자리 19920125
        fmt.Scan(&birthdate)
        //생년월일 값의 유효값 검사
       valid = dateCheck(birthdate)
       if valid == false {
           fmt.Println("생년월일을 8자리로 입력해주세요")
           continue
        }
        
        //5. 연락처
       fmt.Print("연락처: ")
        fmt.Scan(&phone)
       if valid == false {
           fmt.Println("연락처를 입력해주세요")
           continue
        }
        
        //회원가입 완료된 아이디는 ids 배열에 추가합니다.
       ids = append(ids, id)
        fmt.Println("---------- 회원가입 성공 -------------")
        //회원가입 완료된 정보를 보여준다.
        printInfo(id, name, birthdate, phone)
       fmt.Println("종료하려면 0번을 계속하려면 1번을 입력해주세요")
       fmt.Scanf("%c")
       fmt.Scan(&exit)
   }
}
