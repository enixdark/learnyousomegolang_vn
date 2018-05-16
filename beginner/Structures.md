## Structure

- Chúc mừng bạn, vậy đã bạn đã đi được một quãng đường cũng kha khá rồi. từ trước tới giờ bạn đã được tiếp xúc với biến, với cấu trúc điều khiển, hàm, package, một số  kiểu cấu trúc như mảng, slice, map. Với những thứ đã học mà tôi vừa liệt kê ra đủ để  bạn có thể  viết được các chương trình tương đối rồi nhưng với một trường trình, một bài toán phức tạp thì sao, khi mà chúng ta cần thao tác với nhiều dữ liệu hơn, thậm chí với cả những dũ liệu không theo một cấu trúc thì sao. Lúc này bạn sẽ cần phải tự mình định nghĩa riêng cấu trúc để  giải quyết cho từng bài toàn cụ thể. Và cũng là lúc bạn cần tới một khái niệm mới đó là Structure và một số  khái niệm khác nữa Method, Interface và một chút phong cách OOP cùng cách tổ  chức code. 

- Thôi không lan man nữa, trong chương chúng ta sẽ tìm hiều về  Structure. Về  bản chất structure là một kiểu cấu trúc dữ liệu do người dùng tự định nghĩa mà trong đó nó sẽ bao gồm một tập hợp các trường dữ liệu do người đùng đưa ra và được kết hợp thành một khối ( thường là dưới dạng một biến ) thường dược dùng để  lưu trữ thông tin dứ liệu. Hãy hình dung giả sử  bạn có một chiếc máy tính ( computer ), nhưng để  một chiếc máy tính chạy bạn phải có RAM, màn hình ( display screen ), CPU, Main, Mouse, Keyboard, etc. Để  lưu trữ bạn có thẻ tạo ra mỗi biến cùng tên kết hợp một định danh nào đó. Vd như Anh A có một chiếc máy tính gồm CPU 1GB RAM ( ở đây sẽ dùng đơn vị là MB ), CPU intel core i3, màn máy tính Dell 24 inch, Main hãng asus, HDD 1TB ( tính theo đơn bị GB ). Khi đó viết dưới đoạn mã Go sẽ như sau:

```go
// sử  dụng biến thông thường
cpuA := "Intel Core i3"
displayScreenA := "Samsung 24 inch"
ramA := 1024
mainboardA := "Asus"
hddA := 1024

// sử  dụng mảng và coi như thứ tự dịnh nghĩa được sắp xếp theo [displayScreenA, mainboardA, cpuA, ramA, hddA]

computerA := ["Samsung 24 inch", "Asus", "Intel Core i3" , "1024", "1024"] 

// sử  dụng map 

computer := map[string]map[string]string{
    "computerA": map[string]string{
        "displayScreen": "Samsung 24 inch",
        "cpu": "Intel Core i3",
        "hdd": "1024",
        "mainboard": "Asus",
        "ram": "1024",
    }
}
```

- Hm! Bạn thấy đó dùng cách thông thường hoàn toàn có thể  định nghĩa được nhưng chúng khá phức tạp và rối phải không, nhất là trong trường hợp chúng ta có nhiều dữ liệu như vậy thì sao ?

- Trước khi thắc mắc tiếp hãy xem cú pháp và ví dụ sau:

```go
type StructureName struct {  
    fieldName Type
}
```

Vd:

```go
type Computer struct {
    displayScreen string
    cpu string
    hdd int
    mainboard string
    ram int
}

hay 

type Computer struct {
    displayScreen, cpu, mailboard string
    hdd, ram int
}
```

Vd sử  dung:

```go
computerA := Computer{
    displayScreen: "Samsung 24 inch",
    cpu: "Intel Core i3",
    hdd: 1024,
    mainboard: "Asus",
    ram: 1024,
}
```

- Oh thoạt nhìn khá giống kiểu map lồng nhau nhưng mà dữ liệu của các khóa, trường chính xác hơn và nhìn gọn gàng hơn. Đúng vậy đó chính là cấu trúc struct, bằng cách người dùng tự định nghĩa các trường kèm theo các kiểu dư liệu mong muốn, họ có thể  dùng chúng cho việc lưu trữ dữ liệu một cách chính xác, gọn gàng và dễ  quản lí hơn.

