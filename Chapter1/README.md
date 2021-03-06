## 五个例子

go 语言的安装，环境设置，构建命令等。

查看安装的 go 版本： `go version`

设置 GOPATH：`export GOPATH＝$HOME/Hello`

go 编译构建：`go build`

运行 go 程序： `./Hello`

可以通过 `go install` 命令 把生成的 go 程序安装到一个标准的路径中（如 `$GOPATH/bin`)。

这意味着，只需要把这个标准路径（`$GOPATH/bin`）加入到 `PATH` 环境变量中，那么我们安装的所有 go 程序都可以在任意路径运行。

go 生成的程序是 native 的，可以 copy 到没有 go 环境的系统运行。这对于快速部署一些工具特别有用，节约大量的运维成本。

这也是 go 程序的优势所在。

### Hello Who?

通过 `hello.go` 这个程序介绍了 go 语言的代码结构，风格及一些语言特性。

比如开始是个 pachage，接着是 import，可运行的程序必须有 main 函数等。

go 的代码风格是极简的，如每行语句不需要分号，控制结构体（ 如 if 语句和 for
 语句）也不需要圆括号，控制体和函数体均使用花括号（｛ ｝）作为边界符。

 `:=` 是快速变量声明操作符

 `＝` 是赋值操作符

### 大数字——二维切片(slice)

 指出了在函数和方法外部不能使用 `:=` 操作符声明变量，而是通过使用 `var` 关键字和 `＝` 赋值运算符的长声明方式。

 `for range` 的使用。

 char（byte） 转化为 int 的方式：把 char 和 ‘0’ 相减。

### 栈——自定义类型及其方法

每个本地包都需要保存在一个与包名同名的目录下。本地包可以包含子包（如 `path／filepath`）。

一种 for 循环的使用方式。

函数可以返回多值。一个常用的编程范式就是函数或方法可以返回一个正常值和一个错误值（其类型为error）。

自定义类型可以定义方法。方法是一种带有接收器（receiver）的函数。此接收器为自定义类型。

如：

```go
func (stack Stack) Len() int {
    return len(stack)
}
```
指针
>通过在类型名字前面添加一个星号（即星号 *）来声明指针。指针是一个保存了另一个值的内存地址的变量。

接收器声明为一个**指针类型**的主要原因是为了修改此接收器(另外，对复杂自定义类型来说也为了效率)。如下方法：
```go
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}
```
变量 stack 的类型为 *Stack，也就是说变量 stack 保存了一个指向Stack类型值的指针。

通过`解引用`可以获取该指针所指向内存的实际 Stack 值，`解引用`操作通过在变量前加一个星号来完成。

所以，我们写 stack 时，表示 Stack的指针（即 *Stack）。写 *stack 时，表示该指针所指的实际 Stack 值。

星号不同位置的不同含义：
* 在两个数字之间表示乘法，如 x * y。
* 在类型名称前面表示指针，如 *MyType。
* 在变量名称之前表示解引用，如 *V。

`nil` 表示空指针，以及空引用。

### americanise 示例——文件，映射和闭包

这个例子主要讲的知识点为：

* `defer` 的作用：在函数调用结束前执行的语句。`defer` 可以延迟语句，也可以延长函数。
* `err != nil` 这种处理错误的范式。
* 函数作为第一类值可以作为参数传递，也可以作为返回值。并支持闭包。即函数可以作为一个变量声明。
* IO 相关的函数。
* map 数据类型。并强调：map，slice，channel 都必须通过 `make()` 函数创建。

### 从极坐标到笛卡尔坐标——并发

介绍了 go 语言强大的并发机制。channel 和 goroutine

channel 可以 传递数据 和 实现线程同步。

goroutine 使用 go 语句实现异步函数调用。

使用 make 创建 channel：

```
messages := make(chan string, 10)
```
第一个参数表示接收的类型为 字符串，第二个参数表示缓冲区大小为 10。

发送数据到 channel：
```
messages <- "Leader"
messages <- "Follower"
```
当 `<-` 操作符用做二元操作符时，它的左操作数必须是一个 channel，右操作数必须是发往该 channel 的数据。

当 `<-` 操作符用作一元操作符时，它的右边必须是个channel，表示从这个 channel 接收一个值:

```
message1 := <-messages
message2 := <-messages
```

channel 和 goroutine 的知识第七章会详述。