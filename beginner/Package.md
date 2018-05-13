## Package

- Nếu như để  ý bạn thấy rằng hầu hết từ trước tới giờ chúng ta chỉ sử  dụng một file duy nhất cho mỗi chương trình cùng với khai báo tên package trong 
file đó là `main`, và sử  dụng các package được xây dựng sẵn như `fmt`, `time`. Với mỗi một vd, chương trình cụ thể  mà ta thực hiện, nếu như các thành phần logic, đoạn mã trong 
ví dụ mà chúng ta thực hiện đã có từ trước đó ( ví dụ các đoạn mã trong phần control flow trước đó, cơ bản là chúng khá giống nhau có nhiều đoạn không thay đổi gì dù ở file mới ). Chúng 
ta có thể  tránh việc lặp lại này bằng cách sử  dụng duy nhất một file tuy nhiên cách làm này có một hạn chế  đó là  với một dự án lớn chứa hảng trăm nghìn dòng mã thì thực sự 
rất khó khăn trong việc tìm kiếm và quản lí. Do đó có một cách tốt hơn là chia nhỏ các đoạn logic ra ở mỗi file hay mỗi folder ứng với chức năng của nó và import các file/folder đó khi để  tái sử  dụng các đoạn mã, đồng thời cũng giúp cho việc quản lí, tìm kiếm và đọc mã dễ  dàng hơn, cách làm này khá quen thuộc đúng không nhỉ ? Với những ai đã  từng làm quen với khái niệm modularity hay làm việc với các ngôn ngữ như Java, Python, etc thì không lạ gì cả. Tương tự trong Go chúng ta sẽ sử  dụng `package` để  biểu diễn khái niệm này.

- Cú pháp để  khai báo một package sẽ viết như sau:

```go
package name
```

- Trong đó từ khóa `package` luôn được khai báo ở đầu của mỗi file và `name` là tên của package mà bạn muốn đặt. Trong Go bạn có thể  đặt trùng tên của một package ở nhiều file khác nhau tuy nhiên hãy cẩn thận khi khai báo các hàm trong mỗi file này, vì chúng không được phép trùng nhạu. 

- Lưu ý: cách định nghĩa `package` và sử  dụng chúng trong Go có đôi chút khác biệt hơn so với các ngôn ngữ khác, do đó nếu bạn cảm thấy chưa quên thuộc thì đừng lo 
bởi vì chúng ta sẽ từ từ làm quen thông qua các ví dụ cụ thể.

- Trước khi thực hành với `package` chúng ta cần phải cấu trúc lại thư mục một chút. Xem lại chương đầu tiên, thì theo quy ước mọi ứng dụng, dự án mã nguồn của chúng ta lên được đăt trong $GOPATH/src. Trong chương này chúng ta sẽ thực hiện một chương trình máy tính đơn giản cho phép công cộng, trừ, nhân chia, etc. Nhưng thay vì định nghĩa trông một file thì chúng ta sẽ tạo ra một folder khác gọi tên là `cmath`, folder sẽ định nghĩa các file nằm trong package có tên là `cmath` và chứa các hàm có chức năng tính toán. Sau đó package này sẽ được gọi tới trong file chứa hàm main.

- Hãy tạo ra một dự án có cấu trúc như sau:

```go
src 
  |> package
          |> main.go
          |> cmath
                |> add.go
                |> sub.go
                |> div.go
                |> mul.go
bin
  |> package
```

- Trước tiên để  chúng ta sẽ hình dung các công việc cơ bản nhất chúng ta phải làm. Vì yêu cầu của bài toán là thực hiện mô phỏng một chiếc máy tính cá nhân đơn giản do đó trước tiên ta sẽ hình dung nó sẽ làm những gì, đầu tiên đó là thực hiện các phép tính cộng, trừ, nhân, chia cơ bản.

- Để  bắt đầu hãy gõ đoạn mã sau vào file `main.go` trước tiên:

```go
package main

import (
    "fmt"
    "package/cmath"
)

func main() {
    a := 6
    b := 2

    add := cmath.Add(a, b)
    sub := cmath.Sub(a, b)
    mul := cmath.Mul(a, b) 
    div := cmath.Div(a, b)

    fmt.Printf("Tổng của hai số  %d và %d là : %d\n", a, b, add)
    fmt.Printf("Hiệu của hai số  %d và %d là : %d\n", a, b, sub)
    fmt.Printf("Tích của hai số  %d và %d là : %d\n", a, b, Mul)
    fmt.Printf("Thương của hai số  %d và %d là : %d\n", a, b, Div)
}
```

- Dĩ liên nếu lúc này bạn chạy đoạn mã trên sẽ lỗi bởi vì chúng ta không biết thông tin gì về  package `cmath` cũng như các hàm tiện ích mà nó có cả.
Tuy nhiên bởi mình chúng ta đã xác định rõ được mục địch của chương trình do đó chúng ta có thẻ định nghĩa sơ lược các hàm theo ý muốn của chúng ta.