- Lưu ý: Nếu bạn đã học về  C và sau đó là C++ hay một số  ngôn hướng đối tượng OOP khác như Java thì sẽ không lạ gì với cấu trúc struct và Class, thậm chí một số  bạn sẽ hỏi là sao không dùng Class mà lại struct làm gì. Có một điều tôi quên nói với bạn là trong Go không có khái niệm Class, Object do đó cũng sẽ không có OOP. Tuy vậy bạn cũng nhớ rằng OOP chỉ đơn giản là một khái niệm thôi và cơ bản Structure cũng có thể  làm được những gì mà OOP có. Ngoài ra còn có một số phong cách, khái niệm khác như Functional nữa do đó đừng thần thánh hóa OOP làm gì. Đối với Functional tôi sẽ có hẳn một phần tương đối lớn dành riêng cho nó. Còn nếu bạn vân muốn OOP thì đừng lo, bởi vì Go không hẳn là không hỗ  trợ OOP tuy nhiên chúng ta sẽ làm theo một cách khác thay vì định nghĩa Class. Sóm thôi bạn sẽ nhận ra.

### Khởi tạo một structure

- Việc đầu tiên của chúng ta sau khi định nghĩa đó là khởi tạo nó. Áp dụng lại vd phái trên, 
trong Go struct được khai báo và khởi tạo như sau:

```go
var name CustomStruct

name := CustomStruct{ 
    field: value
}
```

Vd: 

```go
package main

import (
	"fmt"
)

type Computer struct {
	displayScreen string
    cpu string
    hdd int
    mainboard string
    ram int
} 

func main() {
	
	computerA := Computer{
		displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
    }

    computerB := Computer{"Samsung 20 inch", "AMD Ryzen 5", 256, "Asrock", 8000}
    
    computerC := Computer{
        displayScreen: "Dell 2̃7 inch",
    }

    fmt.Println("Anh A đang sở hữu chiếc máy tính có cấu hình", computerA)
    fmt.Println("Anh B đang sở hữu chiếc máy tính có cấu hình", computerB)
    fmt.Println("Anh C đang sở hữu chiếc máy tính có cấu hình", computerC)

}
```

- Kết quả:

```go
Anh A đang sở hữu chiếc máy tính có cấu hình {Samsung 24 inch Intel Core i3 1024 Asus 1024}
Anh B đang sở hữu chiếc máy tính có cấu hình {Samsung 20 inch AMD Ryzen 5 256 Asrock 8000}
Anh C đang sở hữu chiếc máy tính có cấu hình {Dell 2̃7 inch  0  0}
```

- Việc khởi tạo một structure đơn giản phải không, thậm chí trong trường hợp bạn quên không khởi tạo giá trị cho các trường phía sau thì Go cũng tự động cung cấp các giá trị mặc định cho chúng.

- Lưu ý: Nếu bạn không xác định rõ các trường cùng giá trị tương ứng khi khởi tạo thì hay nhớ thứ tự sắp xếp các trường mà bạn khai báo một structure cùng kiểu dữ liệu tương ứng của chúng. Nếu không Go sẽ ngay lập tức đưa ra thông báo lỗi về  việc không khớp kiểu dữ liệu đang sử  dụng.

### Lấy thông tin và thay đổi dữ liệu

- Sau khi bạn đã khởi tạo một structure rồi thì việc tiếp theo bạn thường làm đó là lấy, trích xuất tất cả hoặc một thông tin bất kỳ ra để  sử  dụng và đôi khi cần thay đổi lại nó. Việc này rất đơn giản bạn chỉ cần xác định tên của trường mà bạn muốn truy cập
vào và để thay đổi giá trị thì chỉ việc gán lại trường đó với một giá trị nào thôi

```go
package main

import (
	"fmt"
)

type Computer struct {
	displayScreen string
    cpu string
    hdd int
    mainboard string
    ram int
} 

func main() {
	
	computer := Computer{
		displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
    }

    fmt.Println("Bạn đang sở hữu một màn hình", computer.displayScreen)
    computer.displayScreen = "Samsung 27 inch"
    fmt.Println("Oh bạn vừa chuyển sang sử  dụng màn", computer.displayScreen)
}
```

- Kết quả:

```go
Bạn đang sở hữu một màn hình Samsung 24 inch
Oh bạn vừa chuyển sang sử  dụng màn Samsung 27 inch
```

