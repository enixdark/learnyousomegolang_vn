## Câu lệnh điều khiển ( Control flow statements )

- Trong cuộc sống bạn sẽ gặp rất nhiều điều mà chúng ta phải đưa ra một quyết định hay một hành động nào đó, một sự việc lặp đi lặp lại giả sử  như mẹ bạn sai bạn đi chợ mua trứng và cà chua, trước khi đi mẹ bạn dặn rất kỹ rằng ra chợ mua 3 qủa trứng, tuy nhiên nếu như có cà chua thì mua 6 quả. Đó trong trường hợp này bạn sẽ gặp một quyết định là nếu có hay không thì sẽ mua ? 
- Tương tự trong lập trình có rất nhiều bài toán với logic phức tập, ví dự như kiểm tra xem tuổi của có đủ điều kiện xem phim 18+ hay không, giả sử  chúng ta sẽ có điều kiện nếu bạn trên 18 tuổi như vậy bạn đủ đièu kiện xem phim còn ngược lại thì không. Chắc chắn là với những bạn toán như vậy này chúng ta không thể  giải quyết chỉ bằng cách tính toán rồi in ra được. Lúc này chúng ta sẽ cần một cấu trúc điều khiển để  dựa trên những điều kiện mà đưa ra phương pháp giải dúng, chính xác. 

- Trong Go để  dơn giản hóa mọi thứ, sẽ không có quá nhiều cấu trúc khiến cho người lập trình khó khăn trong việc tiếp cận. Do đó nó chỉ cung câp 3 loại cáu trúc chính: if..else, For và Switch.


### Câu lệnh điều kiện if, if else

- if là lệnh đầu tiên trong cấu trúc điều khiển mà chúng ta sẽ giới thiệu trước tiên và nó cũng là cáu trúc được sử  dụng nhiều nhất trong hầu hết các logic. Nó thường được sử  dụng để xác định một điều kiện có đúng hay không, nếu đúng thì sẽ thực hiện logic bên trong phần `{}` của nó. cú pháp biểu diễn của nó sẽ như sau   

```go
if condition {
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

- Nào hãy chạy chương trình và bạn sẽ thấy câu lệnh bên trong biểu thức if sẽ được thưc hiện

```go
Cửa hàng có đủ trứng để  cung cấp cho bạn
```

- Thật dễ  dàng phải không, tuy nhiên ví dụ trên chưa thực sự dúng với những gì chúng ta đã nêu ra hay chính xác là dúng với những gì mà mẹ bạn đã dặn rò. Vì vậy chúng ta cần điều chỉnh dôi chút. Hãy nhớ lại trong chương học về  các kiểu dữ liệu, chính xác là phần đại số  boolean, chúng ta đã học răng các đại số  có thể  kết hợp được với nhau bằng các toán tử  `&&`, `||` và `!`. Nào giờ là lúc chúng ta áp dụng kiến thức mà chúng ta đã học trước đó vào:

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

- "Nếu có thì còn ngược lại thì" , vâng tới đây chúng ta đã có được trọn vẹn logic của một điều kiện. Mà khoan có gì đó thật sự không đúng lắm ? Trong cuộc sống không phải lúc chỉ có hai trường hợp cả, đôi khi bạn sẽ có nhiều lựa chọn để  đưa ra một quyết định. Quay trở lại vừa trường hợp bạn đi chợ mua trứng, để  tăng tính phức tạp lên một chút chúng ta sẽ giả sử  răng mẹ bạn đưa cho bạn 1̀̀5000 để  mua đồ, và bạn đi tới quầy hàng mà mọi khi bạn vẫn hay mua tuy nhiên giá mà quầy A đưa ra cho số  trứng và cà chua vượt quá giá trị tiền mà bạn mang theo, do đó bạn quyết định đi sang các quầy khác để  xem trước khi quyết định mua ( giả sử  rằng tât cả các quầy đều có dư giả số  trứng và cà chua ), các bước sẽ tiến hành như sau:
  + quầy A, tại đây giá của một quả trứng là 4500, một quả cà chua là 1000
  + quầy B, tại đây giá của một quả trứng là 2500, một quả cà chua là 2500
  + quầy C, tại đây giá của một quả trứng là 3000, một quả cà chua là 1000

-

```go
package main

import (
    "fmt"
)

