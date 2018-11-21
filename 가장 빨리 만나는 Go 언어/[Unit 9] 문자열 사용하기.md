- 문자열 사용하기

  - `" "` 사용

  - UTF-8 로 표현할 수 있는 문자

  - 여러 줄로 된 문자열 저장할 때는 백쿼트로 묶어줌

    ```go
      var s7 string = `안녕하세요
      Hello, world!`
    ```

- 문자열 길이 구하기: `len`

  - 한글 문자열은 UTF-8로 저장되기 때문에 길이가 다름 

    - 한글 ⇒ UTF-8로 저장 ⇒0xed, 0x95, 0x9c, 0xea, 0xb8, 0x80

    - 실제 문자열의 길이를 구하려면? ⇒ `RuneCountInString` 함수 사용

      ```go
      package main
          
      import "fmt"
      import "unicode/utf8"
      
      func main() {
          var s1 string = "한글"
          var s2 string = "Hello"
      
          fmt.Println(len(s1)) // 6: UTF-8 인코딩의 바이트 길이이므로 6
          fmt.Println(len(s2)) // 5: 알파벳 5글자이므로 5
          fmt.Println(utf8.RuneCountInString(s1)) // 2: 문자열의 실제 길이를 구함
      }
      ```

- 문자열 연산하기

  - 비교: `==`
  - 문자열 합치기: `+`
  - 배열과 동일하게 `[]` 로 각 문자를 가져올 수 있음(배열은 0으로 시작함)

- 문자열 수정하기

  - 변수에 문자열을 저장한 뒤 내용을 수정할 수 없음 but 다른 문자열로 대입은 가능함

    ```go
    var s1 string = "Chogolang"
    s1 = "another word"
    fmt.Println(sq[0]) // 97: ASCII코드 a
    s1[0] = 'z' // 컴파일 에러
    ```


