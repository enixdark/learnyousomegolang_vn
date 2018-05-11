## Biến ( Variables )

### Biến là gi ?

- Tương tự với các ngôn ngữ lập trình khác, chúng ta cũng sử  dụng biến trong Go để  lưu trữ các giá trị của một loại dữ liệu xác định để  sử  dụng cho chương trình.

- Trong Go, bạn sẽ có một số  cách để  khai báo một biến như sau:

#### Khai báo một biến cùng với kiểu dữ liệu

- Trong Go chúng ta sẽ sử  dụng tù khóa `var` để  khai báo một biến. Nó sẽ có dạng sau

```go
var name type
```

- Trong đó:

  + `var`: là từ khóa để  khai báo một biến.
  + `name`: là tên mà bạn sẽ đặt cho biến đó.
  + `type`: là kiểu dữ liệu mà bạn xác định khi khai báo biến đó.

vd:

```go
package main

import "fmt"

func main() {  
    var money int // khai báo một biến có tên là name thuộc kiểu dữ liệu số  nguyên 
    fmt.Printf("trong ví tôi có: %d VND.", money)
}
```

- Trong đoạn mã trên chúng ta đã khai báo một biến thuộc kiểu dữ liệu số  nguyên thông qua biến `money` bằng lệnh `var money int`. Bởi vì chúng ta không khởi tạo bất kỳ giá trị nào cho biến này do đó Go sẽ luôn thiết lập nó với giá trị là 0 đối với kiểu số nguyên. Dòng tiếp theo chúng ta sử  dụng hàm Printf và định dạng ` %d`  ( định dang này cho biến chúng ta cần truyền vào, in một giá trị thuộc kiẻu số  nguyên ) để  hiển thị giá trị của biến `money` mà chúng ta đã khởi tạo ra màn hình. Kết quả của đoạn mã trên khi bạn chạy chương trình là:

```go
trong ví tôi có: 0 VND.
```

- Trong trương hợp bạn muốn gán một giá trị cho biến với khai báo bạn có thể  làm như sau:

```go
package main

import "fmt"

func main() {  
    var money int 
    fmt.Printf("trong ví tôi có: %d VND.\n", money)
    money = 10000 // gán giá trị cho biến vừa khởi tạo
    fmt.Printf("trong ví tôi có: %d VND.\n", money)
    money = 20000 // tiếp tục gán giá trị cho biến vừa khởi tạo
    fmt.Printf("trong ví tôi có: %d VND.", money)
}
```

- Kết quả của đoạn mà trên sẽ là:

```go
trong ví tôi có: 0 VND.
trong ví tôi có: 10000 VND.
trong ví tôi có: 20000 VND.
```


### Khai báo một biến cùng với khởi tạo giá trị

- Ở phần trên chúng ta đã học cách khai báo biến và gán giá trị cho chúng, tuy vậy có một điều thiếu sót đó là việc khởi tạo giá trị mặc định cho một biến thay vì phải lặp lại việc tạo biến sau đó mới gán biến. Thực tế  việc này làm rất đơn giản:

```go
package main

import "fmt"

func main() {  
    var money int = 10000
    fmt.Printf("trong ví tôi có: %d VND.", money)
}
```

- Nhìn xem, thật sự không có gì khác biệt, chúng ta chỉ thêm một giá trị gán vào khi khởi tạo biên mà thôi. Hãy chạy chương trình và kiểm tra xem bạn sẽ thấy kết quả sau được in ra:

```go
trong ví tôi có: 10000 VND.
```

### Suy luận kiểu ( Type inference )

- Như ban biêt trong các ngôn ngữ lập trình kiểu động ( dynamic type ), việc khai báo kiểu dữ liệu của một biến hầu như không có. Chúng ta chỉ cần khai báo biến và sử  dụng,máy ảo sẽ tự động suy luận và đoán ra được kiểu dữ liệu mà chúng ta sẽ dùng trong chương trình. Mặc dù là ngôn ngữ tĩnh nhưng Go cũng hỗ  trợ kiểu khai báo náy. Bằng cách xác định giá trị khởi tạo của một biến, Go sẽ phân tích và đưa ra kiểu dữ liệu mà biến đó sử  dụng. vì vây trong những trường hợp bạn đã xác đinh rõ giá trị khởi tạo rồi thì có thể  bỏ qua việc khai báo kiểu dữ liệu.

```go
package main

import "fmt"

func main() {  
    var money = 10000
    fmt.Printf("trong ví tôi có: %d VND.", money)
}
```

### Khai báo cùng đa biến 

- Trong một số  ngôn ngữ như C, Python chắc hẳn bạn không lạ khi thấy việc khai báo 
từ hai biến trở trong cùng một câu lệnh. Tương tự Go cũng hỗ  trợ cách làm này như sau:

```go
var name1, name2 [type] = initialvalue1, initialvalue2 
```

