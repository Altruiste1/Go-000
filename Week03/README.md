学习笔记
## 超时控制

服务端埋点

listdirectory 的三种方式

# 避免goroutine泄漏
注意三件事
- 知道什么时候退出
- 能控制什么时候退
- 超时控制

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

## 一些名词
- 染色信息
- COW：copy on write 写时复制

## 计算密集型goroutine生命周期不好管理
IO密集型较容易管理

## chan
channel关闭交给发送者因为channel关闭不影响读
