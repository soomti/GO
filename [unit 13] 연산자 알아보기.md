# unit 13. 연산자 알아보기

```go
var a int = 1 // = 대입
a := 1 // 변수 선언 및 대입

a := 3
b := 2
c := a & b  // a와 b의 AND 연산
c := a | b  // a와 b의 OR 연산
c := a ^ b  // a와 b의 XOR 연산
c := a << 1 // a의 shift 1

c := &a     // a의 메모리 주소를 c에 대입

a := new(int)
*a = 1      // 역참조 연산

c := make(chan int)
go func() {
    c <- 1 // 채널 c에 1을 보냄
}()

a := <- c  // 채널 c에서 값을 가져와서 1을 대입
```