- Nào giờ là lúc mở folder `cmath` lên và thực hiện như sau:

```go
// add.go
package cmath

func Add(a int, b int) int {
    return a + b
}
```

```go
// sub.go
package cmath

func Sub(a int, b int) int {
    return a - b
}
```

```go
// div.go
package cmath

func Div(a int, b int) int {
    if b == 0 {
        return 0
    }
    return a / b
}
```

```go
// mul.go
package cmath

func Mul(a int, b int) int {
    return a * b
}
```

- Nào hãy chạy thử chương trình xem. Trong trường hợp bạn gặp lỗi `cannot find package "cmath" in any of` hãy xem lại đường dẫn bạn thiết lập cho biến $GOPATH và quan trọng là khi sử  dụng từ khóa `import` đối với một package mà bạn tự định nghĩa bạn phải xác định rõ nơi chính xác mà chưa folder/file định nghĩa package đó. Ví dự như trong trường hợp này tôi đặt package `cmath` trong `$GOPATH/src/package/cmath` khi đó 
tôi sẽ phải gọi tới `import custom_package_path` ( ở đây sẽ là `import "package/cmath"`) để  nạp package và các hàm cần thiết:

```go
Tổng của hai số  6 và 2 là : 8
Hiệu của hai số  6 và 2 là : 4
Tích của hai số  6 và 2 là : 12
Thương của hai số  6 và 2 là : 3
```

- Thật tuyệt phải không, bằng cách sử  dụng `package` chúng ta đã chia nhỏ các công việc ra mỗi file riêng khiến các đoạn mã trở lên ngắn gọn hơn. Hơn nữa, chúng tâ hoàn toàn có thể  sử  dụng các chức năng này ở các chương trình khác mà không cần phải viết lại. 

- Tuy vậy chắc hẳn sẽ có rất nhiều điêu bạn muốn hỏi trong đoạn mã trên đúng không ? 
Đầu tiên chắc bạn sẽ thắc mắc là tạo sao các hàm trong package `cmath` tôi lại đặt tên bắt đầu bắt một chữ in hoa. Lí giải do việc này dó là vì trong Go khi sử  dụng chữ in hoa ở đầu đối với một hàm hay một biến chúng sẽ được hiểu như một hàm, biến công khai, tức là bạn có thể  gọi chúng ở bất kỳ đâu trong các package khác bằng lệnh `package.Func`. Ví dụ như ở đoạn mã trên tôi có thể  gọi hàm `cmath.Add` được định nghĩa trong package `cmath` trong pạckage `main`. Còn khi bạn khai báo với một chữ in thường thì chúng sẽ được hiểu là một dạng riêng tư, tức là chỉ có thể  sử  dụng bên trong package đó thôi. Hãy thử  sủa lại:

```go
// add.go
package cmath

func add(a int, b int) int {
    return a + b
}
```

- Sau đó gọi tới hàm `cmath.add` và chạy chương trình, ngay lập bạn sẽ nhận được thông báo `cannot refer to unexported name cmath.add`.


- Tiếp đó có thể  sẽ có một số  bạn hỏi về  việc nạp một package vào thông qua từ khóa `import` và hướng tới đường dẫn chưa package đó. Cách làm này đối với một số  chương trình nhỏ thì không vấn đề  gì nhưng với một chương trình lớn thì thật sự là khó khăn. Vậy có cách nào tốt hơn trong việc này không. Câu trả lời là có! Khi đó bạn sẽ sử  dụng tới một số công cụ quản lí package như godeps, vgo hay glide, nhưng chúng ta sẽ tìm hiểu tới chúng sau này, còn hiện tại tôi muốn giới thiệu cho các bạn một cách nữa để  nạp package thông qua `go get` và `github`.

#### Import package từ GIT

