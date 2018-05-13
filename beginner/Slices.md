## Slice

- Nếu như chưng trước chúng ta đã tìm hiểu về  kiểu mảng ( arrays ) và thực hiện một số  tính năng có trong mảng. 
Trong chương này chúng ta sẽ làm quen với một khái niệm mới gọi là Slice. Thực chất Slice không phải là một kiểu dữ liệu đặc biệt gì, 
chúng chỉ đơn giản là mảng thôi ( hay chính xác là tham chiếu tới mảng ) nhưng thay vì là những mảng đơn thuần, nó được các những người phát triển Go cung cấp cho 
rất nhiều hàm, tiện ích.

- Trước tiên chúng ta sẽ xem cú pháp để  khai báo một slice:

```go
var name []Type

hay 

var name []Type = []Type{}

hay

name := []Type{}

hay 

name := array_variable[start:end]
```

- Oh! Chờ chút có điều gì đó không đúng lắm ở đây. Đúng vậy có một sự thật là trong chương trước tôi có đề  cập tới việc khởi tạo một mảng mà không yêu cầu xác định kích thước. Nhưng tôi quên nói với bạn một điều , đó là cách mà chúng ta sử  dụng `[]` khi khai báo thực chất đó là slice. Nhưng nhìn xem chúng ta vẫn thao tác với slice như với một mảng bình thường đúng ko.

- Đúng vậy tuy nhiên có một chút khác biệt ở đây. Trước tiên Slice về bản chất vẫn là mảng nhưng nó là một kiểu tham chiếu tới mảng. Bạn có thể nhìn sơ đồ sau để  dễ  hình dung hơn:

<br>
<p align="center">
  <img src="slices1.jpg"/>
</p>
<br>  

- Trong đó:
  + Pointer ( Ptr ): con trỏ trỏ tới mảng
  + length ( len ): kích thước, độ dài của slice 
  + capacity  ( cap ): kích thước, độ dài tối đa mà slice có thể  chứa.

- Có vẻ hơi khó hình dung, Do đó chúng ta sẽ làm một số  ví dụ để  có thể  sáng tỏ hơn:

```go
package main

import (
    "fmt"
)

func main() {
	days := [7]string{"Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy", "Chủ Nhật"}

	school := days[0:5]
	off := days[5:]

	fmt.Println("Học sinh phải cắp sách tới trường vào ngày:", school)
	
	fmt.Println("Học sinh được nghỉ học vào ngày", off)
}
```

- Trong ví dụ này chúng ta sẽ khai báo một mảng gồm 7 phần tử  để  chỉ các ngày trong tuần. Tiếp đó chúng ta sẽ tạo ra 2 slice, trong đó một slice sẽ được sử  dụng để  chưa thông tin về  những ngày đi học, ở đây ta gọi là `school` và một slice để  chứa thông tin về  ngày nghỉ học `off`. Cả hai slice này đều  được tạo ra từ mảng gốc `days`. Như bạn thấy chúng ta sẽ dùng `[start:end]` để  tạo ra slice từ mảng góc. Tôi sẽ giải thích thêm chút, trong cấu trúc `[start:end]` này, chúng ta sẽ đùng để  lấy thông tin cả phần từ từ ví trí chỉ mục bắt đầu tức (start) cho tới ví trí chỉ mục mà chúng ta muốn kết thúc `end`, ngoài ra nếu như bạn chỉ xác định một vế  là `start` hay `end` thì vế  còn lại sẽ tương ứng với chỉ mục cuối cùng hay điẻm bắt đầu của mảng gốc. Trong trường hợp không có cả `start` lẫn `end` , khi đó `[:]` sẽ ngầm hiều  là bạn muốn tạo ra một slice bằng với mảng gốc.


- Nào hãy chạy chương trình và xem kết quả nhận được

```go
Học sinh phải cắp sách tới trường vào ngày: [Thứ Hai Thứ Ba Thứ Tư Thứ Năm Thứ Sáu]
Học sinh được nghỉ học vào ngày [Thứ Bảy Chủ Nhật]
```

- Đi học thật là tuyệt vời mà nghỉ học còn tuyệt vời hơn. 

### Kiểu dữ liệu tham chiếu ( reference type )

- Chờ một chút ở phần trên tôi có nhắc tới slice chỉ là kiểu tham chiếu tới mảng thôi đúng không. Hm vậy chắc là nó giống với mảng trong một số  ngôn ngữ khác ? Thử làm một ví dụ chỉnh sửa dữ liệu từ slice xem sao:

