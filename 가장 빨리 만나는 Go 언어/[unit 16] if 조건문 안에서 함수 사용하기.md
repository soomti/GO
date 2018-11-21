# 16. if 조건문 안에서 함수 사용하기

조건문 안에서 함수를 사용하는 방법에 대해 알아보자 

- hello.txt

  ```
  Hello, World! 
  ```

- hello.txt 파일의 내용을 읽어 화면에 출력하는 코드

  ```go
  var b []byte
  var err error
  
  b, err = ioutil.ReadFile("./hello.txt")
  // 에러가 없으면 
  if err == nil {
      fmt.Printf("%s",b)
  }
  ```

  ##### 

이 내용을, 조건식 안에서 함수 실행 뒤 조건 판단하는 방식으로 사용하면 코드를 줄일 수 있다.



- 조건식 안에서 함수 실행 

  ```go
  // 조건식 안에서 ; 로 구분 후 함수 실행 및 조건식 생성 
  if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
      fmt.Printf("%s",b)
  }
  ```

  이렇게 조건식 안에서 생성한 변수는 , `if` 뿐 만 아니라, `else`, `else-if` 에서도 사용할 수 있지만 바깥에서는 사용할 수 없다. 

  ```go
  if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
      fmt.Printf("%s",b) // 변수 b 사용 가능
  } else {
      fmt.Println(err) // 변수 err 사용 가능 
  }
  
  fmt.Println(b) // error
  fmt.Println(err) // error
  ```



------

### package "io/ioutil"

https://golang.org/pkg/io/ioutil/

```go
func ReadFile(filename string) ([]byte, error)
```

파일을 string 으로 읽고, 그에 해당하는 내용값과, 에러값을 반환하는 것 같다. 그래서 

```go
b, err = ioutil.ReadFile("./hello.txt")
```

이렇게 받나봄.

`err == nil` 의 값이 `true`면 성공적으로 값을 받아온것. 

> 근데 여기 eof 를 반환하지 않는다의 뜻을 잘 모르겠움..