- trong đó tôi sử  dụng dấu [] cho kiểu dữ liệu, nó có ý nghĩa như một cách nói việc có thể  dùng hay không khai báo cũng được.

```go
package main

import "fmt"

func main() {  
    var bank, money = "Vnbank", 10000
    fmt.Printf("trong ngân hàng %s có: %d VND.", bank, money)
}
```

- Như bạn thấy ở đoạn mã trên chúng ta sử  dụng hai biến bank, money và khởi tạo giá trị cho chung, mặc dù chúng ta không khai báo kiểu dữ liệu xác định nhưng máy ảo Go vãn có thể suy luận được. Kết quả của đoạn mã trên sẽ là:

```go
trong ngân hàng Vnbank có: 10000 VND.
```

- Tuy vậy đoạn mã trên có một chú ý đó là việc khai báo và khởi tạo các biến trên cũng một dòng gặp một vấn đề  đó là nếu như chúng ta khai báo quá nhiều biến sẽ khiến cho dòng lệnh trở lên xuất xí mà khó nhìn. Dĩ liên là bạn có thể  tách nhỏ việc khai báo các biến ra. Nhưng thực sự không cần phải làm vậy, may mắn là trong Go hỗ  trợ cách giải quyết vấn đề  này. Bằng cách sử  dụng cặp `()` bạn có thể  khiến cho dòng mã trở lên sạch sẽ gọn gàng hơn qua cú pháp:


```go
var (  
    name1 = initialvalue1
    name2 = initialvalue2
    ...
    nameN = initialvalueNss
)
```  

- Két hợp cứ pháp trên:

```go
package main

import "fmt"

func main() {  
    var (
        name   = "A"
        bank   = "Vnbank"
        money int
    )

    fmt.Printf("Anh %s có %d VND trong ngân hàng %s.", name, money, bank)
}
```

- Nào hãy biên dịch và chạy thử  chương trình, ngay lập tức chúng ta sẽ có được dòng chữ được in ra như sau:

```go
Anh A có 0 VND trong ngân hàng Vnbank.
```

### Khai báo rút gọn

- Chúng ta đa học cách khai báo bién rồi, khởi tạo biến, gán giá trị cho một biến hay nhiều biến. Tuy vậy không trong Go còn cung cấp cho ta một cách khai bao biến ngắn gọn hơn. Thay vì sử  dụng từ khóa `var`, chúng ta có thể  sử  dụng toán tử  `:=` để  làm việc này thông qua cú pháp:

```go
name := initialvalue
```

- Chúng ta sẽ chỉnh sửa lại đoạn mã lúc trước kết hợp với cú pháp trên: 


```go
package main

import "fmt"

func main() {  
    bank, money := "Vnbank", 10000
    fmt.Printf("trong ngân hàng %s có: %d VND.", bank, money)
}
```

- Tuyệt! Giờ nhìn chúng trông thật ngắn gọn. Tuy nhiên có một lưu ý bạn cần biết khi sử  dụng cách khai báo này. Đối với việc sử  dụng kiểu rút gọn này nó yêu cầu các biến 
khai báo của chúng ta luôn phải có giá trị khởi tạo. Trong trường hợp sử  dụng đa biến, thì vế  bên phải tức vé giá trị khởi tạo cũng phải cung cấp đúng bằng số  lượng ở vế  bên trái tức vế  biến. 

```go
package main

import "fmt"

func main() {  
    bank, money := "Vnbank" // lỗi sẽ xảy ra bởi vì không cung cáp dủ giá trị
    fmt.Printf("trong ngân hàng %s có: %d VND.", bank, money)
}
```

- Và một khi tạo khởi tạo thì chúng ta không thể  dùng cách viết này với các biến 
đó nữa:

```go
package main

import "fmt"

func main() {  
    bank, money := "Vnbank", 10000 // lỗi sẽ xảy ra bởi vì không cung cáp dủ giá trị
    bank, money = "Vnbank", 10000 // trường hợp này dúng
    bank, money := "V", 10 // trường hợp này sai xảy ra lỗi vì biến đã được tạo ra rồi
    fmt.Printf("trong ngân hàng %s có: %d VND.", bank, money)
}
```

- Ngoài ra còn một Lưu ý là một khi biến đã được khởi tạo và gán giá trị bất kể  là 
biến dó bạn đã được khai báo cùng kiểu dư liệu hay không thì bạn không thể  gán nó với 
giá trị khác được ( Hãy nhớ Go làm một ngữ mang tính đê  cao về  kiểu tính - strong static type )  

## Tổng kết

- Trong chương này bạn đã học về  các cái khai báo biến sử  dụng chúng cho một số ví dụ đơn giản rồi cùng một só kiểu dữ liệu rồi. Trong chương tiếp theo chúng ta sẽ đi sâu vào các kiểu dữ liệu hỗ  trợ trong Go. 