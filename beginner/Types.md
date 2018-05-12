## Các kiểu dữ liệu ( Types )

- Trong chương trước chúng ta đã tìm hiểu về  biến cùng với một vài kiểu dữ liệu như số nguyên, chuối. Trong 
chương này chúng ta sẽ đi chi tiết hơn vè các kiểu dữ liệu hỗ  trợ trong Go.

- Về  cơ bản trong go hỗ  trợ 3 kiểu dữ liệu chính: chuối ( string ), số  học ( numeric ) và boolean , còn một số  kiểu cấu trúc khác chúng ta sẽ đề  cập sau. Trong các kiểu này , đặc biệt kiểu só học sẽ gồm các kiểu con, chúng sẽ được biểu diễn dựa trên sơ đồ  sau:

```go
- bool
- Numeric Types
    int8, int16, int32, int64, int
    uint8, uint16, uint32, uint64, uint
    float32, float64
    complex64, complex128
    byte
    rune
- string
```

### Bool

- Kiểu bool hay còn gọi là đại số  boolean là kiểu dữ liệu gồm hai giá trị `true` và `false` được sử  dụng 
để  biểu thị một logic nào đó mang tính chất đúng hay sai.

```go
package main

import "fmt"

func main() {  
	x := true
	y := false
	z := x && y
	k := x || y
    fmt.Println("x:", x, "y:", y, "z:", z, "k:",k)
}
```

- Trong đoãn mã trên chúng ta đã khởi tạo hai biến x,y với hai giá trị boolean. Tiếp đó hãy để  ý chúng ta sử  dụng hai toán tử 
`&&` ( and ) và `||` ( or ) và hiện với hai biến x, y vừa tạo rồi gán chúng cho hai biến mới z, k. Để  giải thích tri tiết hơn chúng ta 
sẽ đi xét đại số  boolean, như bạn đã biết đại số  boolean gồm hai giá trị `true` và `false`, khi một trong hai giá trị này kết hợp với nhau thông qua 
một toán tử  `!` ( not ), `&&` ( and ), `||` ( or ) sẽ có dạng như sau:

```go
not true -> false
not false -> true
true and true -> true
true and false -> false
false and false -> false
true or true -> true
true or false -> true
false or false -> false
``` 

- Từ biểu thức trên chúng ta sẽ có được kết quả sau khi chạy đoạn mã:

```
x: true y: false z: false k: true m: false n: true
``` 

### Kiểu số học ( Numeric )

- Tiếp theo chúng ta sẽ đi vào tìm hiểu các kiểu dữ liệu số học được liệt kê ở đầu chương.

#### Kiểu số  nguyên ( Integer )

- Nếu như bạn đẻ ý ở những dòng đầu của kiểm số  học, bạn sẽ thấy rằng trong Go phân biệt ra rất nhiều kiểu dữ liệu số nguyên. Trong đó bao gồm hai loại kiểu 
đó là kiểu dữ liệu số  nguyên có đấu ( signed integer ) và kiểu dữ liệu số  nguyên không có dấu ( unsigned integer ) :


##### Các kiểu số  nguyên có dấu ( Signed integers )  

- int8: là kiểu dữ liệu số  nguyên để  hiển với 8 bit dữ liệu

  kích thước: 8 bits
  
  phạm vi: từ -128 tới 127

- int16: là kiểu dữ liệu số  nguyên được biểu diễn với 16 bit dữ liệu

  kích thước: 16 bits

  phạm vi: từ -32768 tới 32767

- int32: là kiểu dữ liệu số  nguyên được biểu diễn với 32 bit dữ liệu

  kích thước: 32 bits

  phạm vi: từ -2147483648 tới 2147483647

- int64: là kiểu dữ liệu số  nguyên được biểu diễn với 64 bit dữ liệu
  
  kích thước: 64 bits

  phạm vi: từ -9223372036854775808 tới 9223372036854775807

- int: là kiểu dữ liệu số  nguyên được biểu diễn từ 32 tới 64 bit dữ liệu, số  bit sẽ phụ thuộc vào nền tảng (     platform ), kiến trúc mà chúng ta sử  dụng. Thông thường chúng ta thường dùng kiểu dũ liệu này đối với một giá   trị số  nguyên mà chúng ta không xác định rõ phạm vi giá trị của nó. 

  kích thước: 32 bits với hệ thống 32 bit và 64 bit với hệ thống 64 bit.
  
  phạm vi: từ -2147483648 tới 2147483647 trong hệ thống 32 bit và giá trị từ -9223372036854775808 tới 9223372036854775807 trong hệ thống 64 bit.