func main() {  
    
    money := 15000

    egg_price_A := 4500 
    tomato_price_A := 1000
    
    egg_price_B := 2500 
    tomato_price_B := 2500

    egg_price_C := 3000 
    tomato_price_C := 1000

    if money - ( egg_price_A * 3 + tomato_price_A * 6 ) >= 0 {
        fmt.Println("Mua hàng ở quầy A")
    } 
    if money - ( egg_price_B * 3 + tomato_price_V * 6 ) >= 0 {
        fmt.Println("Mua hàng ở quầy B")
    }
    if money - ( egg_price_C * 3 + tomato_price_C * 6 ) >= 0 {
        fmt.Println("Mua hàng ở quầy C")
    }
}
```

- Sau một hồi suy luận rốt cục bạn cũng quyết định mua được hàng ở cửa hàng C. Buổi đi chợ của bạn đáng lẽ ra kết thúc tốt đẹp nhưng mà bất chợt cô bán hàng ở quầy A chợt thay đổi đổi lại giá trứng từ 4500 tới 3000 do nhẫn lẫn giữa giá trứng gà ta và trứng gà tây. Oh lúc này giá của cửa hàng A đã bằng với giá của cửa hàng C. Và nếu sử  dụng đoạn chương trình trên thì điều kiện ở cửa hàng A và C đều thỏa mãn, và ngay lập tức màn hình sẽ in ra:
d
```go
Mua hàng ở quầy A
Mua hàng ở quầy C
```

- Thật tệ dường như chúng ta đã không có quyết định cụ thể . Mông lung như một trò đùa, nghĩ hoài không ra. Đây là sẽ lúc bạn đưa ra quyết định, về  cơ bản ngay lập tức trường hợp ở quầy A đã thỏa mãn rồi , và không càn phải quan tâm tới các quầy khác nữa. Để  đưa quyết định  như vậy chúng ta sẽ sử  dụng điều kiện
`if..else if...`: 

```go
package main

import (
    "fmt"
)

func main() {  
    
    money := 15000

    egg_price_A := 3000 
    tomato_price_A := 1000
    
    egg_price_B := 2500 
    tomato_price_B := 2500

    egg_price_C := 3000 
    tomato_price_C := 1000

    if money - ( egg_price_A * 3 + tomato_price_A * 6 ) >= 0 {
        fmt.Println("Mua hàng ở quầy A")
    } else if money - ( egg_price_B * 3 + tomato_price_B * 6 ) >= 0 {
        fmt.Println("Mua hàng ở quầy B")
    } else if money - ( egg_price_C * 3 + tomato_price_C * 6 ) >= 0 {
        fmt.Println("Mua hàng ở quầy C")
    } else {
        fmt.Println("Không đủ tiền thì về  nhà thôi")
    }
}
```

- Kết quả mà ta nhận được:

```go
Mua hàng ở quầy A
```

- Tuyệt vời, kiểu gì về  nhà mẹ cũng sẽ khen. Buổi di chợ tới đây là két thúc tốt đẹp.

- Lưu ý là với điều kiện trong if bạn có thể  để  chúng trong cặp ngoặc `()` vd:

```go
if ( condition ) {
    // logic
}
```
- Hay bạn có thể  kết hợp một số  câu lệnh như việc khởi tạo giá trị đi kèm với điều cũng đc vd:

```go
if i := 1; i > 0 {
    // logic
} 
```

- Tuy nhiên với cách làm trên thì cặp ngoặc `()` không dược phép sử dụng vd:

```go
// một thông báo lỗi sẽ được đưa ra
if (i := 1; i > 0) {
    // logic
} 
```


### For loop

- Bên cạnh việc sử  dụng cấu trúc điệu kiện `If..else`, thì một cấu trúc không kém phần quan trọng 
đó là cấu trúc lặp. Trong Go để  đơn giản hóa chúng ta sẽ chỉ có duy nhất một cấu trúc lặp đó là là `for`
và viết theo cú pháp sau:

```go
for init_value; condition; post {  
    // logic
}
```

vd

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

- Đầu tiên chúng ta sẽ khởi tạo giá trị để  bắt đầu vòng lặp, tại môi vòng lặp sẽ kiẻm tra điều kiện của giá trị và sau mỗi vòng lặp chúng ta sẽ tịnh tiến lên một đơn vị cho tới khi điều kiện không còn được thỏa mãn nữa thì chúng ta sẽ kết thúc vòng lặp.

- Lưu ý: về  cơ bản cú pháp theo chuẩn của vòng lặp for yêu cầu phải có 3 phân đoạn tuy nhiên tùy trường hợp chúng ta có thể  bỏ qua bất kỳ một phân đoạn nào hoặc không cần cũng được, khi đó chúng ta sẽ được một vòng lặp vô hạn vd:

```go
for {
    fmt.Println("Infinity gauntlet")
}
```

- Trong trường hợp giả sử  bạn sử  dụng vòng lặp vộ hạn nhưng lại muốn thoát khỏi vòng lặp vô hạn. Oh đơn giản bạn sẽ nói ngay đó là sử  dụng từ khóa `return` hay hàm `Exit` bên trong vòng lặp, nhưng không cần phải làm vậy. Nếu bạn thường xuyên sử  dụng vòng lặp for ở một số  ngôn ngữ bạn sẽ thấy việc sử  dụng từ khóa `break`. Đúng vậy, bằng cách chèn `break` vào trong vòng lặp nó sẽ ngay lập tức thoát ra khỏi vòng lặp đó. chúng ta xét ví dụ đơn giản bằng cách sử  dụng vòng lặp vô hạn để  in ra 10 dòng chứ "Infinity gauntlet", sau khi đủ 10 dòng chữ nó sẽ lập tức thoát khỏi vòng lặp:

```go
package main

