## Cấu trúc Mảng ( Arrays )

- Bây giờ chúng ta sẽ đi vào tìm hiểu một trong những cấu trúc được sủ dụng nhiều nhất trong các thuật toán đó là cấu trúc mảng. Khái niệm cơ bản của mảng được định là một tâp hợp các phần có cùng kiểu dữ liệu và được xác định rõ kích thước. Trong đó các phần tử  trong mảng được lưu trữ một cách tuần tự và có thể  truy xuất được các phần này dựa trên ví trí của các phần tử  đó được đặt trong mảng. Vd rõ nét nhất của cấu trúc mảng mà ta đã gặp đó là kiểu chuỗi. Về  bản chất kiểu dữ liệu chuỗi chỉ là một tập hợp các ký tự thuộc kiểu byte và có độ dài đã định trước ( như "Name" là một chuồi hay một mảng gồm 4 ký tự ).

- Lưu ý: với mảng điều kiện bắt buộc là chúng phải thuộc cùng một kiểu dữ liệu.

- Tiếp theo, chúng ta sẽ nói về  các phần tử  của mảng. Hãy nhìn hình sau:

<br>
<p align="center">
  <img src="arrays.jpg"/>
</p>
<br>     

- Như bạn thấy ở đây chúng ta có một mảng gồm 5 phần tử  thuộc kiểu số  nguyên, các phần tử  này được đặt lần lượt cạnh nhau. Khi đó ta có thể  nói mảng này thuộc kiểu số  nguyên và có độ dài xác định là 5. Chúng ta hoàn toàn có thể  truy cập vào các phần tử  của mảng này thông qua ví trị chỉ mục ( index ) của chúng. Tuy nhiên có một điều cần lưu ý đó là ví trị bắt đầu của phần tử  đầu tiên luôn quy ước là 0 và ví trị kết thúc của phần tử  cuối cùng luôn luộn bằng với độ dài của mảng trừ đi 1 đơn vị. Hãy ghi nhớ điều này để  tránh nhầm lẫn gây ra sai xót khi thao tác với mảng.

- Để  hiểu rõ hơn chúng ta sẽ làm một ví dụ cụ thể  về  mảng, nhưng trước khi viết chúng ta sẽ xem cách khai báo một mảng trong Go như thế  nào đã:

```go
var name[number]Type 
```

- Ở đây `name` sẽ là tên của biến dược đặt cho mảng đó, `number` là kích thước mà chúng ta sẽ khởi tạo mảng và `Type` là kiểu dữ liệu của mảng.

```go
package main

import "fmt"

func main() {
	var x [3]int 
	var y [3]float32
	var z [3]string
	var e [3]bool
	
	fmt.Println("Giá trị khởi tao ban đầu của mảng số  nguyên là:", x)
	fmt.Println("Giá trị khởi tao ban đầu của mảng số  thực là:", y)
	fmt.Println("Giá trị khởi tao ban đầu của mảng thuộc kiểu chuỗi là:", z)
	fmt.Println("Giá trị khởi tao ban đầu của mảng boolean là:", e)
}
```

- Kết quả:

```
Giá trị khởi tao ban đầu của mảng số  nguyên là: [0 0 0]
Giá trị khởi tao ban đầu của mảng số  thực là: [0 0 0]
Giá trị khởi tao ban đầu của mảng thuộc kiểu chuỗi là: [  ]
Giá trị khởi tao ban đầu của mảng boolean là: [false false false]
```

- Ở đây chúng ta sẽ khởi tạo một vài mảng thuộc kiểu số  nguyên, thực và kiểu chuỗi có kích xác dịnh là 3 và hiẻn thị giá trị của chúng lên màn hình. Bởi vì chúng ta không xác dịnh giá trị khởi tạo cho các phần tử  do đó chúng sẽ sử  dụng giá trị mặc định và được đặt trong ngoặc `[]` như một cách xác dịnh rằng đây là mảng. 