```go
package main

import (
    "fmt"
)

func main() {
	days := [7]string{"Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy", "Chủ Nhật"}

	slice_days := days[0:5]
    
    slice_days[0] = "2"

	fmt.Println("Danh Sách ngày trong tuần:", days)
	
}
``` 

- Hãy chạy chương trình và xem kết quả mà ta nhận được:

```go
Danh Sách ngày trong tuần: [2 Thứ Ba Thứ Tư Thứ Năm Thứ Sáu Thứ Bảy Chủ Nhật]
```

- Oh! Nhìn xem, dữ liệu đã bị thay đổi đổi mặc dù chúng ta không chỉnh sửa gì trong mảng `days` cả. Nếu như lúc trước chúng ta thao tác tương tự điều này với mảng thì không ảnh hưởng gì tới mảng gốc cả bởi vì mảng là kiểu dữ liệu giá trị, nhưng slice thì ngược lại, nó là kiểu dữ liệu tham chiếu do đó khi thay đổi bất kỳ phần tử  nào trong slice thì nó sẽ ảnh hưởng tới mảng gốc mà tạo ra nó. Đây cũng chính là ý nghĩa của Pointer trong sơ đồ ban đâu tôi đã đưa ra.

### Kích thước ( length ) và khả năng chứa ( capacity )

- Qua vd trên bạn chắc hẳn của hiểu được phần nào về  con trỏ Pointer trong sơ đồ  ở đâu chương rồi nhỉ, vậy còn length và capacity thì sao ?

- Ah length ( len ) không có ý nghĩ đặc biệt gì lắm, nó chính là kích thước của slice đó và gọi thông qua hàm `len`. Nhưng ở đây chúng ta chú ý tới capacity, chúng ta sẽ sử  dụng hàm `cap` để  gọi
, nhưng cap thực sự nó là gì ? Để  giả thích trước tiên bạn hãy nhìn sơ đồ  sau:

<br>
<p align="center">
  <img src="slices2.jpg"/>
</p>
<br>   

- Như bạn tháy ở so đồ  này chúng ta có một mảng gồm 6 phần tử  có giá trị từ  10 tới 60 , và một slice trỏ tới 3 phần tử  20, 30, 40 của mảng ( bắt đầu  từ chỉ mục 1 tới chỉ mục 4). Độ dài tương ứng của slice sẽ là 3 nhưng thông số  về  sức chứa của nó lại là 5. Nếu nhìn vào sơ đồ  dễ  dàng có thể  tháy mặc dù slice chỉ trỏ tới 3 phần tử  của mảng nhưng chưa phải là phần tử  cuối cùng, tại ví trí kết thúc của slide, vẫn còn 2 phần tử  nữa. Hay nói cách khác là thông số  về  sức chứa của slice sẽ được tính tại vị trí chỉ mục bắt đầu của slice đó hay của mảng gốc tới điểm, chỉ mục chưa phần từ cuối cùng của mảng gốc. Mục đích của thông số  này là để  gợi ý cho húng ta biết được rằng slice này vẫn có thể  mở rộng hay trỏ tới phạm vi rộng hơn phạm vi hiện tại mà nó đang trỏ trong mảng gốc. Hay nói một cách khác là slice là một kiểu cấu trúc mảng động ( dynamic array ).

- Để  hiểu rõ hơn hãy xét ví dụ sau đây:


```go
package main

import (
    "fmt"
)

func main() {
	days := [7]string{"Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy", "Chủ Nhật"}

	schools := days[0:5]

	fmt.Println("len:", len(schools), "cap:", cap(schools))
    	
	fmt.Println("Học sinh phải cắp sách tới trường vào ngày:", school)
	fmt.Println("Hm nhưng nhà trường chợt nhận ra các em học sinh rất ham hoc vì vậy")

	schools = schools[:6]
	fmt.Println("Học sinh phải cắp sách tới trường vào ngày:", school)
}
```

- Hãy thử  chạy chương trình trên xem, bạn sẽ thấy điều kỳ diệu sẽ xảy ra:

```go
len: 5 cap: 7
Học sinh phải cắp sách tới trường vào ngày: [Thứ Hai Thứ Ba Thứ Tư Thứ Năm Thứ Sáu]
Hm nhưng nhà trường chợt nhận ra các em học sinh rất ham hoc vì vậy
Học sinh phải cắp sách tới trường vào ngày: [Thứ Hai Thứ Ba Thứ Tư Thứ Năm Thứ Sáu Thứ Bảy]
```

- Oh! mặc dù độ dài của slice chỉ là 5 nhưng khi ta thực hiện phép gán để  nới rộng slice ra ( phạm vi mở rộng phải nhỏ hơn hoặc bằng giá trị của cap ), nó hoạt động mà không gặp bất kỳ lỗi nào. 

### Thao tác với slice

- Tương tự như mảng trong phần này chúng ta sẽ đi vào các cách thap tác với slice thường xuyên sử  dụng tới.

#### Hàm khởi tạo

- Như cú pháp ở đầu chương chúng ta đã làm quen với một số  cách để  khởi tạo một slice rồi ,tuy nhiên trong bộ thư viện chuẩn của Go còn hỗ  trợ việc khởi tạo một slice thông qua hàm `make` nữa ( hàm này không chỉ dùng để khởi tạo duy nhất mảng, slice mà nó cùng dùng cho cả map , channel, etc ). 
Hàm này được định nghĩa như sau:

```go
func make([]T, len, cap) []T
```

- Như bạn thấy hàm `make` này sẽ trả về  kết quả là một slice thuộc kiểu dữ liệu T cùng với đó là 
độ dài và sức chưa mà slice đó có khả năng có. Để  hình dung rõ hay xét ví dụ sau:

```go
package main

import (  
    "fmt"
)

func main() {  
    temp := make([]int, 2, 10)
	fmt.Println(temp)
	temp = temp[:10]
	fmt.Println(temp)
}
```

- Ở đây chúng ta sẽ tạo ra một slice `temp` thuộc kiểu số nguyên , có độ dài ban đầu là 2 cùng với giá trị mặc định và nó có thể  mở rộng tới 10 phần tử.  

```go
[0 0]
[0 0 0 0 0 0 0 0 0 0]
```

#### thêm phàn tử

- Như chúng ta biết, mảng là một kiểu cấu trúc có kích thước xác định do đó việc mở rộng mảng gần như không thể  nhưng với slice thì lại khác. Slice có thể  mở rộng kích thước của nó mà không có bất kỳ lỗi nào, và để  làm điều này bạn hãy sử  dụng hàm `append` trong bộ thư viện chuẩn.

- Cú pháp của hàm `append` được đinh nghĩa như sau:

```go
func append(s []T, x ...T) []T  
```

- Chúng ta thấy rằng hàm sẽ nhận hai tham số  trong đó tham số  đầu chính là slice gốc mà bạn muốn thay đổi và tham số  thứ hai hơi đặc biệt một chút, `...T` ở đây là một khái niệm được gọi là Variadic, nó sẽ ám chỉ tới một danh sách các tham số  mà bạn muốn truyền vào. Vì vậy nói ngắn gọn 
thì hàm `append` này nhận hai tham số nhưng thực chất là nó có thể  nhận rất nhiều tham số  và những tham số  này là những giá trị mà bạn muốn đẩy vào trong slice. 

- Hãy xét ví dụ sau:

```go
package main

import (  
    "fmt"
)

func main() {  
    temp := make([]int, 2, 10)
	fmt.Println("Gía trị bạn đâu là:", temp)
	temp = append(temp, 1, 2)
	fmt.Println("Gía trị sau khi gọi hàm append là:", temp)
}
``` 

Kết quả:

```go
Gía trị bạn đâu là: [0 0]
Gía trị sau khi gọi hàm append là: [0 0 1 2]
```
- Tốt! biến slice của chúng ta đã được thêm hai phần tử  mới.


### hàm copy

- slice là kiểu mảng tham chiếu, nếu lấy một biến mới gán tới biến slice gốc thì nó chỉ là 
biến mới này trỏ tới biến slice gốc thôi và chỉnh sửa bất kỳ một trong hai biến này đều ảnh hưởng lân nhau. Trong trường hợp bạn muốn tạo ra mới một biến slice không ảnh hưởng tới biên slice gốc bạn có thể  dùng hàm `copy`. Hàm này sẽ sao chép nội dung của slice gốc và tạo ra một slice mới tách biệt với slide gốc và nếu có chỉnh sửa giá trị gì thì nó cũng không ảnh hưởng lẫn nhau.