- Trước tiên bạn phải khởi tạo một dự án trên [github](https://github.com) cái đã. Sau đó hãy đẩy project này lên repo của dự án mà bạn vừa tạo. Ở đây tội đã đảy source vào trong `github.com/enixdark/learnyousomegolang_vn`
do đó trước tiên tôi sẽ phài gọi lệnh `go get github.com/enixdark/learnyousomegolang_vn` để  lấy mã nguồn về  
đường dẫn nơi chứa các package. sau khi đã tải về  hãy thay thế  bằng đường dẫn nơi chưa github của bạn:


```go
package main

import (
    "fmt"
    "github.com/enixdark/learnyousomegolang_vn/beginner/source_code/src/package/cmath"
)

func main() {
    a := 6
    b := 2

    add := cmath.Add(a, b)
    sub := cmath.Sub(a, b)
    mul := cmath.Mul(a, b) 
    div := cmrm path.Div(a, b)

    fmt.Printf("Tổng của hai số  %d và %d là : %d\n", a, b, add)
    fmt.Printf("Hiệu của hai số  %d và %d là : %d\n", a, b, sub)
    fmt.Printf("Tích của hai số  %d và %d là : %d\n", a, b, mul)
    fmt.Printf("Thương của hai số  %d và %d là : %d\n", a, b, div)
}
```

- Dễ  dàng phải không, nếu để  ý thì cách làm này vẫn chỉ là việc import một package từ trong hệ thống của bạn thôi nhưng thay vì phải tường mình rõ đường dẫn trong thư mục trong hệ thống của thì bạn chỉ cần ngầm định đường dẫn dựa trên đường dẫn của github mà thôi. 

### Hàm khởi tạo ( Init function )

- Ở đoạn mã trên bạn đã biết về  cách khai báo package, nạp package tuy nhiên bạn lại không biết là liệu chính xác package đó được đã nạp hay chưa, và nếu khi nạp một package bạn muốn làm một điều gì đó trước tiên ? Thật may mắn là chúng ta có một hàm danh cho điều này. Trong Go mỗi một package đều có một hàm khởi tạo `init`. hàm này được viết the cú pháp sau:

```go
func init() {  
    // logic
}
```  

- Hàm này sẽ không có bất kỳ tham số  hay giá trị trả về  tuy nhiên nó sẽ luôn được tự gọi đầu tiên khi một package được nạp vào. Nào giờ hãy khởi tạo thêm một file đặt tên là `init.go` và đặt trong `package/cmath`
trong hàm này chúng ta sẽ thực hiện như sau:

```go
package cmath

import (
	"fmt"
)

const VERSION = "0.0.1"

func init() {  
   	fmt.Println("module cmath đã được nạp cùng phiên bản", VERSION)
}
```

- Ở file `init.go` này chúng ta sẽ định nghĩa một hàm init cùng một hằng để  hiển thị phiên bản hiện tại của package này. Bên trong hàm `init` chúng ta sẽ đơn giản là hiển thị thông tin phiên bản hiện tại ra. Sau đó lưu file và chạy chương trình. Bạn sẽ thấy thông báo sau được đưa ra trước tiên:

```go
module cmath đã được nạp cùng phiên bản 0.0.1
``` 

- Như bạn thấy nó sẽ ngay lập tức gọi tới hàm `init` trong package `cmath` mặc dù chúng ta không tường minh gọi hàm này ở bất kỳ đâu cả. Nó sẽ được gọi theo thứ tự như sau nếu một package được nạp vào:
  + đầu tiên hàm các biến định nghĩa bên ngoài hàm sẽ được gọi trước.
  + Sau đó hàm `init` sẽ được gọi.
  + Tiếp sau đó là cá hàm khác định nghĩa trong module sẽ được nạp vào.

- Bằng cách sử  dụng hàm khởi tạo này , chúng ta có thể  chèn các đoạn mã khởi tạo khác, các đoạn thộng bán cần thiết vào trong quá trình nạp package.

- Lưu ý: chúng ta không nhất thiết phải tạo một file `init.go` chỉ để  chưa hàm khởi tạo này, bạn có thể đặt nó ở bất kỳ đâu trong các file khai báo cùng với package `cmath` tuy nhiên để  dễ  dàng nhận, đọc chúng ta lên tách ra và phân định rõ chức năng ứng với tên file sẽ tốt hơn.

### Blank identifier

- Yêu cầu trong Go khá chặt ché, nó không cho phép chúng ta import các package hay khai báo các biến, hàm mà không dùng tới chúng trong chương trình. Đúng có thể  nói cách làm này sẽ mang lại hiệu qua cho lập trình viên hơn và tránh được lãng phí tài nguyên cho chương trình. Tuy vậy không phải cách làm này lúc nào cũng hiệu quả nhất là trong quá trình phát triể và kiểm thử. Gỉa sử  khi bạn đang test một số  chức năng của một package nào đó nhưng tại một thời điểm bạn chưa có ý định sử  dụng và bạn chỉ muốn ẩn nó đi tạm thời. Bạn có thể  comment nó lại hay xóa nó đi, tuy vậy đó chưa phải là cách hay. Trong Go bằng cách sử  dụng ký tự `_` đặt trước. Cú pháp của nó sẽ viết như sau:

```go
_ package_name
```

- Khi đó trình biên dịch sẽ ngầm hiểu rằng package bạn khai báo vẫn được nạp vào nhưng bạn chưa hề  có ý dịnh sử  dụng và có thể  sử  dụng trong tương lai, nó sẽ bỏ qua và không đưa ra bất kỳ một cảnh báo nào cả, và chương trình vẫn hoạt động bình thường. Vd:

```
package main

import (
    "fmt"
    _ "package/cmath"
)

func main() {
    fmt.Println("Work")
}
``` 

- Thật tuyệt phải không!

### Tổng kết

- Như vậy chúng ta đã tìm hiểu được cách sử  dụng package để  quản lí và bảo trì các đoạn mã nguồn trở lên 
tốt hơn rồi. Trong các chương sau chúng ta sẽ ứng dụng dựa trên cấu trúc package này và đi sâu vào các thành phần khác trong Go như các kiểu cấu trúc dữ liệu, xử  lí chuối và tính chất quan trọng mà chúng ta đã đề  cập đàu chương đó là `concurrency`   