- Tiếp theo chúng ta sẽ tiến hành viết một số  ví dụ để  phân tích các kiểu dữ liệu trên:

##### Các kiểu số  nguyên không có dấu ( Unsigned integers )  

- Tương tự giống với trường hợp trong Signed integer, chúng cũng bao gồm các kiểu dũe liệu sau

- uint8: là kiểu dữ liệu số  nguyên để  hiển với 8 bit dữ liệu

  kích thước: 8 bits
  
  phạm vi: từ 0 tới 255

- uint16: là kiểu dữ liệu số  nguyên được biểu diễn với 16 bit dữ liệu

  kích thước: 16 bits

  phạm vi: từ 0 tới 65535

- uint32: là kiểu dữ liệu số  nguyên được biểu diễn với 32 bit dữ liệu

  kích thước: 32 bits

  phạm vi: từ 0 tới 4294967295

- uint64: là kiểu dữ liệu số  nguyên được biểu diễn với 64 bit dữ liệu
  
  kích thước: 64 bits

  phạm vi: từ 0 tới 18446744073709551615

- uint: tương tự kiểu int là kiểu dữ liệu số  nguyên được biểu diễn từ 32 tới 64 bit dữ liệu, số  bit sẽ phụ thuộc vào nền tảng ( platform ), kiến trúc mà chúng ta sử  dụng. Thông thường chúng ta thường dùng kiểu dũ liệu này đối với một giá trị số  nguyên mà chúng ta không xác định rõ phạm vi giá trị của nó. 

  kích thước: 32 bits với hệ thống 32 bit và 64 bit với hệ thống 64 bit.
  
  phạm vi: từ 0 tới 4294967295 trong hệ thống 32 bit và giá trị từ đ tới 18446744073709551615 trong hệ thống 64 bit.
 

```go
package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
  var a int8 = 89
  var a1 uint = 90
  var b int16 = 1111
  var b1 uint16 = 1112
  var c int32 = 1111111
  var c1 uint32 = 1111112
  var d int64 = 1111111111111
  var d1 uint64 = 1111111111112
  
	e := 95

  fmt.Println("giá trị của a là:", a, 
        "\ngiá trị của a1 là:", a1, 
        "\ngiá trị của b là:", b, 
        "\ngía trị của b1 là:", b1, 
        "\ngía trị của c là:", c, 
        "\ngiá trị của c1 là:", c1, 
				"\ngiá trị của d là:", d, 
        "\ngía trị của d1 là:", d1,
        "\ngía trị của e là:", e)

  fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
  fmt.Printf("kiểu dữ liệu của a1 là %T, kích thước của a1 là %d byte\n", a1, unsafe.Sizeof(a1)) 
  fmt.Printf("kiểu dữ liệu của b là %T, kích thước của b là %d bytes\n", b, unsafe.Sizeof(b)) 
  fmt.Printf("kiểu dữ liệu của b1 là %T, kích thước của b1 là %d bytes\n", b1, unsafe.Sizeof(b1)) 
  fmt.Printf("kiểu dữ liệu của c là %T, kích thước của c là %d bytes\n", c, unsafe.Sizeof(c))
  fmt.Printf("kiểu dữ liệu của c1 là %T, kích thước của c1 là %d bytes\n", c1, unsafe.Sizeof(c1)) 
  fmt.Printf("kiểu dữ liệu của d là %T, kích thước của d là %d bytes\n", d, unsafe.Sizeof(d))
  fmt.Printf("kiểu dữ liệu của d1 là %T, kích thước của d1 là %d bytes\n", d1, unsafe.Sizeof(d1)) 
	fmt.Printf("kiểu dữ liệu của e là %T, kích thước của e là %d bytes", e, unsafe.Sizeof(e)) 
}
```

- Ở đoạn mã nãy chúng ta sẽ tién hành lần lượt khai báo các biến thuộc các kiểu dữ liệu có trong signed integer. Sau đó chúng ta sẽ tiến hành lần lượt in ra giá trị của chúng cùng với kích thước ra ngoài màn hình. Ở đây để  hiên thị được chính xác kiểu dữ chúng ta sẽ sử  dụng đinh dạng `%T` trong hàm `Printf` và sử  dụng hàm `Sizeof` có trong package unsafe để  quy ra kích thước mà biến dó đang sử  dụng. Lưu ý là package unsafe không được khuyến cáo sử  dụng trong các chương trình thực tế  tuy nhiên trong chương trình này chúng ta sẽ sử  dụng mục đích là để hiển thị kích thước của các biến.

