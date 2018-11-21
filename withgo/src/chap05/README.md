# chap 05

## example

```go
package main // 맨 위에 패키지를 설정해준다.

import "fmt" // 입출력 패키지. 

func main() {
	fmt.Println("chap 05. Hello World!")
}
```

## 실행파일 생성 명령어 

```go
go build filepath
````
> 이때, go build 를 실행한 위치에 실행 파일이 생성된다. 책에는 go 파일이 생성되는 디렉터리라고 나와있는데 아닌듯. 

```
`go install` 명령 사용시에는, 관련된 패키지 소스 파일을 모두 받아와 컴파일을 실행하고, `bin` 파일에 실행 파일이 생기게 된다. 
```
> go install 실행시 
>  ```
>  go install: no install location for .go files listed on command > line (GOBIN not set)
> ``` 
> go install만 처야된다. 패키지별로 생기는게 아닌듯. 잘 모르겠음 
> go install로 쳣을때 chap05 의 파일 명으로 네이티브 바이너리가 생김. 그래서 패키지 별로 소스파일이 생기나? 해서 test 폴더를 만들어서 실행해보니 안됨 궁금 이거 왜그렇고 어케해야되]
