## Hàm ( Function )

- Hàm là một khái niệm để  chỉ một khối mã hay nói cách khác là một tập hợp các đoạn mã, logic cho một công 
việc xác định ( ví dụ như thực hiện một phép tính tổng , etc ). Một hàm sẽ có thể  nhận một số  dữ liệu đâu vào để  xử  lí logic và trả về  kết quả. dưới đây là hình ảnh minh hoa cụ thể:

<br>
<p align="center">
  <img src="functions.jpg"/>
</p>
<br>

### Định nghĩa hàm

- Trong Go để  khai báo một hàm chúng ta sẽ sử  dụng từ khóa `func` tiếp theo đó là tên hàm cần định nghĩa và các tham số  cùng kiểu dữ liệu xác định cần đưa vào trong hàm, chúng sẽ đặt trong cặp ngoặc `()` ngay sau tên hàm và kiểu dữ liệu cân trả về  hàm, sau đó là nội dung logic cần viết trong `{}`. Nó sẽ có dạng như sau:

```go
func function_name(param type) return_type {
    // logicS
    // return value
}
```

Vd:

```go
package main

import (
    "fmt"
)

func sum(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println(sum(1, 2))
    fmt.Println(sum(2, 4))
}
```

- ở đây chúng ta đinh nghĩa một hàm tính tổng nhận đầu vào là hai tham số  thuộc kiểu dữ liệu số nguyên và kết quả trả về  là kiểu dữ liệu số  nguyên, bên trong hàm này chúng ta sẽ thực hiện phép công hai tham số  đưa đã đưa vào.

- Để  gọi hàm này chúng ta sẽ gọi tên hàm mà chúng ta đã định nghĩa kèm tham số truyền vào:

```go
sum(1 ,2) \\ 3

sum(2, 4) \\ 6
```

- Tuy vậy nếu như bạn để  y hàm `main` của chúng ta nó không hề  nhận bất kỳ tham số  đâu vào nào hay trả về  kết quả nào hết. Đúng vậy lúc trước chúng ta đã nói tới hàm `main` là một hàm dặc biệt do đó có thể  nó là trường hợp duy nhất mà ko cần phải định nghĩa như sơ đồ.

- Ah thật ra không phải vậy, về  cơ bản ở sơ đồ  trên tôi đưa ra một trường hợp tổng quá của hàm thôi. Ngoài ra còn một số trường hợp khác nữa vd như giống vơi hàm `main` không cần xác định tham số  và kết quả trả về  hay một số trường hợp trả về  đa giá trị, etc.


### Hàm không cùng với tham số  và kết quả về

- Như đã đề  cập phía trên chúng ta có rất nhiều loại hàm trong đó hàm không có kết quả trả về  hay đối số  sẽ được viết dưới dạng như sau:


```go
func function_name(){
    // logicS
    // return value
}
```

Vd:

```go
package main

import (
    "fmt"
)

func hello_go() {
    fmt.Println("Hello Go!")
}

func main() {
    hello_go()
}
```

- Hãy chạy chương trình và bạn sẽ nhận được kết quả như dưới đây:

```go
Hello Go!
```

### Hàm cùng trả về  đa giá trị 

- Hầu hết các hàm mà chúng ta đã gặp hay đã làm từ trước tới giờ hay trong các ngôn ngữ khác thì chỉ trả về  duy nhất môt giá trị. Nếu muốn trả về  nhiều giá trị bạn sẽ phải bao bọc dũ liệu trong một cấu trúc hay đối tượng nào đấy ( nếu bạn không hiểu đoạn này thì không sao bởi vì ở các chương sau chúng ta sẽ học về  chúng). Tuy nhiên trong Go hỗ  trợ thêm cho chúng ta một cách nữa để  trả về  nhiều giá trị cùng lúc mà không phải dùng tới một cấu trúc nào hết.

- Bạn chỉ cần dựa theo mẫu sau:

```go
func function_name(param type) (return_type1, return_type2, ... ,return_typeN )  {
    // logicS
    // return value1, value2
}
```

Vd:

```go
package main

import (
    "fmt"
)

func people(name string, age int) (string, int) {
    return name, age
}

func main() {
    fmt.Println(people("Quan", 25))
}
```

- Thật dễ  dàng phải không, bạn chỉ cần xác định các dữ liệu cần trả về  và dặt chúng trong cặp ngoặc `()` là xong. Không có quá gì phức tạp cả và việc trả về  nhiều giá trị này còn có một ý nghĩa khá quan trọng, chỉ một chốc thôi bạn sẽ thấy.

### Hàm trả về  gía trị lỗi

