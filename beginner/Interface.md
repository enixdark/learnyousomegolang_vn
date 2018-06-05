## Giao Tiếp ( Interface )

- Trong OOP, Interface được sử  dụng như một khái niệm để  xác định những hành vi lên có trong một đối tượng. Vd như một đối tượng người hay một đối tượng con vật có những hành vi sau như đi, ngồi, chạy, dừng. Thông thường nó được liệt kê dưới một tập các phương thức ( nhưng chưa thực hiện logic ). Tương tự interface trong Go cũng được sử  dụng với mục dích như vậy nhưng thay vì dùng cho đối tượng thì nó sẽ được dùng cho các kiểu dữ liệu do người dùng định nghĩa. 

- Lưu ý: thông thường với interface quy chuẩn chúng ta thường sử  dụng để  đưa ra các hành vi chung chung dưới dạng các phương thức ( chỉ khai báo tên , tham số  và dữ liệu trả về  thôi mà chưa thực hiện bất kỳ logic nào ) . Sau đó khi một kiểu dữ liệu nào sử  dụng interface này sẽ phải thực hiện tất cả các logic cần làm. Tuy nhiên một số  ngôn ngữ còn hỗ  trợ thực hiện logic ngay trong chính interface đó nhưng với Go chúng ta vẫn tuân theo quy chuẩn ban đầu thôi. 

- Cú pháp của interface trong Go được viết như sau:

```go
type InterfaceName interface {  
    methodName Type
}
```

- Hãy xét ví dụ sau để  hiểu rõ:

```go
package main

import (
    "fmt"
)

type People struct {
    name string
    age int
}

type Action interface {
    run() People
    stop() People
    walk() People
}

func (p People) run() {
    fmt.Printf("Anh %s đang chạy\n", p.name)
}

func (p People) walk() {
    fmt.Printf("Anh %s đang đi bộ\n", p.name)
}

func (p People) stop() {
    fmt.Printf("Anh %s đang đứng yên\n", p.name)
}

func main() {
    p := People{ name: "Quan", age: 25 }
    var a Action = p
    a.walk()
    a.run()
    a.stop()
}
```
- Kết quả 

```go
Anh Quan đang đi bộ
Anh Quan đang chạy
Anh Quan đang đứng yên
```

- Oh chắc bạn sẽ ngạc nhiên và hỏi tại sao lại như vậy đúng không ? Đừng vội tôi sẽ giải thích. Ở đoạn ví dụ này chúng ta viết một chương trình đơn giản mô phỏng hành vi của một người gồm có dừng, đi bộ và chạy. Trước tiên chúng ta sẽ khai báo một kiểu cấu trúc có tên là People chứa thông tin tên và tuổi của một người, tiếp đó lúc này chúng ta sẽ dịnh nghĩa một interface gọi là Action, interface này sẽ mô tả các hành vi của người thông qua các phương thức. Như bạn thấy chúng ta không có bất kỳ đoạn logic nào bên trong cả, đơn giản chỉ là khai báo thôi. Tiếp theo dựa vào các phương thức có trong interface chúng ta thực hiện các phương thức ứng với kiểu cấu trúc People. Trong hàm `main` lúc này bạn có để  ý rằng chúng ta có thể  gán một kiểu thuộc `interface` tới 1 biến có kiểu  cấu trúc không. Nó không hề  có lỗi nào cả mặc dù hai kiểu này hoàn toàn khác nhau, đúng nó không hề  lỗi bởi vì khi bạn khai báo thực hiện hết các phương thức đã được định nghĩa trong interface thì bạn hoàn toàn có thể  gán với một kiểu interface tới một kiểu cấu trúc như một kiểu đại diện. Sau đó khi gọi các phương trong interace Go sẽ tìm kiếm và gọi tới các phương thức trong cấu trúc.