import (
    "fmt"
)

func main() {  
    
    i := 0

    for {
        i++
        if i > 10 {
            break
        }
        fmt.Println("Infinity gauntlet")
    }
}
```

- Chạy đoạn mã trên bạn sẽ thấy sau 10 lần in ra dòng chữ "Infinity gauntlet", nó sẽ tự dộng thoát khỏi vòng lặp và kết thúc chương trình.

- Tiếp theo còn một từ khóa nữa trong vòng lặp mà chúng ta chưa để  cập đén dó là từ khóa `continue`.
Gia sử  như bạn không muốn bất kỳ một phép lặp nào trong khoảng từ 1 đến 100 in ra dòng chữ "Infinity gauntlet", tât nhiên đẻ thực hiện điều này bạn có thể  dùng điều kiện `if..else` , kiểm tra lần lặp trước đó từ 1 tới 100 xem, nếu thỏa mãn thì không thực hiện phép in ra màn hình. Tuy nhiên có một cách đơn giản hơn việc sử  dụng `if..else` đó là sử  dụng `continue`, bằng cách này nó sẽ bỏ qua vòng lặp và chuyển sang vòng lặp tiép theo. 

Vd:

```go
package main

import (
    "fmt"
)

func main() {  
    
    i := 0

    for {
        i++
        if i <= 100 {
           continue
        }
        fmt.Println("level loop:", i)
        fmt.Println("Infinity gauntlet")
        break
    }
}
```  

- Ngoài các trường hợp trên ra còn một trường hợp sử  dụng vòng for nữa với các cấu trúc như mảng, tuy nhiên chúng ta sẽ nói tới ở các chương sau này. 

### Switch

- Trong phần cuối cùng này chúng ta sẽ đi vào tìm hiểu về  switch. Switch về  mặt ý nghĩa khá giống với cấu trúc điều kiện `if`. Tuy nhiên nó sẽ có một số  trường hợp sử  dụng switch xem đem lại hiệu quả và ngắn gọn hơn cho đoạn mã của bạn. cú pháp của switch được viết như sau:

```go
switch (var) {
    case condition:
        // logic
        [break]
    ....
    default:
        // logic
}
```

- Như chúng ta thấy với switch chúng ta sẽ phải xác định một giá trị truyền vào trước, khi giá trị này được truyền vào nó sẽ so sánh các điều kiện trong case, nếu như khớp nó sẽ thực hiện logic ở trong đièu kiện case đó, ngược lại nó sẽ thực hiện ở trường hơn mặc định ( default ). 
- Lưu ý là trường hợp default có thể  có hoặc không cũng được:

- Trước tiên để  đơn gian chúng ta sẽ xét một vị dụ tìm ra số  may mắn trong khoảng từ 1 tới 10:

```go
package main

import (
    "fmt"
	"math/rand"
	"time"
)

func main() {  
	rand.Seed(time.Now().UnixNano())
	
    lucky_number := rand.Intn(10)
    
    switch (lucky_number) {
        case 1:
            fmt.Println("1 là số  may mắn")
        case 2:
            fmt.Println("2 là số  may mắn")
        case 3:
            fmt.Println("3 là số  may mắn")
        case 4:
            fmt.Println("4 là số  may mắn")
        case 5:
            fmt.Println("5 là số  may mắn")
        case 6: 
            fmt.Println("6 là số  may mắn")
        case 7:
            fmt.Println("7 là số  may mắn")
        case 8:
            fmt.Println("8 là số  may mắn")
        case 9:
            fmt.Println("9 là số  may mắn")
        case 10:
            fmt.Println("10 là số  may mắn")
    }
}
```

- Ở đây chúng ta phải import package `math/rand` vào để  sử  dụng hàm sinh ra một số nguyên ngãu nhiên. Và bởi vì chúng ta cần mỗi kết quả chạy hàm là một giá trị ngẫu nhiên do đó ta phải đựt
lời gọi hàm hàm `rand.Seed(time.Now().UnixNano())` vào trước. Sau khi chạy bạn sẽ thấy dựa trên giá trị số  nguyên ngầu nhiên được sinh ra trong khoảng từ 1 tới 10 sẽ khớp với một trường hợp trong case.

- Nhưng trong trường hợp nếu như chúng ta xác định trước một số  may mắn thì sao, khi đó các số  khác đều không phải só may mắn và việc viết mỗi lần trùng lặp in ra thông báo như vậy thật
lãng phí. Để  rút gọn và nhìn đoạn mã sạch sẽ hơn, trong switch cho phép chúng ta được phép sử  dụng đa điều kiện trong cùng một case:

```go
package main

