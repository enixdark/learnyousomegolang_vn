## Map

- Khi trên tay bạn đang có một cuốn sách mới, một cuốn sách mà bạn chưa hề  được nghe qua, việc đầu tiên bạn tìm đến đó là gì ? 
Giở phần giới thiệu ra ? Lật ngược cuốn sách xem có phần tóm sơ lược ? Gọi điện thoại cho người thân hay lập tức lên mạng tra thông tin ?
Oh không đó không phải là điều tôi muốn nói, cơ bản là bạn sẽ lật tới phần mục lục và xem các nội dung mà cuốn sách đó sẽ có. Đúng vậy điều tôi muốn nhấn mạnh đó 
là bạn sẽ dựa vào thông tin các trang các phần mục lục và tìm tới chương, phần mà bạn muốn đọc trong sách. Hình như hơi lạc chủ đề  rôi, vậy thì điều này 
liên quan gì tới cấu trúc `Map` mà chúng ta sẽ tìm hiểu. Hm có, có liên quan, hãy hình dùng mỗi phần ,chương ở mục lục như một khóa trong sách và hướng tới một 
giá trị là số  trang mà bạn sẽ tìm tới Hoặc nếu đó là một cuốn từ điển, thì mỗi chữ cái ( trong bảng chữ cái latin ) đều ứng với ví trị trong sách mà bạn sẽ tìm tới tra được từ. Những thông tin về  chương, chữ cái như vậy được gọi là khóa và mỗi số trang hướng tới được coi là một giá trị, khi đó ta sẽ có được từng cặp khóa-giá trị và đó chính là cấu trúc của `Map` mà chúng ta sẽ học trong chương này.

- `Map` được định nghĩa là một tập hợp các cặp khóa-giá trị không theo thứ tự ( không cần phải sắp xếp thứ tự các khóa ), trong đó mỗi khóa là duy nhất và dựa vào khóa ta có thể  tìm được một giá trị tương ứng ( các giá trị không yêu cầu phải là duy nhất, hai khóa có thế  có cùng một giá trị vẫn thỏa mãn ).

- Cú pháp của `map` được viết như sau:

```go
var name map[KeyType]ValueType
```

- Trong đó:
  + `name`: là tên của biến map mà bạn muốn đặt
  + `KeyType`: là kiểu dữ liệu của khóa mà bạn muốn sử  dụng
  + `ValueType`: là kiểu dữ liệu của giá trị mà bạn muốn sử  dụng

Vd:

```go
var book map[string]int
```

- Ở đây tôi định nghĩa một biến map có tên là 'book' với khóa thuộc kiểu chuỗi và giá trị liên kết với khóa thuộc kiểu số nguyên.

### Khởi tạo map

- Tiếp theo chúng ta sẽ tìm hiểu cách để  tạo ra một map. Ở các phần trước chúng ta đã học, hầu hết khi khai báo các kiểu dữ liệu, nếu như bạn không xác định rõ giá trị khởi tạo thì mặc định Go sẽ tạo ra gía trị ứng với kiểu dữ liệu đó. Nhưng với map chúng ta không thể  dùng cách này được. Nó yêu cầu một là bạn phải khởi tạo trước giá trị hoặc là bạn khởi tạo nhưng sử  dụng tới hàm `make`:

```go
package main
import (
	"fmt"
)

func main() {
	var init map[string]int = make(map[string]int)
	data := map[string]int{
		"A": 1,
		"B": 2,
		"C", 3
	} 
	
	fmt.Println("Map được khơi tạo cùng với giá trị mặc định là:", init)
	fmt.Println("Map được khơi tạo cùng với giá trị xác định trước là:", data)
}
```
- Chạy đoạn mã trên bạn sẽ nhận được:

```go
Map được khơi tạo cùng với giá trị mặc định là: map[]
Map được khơi tạo cùng với giá trị xác định trước là: map[A:1 B:2 C:3]
```

- Lưu ý: bạn hoàn toàn có thể  sử  dụng cách khai báo thông thường `var init map[string]int`, Go cũng sẽ tạo ra một map rỗng, tuy nhiên bằng cách này biến của bạn thực chất chỉ đang trỏ nil thay vì một map rỗng do đó không thể  sử  dụng được các chức năng của map.

### Thao tác với map

#### Lấy, thêm , chỉnh sửa , xóa cặp khóa-giá trị

- Để  thực hiện việc thêm một dữ liệu mới hay thay đổi khóa dữ liệu rất đơn giản, bạn chỉ cần dựa theo cú pháp sau:

```go
name[Key] = Value
```

- Khá đơn giản phải không, bạn chỉ cần xác định rõ một từ khóa và giá trị bạn muốn thêm vào bằng cú pháp trên là được. Trong trường hợp bạn muốn chỉnh sửa giá trị của một khóa đã có trong `map` thì cũng tương tự như việc bạn tạo mới một dữ liệu thôi, bạn chỉ cần gán lại giá trị của khóa đó là được.

```go
package main
import (
	"fmt"
)

func main() {
	var init map[string]int
	
    fmt.Println("Map được khơi tạo cùng với giá trị mặc định là:", init)
    
    init["A"] = 1
    fmt.Println("Map sau khi được thêm dữ liệu:", init)
    
    init["A"] = 10
    fmt.Println("Map sau khi sửa đổi dữ liệu:", init)
}
```

- Kết quả:

```go
Map được khơi tạo cùng với giá trị mặc định là: map[]
Map sau khi được thêm dữ liệu: map[A:1]
Map sau khi sửa đổi dữ liệu: map[A:10]
```

