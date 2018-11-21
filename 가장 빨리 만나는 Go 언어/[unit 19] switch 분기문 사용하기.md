# 19. switch 분기문 사용하기

switch 를 사용하여 더 간단하게 조건을 표현할 수 있다. 

##### 특징

다른 프로그래밍과 달리 case 에서 break 를 생략한다 

각 `case` 는 순서대로 값을 판단하고, 값이 일치하면 해당 코드 실행한 뒤 `switch` 분기문을 중단한다. 

다 해당되지 않을 경우 `default:` 를 실행한다. 

case 는 숫자 뿐만 아니라 문자열도 값으로 사용할 수 있다.

- **switch 변수 { case 값: 코드 }**

  ```go
  switch 변수 {
    case a:
      // 값1
    case b:
      // b
    case c:
      // c
    default:
      // default
  }
  ```

- sample

  ```go
  i := 3
  switch 변수 {
    case 0:
      fmt.Println(0)
    case 1:
      fmt.Println(1)
    case 2:
      fmt.Println(2)
    case 3:
      fmt.Println(3)
    case 4:
      fmt.Println(2)
    default:
      fmt.Println(-1)
  }
  ```

  ##### 실행결과

  ```
  3
  ```



## 19.1 break 사용하기

if 조건문 안에서 break 키워드를 사용하여 문장 실행을 중단시킬 수 있다. 

- 문장 실행 중단 

  ```go
  s := "Hello"
  i := 2
  
  switch i {
  	case a:
     	  fmt.Println(1)
      case 2:
        if s == "Hello" {
          fmt.Println("Hello 2")
  		break // siitch 분기문 실행을 중단한다. 
        }
        fmt.Println(2) // 출력 안됨 
  	default :
  	  fmt.Println(-1)
  }
  ```

  ##### 실행결과

  ```
  Hello 2
  ```


## 19.2 fallthrough 사용하기

**fallthrough** 특징

case 문장 실행 뒤 아래 case 문장을 실행하고 싶을 때 사용한다. 

단, 맨 마지막 case 에는 사용할 수 없다

- example 

  ```go
  i := 3
  switch i {
  	case 4:
        fmt.Println("4 이상")
        fallthrough
      case 3:
        fmt.Println("3 이상") // 변수와 일치 
        fallthrough 
      case 2:
        fmt.Println("2 이상") // 실행
        fallthrough 
      case 1:
        fmt.Println("1 이상") // 실행
        fallthrough 
      case 0:
        fmt.Println("0 이상") // 실행 
        // 마지막 case 에서는 fallthrough 사용 불가 
  }
  ```

  ##### 실행결과

  ```
  3이상
  2이상
  1이상
  0이상
  ```


## 19.3 여러 조건을 함께 처리하기

여러 조건을 같은 문장으로 처리하고 싶은경우 case 에서 `,` 로 값을 구분해준다.

- __case 값1,값2,값3__

  ```go
  i := 3
  
  switch i {
      case 2, 4, 6:
  	    fmt.Println("짝수")
      case 1, 3, 5:
  	    fmt.Println("홀수")
  }
  ```

  ##### 실행결과

  ```
  홀수 
  ```


## 19.4 조건식으로 분기하기 

**switch** 키워드 이후 변수 지정 없이, case 에서 조건식으로만 문장을 실행할 수 있다.

- example

  ```go
  i :=3
  
  switch {
      case i==2 && i==4 && i==6:
  	    fmt.Println("짝수")
      case i==1 && i==3 && i==5:
  	    fmt.Println("홀수")
  }
  ```


분기문 안에서 함수 실행 후 결괏값으로 분기를 할 수 있다. 

이때는 호출 하고 `;` 를 붙혀 주어야 한다 

case 에서는 값으로 분기할 수 없고, 조건식만 사용할 수 있다. 



- example 

  ```go
  func main() {
    rand.Seed(time.Now().UnixNano())
    switch i:= rand.Intn(10); {
      case i >= 3 && i <6:
        //
      case i == 9:
        //
      default:
        //
  }
  ```
