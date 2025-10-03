
当 panic 被调用时，当前函数的执行会立即停止，然后开始执行该函数中已经 defer 的函数。之后，panic 会向上传播到调用栈的上一层函数，重复这个过程，直到整个 goroutine 终止。
Go 不支持传统的 try-catch 异常机制，但可以通过 defer + recover 来捕获 panic，从而防止程序崩溃。

最佳实践建议：
- 尽量使用 error 而不是 panic 来处理常规错误。
- 只在真正“不可能”发生的情况下使用 panic。
- 在库代码中应避免使用 panic，而应返回 error。
- 在 main 函数或 goroutine 的最外层可以使用 defer + recover 做兜底处理，防止程序意外退出。