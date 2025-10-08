# Go Range 操作字符串、映射（map）以及通道实例

这个示例演示了Go语言中range关键字在操作字符串、映射（map）以及通道时的使用方法。

## 运行示例

```bash
cd range-string-map-channel
go run main.go
```

## 示例内容

1. **range操作字符串** - 展示如何使用range正确遍历包含Unicode字符的字符串
2. **range操作映射（map）** - 展示如何使用range遍历map，包括获取键值对、只获取键、只获取值
3. **range操作通道** - 展示如何使用range从通道接收数据直到通道关闭
4. **生产者-消费者模式** - 展示range操作通道在实际应用中的使用

## 代码说明

### 1. range操作字符串

```go
str := "Hello 世界"
for i, char := range str {
    fmt.Printf("字节索引: %d, 字符: %c, Unicode码点: %U\n", i, char, char)
}
```

使用range遍历字符串时，第一个返回值是字节索引，第二个返回值是Unicode码点。这与传统的按字节遍历不同，能正确处理多字节的Unicode字符。

### 2. range操作映射（map）

```go
students := map[string]int{
    "张三": 85,
    "李四": 92,
    "王五": 78,
    "赵六": 96,
}

// 遍历键值对
for name, score := range students {
    fmt.Printf("姓名: %s, 成绩: %d\n", name, score)
}

// 只遍历键
for name := range students {
    fmt.Printf("姓名: %s\n", name)
}

// 只遍历值
for _, score := range students {
    fmt.Printf("成绩: %d\n", score)
}
```

注意：map的遍历顺序是不确定的，Go语言不保证每次遍历的顺序相同。

### 3. range操作通道

```go
ch := make(chan string, 5)

// 发送数据的goroutine
go func() {
    defer close(ch) // 发送完数据后关闭通道
    for i := 1; i <= 5; i++ {
        ch <- fmt.Sprintf("消息%d", i)
    }
}()

// 使用range从通道接收数据，直到通道关闭
for msg := range ch {
    fmt.Printf("接收到: %s\n", msg)
}
```

使用range操作通道会持续接收数据，直到通道被关闭。这是一个非常有用的模式，可以简化通道的读取操作。

### 4. 生产者-消费者模式

示例展示了如何使用range操作通道实现生产者-消费者模式，这是Go语言中常见的并发模式。

## 运行结果示例

```
=== range操作字符串 ===
字符串: Hello 世界
字节索引: 0, 字符: H, Unicode码点: U+0048
字节索引: 1, 字符: e, Unicode码点: U+0065
字节索引: 2, 字符: l, Unicode码点: U+006C
字节索引: 3, 字符: l, Unicode码点: U+006C
字节索引: 4, 字符: o, Unicode码点: U+006F
字节索引: 5, 字符:  , Unicode码点: U+0020
字节索引: 6, 字符: 世, Unicode码点: U+4E16
字节索引: 9, 字符: 界, Unicode码点: U+754C

=== range操作映射（map） ===
学生成绩:
姓名: 张三, 成绩: 85
姓名: 李四, 成绩: 92
姓名: 王五, 成绩: 78
姓名: 赵六, 成绩: 96

=== range操作通道 ===
接收通道中的消息:
接收到: 消息1
接收到: 消息2
接收到: 消息3
接收到: 消息4
接收到: 消息5
通道已关闭，range循环结束
```

## 注意事项

1. **字符串遍历**：range遍历字符串时，索引是字节索引，不是字符索引
2. **Map遍历顺序**：map的遍历顺序是不确定的
3. **通道range**：range操作通道会阻塞直到有数据或通道关闭
4. **关闭通道**：使用range操作通道时，必须在发送端关闭通道以结束循环