# 17. for 반복문 사용하기

go 언어는  반복문이 `for` 만 존재한다.

- **for 초깃값; 조건식; 변화식 {}**

  ```go
  for i:=0; i < 5; i++ {
      fmt.Println(i);
  }
  ```

  실행결과

  ````
  0
  1
  2
  3
  4
  ````

##### 특징

반복문 조건식에 괄호**()** 를 사용하지 않는다. 

중괄호 **{}** 를 다음줄에 작성하거나, 생략시 컴파일 에러가 발생한다.

- compile error 발생 

  ```go
  // case 1 
  // 중괄호 표기 실수
  
  for i:=0; i < 5; i++ 
  { 
      fmt.Println(i);
  }
  
  // case 2 
  // 중괄호 생략
  for i:=0; i < 5; i++ 
      fmt.Println(i); // error!
  ```



for 키워드에 조건식만 사용하면 while 처럼 동작한다.

이때, 변화식은 중괄호 안해서 처리해야 한다. 

- **for 조건식 {}**

  조건식을 설정하지 않거나, 변화식이 없는 경우 무한루프 발생을 주의해야 한다. 

  ```go
  i := 0
  for i < 5 {
      fmt.Println(i)
      i = i + 1 // i++
  }
  ```

  무한루프 발생

  ```go
  for {
      fmt.Println("Hello, world!")
  }
  ```


## 17.1 break 사용하기

for 키워드에 아무것도 설정하지 않으면 무한루프가 된다.

반복문을 중단하고싶을 땐 중괄호 블록 안에서 조건을 정하고 `break` 키워드를 사용하면 된다. 

- break 를 통한 반복문 중단

  ```go
  i := 0 
  for {
      if i > 4 {
         break 
      }
      fmt.Println(i)
      i = i+1
  }
  ```

  ##### 실행결과

  ```
  0
  1
  2
  3
  4
  ```

- 레이블 지정

  break 키워드에 레이블을 지정할 수 있다. 

  ##### 레이블

  레이블을 만나서 빠져나오면, 레이블 밑의 for 반복문을 더 이상 실행하지 않고 완전히 중단한다. 

  ```go
  Loop:
      for i := 0; i < 3 ; i++ {
          for j :=0; j < 3; j++ {
              if j == 2 {
                  break Loop 
              }
              fmt.Println(i,j)
          }
      }
      fmt.Println("Hello, world!")
  ```

  ##### 실행결과

  ```
  0 0
  0 1 
  Hello, world!
  ```

  레이블과 for 키워드 사이에 다른 코드가 들어가면 안된다. 컴파일 에러 발생!

## 17.2 continue 사용하기

특정 부분 이하는 실행하지 않고 넘어갈 경우 `continue` 키워드 사용한다.

- countinue 사용 

  특정 부분 이하는 실행되지 않고 레이블 바로 아래부터 실행한다.

  ```go
  for i := 0; i < 5; i++ {
      if i ==2 {
          continue
      }
      fmt.Println(i)
  }
  ```

  continue 키워드에도 레이블을 지정할 수 있다.

  ##### 실행결과

  ```
  0
  1
  2
  3
  4
  ```

- 레이블 지정 

  ```go
  Loop: 
  	for i := 0; i < 3; i++ {
          for j:=0; j<3; j++ {
              if j==2 {
                  continue Loop
              }
              fmt.Println(i,j)
          }
  	}
  	 fmt.Println("Hello World!")
  ```

  ##### 실행결과 

  ```
  0 0
  0 1
  1 0
  1 1 
  2 0 
  Hello, world!
  ```

  레이블과 반복문 사이에는 다른 코드가 있으면 안된다. 컴파일 에러!



## 17.3 반복문에서 변수 여러개 사용하기

반복문의 변화식에서 여러 변수를 처리하면 병렬 할당을 사용해야 한다.

- 병렬할당

  ```go
  for i,j := 0,0; i < 10; i,j = i+1,j+2 {
      fmt.Println(i, j)
  }
  ```

  병렬 할당이 아닌, 그대로 변화식을 나열하는 경우 컴파일 에러가 발생

- compile error

  ```go
  for i,j := 0,0; i < 10; i++,j+=2 { // compile error!
      fmt.Println(i, j)
  }
  ```


