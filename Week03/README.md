学习笔记
[toc]
## goroutine
注意三件事
- 知道什么时候退出
- 能控制什么时候退
- 超时控制 （利用context）

## memory model
### 内存重排
happen before
在一个 goroutine 中，读和写一定是按照程序中的顺序执行的。即编译器和处理器只有在不会改变这个 goroutine 的行为时才可能修改读和写的执行顺序。由于重排，不同的goroutine 可能会看到不同的执行顺序。例如，一个goroutine 执行 a = 1;b = 2;，另一个 goroutine 可能看到 b 在 a 之前更新。
### mem barrier  ---内存屏障
锁
cpu内存重排 mem内存重排 go mem model


[go内存模型](https://www.jianshu.com/p/5e44168f47a3)

## context
- context尽量显式地放入到函数签名中，而不是放到结构体里

## errgroup
https://pkg.go.dev/golang.org/x/sync/errgroup<br>
Package errgroup provides synchronization, error propagation, and Context cancelation for groups of goroutines working on subtasks of a common task.

- 其中errgroup包中有两个地方会有点坑
1. func WithContext(ctx context.Context) (*Group, context.Context)<br>
返回的context作用域仅仅在WithContext里,使用这个值无实际意义
2. func (g *Group) Go(f func() error) <br>
 这个方法里启用了一个新的goroutine，但是未对可能出现的panic进行recover,因此使用这个函数时需慎用，或者在这个方法的goroutine里手动加入recover使用


## 一些名词
- 染色信息
- COW：copy on write 写时复制

## 计算密集型goroutine生命周期不好管理
IO密集型较容易管理

## chan
channel关闭交给发送者因为channel关闭不影响读
