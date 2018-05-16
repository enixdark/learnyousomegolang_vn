## Method

- Ở phần trước chúng ta đang tìm hiểu về  kiểu cấu trúc, định nghĩa một số  cấu trúc và sử  dụng chúng. Nhưng các thao tác của chúng ta con khá hạn chế  Do đó trong chương này chúng ta sẽ mở rộng ra với hàm. Nào hãy xét ví dụ sau:

```go
package main

import (
    "fmt"
    "errors"
)

type People struct {
    name string
    age int
}

func find(people People, peoples []People) (error, People) {
    for _, p := range peoples {
        if people.name == p.name {
            return nil, p
        }
    }
    return errors.New("Không tìm thấy người có tên là " + people.name), People{}
}

func main() {

    peoples := []People{
        People{ name: "Quan", age: 25},
        People{ name: "Quy", age: 23 },
        People{ name: "Que", age: 18 },
    }


    err, people := find(People{ name: "Quan" }, peoples)

    if err == nil {
        fmt.Println("Thông tin về  người cần tìm là", people)
    }
}
```

- Ở ví dụ trên chúng ta sẽ tạo ra một chương trình dể  tìm kiếm về  thông tin của một người trong một danh sách có sẵn thông qua tên của họ. Để  làm điều này chúng ta sẽ cần một struct gọi là People cho việc lưu trữ thông tin của một người ( Ở đây chúng ta chỉ cân lưu trữ thông tin đơn giản gồm tên và tuổi thôi, đối với các ứng dụng phức tạp thì có thể  lưu trữ nhiều trường hơn). Tiếp đó để  tìm tên của một người chúng ta sẽ tạo ra một hàm để  xử  lí việc này. Logic của hàm này không có gì phức tạp cả, như bạn thấy ở đoạn code trên thôi chỉ đơn giản là xác dịnh tham số  tên và danh sách càn truyền vào thôi. Nhưng có một cách làm tốt hơn thế  này nhiều, trước hết hãy nhớ lại kiến thức hàm, phương thức ( method ) trong về  OOP chút , còn nếu như ban chưa biết gì về  OOP thì đừng lo tôi sẽ giải thích ngay thôi. Trong OOP, một lớp của bạn sẽ bao gồm các thuộc tính ( property ) để  chỉ dữ liệu và các hàm là các hành vi mà một đối tượng có thể  có. Vd như một lớp Người ( class People ) có thuộc tính tên ( name ), tuổi ( age ) và có các hành động như chào, hát  biểu diễn thông qua các phương thức hello, sing và gọi bằng cách `object.func`, etc. Về  cơ bản các phương thực bản chất chỉ là các hàm nhưng chúng ngắn chặt với đối tượng đó, chỉ có thể  sử  dụng cùng với đối tượng đó.

- Mặc dù như đã nói trước, Go không phải là ngôn ngữ lập trình hướng đối tượng ( OOP ) do đó bạn sẽ không thể  có class hay object. Nhưng cũng đừng quên tôi có nhấn mạnh rằng OOP chỉ đơn giản là một khái niệm và nó không nhất thiết phải gắn chặt với class. Bản thân struct cũng có thể  làm được điều này và cũng có method như trong OOP. Hãy xem ví dụ sau:

```go
package main

import (
    "fmt"
    "errors"
)

type People struct {
    name string
    age int
}

func (p People) find(peoples ...People) (error, People) {
    
    for _, people := range peoples {
        if p.name == people.name {
            return nil, p
        }
    }
    return errors.New("Không tìm thấy người có tên là " + p.name), People{}
}

func main() {

    peoples := []People{
        People{ name: "Quan", age: 25},
        People{ name: "Quy", age: 23 },
        People{ name: "Que", age: 18 },
    }

    p := People{ name: "Quan" }
    err, people := p.find(peoples...)

    if err == nil {
        fmt.Println("Thông tin về  người cần tìm là", people)
    }
}
```

- Kết quả:

```go
Thông tin về  người cần tìm là {Quan 0}
```