- Bạn thấy đó không có gì phức tạp cả thậm chí còn dễ  dàng hơn khi so sánh với mảng ( array ) và slice. Với thao tác thêm và chỉnh sửa không có khó khăn gì hết nhưng đối với thao tác lấy dữ liệu từ một khóa sẽ có một chút phức tạp. Đầu tiên nếu như khóa đó tồn tại trong `map` rồi thì bạn chỉ cần lấy nó ra là xong, nhưng giả sử nếu khóa chúng ta muốn tìm không tồn tại thì sao ?

```go
package main
import (
	"fmt"
)

func main() {
	data := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

    fmt.Println("A -> ", data["A"])
	fmt.Println("D ->", data["D"])
}
```

Kết quả:

```go
A -> 1
D -> 0
```

- Nhìn xem ở ví dụ trên, mặc dù khóa D không tồn tại nhưng khi ta lấy dữ liệu từ khóa đó ra Go sẽ tự động gán nó với giá trị khởi tạo ứng với kiểu dữ liệu đó. Tệ nhỉ, nếu như trường hợp chúng ta cũng có một khóa khác với giá trị bằng với giá trị khởi tạo thì thật khó phân biệt đâu là khóa đó có tồn tại hay không trong 'map'. Đúng vậy, do đó để  lấy thông tin của một khóa trong 'map' chúng ta cần chỉnh sửa lại một chút với cú pháp sau:

```go
value, ok := map[Key]
```

- Khi lấy một thông tin từ một 'map', ngoài giá trị liên kết với khóa ra, nó sẽ còn đưa cho chúng ta biết thông tin về  một khóa có tồn tại hay không trong map. Nếu thì giá trị sẽ trả về  là `true` còn ngược lại là `false` và từ đó bạn có thể  sử  dụng giá trị này để  đưa ra kết luận.

Vd:

```go
package main
import (
	"fmt"
)

func main() {
	data := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	_, ok := data["D"]
	if ok != true {
        fmt.Println("Khóa D không tồn tại trong map")
	}
    
}
```

Kết quả:

```go
Khóa D không tồn tại trong map
```

- Cuối cùng khi bạn không muốn sử  dụng khóa đó nữa bạn có thể  xóa bỏ nó bằng hàm `delete(map, key)`. Hàm này sẽ nhận đầu vào map và khóa mà bạn muốn xóa trong map đó. Trong trường hợp nếu khóa đó không tồn tại thì nó sẽ bỏ qua mà không đưa ta bất kỳ cảnh báo, lỗi nào cả.

```go
package main
import (
	"fmt"
)

func main() {
	data := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	fmt.Println("Dữ liệu trước khi xóa", data)
	
	delete(data, "A")
	delete(data, "B")
	delete(data, "A")

	fmt.Println("Dữ liệu sau khi xóa", data)

}
```

Kết quả:

```go
Dữ liệu trước khi xóa map[A:1 B:2 C:3]
Dữ liệu sau khi xóa map[C:3]
```

### Map là kiểu dư liệu tham chiếu

- Trong phàn trước khi chúng ta học về  mảng và `slice` chúng ta đã được biết mảng là kiểu dứ liệu giá trị trong khi `slice` lại là kiểu tham chiếu. Giống với `slice`, `map` cũng là kiểu dữ liệu tham chiếu:

- Xét ví dụ sau:

```go
package main
import (
	"fmt"
)

func main() {
	var data map[string]int
	origin := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	} 

	data = origin 
	fmt.Println("giá trị của origin là:", origin)
	
	data["D"] = 4
	fmt.Println("giá trị của origin sau khi chỉnh sửa tử  biến data là:", origin)
}
```

- Kết quả:

```go
giá trị của origin là: map[A:1 B:2 C:3]
giá trị của origin sau khi chỉnh sửa tử  biến data là: map[B:2 C:3 D:4 A:1]
```

- Như bạn thấy thây đổi dữ liệu ở biến data kéo theo biến origin cũng thay đổi theo.

#### kích thước và lặp

- Hãy nhìn các đoạn thông tin được in ra trong map , một cặp khóa giá trị được đặt trong `[]`, hình dung một cách đơn giản bạn có thể  coi mỗi cặp khoa-giá trị như một phần tử  của mảng một chiều, và vì là hình dung nó như một mảng do đó chúng ta hoàn toàn có sẽ sử  dụng được các chức năng như tính kích thước của một mảng và sử dụng range để  duyệt mảng đó:

```go
package main
import (
	"fmt"
)

func main() {
	data := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	} 

	fmt.Println("Độ dài của map là:", len(data))
	fmt.Println("Các giá trị có trong map là")
	for key, value := range data {
        fmt.Println(key, " -> ", value)

    }
}
```

- Kết quả:

```go
Độ dài của map là: 3
Các giá trị có trong map là
A  ->  1
B  ->  2
C  ->  3
```

- Lưu ý: như ở đầu chương tôi đã nói về  việc map là một cấu trúc dữ liệu không dựa trên thứ tự từ khóa do đó 
có thể  kết quả ở trên sẽ khác thứ tự so với kết quả chạy ở chương trình của bạn nhưng tổng két quả in ra các từ khóa và giá trị vẫn vậy.

### Tổng Kết

- Bạn đã làm quen với cấu trức dữ liệu map, biết được khái niệm và thao tác với map rồi. mặc dù còn rất nhiều các thao tác khác vơi map nữa tuy vậy chung ta không thể  bao quát hết được do đó tôi chỉ giới thiệu cơ bản những thao tác mà chúng ta hay dùng thôi. 