### Chỉ  mục ( Index )

- Chúng ta thường sử  dụng biến để  lưu trữ một thông tin dữ liệu nào đó để  sử  dụng trong một số    nhiệm vụ cụ thể  ví dụ như lưu trữ một biến tổng để  tính tổng một dãy số  nguyên hay lưu trữ một số ngẫu nhiên bất kỳ và dùng nó để  xác định xem nó có nằm trong số  những số  may mắn hay không. Tất cả những trường hợp trên chỉ đơn giản là lưu trữ một giá trị mà thôi. Trong trường hợp chúng ta muốn lưu một đoạn dữ liệu như muốn lưu trữ tổng tính được qua các bước thì sao. Vâng chúng ta có thể  dùng tới kiểu mảng array, khi đó mỗi phàn tử  trong mảng khi đó sé mang các giá trị qua từng bước tính tổng. tuy nhiên để  truy cập vào từng phần tử  này thay và thôi nó ta cần phải phải dùng tới một khái niệm trong mảng gọi là chỉ mục ( index ). chỉ mục sẽ cho ta biết vị trí phần tử  của mảng đó để  truy cập và trích xuất cũng như thay đổi thông tin của một phần tử.

```go
// lệnh này sẽ cho phép ta truy cập để  lấy thông tin của phàn tử 
array[index]

// lệnh này sẽ cho phép thay đổi giá trị của phần tử
array[index] = value 
```

- Xét ví dụ:

```go
package main

import "fmt"

func main() {

	var stores [10]int

	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
		stores[i-1] = sum 
	}

	fmt.Println("kết quả lưu trữ tổng qua các lần gọi là:", stores);
}
```

- Kết quả sau khi chạy:

```go
kết quả lưu trữ tổng qua các lần gọi là: [1 3 6 10 15 21 28 36 45 55]
```

- Như bạn thấy ở đấy chúng ta se sử  dụng một một mảng gồm 10 phần tử  để  lưu trữ cấc kết quả tính tổng tích lũy được từ 1 tới 10. Chúng ta sẽ sử  dụng một vòng lặp để  chạy từ  1 tới 10 và dùng một biến để  lưu trữ tổng tính được. Bởi vì quy ước của mảng luôn bắt đầu từ vị trí 0 do đó để  gán lại giá trị cho phần tử  chúng ta sẽ giảm đi một đơn vị ở biến `i`. Cuối cùng chúng ta được kết quả như trên.

### Khởi tạo mảng cùng giá trị xác định

- Ở các ví dụ trước khi khai báo một mảng cùng kiểu dữ liệu, chương trình sẽ khởi tạo các giá trị mặc định tương tương với kiểu dữ liệu đó. Tuy nhiên mỗi lần khởi tạo xong, nếu như chúng ta muốn có được một mảng chứa các giá trị như mong muốn thì phải thực hiện thao tác truy cập vào từng phần tử  để  thay đổi giá trị . Cách làm này rất mất thời gian, thay vào đó chúng ta có thể  khởi tạo mang cũng các giá trị mà chúng ta mong muốn trong  mảng đó theo cú pháp sau:

```go
var name [number]int = [number]int{value_1, value_2, ..., value_number }

hay

name := [number]int{value_1, value_2, ..., value_number }
```

Vd: 

```go
package main

import "fmt"

func main() {

	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(nums);
}
```

Kết quả:

```go
[1 2 3 4 5 6 7 8 9 10]
```

- Lưu ý: Không nhất thiết bạn phải xác định rõ giá trị của các phần tử  trong một mảng có kích thước N, thay vào đó bạn có thể  dựa khởi tạo số  lượng phần tử  mà muốn khởi tạo giá trị cũng đc Vd:

```go
package main

import "fmt"

func main() {

	nums := [10]int{1, 2}

	fmt.Println(nums);
}
```

Kết quả:

```go
[1 2 0 0 0 0 0 0 0 0]
```

- Như bạn thấy, chúng ta có thể  cung cấp các giá trị cho một vài phần tử  cũng đc, còn những phần tử còn lạ sẽ tự động nhận giá trị mặc định.

- Đối với trường hợp khởi tạo giá trị mảng, đôi khi việc phải xác định kiểu kích thước sau đó là đếm số lượng phần tử  mảng mình đã điền vào trong cặp ngoặc `{}` cũng khiến bạn cảm thấy chản nản, vì vậy để  đơn giản bạn hoàn toàn có thể  không cần phải xác kích thước mà chỉ cần điền thẳng các giá trị càn khởi tạo vào. Để  thực hiện đièu này bạn chỉ cần thay thế  số  kích thước mà bạn muốn xác định khi khởi tạo mảng bằng `...` hay để  trống:

   ```go
package main

import "fmt"

func main() {

	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	names := []string{"Hello", "Go", "!"}

	fmt.Println(nums);
	fmt.Println(names);
}
```

Kết quả:

```go
[1 2 3 4 5 6 7 8 9 10]
["Hello" "Go" "!"]
```

- Thật tuyệt phải không, bạn khong cần phải xác định trước kích thước mảng cần tạo làm gì hết, chỉ cần cung cấp các giá trị được khởi tạo thôi, còn lại công cụ bien dịch sẽ tự động phân tích và đưa ra kích thước mảng cho bạn.

### Thao tác với mảng

- Tiếp sau đây chúng ta sẽ đi sâu hơn về  mảng và thực hiện một số  chức năng trong mảng:

##### Kích thước mảng 

- Trước tiên tôi phải nhấn mạnh một điều là với mảng, kích thước là cố  định do đó bản không thể  mở rộng hay thu hẹp mảng được và bạn cũng không thể  gán mảng này với mảng kia khi mà chúng khác nhau về  mặt kích thước được Vd:

```go
a := [3]int{1,2,3}
b := [4]int{1,2,3,4}

b = a // lỗi
```

- Tuy nhiên trong trường hợp bạn không xác định rõ kích thước thì lại được:

```go
a := []int{1,2,3}
b := []int{1,2,3,4}

b = a // ok
```

#### Mảng là kiểu giá trị ( value types )

- Nếu trong một số  ngôn ngữ lập trình như Java, Python, etc. Mảng thường được nhìn nhận là một kiểu dữ liệu tham chiếu ( reference types ). nếu như bạn gán một biến mới gọi là `a` tới một biến mảng đã tạo trước đó gọi là `b` thì trình biên dịch sẽ ngầm định biến a này trỏ/tham chiếu tới biến b kia, và nếu bạn thay đổi giá trị phần tử  bất kỳ trong biến a hay b thì giá trị của cả hai cũng thay đổi như nhau. Nhưng trong Go thì mảng chỉ là một kiểu giá trị, và nếu bạn lấy biến a gán tới biến b, thì nó chỉ đơn giản là biến a sẽ nhận một giá trị được sao chép từ biến b. bạn thay đổi bất cứ phần tử  nào trong biến a cũng không ảnh hưởng tới biến b và ngược lại:

```go
package main

import "fmt"

func main() {

	b := [5]int{1,2,3,4,5}
	a := b

	a[1] = 0

	fmt.Println(a)
}
```

- Kêt quả sẽ là:

```go
[1 0 3 4 5]
```

#### for, range, len

- Với kiểu mảng, khi bạn muốn truy cập vào tất cả các phần tử  có trong mảng đó bạn có thể  sử  dụng vòng lặp for: 

Vd:

```go
package main

import "fmt"

func main() {

	nums := [5]int{1,2,3,4,5}
	
	for i := 0; i < 5; i++ {
        fmt.Println(i)
	}
}
```

Kết quả:

```go
1
2
3
4
5
```

