# unit 11. 상수 사용하기

```go
// const 키워드로 상수 생성
const age int = 1

// 반드시 선언과 동시에 초기화
const address = "Suji" // string
const a, b int = 2, 3 // a = 10, b = 20
const (
	c, d int = 4, 5 // c = 4, d = 5
    e, f = 6, "Hong" // e = 6, f = "Hong"
)

const score int // error (초기화 X)
age = 20 // error (값 변경 X)
name = "Grace" // error (값 변경 X)
```

