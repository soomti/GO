## **[Unit 7] 변수 사용하기**

- 변수 선언 방법 2가지 

  - `var` 키워드 사용하는 방식

    ```go
      var i int
      var s string
      
      var age int = 10
      var name string = "Maria"
    ```

    - 자료형이 변수명 뒤에 옴
    - 변수명은 문자, 숫자로 이루어짐 
      - 단, 숫자로 시작할 수 없음

  - 자료형을 생략하는 방식

    - 자료형은 자료형을 대입하는 값의 자료형으로 결정됨

    - 반드시 초깃값을 대입해야 함

      var age = 10 var name = "Maria" var address // 컴파일 에러