- Oh! Nhìn xem bạn có thể  gọi trực tiếp hàm từ một biến struct mà không có bất kỳ thông báo lỗi nào hết. Những hàm như vậy được gọi chung là phương thức ( method ). Phương thức là cách để  mô tả hành vi của một kiểu dữ liệu xác định ( Receiver Type - ở đây là đối với kiểu tự định nghĩa trông qua từ khóa `type` ). Trong go phương thức về  cơ bản một là một hàm nhưng trước từ khóa `func` và tên hàm sẽ có thêm một tham số  ở giữa , tham số  này trong Go còn được gọi là tham số  tiếp nhận ( receiver ). 

- Vậy chính xác tham số  tiếp nhận là gì ( receiver ), hiểu đơn giản thì nó sẽ cho biết được mối quan hệ giữa hàm và kiểu dữ liệu định nghĩa của tham số  đấy. Hãy nhìn ví dụ People, bạn có thấy khi ta khai báo một tham số  `(p People)` ở giữa `func` và tên phương thức, lúc này nó sẽ được coi như là một hành vi gắn với kiểu cấu trúc People đó và bất cư khi nào bạn gọi hàm từ một biến thuộc kiểu People thì tham số  này sẽ được coi chính là biến đó ( hãy hình dung nó giống như `self`, `this` trong một số  ngôn ngữ khác ấy, nó tương tự như một thể  hiện của một đối tượng).        

- Cú pháp cho phương thức sẽ được định nghĩa như sau:

```go
type name Type

//thông thương sẽ là 
type T struct {
    field Type
}

func (n T) FuncName() {}
func (n *T) FuncName() {}
```

- Lưu ý: đó là khi khi khai báo một kiểu cấu trúc với các trường có tên xác định thì khi định nghĩa phương thức cho kiểu cấu trúc đó hay tránh việc dịnh nghĩa khai báo một tên trùng với tên của một số  trường trong đó, Vd:

```go     
type T struct {
    Name Type
}

func (n T) Name() {} // ngay lập túc sẽ xuất hiện lỗi type T has both field and method named Name
```

### tham số  tiếp nhận dạng tham trị và tham chiếu ( value reicever  & reference or pointer receiver  )

- Trong cú pháp trên ta thấy có hai trương hợp khai báo phương thức, trong đó trường hợp đầu là khai báo theo tham trị và trường hợp thứ hai là tham chiếu, về  bản chất chúng không khác nhiều so với các ví dụ mà chúng ta đã làm từ đầu chương tới giờ rồi, tuy nhiên để  hiểu rõ về  cách sử  dụng hơn chút chúng ta sẽ đi xét ví dụ sau:

```go
package main

import (  
    "fmt"
)

type People struct {
    name string
    age int
}

func (p People) changeName1(name string) {  
    p.name = name
}

func (p *People) changeName2(name string) {  
    p.name = name
}

func main() {  
    p := People{
        name: "Quan",
        age:  25,
    }
    fmt.Println("Thông tin của một người: ", p)
    
    p.changeName1("Cong")

    fmt.Println("Thông tin sau khi gọi một phương thức với tham trị để  thay đổi tên là: ", p)

    (&p).changeName2("Cong")
    // p.changeName2("Cong") // Trong kiểu cấu trúc bạn hoàn toàn có thể  sử  dụng cách thông thường với kiểu con trỏ tham chiếu 
                             // Go sẽ nhìn nhận và tự động chuyển đổi sang thành (&p)
     
    fmt.Println("Thông tin sau khi gọi một phương thức với tham chiếu để  thay đổi tên là: ", p)
}
```

- Kết quả:

```go
Thông tin của một người:  {Quan 25}
Thông tin sau khi gọi một phương thức với tham trị để  thay đổi tên là:  {Quan 25}
Thông tin sau khi gọi một phương thức với tham chiếu để  thay đổi tên là:  {Cong 25}
```

- Như bạn thấy về  ý nghĩa sử  dụng không có nhièu thay đổi ngoại trừ cách gọi hàm từ một phương thức thuộc kiểu tham trị và một phương thức thuộc tham chiếu.

