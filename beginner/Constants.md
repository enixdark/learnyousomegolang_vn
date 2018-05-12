## Hằng số ( Constants )

- Gỉa sử  bạn đang làm một bài toán về  tính chu vi ( P ) và diện tích ( S ) của một hình tròn,
Nào hãy lục lại kiến thức toán học mà trước đó chúng ta đã học, để  tính diện tích và chi vi của một hình tròn, chúng ta cần biết trước số  Pi, và giờ chúng ta sẽ đặt một biến gọi là P thuộc kiểu dữ liệu số  thực và khởi tạo nó với giá trị là `3.14`. Oh như vậy là ta đã có , giờ việc cần làm tiếp theo là định nghĩa bán kính của hình tròn và thực hiện phép tính với số  Pi. 
- Thật dễ  dàng phải không, ah mà khoan nhưng ở đoạn dưới dó tôi lại sử  dụng P để  gán với giá trị chu vi vừa tính dược thì sao. Oh như vậy là biến P đã thay đổi và chúng ta không có giá trị số Pi nữa rồi. Thật tệ, chúng ta phải thay đổi lại tên của biến Pi thôi, nhưng một lúc sau chúng ta lại vô tình gán nhầm giá trị biến Pi. Hm lúc này việc cần làm đó là chúng ta cần định nghĩa một biến có giá trị không thay đổi lúc này. Nhưng biến như vậy sẽ gọi là hằng số .

```go
package main

import (  
    "fmt"
)

func main() {  
    var P float32 = 3.14
	var R float32 = 2.0 
	var S float32
    P = 2 * P * R
    S = P * R * R // logic error 
	fmt.Println("P:",P, "S:", S)
}
```

- Hăng sô được sử  dụng để  định nghĩa cho một biến có giá trị không thay đổi từ lúc bắt đầu khởi tạo nó. 

Trong Go nó sẽ được khai báo như sau:

```go
const Var [type] = Value
```

```go
package main

func main() {  
    const P = 3.14
    var R int32 = 1 
    P = 2 * P * R // error: cannot assign to P
}
```

- Như bạn thấy chúng không khác nhiều so với cách chúng ta khai báo biến thông thg, ngoài việc đơn giản là khai báo cùng với từ khóa `const` và giá trị khởi tạo là được. tuy nhiên ở đấy có một điều cần ghi nhớ dó là trong hằng số  được chia ra làm 2 loại đó là hằng số  cùng với kiểu dữ liệu xác định ( type contasnt ) và những hằng số  bất định ( untype contasnt ). Vậy sự khác nhau giữa hai loại hằng số  này là gì ? Sớm thôi chúng ta sẽ tìm hiểu về  nó.

- Trong Go dược phân loại thành rất nhiều dạng hằng số  vd như boolean constants, rune constants, integer constants, floating-point constants, complex constants, and string constants. Trong đó với kiểu rune, integer,floating-point, and complex constants sẽ được quy ước chung thành numeric constant

- Lưu ý: Có một điều cần biết đó là khi khai báo một hằng chúng ta phải luôn xác định rõ giá trị khởi tạo trước khi biên dịch, trong trường hợp hằng được tạo ra trong quá trình chạy chương trình sẽ không được phép. 

```go
package main

func main() {  
    var R int32 = 1 
    const P = 3.14 * 2 * R // error: const initializer 3.14 * 2 * R is not a constant
}
``` 

#### Hằng số  xác định và không xác định ( Type Constant và Untype Constant )

- Ở phần trên chúng ta đã nghe qua về  hằng số  xác định ( type constants )  và hằng số  bất định ( untype constants ). Việc giải thích chúng theo ý thuyết tương đối khó hình dung lên bằng phương pháp thực hành qua ví dụ chúng ta sẽ có thể hiểu tốt hơn.


#### Type constant

- Đầu tiên hằng số  xác định ( type constants ) chúng là những hằng số  mà chúng ta đã xác định sẵn kiểu dữ liệu khởi tạo, vd:

```go
const a int = 1
const name string = "Quan"
```

- chìn chung không khác nhiều so với việc khai báo một biến cùng kiểu dữ liệu lắm.

### Untype constant