import (
    "fmt"
)

func main() {  
	rand.Seed(time.Now().UnixNano())
	
    lucky_number := 7
    
    switch (lucky_number) {
        case 1, 2, 3, 4, 5, 6, 8, 9 ,10:
            fmt.Println("đây không phải là số  may mắn")           
        case 7:
            fmt.Println("7 là số  may mắn")
    }
}
```

- Đó bạn thấy không thấy gọn gàn. Ngoài ra bạn cũng có thể  sử  dụng điều kiện kiểm tra trong case.


- Tiếp theo chúng ta sẽ sử  dụng case đẻ viết lại ví dụ ở đi chợ mua hàng ở trên:

```go
package main

import (
    "fmt"
)

func main() {  
    
    money := 15000

    egg_price_A := 3000 
    tomato_price_A := 1000
    
    egg_price_B := 2500 
    tomato_price_B := 2500

    egg_price_C := 3000 
    tomato_price_C := 1000

    switch (true) {
        case money - ( egg_price_A * 3 + tomato_price_A * 6 ) >= 0:
            fmt.Println("Mua hàng ở quầy A")
        case money - ( egg_price_B * 3 + tomato_price_B * 6 ) >= 0:
            fmt.Println("Mua hàng ở quầy B")
		case money - ( egg_price_C * 3 + tomato_price_C * 6 ) >= 0:
            fmt.Println("Mua hàng ở quầy C")
        default: 
            fmt.Println("Không đủ tiền thì về  nhà thôi")
    }
}
```

- Về  mặt logic đoạn mã trên không có thay dổi so với đoạn mã sử  dụng `if..else`, và trông có vẻ phức tạp hơn nhưng trong một só bài toán khác nó có thể  gọn hơn. 
Tóm lại là tùy theo cách dùng mà bạn có thể  quyết định sử  dụng if hay switch.

- Lưu ý nếu bạn so sánh switch trong một số  ngôn ngữ khác như Java, C bạn thấy hầu hết chúng có thể  một dòng `break` phía dưới mỗi `case` để thông báo chám dứt switch, với go bạn vẫn  ó thể  sử  dụng tuy nhiên không cần thiết phải chèn vào vì go sẽ ngâm định là sau khi thực hiện logic trong `case` nó sẽ tự động thoát ra. Tuy vậy trong các ngôn ngữ khác nếu không dùng đặt `break` chúng sẽ giống như trường hợp ta sử  dụng liên tục các cấu trúc `if`, hay nói một cách cụ thể  là dù có thoải mãn điều kiện hay không thì nó vẫn sẽ kiểm tra các `case` khác. Thực tế  trong Go cũng hỗ  trợ đièu này, có đièu bạn sẽ phải tường minh bằng cách sử  dụng từ khóa `fallthrough` và dặt chúng vào các trường hơp xác dịnh vd:

```go
package main

import (
    "fmt"
)

func main() {  
    
    money := 15000

    egg_price_A := 3000 
    tomato_price_A := 1000
    
    egg_price_B := 2500 
    tomato_price_B := 2500

    egg_price_C := 3000 
    tomato_price_C := 1000

    switch (true) {
        case money - ( egg_price_A * 3 + tomato_price_A * 6 ) >= 0:
            fmt.Println("Mua hàng ở quầy A")
            fallthrough
        case money - ( egg_price_B * 3 + tomato_price_B * 6 ) >= 0:
            fmt.Println("Mua hàng ở quầy B")
		case money - ( egg_price_C * 3 + tomato_price_C * 6 ) >= 0:
            fmt.Println("Mua hàng ở quầy C")
    }
}
```
  
- Chạy cậu lệnh này ta sẽ được kết quả:

```go
Mua hàng ở quầy A
Mua hàng ở quầy B
Mua hàng ở quầy C
```

- Oh tại sao tôi nói dùng cẩn thận đó là bởi khi ngay khi sử  dụng `fallthrough` , nó sẽ nhảy tới `case` tiếp theo và thực hiện ngay đoạn logic mà không cần quan tâm tới 
điều kiện đúng hay sai. 


### Tổng kêt

- Vậy là chúng ta đã nắm được các cậu lệnh điều khiển trong Go rồi cũng nhưng thao tác với một số  trường hợp đặc biệt, Trong chương tiếp theo chúng ta 
sẽ tìm hiểu vê hàm và cách định nghĩa một hàm cũng như gọi hàm đó như thế  nào.  