- Go 언어는 정수, 실수 (부동소수점), 복수수를 지원

- 정수

  - 8진수: 숫자 앞에 0을 붙임

  - 16진수: 숫자 앞에 0x 또는 0X를 붙임

    ```go
    var num1 int = 32
    var num2 int = 0723 // 8진수로 저장
    var num3 int = 0x32fa2c75 //16진수로 저장
    ```

  - 소수점을 사용하거나 지수 표기법으로 표기할 수 있음
    - 지수 표기법 
      - ex) 3.14 = 314 * 10-2 = 314e-2 또는 314E-2
  - 변수에 저장될 때는 부동소수점 방식을 사용

  ```go
  // 소수점 사용
  var num1 float32 = 0.1
  
  // 지수 표기법
  var num2 float32 = 1e7
  var num3 float64 = .12345E+2
  ```

  - 부동소수점 반올림 문제: 수학적으로 실수는 무한히 많은데 컴퓨터는 2진수를 사용하므로 2진수로 실수를 정확하게 표현하기 어려움

    ```go
    package main
    
    import "fmt"
    
    func main() {
    var a float64 = 10.0
    
    for i := 0; i < 10; i++ {
    a = a - 0.1
    }
    
    fmt.Println(a)        // 9.000000000000004: 반올림 오차 발생
    fmt.Println(a == 9.0) // false: 실수는 ==로 비교하면 안됨
    }
    ```

    - `a == 9.0` 으로 비교하면 잘못된 결과가 나옴!!

    - ⇒ 다음과 같이 비교하기 (연산한 값과 비교할 값의 차이를 구한 뒤 머신 입실론보다 작거나 같은지 확인)

      ```go
        package main
          
          import "fmt"
          import "math"
          
          func main() {
          	var a float64 = 10.0
          
          	for i := 0; i < 10; i++ {
          		a = a - 0.1
          	}
          
          	fmt.Println(a) // 9.000000000000004
          
          	const epsilon = 1e-14                   // Go 언어 머신 입실론
          	fmt.Println(math.Abs(a-9.0) <= epsilon) // true: 연산한 값과 비교할 값의 차이를 구한 뒤
          	                                        // 머신 입실론보다 작거나 같은지 비교
          }
      ```

    - 머신 입실론?

      - 어떤 실수를 가장 가까운 부동소수점 수로 반올림했을 때 상대 오차는 항상 머신 입실론 이하이므로 입실론 보다 차이가 작으면 두 실수는 같은 값
      - 차이가 음수가 나올 수 있으므로 절댓값으로 만들어 비교

- 복소수

  ```go
    var num1 complex64 = 1 + 2i //실수부 1, 허수부 2
    var num2 complex64 = complex(1, 2) // num1 == num2
  ```

  - 실수부 + 허수부 형태
  - 허수부 마지막에 `i` 를 붙임
  - 소수점 및 지수 표기법으로 표현
  - 복소수는 +로 표현하는 대신 complex함수를 사용할 수도 있음

- 바이트

  ```go
    var num1 byte = 10 // 10진수 저장
    var num2 byte = 0x32 // 16진수 저장
    var num3 byte = 'a' // 문자 저장
    
    var num1 byte = "a"  // 컴파일 에러
    var num2 byte = 'ab' // 컴파일 에러
    var num3 byte = "ab" // 컴파일 에러
  ```

  - 16진수, 문자 값으로 저장
  - 바이너리 파일에서 데이터를 읽고 쓰거나, 데이터를 암호화, 복화하할 때 사용
  - `"` 사용할 수 없음, 문자열은 byte에 저장 못함

- 룬

  - 유니코드 문자를 저장할 때 사용

  - 사용법 

    - `' '` 로 묶어주기

    - `\u` 나 `\U` 를 사용 (`\U` 는 값을 16진수 8자리로 맞춰주어야 함)

      ```go
      var num1 uint16 = 10
      var num2 uint32 = 80000
      
      fmt.Println(num1 + uint16(num2)) // 14474: uint32에서 uint16으로 변환하면서 넘치는 값을 버림
      ```

    - byte와 같이 문자열, `"` 사용 불가

- 숫자 연산하기

  - 덧셈(+), 뺄셈(-), 곱셈(*), 나눗셈(/), 나머지(%), 시프트(<<, >>), 비트 반전(^) 연산자를 사용할 수 있음

  - 주의! 

    - 연산시 서로 자료형 다르면 컴파일 에러 발생 ⇒ 명시적으로 자료형을 변환해주기

    - 실수(float32)를 정수(int)로 변환하면 소수점 이하 값을 **버림**

    - 크기가 작은 자료형으로 변환하면 넘치는 값 버림

      ```go
      var num1 uint16 = 10
      var num2 uint32 = 80000
      
      fmt.Println(num1 + uint16(num2)) // 14474: uint32에서 uint16으로 변환하면서 넘치는 값을 버림
      ```

- 오버플로우 언더플로우 

  - 오버플로우: 각 자료형에서 저장할 수 있는 최대 크기를 넘어섬
  - 언더플로우: 최소 크기보다 작아짐

- 변수의 크기 구하기: `unsafe` 패키지의 `Sizeof` 함수를 사용

  - 크기는 **바이트** 단위

```go
package main
    
import "fmt"
import "unsafe" // Sizeof 함수를 사용하기위해 unsafe 패키지 사용

func main() {
    var num1 int8  = 1
    var num2 int16 = 1
    var num3 int32 = 1
    var num4 int64 = 1

    fmt.Println(unsafe.Sizeof(num1)) // 1
    fmt.Println(unsafe.Sizeof(num2)) // 2
    fmt.Println(unsafe.Sizeof(num3)) // 4
    fmt.Println(unsafe.Sizeof(num4)) // 8
}
```