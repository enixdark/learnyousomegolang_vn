## Singleton pattern ?

- Ở giải đấu bóng đá AFF cup quy định trước sau và sau khi kết thúc trận đấu, mỗi đội tuyển đều phải cử  ra huấn luận viên trưởng và một cầu thủ 
để  tham gia họp báo trả lời các vấn đề . Nếu ko họ sẽ bị liên đoạn bóng đá thế  giới phạt. Tất nhiên là ko đội bóng nào muốn bị phạt cả, trong quy định này ta thấy việc cử  ra một cầu thủ là đièu bắt buộc nhưng ko yêu cầu luôn là một cầu thủ chỉ định mà thay vào đó bất cứ ai là cầu thủ tham dự trận đấu đều có thể  tham gia, đó có thể  là thủ môn, tiền đạo hay hậu vệ. Tuy nhiên có một lưu ý ở đây đó là huấn luận viên trưởng của đội bóng, đây là người mà ko thể  thay thế  được. Thật vậy, nhưng điều đó thì liên quan gì tơi mẫu Singleton. Oh trước tiên chúng ta sẽ đi vào khái niệm của mẫu này.

## Định nghĩ từ GoF

- Đảm bảo một lớp chỉ có một thể  hiện và được sử  dụng như một biến toàn cục

## Khái niệm 

- Ở đây chúng ta hiểu đơn giản đó là một lớp sẽ chỉ có duy nhất một thể  hiện, đối tượng được tạo ra, không được phép có nhiều đối tượng của cùng một lớp đó được. Một khi một đối tượng ( tạm gọi là A ) được tạo ra từ lớp đó thì ở lần tạo đối tượng tiếp theo từ chính lớp đó vẫn là đối tượng A. Tính chất này đảm bảo giảm bớt được sự lãng phí khi tạo ra nhiều đối tượng không cần thiết và dễ  dãng bảo trì.

## Mục đích sử  dụng:

- Chúng ta thường sử  dụng mãu này khi:

+ Cần một biến, đối tượng duy nhất, chia sẻ giá trị trong chương trình.
+ Giới hạn việc sinh ra đối tượng của một kiểu dữ liệu, lớp trong một chương trình.   

- Để  mình họa, thông thường trong trong công việc bạn sẽ sử  dụng khi:

+ Khi thao tác với cơ sở dữ liệu ( Database ), bạn muốn sử  dụng cùng kết nối tới database để  thực hiện các câu lệnh truy vấn ( query ).
+ Khi kết nối tới server qua ssh.
+ Khi sử  dụng một cache object để  lưu trũ thông tin.

## Ví dụ:

- Tiếp theo chúng ta sẽ đi vào ví dụ cũ thể  của mâu Singleton, ở đây tôi sẽ sử  dụng ví dụ ở đầu chương để  mình họa chương trình.
Trước tiên chúng ta sẽ tạo ra một file `singleton.go` và gõ vào file này các đoạn mã sau:

```go
package singleton

// import "fmt"

type struct Coach {
        name string
        age int
}

func (c *Coach) IncreaseAge() int {
        return 0
        // c.age = c.age + 1
        // return c.age
}

func (c* Coach) GetAge() int {
        return 0
        // return c.age
}

func (c* Coach) GetName() string {
        return ""
        // return c.name
}

func GetCoach() *Coach {
        // if (coach == nil ) {
        //     coach = &Coach{"A", 50}
        // }

        // return coach
        return nil
}

var coach *Coach
```

- Ở đây chúng ta sẽ định nghĩa `Coach` thuộc kiểu cấu trúc gồm hai trường `age` và `name` để  chi tuổi và tên của huấn luận viện mà chúng ta cân tạo ra. Tiêp đó chúng ta thêm vào một số  chức năng để  hiển thị tên và tuổi cùng với tăng số  tuổi của huấn luyện viên lên 1 sau mỗi lần gọi. Tiếp theo chúng ta sẽ đi thưc hiện viết các đoạn unittest ứng với logic của các chức năng được đưa ra.

- Hãy tạo ra một file với tên là `singleton_test.go` và định nghĩa mỗi chứng năng tương ứng với hàm test sau:

```go
package singleton

import "testing"

func TestGetCoach(t *testing.T) {
        c1 := GetCoach()
        if c1 == nil {
                t.Error("expected pointer to Singleton after calling GetCoach(), not nil")
        }

        c2 := GetCoach()

        if c1 != c2 {
                t.Error("expected between c1 and c2 are'diff")
        }

        name := c1.GetName()

        if name != "A" {
                t.Error("expected between name of coach instance 's not A")
        }

        age := c1.GetAge()

        c1.IncreaseAge()

        ageNew := c1.GetAge()

        if age + 1 != ageNew {
                t.Error("expected the age of coach doesn't increase")
        }
}
```

- Nào giờ hãy thử  chạy đoạn mã trên và tất nhiên là bạn sẽ gặp lỗi khi chạy chúng:

- Tiếp theo chúng ta sẽ tiến hành viết logic cho chúng.

```go
func (c *Coach) IncreaseAge() int {
        c.age = c.age + 1
        return c.age
}

func (c* Coach) GetAge() int {
        return c.age
}

func (c* Coach) GetName() string {
        return c.name
}

func GetCoach() *Coach {
        if (coach == nil ) {
            coach = &Coach{"A", 50}
        }

        return coach
}
```

- Nào trước tiên chúng ta sẽ phân tích một chút, như bạn biết về  định nghĩ lúc trước của mãu Singleton, chúng ta sử  dụng mẫu nãy để  tạo một đối tượng và sử  dụng nó trong toàn bộ chương trình. Thông thường với ngôn ngữ như Java hay C++ chúng ta sẽ sử  dụng kiểu dữ liêu `static` trong `class` sau đó định nghĩ một phương thức cũng thuộc kiểu `static` và gọi tới để  lấy thể  hiện của class, tuy nhiên với Go, ta biêt rằng Go không phải là ngôn ngữ OOP, nhưng chúng ta hoàn toàn có thể  làm điều này bằng cách sử  dụng `struct` cùng con trỏ.

- Ở đây đầu tiên ta sẽ khai báo một `struct` Coach và khai báo con trỏ thuộc `struct` dùng làm biến global phía bên ngoài, Bằng cách này chúng ta có thể  sử  dụng biến coach ở hầu hết các nơi trong chương trình:

```go
type struct Coach {
        name string
        age int
}

var coach *Coach
```

- Bởi vì chúng ta mới chỉ khai báo biến `coach` như một con trỏ của `struct` Coach do đó mặc định nó luôn trỏ tới `nil`, việc tiếp theo chúng ta cần đó là tạo ra giá trí cho biến này thông qua hàm `GetCoach`:

```go
func GetCoach() *Coach {
        if (coach == nil ) {
            coach = &Coach{"A", 50}
        }

        return coach
}
```

- Để  ý ở đây hàm `GetCoach` sẽ chỉ tạo ra duy nhất một thể  hiện của `struct` Coach thông qua việc gán con trỏ `coach` với địa chỉ của thể  hiện đó và trả lại thông tin con trỏ đó. Và nếu các lân gọi sau bạn sẽ luôn nhận được cùng một con trỏ và địa chỉ của đối tượng mà ta đã tạo ra trước đó, Thật vậy nó đúng với những già mà khái niệm đã đưa ra về  mấy Singleton.

- Tiếp đó việc còn lại là định nghĩa logic cho các hàm của `struct` Coach:

```go
func (c *Coach) IncreaseAge() int {
        c.age = c.age + 1
        return c.age
}

func (c* Coach) GetAge() int {
        return c.age
}

func (c* Coach) GetName() string {
        return c.name
}
```

- Giờ hay chạy lại test , bạn sẽ nhận được kết quả chính xác với những gì mà đã định nghĩa trong case test:

```
=== RUN   TestGetCoach
--- PASS: TestGetCoach (0.00s)
PASS
ok   
```

## Kết luận

- Như vậy là chúng ta đã tìm hiểu về  mẫu đầu tiên trong `Creational Patterns`, Đây là một trong những mẫu `Creational` được sử  dụng nhiều nhất và thường dùng đẻ hỏi trong các cuộc phỏng vấn