- Bởi vì khi khởi tạo một mảng chúng ta luôn xác định trước kích thước cần khởi tạo do đó không khó gì để  xác định điều kiện dừng và chỉ mục có thể  truy cập được của mảng. Nhưng nếu như bạn khởi tạo một mảng cũng với kích thước không xác định thì sao ? Nếu như mảng đó có quá nhiều phần tử  mà bạn không thể  đếm được ? 

- Lúc này chúng ta sẽ cần tới một công cụ hỗ  trợ mảng cho phép chúng ta duyệt các phần tử  của mảng mà không cần quan tâm tới kích thước của mảng đó là toán tử  `range`.  

```go
package main

import "fmt"

func main() {

	nums := []int{1,2,3,4,5}
	
	for index, value := range nums {
        fmt.Println(index, value)
	}
}
```

- Đây là một trong những cách để  duyệt các phần tử  trong một mảng tối ưu nhất, hơn nữa `range` cũng cung cấp cho chúng ta cả ví trị chỉ mục khi duyệt mảng. trong trường hợp bạn không quan tâm tới một trong hai giá trị trên bạn có thể  dùng ký tự `_` để  bỏ qua.

- Nếu như bạn không muốn sử  dụng `range` để  duyệt mảng mà vẫn muốn sử  dụng cách thông thường của vòng lặp for. Bạn có thể  dùng hàm `len` để  chỉ ra kích thước của mảng:

```go
package main

import "fmt"

func main() {

	nums := []int{1,2,3,4,5}
	
	for i := 0; i < len(nums); i++ {
        fmt.Println(i)
	}
}
```

#### Mảng đa chiều 

- Nếu như bạn là một nhà toán học, và bạn đang xử  lí một bài toán ma trận tuy nhiên với cách khai báo mảng như hiện tại khiến bạn gặp khó khăn trong việc lưu trữ và tính toán, đưng lo lắng vì mảng cũng  cung cấp cho bạn cách xử  lí dữ liệu dạng ma trận. Hãy để ý bản chất của mảng chỉ đơn thuần là một tập hợp các phần tử  của cấu trúc cùng một kiẻu dữ liệu mà thôi, vậy nếu bạn coi mảng cũng là một kiểu dữ liệu thì tại sao lại không thể  lồng các mảng vào với nhau được, hoặc thậm chí mảng của mảng của mảng. Những mảng như vậy chúng da sẽ gọi chung là mảng đa chiều.
Để  khai báo chúng cũng rất đơn giản ứng với số chiều bạn muốn tạo ra bạn chỉ càn chèn thêm `[]` vào trước kiểu  dữ liệu thôi. 

Vd: 

```go
var two_dimension [3][3]int // mảng hai chiều

var three_dimension [3][3]int // mảng ba chiều
```

- Áp dụng với ma trận của nhà toán học trên chúng ta sẽ khởi tạo như sau:

```go
package main

import "fmt"

func main() {

	maxtrix := [][]int{ { 1,1,1 }, { 1,1,1 } } 

    fmt.Println(maxtrix)
}
```

Kết quả:

```go
[[1 1 1] [1 1 1]]
```

- Đơn giản phải không, bạn cũng có thể  thêm nhiều chiều hơn nữa cho mảng. Nhưng trong thực tế  thường thì chúng ta chỉ xét tới mảng 3 chiều là hết, néu hơn sẽ rất phức tạp 
và khó nhìn, khi đó chúng ta lên sử  dụng một số  3rd package.


### Tổng kết

- Trong chương này chúng ta đã tìm hiểu cấu truc mảng và xét một vài ví dụ khi thao tác với mảng. Mặc dù chưa thể  bao quát hết được các hàm hoạt động với mảng 
tuy nhiên đừng hấp tấp với vì trong chương tới chúng ta sẽ đi sâu hơn vào việc thao tác với mảng cùng với đó là một khái niệm mới về  cấu trúc slice.

