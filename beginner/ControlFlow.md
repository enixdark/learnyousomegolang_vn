## Câu lệnh điều khiển ( Control flow statements )

- Trong cuộc sống bạn sẽ gặp rất nhiều điều mà chúng ta phải đưa ra một quyết định hay một hành động nào đó, một sự việc lặp đi lặp lại giả sử  như mẹ bạn sai bạn đi chợ mua trứng và cà chua, trước khi đi mẹ bạn dặn rất kỹ rằng ra chợ mua 3 qủa trứng, tuy nhiên nếu như có cà chua thì mua 6 quả. Đó trong trường hợp này bạn sẽ gặp một quyết định là nếu có hay không thì sẽ mua ? 
- Tương tự trong lập trình có rất nhiều bài toán với logic phức tập, ví dự như kiểm tra xem tuổi của có đủ điều kiện xem phim 18+ hay không, giả sử  chúng ta sẽ có điều kiện nếu bạn trên 18 tuổi như vậy bạn đủ đièu kiện xem phim còn ngược lại thì không. Chắc chắn là với những bạn toán như vậy này chúng ta không thể  giải quyết chỉ bằng cách tính toán rồi in ra được. Lúc này chúng ta sẽ cần một cấu trúc điều khiển để  dựa trên những điều kiện mà đưa ra phương pháp giải dúng, chính xác. 

- Trong Go để  dơn giản hóa mọi thứ, sẽ không có quá nhiều cấu trúc khiến cho người lập trình khó khăn trong việc tiếp cận. Do đó nó chỉ cung câp 3 loại cáu trúc chính: if..else, For và Switch.


### Câu lệnh điều kiện if, if else

- if là lệnh đầu tiên trong cấu trúc điều khiển mà chúng ta sẽ giới thiệu trước tiên và nó cũng là cáu trúc được sử  dụng nhiều nhất trong hầu hết các logic. Nó thường được sử  dụng để xác định một điều kiện có đúng hay không, nếu đúng thì sẽ thực hiện logic bên trong phần `{}` của nó. cú pháp biểu diễn của nó sẽ như sau   

```go
if ( condition ){
    // logic
}
```

- Bây giờ chúng ta sẽ áp dụng nó vào trong ví dụ ở đâu chương mà chúng ta đã nhắc tới:

```go
package main

import "fmt"

func main() {  
	eggs := 6 

	if eggs > 3 {
		fmt.Println("Cửa hàng có đủ trứng để  cung cấp cho bạn")
	}
}
```

- Nào hãy chạy chương trình và bạn sẽ thấy cậu lệnh bên trong biểu thức if sẽ được thưc hiện

```go
Cửa hàng có đủ trứng để  cung cấp cho bạn
```

- Thật dễ  dàng phải không, tuy nhiên ví dụ trên chưathực sự dúng với những gì chúng ta đã nêu ra hay chính xác là dúng với những gì mà mẹ bạn đã dặn rò. Vì vậy chúng ta cần điều chỉnh dôi chút. Hãy nhớ lại trong chương học về  các kiểu dữ liệu, chính xác là phần đại số  boolean, chúng ta đã học răng các đại số  có thể  kết hợp được với nhau bằng các toán tử  `&&`, `||` và `!`. Nào giờ là lúc chúng ta áp dụng kiến thức mà chúng ta đã học trước đó vào:

```go
package main

import "fmt"

func main() {  
    eggs := 6 
    tomatos := 1

	if eggs > 3 && tomatos > 0 {
		fmt.Println("Bán cho tôi 6 quả, ờ mà ̉6 quả gì nhi ?")
	}
}
```  

- Tuy nhiên trong trường hợp nếu như không thỏa mãn điều kiện thì sao. Thông thường sẽ có hai cách, cách đầu tiên đó là 
ngay bên dưới điều kiện `if` bạn sẽ thực hiện logic đối với trường hợp không thỏa mãn. Nhưng trái lại nếu nó thỏa mãn thì ngay lập tức sau khi thực hiện trong logic thỏa mãn chúng ta phải dừng hay nói cách khác là thoát khỏi chương trình . Để  thoát khỏi chúng ta có thể  sử  dụng từ khóa `return` hoặc hàm `Exit` trong pạckage `os` vd:

```go
package main

import (
    "fmt"
    "os"
)

func main() {  
    eggs := 6 
    tomatos := 1

	if eggs > 3 && tomatos > 0 {
        fmt.Println("Bán cho tôi 6 quả, ờ mà ̉6 quả gì nhi ?")
        os.Exit(0)
    }
    fmt.Println("Thôi thì đằng nào không có gì để  mua thì mua tạm que kem ăn vậy.")
}
```

- Tuy giả sử  chúng ta sau khi mua hàng xong còn muốn cám ơn người bạn hàng, và không thực sự muốn thoát khỏi chương trình. Lúc này chúng ta sẽ sử  dụng mệnh để  còn lại trong `if` đó là `else`:

```go
package main

import (
    "fmt"
    "os"
)

func main() {  
    eggs := 6 
    tomatos := 1

	if eggs > 3 && tomatos > 0 {
        fmt.Println("Bán cho tôi 6 quả, ờ mà ̉6 quả gì nhi ?")
    } else {
        fmt.Println("Thôi thì đằng nào không có gì để  mua thì mua tạm que kem ăn vậy.")
    }

    fmt.Println("Cám ơn.")
}
```

- "Nếu có thì còn ngược lại thì" , vâng tới đây chúng ta đã có được trọn vẹn logic của một điều kiện. Mà khoan có gì đó thật sự không đúng lắm ? Trong cuộc sống không phải lúc nào ch 