- Nào hãy biên dịch và chạy chương trình bạn sẽ thấy thông báo được dưa ra:

```go
giá trị của a là: 89
giá trị của a1 là: 90
giá trị của b là: 1111
gía trị của b1 là: 1112
gía trị của c là: 1111111
giá trị của c1 là: 1111112
giá trị của d là: 1111111111111
gía trị của d1 là: 1111111111112
gía trị của e là: 95
kiểu dữ liệu của a là int8, kích thước của a là 1 byte
kiểu dữ liệu của a1 là uint, kích thước của a1 là 8 byte
kiểu dữ liệu của b là int16, kích thước của b là 2 bytes
kiểu dữ liệu của b1 là uint16, kích thước của b1 là 2 bytes
kiểu dữ liệu của c là int32, kích thước của c là 4 bytes
kiểu dữ liệu của c1 là uint32, kích thước của c1 là 4 bytes
kiểu dữ liệu của d là int64, kích thước của d là 8 bytes
kiểu dữ liệu của d1 là uint64, kích thước của d1 là 8 bytes
kiểu dữ liệu của e là int, kích thước của e là 8 bytes
```

- Như bạn thấy ứng với mỗi giá trị chúng ta khai báo so khá khớp với bản liệt kê phía trên, đặc biệt trong trường hợp biến e, là biến chúng ta sử  dụng chức năng tự động suy kiểu dữ liệu do dó mặc định của nó sẽ là kiểu `int` và ứng với hệ thống tôi đang sử  dụng là 64 bit.

- Đối với trường hợp bạn gán một giá trị vượt qua phạm vi mà kiểu dữ liệu đó cho phép, công cụ bien dịch sẽ ngay lập tức đưa ra lỗi:

```go
package main

import (  
    "fmt"
)

func main() {  
	var a int8 = 891 \\ lỗi constant 891 overflows int8 sẽ được đưa ra
	fmt.Println("giá trị của a là:", a)
}
```

### Kiểu số  thực ( Floating point )

- Trong một số  ngôn ngữ khác chúng ta thấy kiểu thông thường sẽ chia làm hai kiểu số  thực đó là `float` và `double` để  biểu diên phạm vi từ 32 bit và 64 bit, tuy nhiên trong Go để  tránh nhầm lẫn khi sử  dụng chúng sẽ được quy về  float kèm sau đó là các số  bit thể  hiện:

- float32: kiểu số  thực với kích thước 32 bit 
- float64: kiểu số  thực với kích thước 64 bit 

```go
package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
  var a float32 = 89.1
  var b float64 = -909.2


  fmt.Println("giá trị của a là:", a, 
        "\ngiá trị của b là:", b)

  fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
  fmt.Printf("kiểu dữ liệu của b là %T, kích thước của b là %d bytes\n", b, unsafe.Sizeof(b)) 
}
```

### Kiểu số  phức ( Complex type )

- Trong toán học ngoài số  nguyên và số  thực ra chúng ta còn áp dụng số  phức để  giả quyết các bài toán về tích phân, hình học. Chúng thường dược viết dưới dạng:

```go
a + bi
```

- Tương tự với các kiểu dữ liệu trước chúng cũng được phân làm hai loại:

  + complex64: đây là kiểu dữ liễu số  phức trong đó phần thực và phần ảo sử  dụng kiểu dữ liệu số  thực 32 bit
  + complex128: đây là kiểu dữ liễu số  phức trong đó phần thực và phần ảo sử  dụng kiểu dữ liệu số  thực 64 bit

- Để  khai báo và khởi tạo giá trị cho chúng bạn có thể  viết như sau