### Tham chiếu và tham trị

- Hãy thử  gán một biến structure tới một structure đã có sẵn xem:

```go
package main

import (
	"fmt"
)

type Computer struct {
	displayScreen string
    cpu string
    hdd int
    mainboard string
    ram int
} 

func main() {
	
	computer := Computer{
		displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
	}
	
	var clone Computer = computer
	fmt.Println("Trạng thái ban đâu là", computer)

	clone.hdd = 256
	fmt.Println("Trạng thái máy tính lúc sau là", computer)

	var pointer* Computer = &computer
	(*pointer).hdd =  512
    fmt.Println("Trạng thái máy tính lúc sau là", computer)
    fmt.Println("HDD khi thay đổi là", pointer.hdd)

	var phdd *int = &computer.hdd
	*phdd = 2048
	fmt.Println("Trạng thái máy tính tiếp sau là", computer)
}
```

- Kết quả:

```go
Trạng thái ban đâu là {Samsung 24 inch Intel Core i3 1024 Asus 1024}
Trạng thái máy tính lúc sau là {Samsung 24 inch Intel Core i3 1024 Asus 1024}
Trạng thái máy tính lúc sau là {Samsung 24 inch Intel Core i3 512 Asus 1024}
HDD khi thay đổi là 512
Trạng thái máy tính tiếp sau là {Samsung 24 inch Intel Core i3 2048 Asus 1024}
```

- Dựa vào kết quả của ví dụ trên ở đoạn đầu tiên ta thử  gán một biến mới tới một structure đã có sẵn và thay đổi thông tin một trường trong biến mới này. Không có gì thay đổi cả do đó ta có thể  khẳng định là phép gán trong Go đối với một structure là một kiểu dữ liệu tham trị. Và để  sử  dụng tham chiếu, chúng ta sẽ dùng tới con trỏ ( pointer ), tuy vậy có một điều chú ý ở đây là dể con trỏ truy cập vào thông tin của structure mà nó trỏ tới bạn sẽ phải bao nó trong cặp ngoặc `()` thay vì sử  dụng trực tiếp `*`, tuy nhiên khi chỉ muốn truy cập vào một trường cụ thể  từ con trỏ thì bạn lại chỉ cần xác định trường cang truy cập là đủ ( đây là một cách viết rút gọn bên cạnh cách viết sử  dụng `*`). Ngoài ra ở câu lệnh cuối cùng tôi muốn nhấn mạnh đó là hãy nhớ structure chỉ là một kiểu cấu trúc gồm một tập các kiểu dữ liệu khác hợp lại thông qua mỗi trường lên mỗi một trường cũng có coi là một biến và sử  hữu vùng nhớ riêng của mình, vì vậy không có gì lạ nếu bạn dùng một con trỏ cùng kiểu dữ liệu trỏ tới trường đó cả.


### Cấu trúc nặc danh ( Anonymous Structure )

- Một điều tôi quên không nói với bạn về  kiểu cấu trúc ( struct ), thông thường bạn sẽ sử  dụng từ khóa `type`, từ khóa này cho phép chúng ta tự định nghĩa một kiểu dữ liệu dựa trên một kiểu đã có sẵn. Tuy vậy bạn hoàn toàn có thể  sử  dụng `var` đối với kiểu `struct`. Tôi đã quên không nói là kiểu cấu trúc mặc dù là kiểu do người dùng tự dịnh nghĩa nhưng từ khóa `struct` ở đây cũng là một kiểu dữ liệu. 

- Và trong một số  trường hợp nếu như bạn không muốn xác định trước cấu trúc bạn muốn định nghĩa mà chỉ muốn định nghĩa và dùng nó tại một thời điểm xác định thì có thể  làm như sau:

```go
package main

import (
	"fmt"
)

func main() {
	
	var computerA struct {
        displayScreen string
        cpu string
        hdd int
        mainboard string
        ram int
    } := struct {
        displayScreen string
        cpu string
        hdd int
        mainboard string
        ram int
    } {
        displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
    }

    computerB := struct { string int } { "Samsung 20 inch", 256 }

    fmt.Println("Chiếc máy tính A có thông số  là:", computerA)
    fmt.Println("Chiếc máy tính B có thông số  là:", computerB)
    fmt.Println("HDD của chiếc máy tính B là:", computerB.int)
}
```

