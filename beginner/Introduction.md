## Ngôn ngũ lập trình Go là gi ?

- Golang hay còn được gọi với cái tên quên thuộc Go là một ngôn ngữ lập trình biên dịch kiểu tĩnh ( `compiled - statically typed ` ) được tạo ra bởi Google và lần đầu tiên ₫ược công bố public vào 10/11/2009.

- Google tạo ra ngôn ngữ ngày với mục đích phát triển và tạo ra các hệ thống, ứng dụng web mang tính săn sàng cao ( `high availability` ) và khả năng mở rộng một cách dễ  dàng hơn ( `scability` ).

## Vì sao lại chọn Golang ?

- Hãy để  ý rằng hiện này có hàng tá ngôn ngữ khác, có thể  nhắc đến như Python, Ruby, Java, Scala, Node.js ,etc. Những ngôn ngữ này có lịch sử  phát triển khá lâu và chúng 
được ứng dụng thực tế  ở rát nhiều các công ty, và dùng trong các hệ thống lớn và được cộng đồng hỗ  trợ rất mạnh. Vậy vì sao chúng ta lại chọn Go, một ngôn ngữ có tuổi đời hình thành còn rất trẻ để phát triển cho các service của ban ?, Lí giải cho việc làm này chúng ta sẽ xem xét các tính năng có trong Go sau:

+ Đầu tiên tính `Concurrency` là một phần của Go, mặc dù trong rất nhiều ngôn ngữ khác 
chúng ta có thể  sử  dụng Thread, Process để  tạo ra các chưng trình như vậy nhưng thực tế  là việc quản lí và viết chúng khá phức tạp. Vâng bạn có thể  nói là chúng ta hoàn toàn có thể  dùng cá 3rd module/libs để  hỗ  trợ tính chất này dễ  dàng hơn, nhưng bạn sẽ phải cài đặt thêm, trong khi với Go chúng ta không cần phải làm điều này. Chúng ta có thể  dễ  dàng viết các chương trình đa luồng thông qua goroutine và channel sẵn có trong Go. 
+ Tiêp theo, như đã đề  cập ở đâu chương Go là ngôn ngữ lập trình biên dịch kiểu tĩnh, do đó tất cả các đoạn mã nguồn của Go sẽ được biên dịch sang dạng thuần nhị phần ( native binary ). Trong khi đó với các ngôn ngữ thông dịch khác như Python, Ruby, Node.js không có ( mặc dù bạn có thể  sử  dụng một số  công cụ biên dịch khác nhưng trong phạm vi này chúng ta chỉ nói tới thuần ngôn ngữ hỗ  trợ, không tính tới các 
công cụ bên ngoài ).
+ Các yêu cầu về  mặt kỹ thuật ( Type, Variable, Functions, etc ) trong Go tương đối đơn giản ko có bất kỳ nhiều yêu cầu phức tạp như một số  ngôn ngữ khác ( Node.js với ES6 ). Tất cả chỉ gói gon trong một [trang giấy](https://golang.org/ref/spec) và không có quá nhiều tính chất phức tạp. Để  làm quen với Go bạn không mất quá nhiều thời gian ( Tin tôi đi ).
+ Một điều nữa đó là công cụ biên dịch của Go hỗ  trợ liên kết tính ( static link ), do đó tất cả các mã của một project có thể  liên kết lại thành một file nhị phân duy nhất và có thể  chuyển tới các máy khác nhau hoặc đẩy lên server, cloud dễ  dàng và không cần phải quan tâm tới vấn đề  cài đặt các thư viện, module phụ thuộc.

## Cài đặt

- Go có thể  chạy được trên hầu hết các hệ điều hành hiện này ( OS, Linux, MacOSX ), 
bạn có thể  lên trang chủ của Go và tải các file nhị phận chứ công cụ Go đã được biện dịch sẵn về  tại https://golang.org/dl/ . Hay để  tién hành cài đặt go trên các hệ đièu hành khác nhau ta thực hiện như sau:

### Đối với Linux 

#### Đối với hệ điều hành Linux nói chung 

- Bạn có thể  sử  dụng [gvm](https://github.com/moovweb/gvm) một công cũ quản lí các phiên bản của go.

#### Đối với file nhị phân 

- Bạn có thể  sử  dụng `yum` với Centos hay `apt` với Ubuntu và `wget` để  cài đặt:

- với Centos sử  dụng

```
yum clean all

yum update
```

- với Ubuntu sử  dụng 

```
apt-get update

apt-get -y upgrade
```

- Tiếp đó hãy tải phiên bản go mà bạn muón sử  dụng về 


```
wget https://dl.google.com/go/go_version.linux-amd64.tar.gz
```

- sau khi đã tải xong, hãy giải nén file này ra và chuyển nó tới folder bạn muốn chứa, ở đây toi sủ dụng /usr/local làm thư mục chưa công cụ biên dịch cho go. Nếu không thích bạn hoàn toàn có thể  đặt ở thư mục nào khác tùy bạn.


```
tar -xvf go_version.linux-amd64.tar.gz

mv go /usr/local
```

##### bước tiếp theo chúng ta sẽ tiến hành thiết lập môi trường

- Hãy tạo một biến môi trường có tên là GOROOT và gán nó tới vị trí thư mục của go mà bạn vừa tạo:

```
export GOROOT=/usr/local/go
```

- Tiếp đó chúng ta sẽ thiết lập đường dẫn tới project mà chúng ta sẽ sử  dụng để  viết chương trinh go:

```
export GOPATH=$HOME/Projects/Proj1
```

- cuối cùng để  có thể  nhận được thông tin về  thư mục go và sử  dụng trong hệ thống hãy gán lại biến PATH

```
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

- Cuối cùng để  tránh mất thông tin về  vị trí của thư mục go và sử  dụng cho các lần dăng nhập tiếp theo hãy đặt các đoạn mã trên vào trong file ~/.profile hay ~/.bashrc.

- Hãy kiểm tra xem là chúng ta đã thiết lập chính xác chưa, và go hoạt động chưa nếu chưa hãy làm lại các bước trên. Tiếp theo sau khi đã cài đặt thành công chúng ta 
sẽ tiến hành viết một chương trình làm quen với Go.