- Lưu ý: có điều một quan trọng cần hiểu giữa việc sử  dụng tham số  tiếp nhận dạng tham trị với kiểu con trỏ tham chiếu. Như bạn biết trong kiểu cấu trúc với một khai báo thường thì nó sẽ là kiểu tham trị còn với con trỏ sẽ là kiểu tham chiếu. Tùy theo nhu cầu bài toán, ví dụ như trong các bài toán về  sử  lí đa luồng hay tái sử  dụng tài nguyện thì khi đó quyết định sử  dụng kiểu tham trị hay tham chiếu sẽ đảm bảo tính an toàn, ổn định hay tối ưu hơn. Nếu như bạn vẫn cảm thấy khá khó hiểu thì đừng lo bởi vì ai cũng vậy thôi, và hãy yên tâm tôi sẽ cố một chương nâng cao để  dành riêng cho một số  khái niệm hơi mang tính trừu tượng chút.

### Phương thức và hàm

- Oh! Ngạc nhiên đúng không, vốn dĩ dựa vào khái niệm và cách khai báo bạn đã biết là phương thức bản chất là hàm rồi vậy tôi còn viết vào đây làm gì nữa và bạn cảm thấy giống như tôi đang trêu bạn ấy, thật phí mất mày phút khi đọc đoạn tới đoạn này. Oh đúng một phần nhưng mục đích tôi nhấn mạnh phương thức là hàm bởi vì lí do sau:

+ Chúng có thể  cùng tên: Đối với hàm trong một package, chúng ta không thể  có hai hàm trùng tên được nhưng một phương thức cùng tên với một hàm thông thường thì sao, nó hoàn toàn hợp lệ thậm chí các phương thức cùng tên nhưng khác kiểu tham số  tiếp nhận cũng hợp lệ

Vd: 

```go
package main

import (  
    "fmt"
)

type People struct {
    name string
    age int
}

type Animal struct {
    name string
}

func (a Animal) Show() {
    fmt.Println(a)
}

func (p People) Show() {  
    fmt.Println(p)
}

func Show(p People) {
    fmt.Println(p)
}

func main() {  
    p := People{
        name: "Quan",
        age:  25,
    }
    
    p.Show()
    Show(p)
    
    a := Animal{ name: "Lulu" }
    a.Show()
}
```

- Đó là trường hợp bạn cần lưu ý thôi,

### Liên hệ giữa truòng nặc danh và phương thức 

- Như bạn biết một cấu trúc nếu có một trường thuộc kiểu cấu trúc ( ta sẽ gọi tạm là cấu trúc con ) mà bạn không xác định tên cho trường đó thì bạn có thể  dùng ngay chính tên của cấu trúc đó. Mặc dù cách này có hạn chế  là khi khởi tạo bạn phải có một dòng để  khởi tạo riêng cho cấu trúc này nhưng đối với phương thức nó lại có một lợi thế, giả sử  ở cấu trúc con đó có dịnh nghĩa một phương thức, thông thường thì phương thức sẽ dược gọi từ một biến thuộc kiểu cấu trúc đó, nếu bạn muốn từ  cấu trúc góc gọi tới phương thức kia thì bạn phải gọi tới biến của cấu trúc con rồi mới gọi tới phương thức liên kết với cấu trúc con đó. Tuy nhiên với trường nặc danh kiểu cấu trúc cho phép bạn từ cấu trúc gốc gọi trực tiếp tới phương của cấu trúc con. Hãy nhìn ví dụ sau:

```go
package main

import (
    "fmt"
)

type Car struct {
	color string
}

type People struct {
    name string
	age int
	Car
}

func (c Car) Show() {
    fmt.Println("Màu của xe là mau", c.color)
}

func main() {

	p := People{ name: "Quan", age: 25 }
	p.Car = Car{ color: "blue"}
    p.Show()
}
```

- Bạn thấy đó chúng ta không cần phải tường mình rõ fõ cấu trúc Car rồi mới tới phương thức của nó, thay vào dó chúng ta gọi trực tiếp từ cáu trúc People chứa cấu trúc Car được định nghĩa như một trường nặc danh.


### Tổng kêt

- Mặc dù Go không phải là ngôn ngữ lập trình hướng đối tượng OOP, không có class, không có object tuy nhiên không có nghĩa là không thể  áp dụng được khái niệm của OOP vào Go. Thông qua struct và method chúng ta hoàn toàn có thể  sử  dụng được tính chất cơ bản của OOP. Tuy vậy trong OOP còn khá nhiều tính chất khác nữa. Trong các chương sau bạn sẽ dần dần khám phá ta cách áp dụng OOP trong Go như thế  nào.