```
var a complex64 = 1 + 2i
a := 1 + 2i
```
- Nhìn xem thật đơn giản phải không! Ngoài ra thông thường trong một sô ngôn ngữ chúng thường sẽ được tạo ra dưới bằng cách tạo một đối tượng từ một lớp hay một hàm.
Trong Go bạn có thể gọi tới hàm `complex(r, i FloatType)` được xây dừng trong [complex](https://golang.org/pkg/builtin/#complex) để  khởi tạo 
một số phức. Hàm này sẽ nhận 2 tham số đầu vào tương ứng với phần thực và phần ảo mà bạn muốn khởi tạo và trả về  một giá trị là kiểu số phức phụ thuộc vào kiểu dữ liệu số thực mà bạn truyền vào, vd nếu bạn khởi tạo cùng 2 giá trị thuộc kiểu  float32 nó sẽ trả về  kiểu số phức complex64, tương tự với hai kiểu số  thực float64 thì complex128. Tuy nhiên hai giá trị truyền vào này phải có cùng kiểu dữ liệu, nếu không trình biên dịch sẽ ngay lập tức thông báo lỗi `mismatched types float32 and float64`.

- Bây giờ chúng ta sẽ xét một ví dụ cụ thể:

```go
package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
    var a complex64 = 1 + 2i
    var b float32 = -909.2
    var c float64 = 11111

    var d = complex(b, b)
    var e = complex(c, c)

    fmt.Println("giá trị của a là:", a, 
				"\ngiá trị của b là:", b, 
				"\ngía trị của c là:", c, 
				"\ngiá trị của d là:", d, 
				"\ngía trị của e là:", e)
	
	fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
	fmt.Printf("kiểu dữ liệu của b là %T, kích thước của b là %d bytes\n", b, unsafe.Sizeof(b)) 
	fmt.Printf("kiểu dữ liệu của c là %T, kích thước của c là %d bytes\n", c, unsafe.Sizeof(c))
	fmt.Printf("kiểu dữ liệu của d là %T, kích thước của d là %d bytes\n", d, unsafe.Sizeof(d))
	fmt.Printf("kiểu dữ liệu của e là %T, kích thước của e là %d bytes", e, unsafe.Sizeof(e)) 
}
```

- Ở đây tôi sẽ sử  dụng kiểu khai báo thuần và khai báo thông qua việc dùng hàm `complex`. Sau khi biên dịch và chạy chương trình
bạn sẽ thấy kết quả sau được in ra:

```go
giá trị của a là: (1+2i)
giá trị của b là: -909.2
gía trị của c là: 11111
giá trị của d là: (-909.2-909.2i)
gía trị của e là: (11111+11111i)
kiểu dữ liệu của a là complex64, kích thước của a là 8 byte
kiểu dữ liệu của b là float32, kích thước của b là 4 bytes
kiểu dữ liệu của c là float64, kích thước của c là 8 bytes
kiểu dữ liệu của d là complex64, kích thước của d là 8 bytes
kiểu dữ liệu của e là complex128, kích thước của e là 16 bytes
```

### Byte

- Về  cơ bản nó chỉ là là một định danh của kiểu dũ liệu uint8 mà thôi. Mục đính chính của kiểu này chỉ đơn giản được dùng để  
biểu diễn các giá trị của chuỗi dữ liệu.
```go
package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
    var a byte = 89
    fmt.Println("giá trị của a là:", a, )
    fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
}
```

- Kết quả nhận dược sau khi bạn chạy chương trình:

```go
giá trị của a là: 89
kiểu dữ liệu của a là uint8, kích thước của a là 1 byte
```

### Rune

- Tượng tự kiểu byte, rune cũng là một đinh danh của kiểu dữ liệu int32

```go
package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
    var a rune = 89
    fmt.Println("giá trị của a là:", a, )
    fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
}
```

- Kết quả nhận được sau khi chạy chương trình:

```go
giá trị của a là: 89
kiểu dữ liệu của a là int32, kích thước của a là 4 byte
```

### Kiểu chuõi ( String )

- Nếu bạn đã từng học một số ngôn ngữ khác thì chắc hẳn không quá xa lạ gì với kiểu dữ liệu chuỗi rồi, về  bản chất kiểu 
chuối ( string ) chỉ là một tập hợp các ký tự mà mỗi ký tự là đại diện cho một kiểu dữ liệu , thông thường trong một số  ngôn
ngữ chúng sẽ được định nghĩa là kiểu dữ liệu `char`, tuy nhiên bản chất của mỗi ký tự chỉ đơn giản là môt giá trị trong phạm vi kích thước 8 bit mà thôi ( ngoài ra với ký tự khác như [unicode](https://golang.org/pkg/unicode/) sẽ có cách biểu diễn khác một chút ). 

- Xét ví dụ:

```go
package main

import (  
	"fmt"
	"unsafe"
)

func main() {  
  first_name := "Cong"
  last_name := "Quan"
	fmt.Printf("Tên của tôi là %s %s\n", first_name, last_name)
	fmt.Printf("Độ dài là %d\n", len(first_name)) 
    fmt.Printf("kiểu dữ liệu là %T, kích thước là %d byte\n", 
                first_name, 
                unsafe.Sizeof(first_name)) 
    fmt.Printf("kiểu dữ liệu của một phần tử  là %T, kích thước của một phần tử  là %d byte\n", 
             first_name[0], 
             unsafe.Sizeof(first_name[0])) 
}
```

- Ở đâu chúng ta sẽ khai báo hai giá trị thuộc kiểu chuối sau đó hiển thị chúng, Như tôi đã nói ở trên bản chất của kiểu chuỗi là một tập hợp các ký tự mà mỗi ký tự được biểu diễn dưới dạng một byte ( 8 bit ), chúng ta sẽ sử  dụng hàm `len` để  hiển thị độ dài của một chuỗi và dùng toán tử  `[]` kèm một giá trị số  nguyên để  truy cập vào một ký tự trong chuỗi ký tự. Sau đó tương tự như các ví dụ trên chúng ta sẽ hiển thị thông tin về  kiểu dữ liệu đó ra màn hình.

- Nào hãy chạy đoạn mã trên và bạn sẽ nhận được kết quả:

```go
Tên của tôi là Cong Quan
Độ dài là 4
kiểu dữ liệu là string, kích thước là 16 byte
kiểu dữ liệu của một phần tử  là uint8, kích thước của một phần tử  là 1 byte
```

- Đó đúng có vể  đúng như những gì tôi nói phải không. 


### Giá tri mặc định ( Zero Value )

- Trong những ví dụ đâu tiên chúng ta đã thấy việc khai báo dữ liệu mà không kèm với khởi tạo giá trị, tuy nhiên khi thực hiện 
in ra màn hình chúng ta vẫn thấy chúng có giá trị. Những giá trị như vậy là những giá trị mặc định được tạo ra khi chúng ta khai báo một kiểu dữ liệu xác định mà không xác định giá trị cần tạo.

- các giá trị mặc định của các kiểu được dựa trên sơ đồ  sau:

```go
numeric: 0
bool: false
string: ""
```

### Chuyển kiểu ( Type Conversion )

- Vậy là các kiẻu dữ liệu cơ bản chúng ta đã tìm hiểu rồi, tuy vậy có một điều mà chúng ta chưa để  cập tới đó là giả sử  như bạn đa có một biến với kiểu dữ liệu int32, và bạn sử  dụng biến này để  nhận kết quả của một phép chia, nhưng giá trị của bạn tạị một thời điểm lại là kiểu số  thực, có điều như bạn biết Go là ngôn ngữ lập trình yêu cầu khắt khe về  mặt kiểu dữ liệu tĩnh ( strong static type ) do đó nó không thể  tự động chuyển kiểu đc vậy thì phải làm như thé nào trong trường hợp đó. 

- Vd:
```go
package main

import (  
    "fmt"
)

func main() {  
    var result int32
    result = 4 / 2 // ok

    result = 4 / 2.0 // error
    fmt.Println(result)
}
```

- Hm kết quả của phép chia thứ hai trên ta được két quả là một số  thực và rõ ràng khống thể  gán một giá trị thuộc kiểu số thực
cho biến `result` được. Do đó công cụ biên dịch phát hiện và thông báo lỗi, tuy nhiên chúng ta có thể  sửa lại một chút để  tránh lỗi này bằng cách chuyển kiểu của kết quả nhận được được phép chia:

```go
package main

import (  
    "fmt"
)

func main() {  
    var result int32
    result = 4 / 2 // ok

    result = 4 / int32(2.0)  // ok
    fmt.Println(result)
}
```
Như bạn thấy bằng cách sử  dụng chuyển kiểu thông quả sử  dụng `type` + `(value)`, chúng ta có thê chuyển đổi kiểu dữ liệu sang dạng kiểu dữ liệu mong muồn tuy nhiên không phải kiểu dữ liệu nào cũng chuyển đổi dược do đó tránh lạm dụng chúng.


### Tổng kết

- Trong chương này chúng ta đã làm quen với các kiểu dũ liệu cơ bản trong Go rồi, chúng ta cũng học được cách thao tác với một số kiểu và chuyển đổi kiểu. Trong các chương sau chúng ta sẽ đi sâu hơn vào một số  kiểu dữ liệu này. 