- Kết quả:

```go
Chiếc máy tính A có thông số  là: {Samsung 24 inch Intel Core i3 1024 Asus 1024}
Chiếc máy tính B có thông số  là: {Samsung 20 inch 256}
HDD của chiếc máy tính B là: 256
```

- Những cấu trúc như trên còn được gọi là cấu trúc nặc danh ( anonymous structure ). Ngoài ra còn có một lưu ý ở đẩy bạn cũng có thể  sử  dụng nặc danh hay chính xác là không cần phải xác định tên cho các các trường ( hay còn gọi là trường nặc danh - anonymous field ), tuy nhiên khi sử  dụng cách này nó chỉ cho phép duy nhất dữ liệu khác nhạu tồn tại mà thôi. Giả sử  nếu bạn khai báo struct { string string } ngay lập tức Go sẽ thông báo lỗi, đó là bởi vì không có tên xác định do đó việc truy cập vào sẽ phải sử  dụng kiểu dữ liệu như một khóa xác định cần truy cập.

### Cấu trúc lồng ( Nested struct )

- Một lần nữa tôi muốn nhấn mạnh là kiểu cấu trúc ( struct ) là kiểu dữ liệu do người dùng tự định nghĩa, tức là chúng ta hoàn toàn có thể  chủ động đưa ra một cấu trúc mà chúng ta mong muốn. Trong các ví dụ trên chúng ta mới chỉ định nghĩa các trường là các kiểu dư liệu đã có sẵn ( một số  kiểu dữ liệu cáu trúc như slice, map chúng ta chưa sử  dụng nhưng chúng định nghĩa không khác nhau là mấy ). Tuy nhiên chúng ta chưa nhắc gì về  vấn đề  định nghĩa một trường như một cấu trúc cả:

```go
package main

import (
	"fmt"
)

type Computer struct {
	displayScreen string
    cpu string
    hdd int
    mainboard string
    ram int
} 

type Product struct {
    computer Computer
    price int
}

func main() {
	
	computer := Computer{
		displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
	}
	
	product := Product{
        computer: computer,
        price: 12000000
    }
	fmt.Println("Thông tin sản phẩm:", product)
}
```
- Kết quả

```go
Thông tin sản phẩm: {{Samsung 24 inch Intel Core i3 1024 Asus 1024} 12000000}
```

- Bằng cách này chúng ta có thể  tạo ra những cáu trúc phức tạp hơn vd như những đoạn dữ liệu JSON lồng nhau.

- Lưu ý: ở phần cấu trúc nặc danh chúng ta đã biết về  cách định nghĩa một cáu trúc mà không cần phải xác định tên của các trường, tuy nhiên cách làm này khá hạn chế  khi chỉ được phép định nghĩa duy nhất một kiểu dữ liệu và truy cập bởi chính tên kiểu dữ liệu đó, tuy vậy với struct nó dễ  dàng hơn một chút:

```go
package main

import (
	"fmt"
)

type Computer struct {
	displayScreen string
    cpu string
    hdd int
    mainboard string
    ram int
} 

type Product struct {
	price int
	Computer
}

func main() {
	
	computer := Computer{
		displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
	}
	
	product := Product{ price: 12000000 }
	product.Computer = computer
	
	fmt.Println("Thông tin sản phẩm:", product)
}
```

- Ỏ đây chúng ta đã  sử  dụng một trường nặc danh có kiểu Computer. Khác với trường nặc danh bởi kiểu dũ liệu sẵn có trước đó, bởi ví kiểu cấu trúc là kiểu do người dùng đặt ra do đó dựa vào tên của chúng đôi khi chúng ta cũng biết được ý nghĩa. Do đó sử  dụng trường nặc danh đối với kiểu cấu trúc vẫn cho chúng ta biết được ý nghĩa thật sự của trường nặc danh này. Những trường như vậy chúng ta gọi là 
`Promoted field`.


### Tổng kết:

- Vậy là chúng ta đã tìm hiểu được struct, khái niệm, cách định nghĩa và sử  dụng struct cũng như ứng dụng nó như thế  nào rồi. Tuy vậy chúng ta mới chỉ đi được một nửa thôi. Trong chương tiếp chúng ta sẽ tìm hiểu một khái niệm mà có lẽ các tín đồ  của OOP rất quên thuộc. 
