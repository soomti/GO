# unit 12. 열거형 사용하기

```go
const (
	Sunday       = 0
	Monday       = 1
	Tuesday      = 2
	Wednesday    = 3
	Thursday     = 4
	Friday       = 5
	Saturday     = 6
	numberOfDays = 7
)
fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)
fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)
```