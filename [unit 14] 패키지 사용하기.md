# unit 14. 패키지 사용하기

```go
// import 키워드를 사용하여 패키지 사용
import "package"

// 여러 패키지를 사용할 경우 ()로 묶어서 사용
import (
	"fmt"
    "runtime"
)

// 패키지 이름 앞에 . 을 사용하여 전역 패키지로 사용 가능
import . "fmt"
func main() {
    Println("Hello, world!")
}

// 패키지 이름 앞에 별칭을 사용 가능
import f "fmt"
func main() {
    f.Println("Hello, world!")
}

// 패키지를 import 한 후 사용하지 않으면 컴파일 에러가 나므로 _ 키워드로 컴파일 에러가 발생하지 않도록 할 수 있음
import _ "fmt"
```

