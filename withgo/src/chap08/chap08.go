package main

import "fmt"
import "math"

func main() {

	// 8.1
	var num1 int = 32
	var num2 int = -16
	var num3 int = 0723 // 8
	var num4 int = 0x32fa2c75

	// 8.2

	var a float64 = 10.0 // float 는 default 가 없는듯! 
	// 실수는 == 하면 안됨. 이진수는 실수 계산이 어렵기 떄문. 
	// 이럴 때는 앱실론을 사용한다

	for i := 0; i < 10 ; i ++ {
		a = a - 0.1
	}

	fmt.Println(a) // 9.00000004

	fmt.Println(a == 9.0) // false

	// go 언어 머신 엡실론 계산기 엡실론(Machine epsilon)은 부동소수점 연산에서 반올림을 함으로써 발생하는 오차의 상한이다. 
	// 이 값은 컴퓨터 과학에서 컴퓨터 연산을 통한 수치 해석을 특징짓는다.
	const epsilon = 1e-14 

	fmt.Println(math.Abs(a-9.0) <= epsilon) // true 면 같은 값 . 음수일수도 있으므로 절대값을 붙힌다. 


	// 8.3
	var num5 complex64 = 1 + 2i // 실수부 1 허수부 2i

	var r1 float32 = real(num5) // return 1
	var r2 float32 = imag(num5) // return 2

	fmt.Println(r1,r2)

	// 8.4 byte == 이게 char?
	// 바이너리 파일에서 데이터를 읽고 쓰거나 암호화 복호화할 때 사용
	var num6 byte = 10
	var num7 byte = 0x32 
	var num8 byte = 'a'

	fmt.Println(num6,num7,num8)

	// //error
	// var num9 byte = "a" // compile error
	// var num9 byte = 'ab' // compile error
	// var num9 byte = "ab" // compile error

	// 8.5 rune 유니코드 문자 저장시 사용

	var r4 rune = '한'
	var r5 rune = '\ud55c' // 한
	var r6 rune = '\U0000d55c' //한

	fmt.Println(r4,r5,r6)

	// 8.6 숫자 연산
	// +, -, *, /, %, >> , <<

	var c uint8 = 3
	var d uint8 = 2

	fmt.Println(c+d,c-d,c*d,c/d,c%d,c<<d,c>>d,^c)

	// 명시적 자료형 변환
	fmt.Println(float32(num1))

	// 변수 크기 구하기
	// unsafe.sizeOf()


}
