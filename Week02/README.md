学习笔记
[toc]
# error
## panic和error的使用场景
panic会造成进程被终止，只有在极少数场景需要使用到，一般需要用到panic的场景有
- main函数初始化失败，需要panic
- 项目启动初始化配置文件时，初始化配置文件失败，需要panic
- 项目启动初始化db时，初始化失败，需要panic

error的场景
能预期到的结果，虽然并不是想要的，使用error
## error的设计
- sentinel error
- Error types
- Opaque errors
这是常见的几种error设计
sentinel error可以预定义一些特定值作为error使用，使用方便，调用起来也很简单，他会成为api的一部分，并且在两个包之间形成依赖，当error需要重构时，给调用者带来麻烦

### Error Types
Error type 是实现了 error 接口的自定义类型。例如 MyError 类型记录了文件和行号以展示发生了什么,
相比sentinel error，它保存了更多的上下文信息
但是想要获取完整的错误信息就必须使用断言判断error的正确类型，这使得error自定义的error暴露出来，使Api变得脆弱
建议：是避免错误类型，或者至少避免将它们作为公共 API 的一部分。

### 最灵活的error处理策略Opaque errors
不透明错误处理，因为虽然您知道发生了错误，但您没有能力看到错误的内部。作为调用者，关于操作的结果，您所知道的就是它起作用了，或者没有起作用(成功还是失败)。

## 获取error的堆栈信息
github.com/pkg/errors
errors.Wrap(err,"error msg") //保存错误的堆栈信息
errors.WithMessage(err,"failed") //添加错误信息
errors.UnWrap(err)//剥离外层的error
errors.Cause(err)  // 返回最原始的错误