- Điểm nhấn ở day chính là hằng số  bất định, cơ bản bạn biết rằng kiểu dữ liệu trong Go khác chặt chẽ, néu bạn đã định nghĩa một biến thuộc kiểu số nguyên rồi thì không thể  sư  dụng giá trị của nó sang một biến khác với kiểu số  thực hay một kiểu số  học nào khác. 

- Tuy nhiên trường hợp đó chỉ đúng với khai báo biến thông thường, với hằng số  bất định ( untype constant ) sẽ được ngầm dịnh kiểu dữ liệu và việc ngâm định này thường không xác định chính kiểu dữ liệu mà nó sẽ sử  dụng. Tại sao tôi nói như vậy ? Chúng ta sẽ xét ví dụ sau đây:

```go
package main

import (  
	"fmt"
	"unsafe"
)

func main() {  
	
	const a  = 1
	var b int8 = a
	var c int32 = a
	var d float32 = a
	var e complex64 = a

	type another_string_type string
	const text = "untype constant"
	var f another_string_type = text

	fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
	fmt.Printf("kiểu dữ liệu của b là %T, kích thước của b là %d byte\n", b, unsafe.Sizeof(b)) 
	fmt.Printf("kiểu dữ liệu của c là %T, kích thước của c là %d byte\n", c, unsafe.Sizeof(c)) 
	fmt.Printf("kiểu dữ liệu của d là %T, kích thước của d là %d byte\n", d, unsafe.Sizeof(d)) 
	fmt.Printf("kiểu dữ liệu của e là %T, kích thước của e là %d byte\n", e, unsafe.Sizeof(e)) 
	fmt.Printf("kiểu dữ liệu của f là %T, kích thước của f là %d byte\n", e, unsafe.Sizeof(e))
}
```



- Kết quả sau khi chạy sẽ dược:

```go
kiểu dữ liệu của a là int, kích thước của a là 8 byte
kiểu dữ liệu của b là int8, kích thước của b là 1 byte
kiểu dữ liệu của c là int32, kích thước của c là 4 byte
kiểu dữ liệu của d là float32, kích thước của d là 4 byte
kiểu dữ liệu của e là complex64, kích thước của e là 8 byte
kiểu dữ liệu của c là main.another_string_type, kích thước của c là 16 byte
```

- Oh! nhìn xem có một sự ngầm định khi gán kiểu dữ liệu với hằng số  bất định, điều mà bạn bạn làm với toán tử  `:=` không được nhưng với hằng số  thì lại được.
- Ngoài ra có một chút gì đó mới mẻ ở đây đó là tôi sử  dụng từ khóa `type`, bằng cách sử  dụng từ khóa này 
sẽ cho phép chúng ta định nghĩa ra một kiểu dữ liệu mới dựa trên các kiểu dữ liệu thuần hay những kiểu dữ liệu đã định nghĩa từ trước, trong các chương sau chúng ta sẽ hiểu rõ hơn cách sử  dụng từ khóa này.
Ở đây tôi đã định nghĩa ra một một hằng số  số  học ( numeric ) là a với một giá trị số nguyên nhưng thực chất là một kiểu chưa được xấc dịnh, sau đó bắt đàu định nghĩa các kiểu dũ liệu số học khác kèm với khởi tạo giá trị từ hằng số  `a`, tiếp dó tôi tự định nghĩa một kiểu dữ liệu dựa trên kiểu dữ liệu chuỗi và tạo ra một hằng số  với giá trị là một chuỗi, sau đó tạo ra một biến thuọc kiểu dữ liệu mà tôi tự dịnh nghĩa và gán nó với giá trị của hằng số  chuỗi vừa tạo. Thật tuyệt là chương trình chạy mà không gặp bất cứ vấn đề  lỗi gì.     

- Vậy Lợi ích mà chúng ta có được khi sử  dụng hằng số :
  + Chúng ta sẽ giảm được rủi ro gặp phải trong logic nếu như một biến chó thể chỉnh sửa được 
  + Đối với kiểu số  học.


### Tổng kết

- Trong chương này chúng ta đã làm quen với hằng số , cách định nghĩa một hằng số  dùng cho một giá trị không đối và nắm rõ hơn về  hằng số  xác định ( type constant ) và hằng số  bất dịnh ( type constant ) thông qua một vài ví dụ. Ngoài ra cũng làm quen với một từ khóa mới cho viẹc định nghĩa một kiểu dữ liệu riêng.

   