- Khi xử  lí logic một hàm hay nói cách khác là khi viết một chương trình bạn chắc chắn không thể  không bắt gặp các lõi không mong muốn xảy ra, để  tránh điều này chúng ta sẽ thường sử  dụng try..catch , throw để  bắt và đưa ra các thông báo lỗi tuy vậy cách làm này thường không phải là lựa chọn hay thay vào đó chúng ta vẫn có thể  xử  lí một hàm và trả về  kết quả lỗi nếu có.

- Trong Go, chúng ta sẽ thường xuyên thấy sự kết hợp việc sử  dụng hàm đa giá trị trả về  dùng với một kêt quả lỗi:

Vd:

```go
package main

import (
	"fmt"
	"errors"
)

func divide(a float32, b float32) (error, float32) {
    if b == 0 {
        err := errors.New("Không thể  thực hiện được phép chia")
        return err, 0.0
    }
    return nil, a / b
}

func main() {
    err, result := divide(1, 0)
    if err != nil {
        fmt.Println(err)
    }
    err, result = divide(1, 1)

    fmt.Println(result)
}
```

- Thực hiện chương trình trên ta sẽ được kết quả 

```go
Không thể  thực hiện được phép chia
1
```

- Thật tuyệt vời phải không ?, bằng cách làm này chúng ta đơn giản hóa được mọi việc, chỉ cần xác định xem có lỗi hay không thông qua dư liệu trả về  thay vì phải lồng vào trong một khối xử  lí lỗi.

### Định danh trống ( Blank Identifier )

- Mặc dù với ví dụ trên chúng ta đã thấy được mặt thú ví của việc sử  dụng đa giá trị trả về  từ một hàm. Nhưng 
có một điều đáng chú ý ở đây là đôi khi một số  kết quả chúng ta chỉ quan tâm tới một giá trị thôi còn lại có thể  không sủ dụng tới. Tuy nhiên bởi vì Go làn ngôn ngữ yêu cầu chặt chẽ, do đó nó yêu cầu tất cả các biến được khai báo
đèu phải sử  dụng. Như vậy giả sử  tôi có một hàm trả về  phần tử  đầu tiên của một cặp giá trị truyền vào, và tôi chỉ cần lấy phần tử  đầu tiên ra, nhưng bởi vì hàm này sử  dụng đa giá trị truyền vào trong đó giá trị đàu tiên dùng để  biết được lỗi của chương trình. Tuy vậy với một hàm chỉ đơn giản là trả về  giá trị sẽ không có lỗi. vậy tôi có cách nào để  làm mà không cần phải sử  dụng biến lỗi kia không ? Thật may mắn là các nhà phát triển của Go đã lường trước và hỗ  trợ đièu này. Bằng cách sử  dụng ký tự `_` đặt vào ví trí các gía trị mà chúng ta muốn bỏ qua, không quan tâm tới.

Vd:



```go
package main

import (
    "fmt"
)

func head(a int, b int) (error, int) {
    return nil, a
}

func main() {
    _, value := head(1,2)
    fmt.Println(value)
}
```

- Giờ hãy chạy chương trình và chắc chắn sẽ không có bất kỳ lõi nào cả.

### Hàm không xác tham số  ( Variadic Function )

- Tiếp theo chúng ta sẽ tìm hiểu về  một trường hợp khá hay được sử  dụng trong hàm.

- Nhin vào các ví dụ trước bạn thấy một hàm có thể  không có hoặc có nhiều tham số, nhưng nhiều tham số  là bao. Thật ra 
có nhiều tới đâu thì bạ vẫn phải xác định trước số  lượng tham số  và kiểu dữ liệu của các tham số  mà bạn muốn. Giả sử  một ví dụ trong thực tế  như sau, cánh dồng ngô của bạn đã tới mùa thu hoặc và nhiệm vụ của bạn là thu hoạch lấy chỗ  ngô này, dĩ liên là muốn thu hoặc thì bạn có cái gì đó để  đựng. Vì vây bạn đã lấy một chiếc túi ni lông loai nhỏ để  dựng, nhưng vì vường ngô của bạn có quá nhiều ngôn, chiếc túi ni lông hiện tại chỉ đựng được hai bắp ngô do đó bạn lại lấy túi loại ni lông to hơn, nhưng rồi lại đầy, tiếp đó bạn lại lấy bao tải nhưng rồi cũng lại như cũ, tới một mức độ nó lại đầy:

Vd:

```go
package main

import (
    "fmt"
)

func badge1(a string) {
    fmt.Printf("Bắp %s đã được bỏ vào\n", a)
}

func badge2(a string, b string) {
	fmt.Printf("Bắp %s đã được bỏ vào\n", a)
	fmt.Printf("Bắp %s đã được bỏ vào\n", b)
}

func badge5(a string, b string, c string, d string, e string) {
	fmt.Printf("Bắp %s đã được bỏ vào\n", a)
	fmt.Printf("Bắp %s đã được bỏ vào\n", b)
	fmt.Printf("Bắp %s đã được bỏ vào\n", b)
	fmt.Printf("Bắp %s đã được bỏ vào\n", d)
	fmt.Printf("Bắp %s đã được bỏ vào\n", e)
}


func main() {
	badge1("1")
	badge2("2", "3")
	badge5("4", "5", "6", "7", "8")
}
```

- Như bạn thấy cách làm này thật sự tồi khi cứ phải chuẩn bị các loại túi để  đựng từng bắp ngô lẻ tẻ như vậy nhỉ. Còn nếu như chúng ta sử  dụng một bao taỉ cực lớn, uk thì sẽ giải quyết được vấn đề  thu hoạch ngô mà hhông phải quan tâm quá nhiều việc túi sẽ đầy nhưng nếu thế  giả sử  dụng khi thu hoạch hết ngô rồi mà bạn vẫn còn dữ ra một khoảng túi thì sao, như vậy là quá lãng phí đúng ko ? Chưa kể  tới trong Go không có phép chúng ta lãng phí những thứ mà không sử  dụng. Vậy thì làm cách khác dó là dùng mảng, đung mảng có thể  lưu trữ được rất nhiều dữ liệu và bằng cách dùng mảng chúng ta chỉ cần có một hàm thôi là đủ, nhưng bạn hãy nhớ rằng mảng là một cấú trúc có kích thước xác định do đó nếu muốn bạn sẽ phải đếm tổng bắp ngô mà bạn có được bao nhiêu trước khi quyết định nhét ngô vào, mà đếm ngôn thì có khi tới sáng mai mất. Nếu mảng yêu cầu kích thước xác định trước không được thì cúng ta dùng slide sẽ đc, slice là một cấu trúc mảng động mà. Đúng nhưng bạn phải nhớ rằng slice là cấu trúc mảng động nhưng động trong một khoảng xác định thôi, tức là bạn vẫn phải xác định cái mức dộ, phạm vi của động đó. Cách này cũng không được, cách kia cũng không xong vậy thì còn cách nào khác nữa ?

- Còn, còn rất nhiều lên đừng vội nản, bởi vì phạm vi kiến thức của chúng ta đang tích lũy dần lên, và có rất nhiều khái niệm chúng ta chưa biết tới thôi. Một trong những khái niệm đó là Variadic Function. Vậy Variadic Function là gì ? Variadic Function là một cách sử  định nghĩa tham số  truyền vào trong hàm. Bạn sẽ chỉ cần quan tâm tới kiểu dữ liệu bạn muốn chuyền vào thôi còn lại số  lượng bạn không cần phải bận tâm, tự Go sẽ cân nhắc cho chúng ta. Cũng gióng như bạn có một cái túi thần kỳ, co dãn bao nhiêu tùy ý ấy. 

- Cú pháp của Variadic sử  dụng trong hàm được viết như sau:

```go
...T
// trong hàm

func (...T) return_Type {
    // logic
}
```

- Để  hiểu rõ hơn chúng ta hãy làm lại ví dụ trên với Variadic:

```go
package main

import (
    "fmt"
)

func badge(b ...string) {
    fmt.Printf("Các Bắp đựng là", b)
    for _, value := range b {
        fmt.Printf("Bắp %s đã được bỏ vào\n", value)
    }
}


func main() {
	badge("1")
	badge("2", "3")
	badge("4", "5", "6", "7", "8")
}
```

- Két quả:

```go
Các Bắp đựng là [1]
Bắp 1 đã được bỏ vào
Các Bắp đựng là [2 3]
Bắp 2 đã được bỏ vào
Bắp 3 đã được bỏ vào
Các Bắp đựng là [4 5 6 7 8]
Bắp 4 đã được bỏ vào
Bắp 5 đã được bỏ vào
Bắp 6 đã được bỏ vào
Bắp 7 đã được bỏ vào
Bắp 8 đã được bỏ vào
```

- Thật tuyệt vời phải không, chúng ta không cần phải xác định rõ số  lượng giá trị, biến cần chuyền vào gì hết. Ngoài ra hãy để  ý, tôi đã thêm một hàm hiển thị biến truyền vào ở dòng đầu tiên và bạn thấy nó chúng quen thuộc, không. Vâng nó có thể  coi là kiểu cấu trúc mảng, hay chính xác là slice, khi bạn truyền các giá trị vào , Go sẽ xác định các giá trị đó và đống gói nó vào trong môti slide do đó chúng ta có thể  dùng được các hàm thao với kiểu mảng, slice và không cần chúng ta phải xác định rõ kích thước của mảng. 

