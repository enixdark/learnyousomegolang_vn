## Hello Go!

- Đây là phần tiếp theo của loạt hướng dẫn vê Go. Trong bài học này chúng ta sẽ 
tìm hiểu cách để  viết một chương trình đơn giản trong Go. Yêu cầu trước đó là chúng ta đã cài đặt thành công Go rôi, nếu bạn chưa cài thì có thể  xem lại chương trước.

- Đối với bất kỳ lâọ trình viên nào, khi mới bắt đầu làm quen với một ngôn ngữ lập trình mới chúng ta đều bắt đầu từ một chương trình đơn giản nhất. Vâng một chương trình 'Hello World' chắc không xa lạ gì hết, nó chỉ đơn giản là hiển thị dòng chữ này lên màn hình console thôi. Chúng ta sẽ thực hiện một chương trình như vậy nhưng sẽ thay chữ 'Hello' bằng 'Go', nào hãy mở trình soạn thảo ( editor ) lên và bắt đầu viết. 
Nếu bạn không quen thuộc với trình soạn thảo thì nó chỉ đơn giản là một file lên và gõ các dòng chữ càn thiết thôi tuy nhiên để  đơn giản tôi gợi ý bạn sử  dụng một số  công cụ soan thảo như [Sublime](https://www.sublimetext.com/) hay [Visual Code](https://code.visualstudio.com/) kèm với plugin/extension sẽ giúp bạn lập trình dễ  dàng hơn.

#### Chương trình đầu tiên

- Lưu ý: Trước khi bắt đầu một chương trình trong Go, hãy thiết lâp biến môi trường GOPATH trước tiên để  tránh việc chương trình không nhận thư mục khi biên dịch.

- Tiếp đó trong thư mục $GOPATH/src, hãy tiến hành tạo một project mới và đặt tên là hellogo.

```
go
|> src
    |> hellogo
```

- Trong thư mục này hãy tạo ra hai file và đặt tên là `hellogo.go` và `hellogo_test.go`, trong đó file `hellogo.go` sẽ chưa logic chương trình:


```
go
|> src
    |> hellogo
            |> hellogo.go
```

Tiếp đó hãy gõ doạn mã sau vào trong file `.go` mà bạn vừa tạo:

```
package main 

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go!");
}

```

Trong đó: 

- package main: mỗi file go mà bạn sử  dụng đều phải bắt đầu cùng với `package` và tên của package đó, chúng sẽ được sử  dụng cho việc tái sử  dụng trong các file, package khác. Ở đây chúng ta dùng `main` là bởi vì chương trình của chúng ta chỉ có một file duy nhất, không phức tạp.

- import: Tiếp đó để  sử  dụng các hàm tiện tích từ các `package` khác bên trong packge, file hiện tại, bạn sử  dụng từ khóa `import` kèm theo đó là tên `package` cần sử  dụng. Lưu ý, bạn không càn phải sử  dụng cặp ngoặc `()` nếu chỉ có một package , tuy nhiên để  dễ  nhìn và gon gang hơn sau này thì tôi đề  nghị bạn lên sử  dụng cặp ngoặc này bất kể  là có một package. Ngoài ra trong Go, các `package` khi sử  dụng với `import` bắt buộc trong file đó phải gọi tới nó ít nhất một lần nếu không trình biên dịch sẽ thông báo lỗi. Ở đây chúng ta khai báo sư  dụng package `fmt`, package này sẽ chưa các hàm tiện ích giúp cho việc nhập xuất ra màn hình dễ  dàng hơn.

- func main(): Cuối cùng giống nhiềù ngôn ngữ lâp trình kiểu tĩnh khác như C, Java, mỗi chương trình luôn cần có một entrypoint cho việc thực thi. Ở đây chúng ta dùng hàm main như một hàm đặc biệt để  thông báo cho chương trình biết dược sẽ bắt đầu thực hiện từ đâu. Lưu ý: hàm main luôn đặt bên trong package main và cặp ngoặc `{}` được dùng để  cho biết phạm vi bắt đầu và kết thúc của một hàm.

- fmt.Println("Hello Go"): như trên chúng ta đã đề  cập việc sử  dụng các hàm tiện ích từ một package. Ở đây chúng ta sử  dụng package fmt và gọi tới hàm Println, hàm này sẽ xuất thông tin ra màn hình.

#### Chạy chương trình

- như vậy chúng ta đa hoàn thành viết một chương trình đơn giản rồi, giờ hãy tiến hành biên dịch chương trình mà bạn vừa viết và thực thi nó:

- Để  biên dịch và chạy chương trình hãy gõ

```
go $GOPATH/src/hellogo/hellogo.go
```

- Nếu không có bất kỳ lỗi gì bạn sẽ nhìn thầy dòng chữ 'Hello Go!' được in ra màn hình ngay lập tức.

#### Tổng kết

- Như vậy trong chương này chúng ta đã tìm hiểu về  viết một chương trình đơn giản rồi,
ở chương tiếp theo chúng ta sẽ đi vào tìm hiểu cách sử  dụng biến ( variable ) trong Go như thế  nào.