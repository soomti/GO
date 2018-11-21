package main
import _ "fmt" // 미 사용시 앞에 이렇게 

func main()  {

	var varname = 10
	varname2 := 10
	_ := 10 // 이건 안되는듯
	_ = varname // 이건 됨 
}