#### Ứng dụng:

- Trong thực tế  Variadic dược ứng dụng khá nhiều, không đâu xa bạn có nhớ hàm `append` mả chúng ta sử  dụng để  thêm một hay nhiều phần tử  mới vài slice không, hãy nhìn cú pháp nó được định nghĩa bạn sẽ thấy:

```go
func append(slice []Type, elems ...Type) []Type  
```

- Đó là một ví dụ về  sử  dụng Variadic. Hay một hàm mà trước này chúng ta vẫn thường sử  dụng để  hiển thị ra màn hình. Vâng chắc hẳn lúc này bạn cũng nhận ra hàm mà tôi muốn nói, đó chính là hàm `Println` trong package `fmt` và nó được định nghĩa như sau `func Println(a ...interface{}) (n int, err error)`. Như bạn thấy nó nhận chỉ duy nhất một tham số  là Variadic
interface, có thể  bạn lúc này đang thắc mắc interface là gì nhưng giờ chưa phải lúc, sớm thôi chúng ta sẽ động tới nó.

- Ở trên là một trong nhiều ví dụ tôi muốn chỉ ra để  bạn thấy rằng Variadic được ứng dụng rộng rãi thế  nào trong thực tế. Còn giờ sẽ có một số thao tác mà tôi muốn bạn hiểu rõ hơn khi làm việc với variadic.

#### tham số  mảng, slice

- Điều trước tiên chúng ta có thể  nhấn mạnh là variadic cũng là mảng, slice vậy ngược lại bạn cũng có thể  dùng mảng và slice như một tham số  truyền vào như một variadic. Đúng vạy hãy xét ví dụ sau:

```go
package main

import (  
    "fmt"
)

func count(number int, numbers ...int) int {  
	number_count := 0
		
    for _, v := range numbers {
        if v == number {
            number_count++
        }
    }
    return number_count
}

func main() {  
	numbers := []int{1,2,3,4,5,6,5,4,3,2,1}
	number := 5
	total := count(number, numbers...)
	
	fmt.Println("Tổng giá trị ̀5 xuất hiện trong mảng là ", total)

    total = count(6, numbers...)
	fmt.Println("Tổng giá trị ̀6 xuất hiện trong mảng là ", total)

	total = count(0, numbers...)
	fmt.Println("Tổng giá trị ̀0 xuất hiện trong mảng là ", total)
}
```

- Ở ví dụ này chúng sẽ sẽ định nghĩa một hàm có chức năng tính tổng số  lần xuất hiện của một giá trị xác định trong một mảng. Hàm này chúng ta sẽ gọi là hàm `count` và nhận đâì vào hai tham số  trong đó tham số  đàu tiên là giá trị ta muốn tìm kiếm và tham số  thứ hai là variadic thuộc kiểu số  nguyên. Tuy nhiên ở trong phần `main` chúng ta sẽ không truyền vào hàm `count` với nhiều tham số  mà thay vào đó sẽ truyền vào với một slice được xác định trước. Lưu ý là ở đây khi gọi hàm cùng các tham số  tôi đã  dùng một cách khá lạ `count(6, numbers...)`, ở đây bởi vì tham số  chỉ châp nhận variadic do đó bạn không thể  truyền mảng một cách thông thường được, do đó buộc phải phải chuyển đổi slice sang từng phần tử  bằng `slice...` ( gọi là chuyển đổi kiểu nhưng thực tế  là nó sẽ truyền trực tiếp slice vào mà cần cần kiểm tra để  đóng gói các phần tử  thành slice).

- Hãy thử  chạy xem:

```go
Tổng giá trị ̀5 xuất hiện trong mảng là  2
Tổng giá trị ̀6 xuất hiện trong mảng là  1
Tổng giá trị ̀0 xuất hiện trong mảng là  0
```

- Chính xác.

### Tổng kết

- Như vậy xuyên suốt chương này chúng ta đã tìm hiểu về  hàm rồi, đồng thời chúng ta cũng học về  các cách sử  dụng hàm với một hay nhièu giá trị trả về  cùng như việc sử  dụng định danh trống để  bỏ qua giá trị mà chúng ta không sử  dụng. Ngoài ta chúng ta cúng biết được về  một khái niệm mới đó là Variadic và biết được ứng dụng của nó trong thực tế.
Trong chương Tiếp theo chúng ta sẽ tìm hiểu về  package để  biêt rõ mối quan hệ giữa hàm với package như thế  nào. 