- Vậy cách làm này có ích gì, nếu nó là một cách đại diện thôi thì thật thừa thãi. Đừng vội hãy nhìn nhận ví dụ sau sẽ hiểu tạo sao lại có interface:

```go
package main

import (
    "fmt"
)

type People struct {
    name string
    age int
}

type Cat struct {
    name string
}

type Action interface {
    run() 
    stop() 
    walk()
}

func (p People) run() {
    fmt.Printf("Anh %s đang chạy\n", p.name)
}

func (p People) walk() {
    fmt.Printf("Anh %s đang đi bộ\n", p.name)
}

func (p People) stop() {
    fmt.Printf("Anh %s đang đứng yên\n", p.name)
}

func (c Cat) run() {
    fmt.Printf("Con mèo %s đang đứng yên\n", c.name)
}

func (c Cat) walk() {
    fmt.Printf("Con mèo %s đang đi bộ\n", c.name)
}

func (c Cat) stop() {
    fmt.Printf("Con mèo %s đang đứng yên\n", c.name)
}

func main() {
    p := People{ name: "Quan", age: 25 }
    c := Cat{ name: "Meww" }
    a := []Action{p, c}
    for _, d := range a{
        d.walk()
        d.run()
        d.stop()
    }
}
```
- Kết quả:

```go
Anh Quan đang đi bộ
Anh Quan đang chạy
Anh Quan đang đứng yên
Con mèo Meww đang đi bộ
Con mèo Meww đang đứng yên
Con mèo Meww đang đứng yên
```

- Điều đặc biệt của Interface năm ở đây. bạn chi cần khai báo một giao tiếp interface chung và thực hiện các hành vi logic cho cả cấu trúc Cat và People là có thê thông qua interface đó đại diện cho cả hai rồi, thậm chí là cho nhiều cấu trúc hơn nữa.

### Tham số  Interface và interface nặc danh

- Ở đoạn code trên ta thấy được việc dùng interface như một cách định danh chúng cho các kiểu cấu trúc mà có chứa các phường thức của interface rồi. Thêm nữa bản thân interface cũng coi là một kiểu dữ liệu, do đó không có lí do gì mà chúng ta lại không sử  dụng như một tham số  cả. Từ ví dụ trên chúng ta có thể  taí cấu trúc lại như sau:

```go
func AllAction(a Action) {
    a.walk()
    a.run()
    a.stop()
}

func main() {
    p := People{ name: "Quan", age: 25 }
    c := Cat{ name: "Meww" }
    a := []Action{p, c}
    for _, d := range a{
        AllAction(a)
    }
}
```

- Kêt quả vẫn vậy nhưng nó gọn hơn một chút thôi :). Thật ra mục đích tôi muốn nhấn mạnh hơn một chút về  điều này, Như cú pháp bạn tháy khi khai báo một interface thông thường sẽ kèm theo ít nhất một phương thức bên trong nó đúng không. Đúng nhưng mà đó chỉ là trường hợp chung thôi còn lại bạn vẫn có thể  khai báo một interface mà không cùng với bất kỳ phương thức nào bên trong cả, khi đó nó sẽ được coi là một interface rỗng.

```go
type InterfaceName interface {}
```

- Cú pháp chỉ đơn giản vậy thôi nhưng nó sẽ dùng như thế  nào khi mà không có bất kỳ một phương thức nào để  gọi cho dù với cách này nó có thể  đại diện cho rất nhiều các kiểu dữ liệu. Vậy bạn còn nhớ trong chương cấu trúc chúng ta đã làm quen với một cấu trúc nặc danh ( anonymous struct ) rồi chứ, bởi vì cấu trúc này chỉ được tao tại một thời điểm xác định và không có tên định danh cho nó lên thường không thể  được dùng làm tham số  truyền vào hàm. Nhưng đó chỉ là ở chương trước thôi còn khi bạn đã làm quen với interface, đặc biệt là interface rông thì bạn có thể , để  rõ ràng nhất hãy xem ví dụ sau đây:

```go

```

