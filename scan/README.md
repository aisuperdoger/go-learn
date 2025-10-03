在Go语言中，标准输入（`stdin`）是程序与用户交互的重要方式。Go语言提供了多种从标准输入读取数据的方法，主要集中在 `fmt` 和 `bufio` 包中。以下是常用的几种方法及其特点：

### 1. 使用 `fmt.Scan` 系列函数

这是最简单直接的方法，适合读取格式化的输入。

*   **`fmt.Scan(&a, &b, ...)`**
    *   **功能**：从标准输入读取数据，以空格、制表符或换行符作为分隔符。
    *   **特点**：
        *   按空格分割，不能读取包含空格的字符串。
        *   读取成功后返回读取的项目数量和错误信息。
        *   如果输入格式不匹配（如期望整数但输入了字母），会报错。
    *   **示例**：
        ```go
        var name string
        var age int
        fmt.Print("Enter name and age: ")
        fmt.Scan(&name, &age)
        fmt.Printf("Name: %s, Age: %d\n", name, age)
        ```

*   **`fmt.Scanf(format, &a, &b, ...)`**
    *   **功能**：根据指定的格式字符串读取输入，功能更强大。
    *   **特点**：可以指定输入的格式，例如 `fmt.Scanf("%s %d", &name, &age)`。
    *   **示例**：
        ```go
        var x, y int
        fmt.Scanf("%d,%d", &x, &y) // 输入 "10,20"
        ```

*   **`fmt.Scanln(&a, &b, ...)`**
    *   **功能**：与 `Scan` 类似，但在遇到换行符时停止扫描。
    *   **特点**：只能读取一行中的数据，遇到换行即结束。

### 2. 使用 `bufio.Scanner` (推荐)

这是最常用且推荐的方法，特别适合按行读取输入。在处理大量输入数据或对性能要求高的场景下，fmt.Scan 的 I/O 效率太低，容易成为性能瓶颈导致超时；而 bufio 通过缓冲机制提供了高效的 I/O 操作，是更优的选择。
-  bufio先从操作系统一次性读取一大块数据到内存中的缓冲区（buffer）。后续的读取操作（如 Scanner.Scan()）直接从这个快速的内存缓冲区中获取数据。

*   **`bufio.NewScanner(os.Stdin)`**
    *   **功能**：创建一个从标准输入读取的扫描器。
    *   **特点**：
        *   **高效**：内部使用缓冲，性能好。
        *   **简洁**：API 简单，`Scan()` 读取一行，`Text()` 获取内容。
        *   **灵活**：默认按行分割，但可以通过 `Split` 函数自定义分隔符。
        *   **安全**：自动处理换行符。
    *   **示例**：
        ```go
        package main

        import (
            "bufio"
            "fmt"
            "os"
        )

        func main() {
            scanner := bufio.NewScanner(os.Stdin)
            fmt.Print("Enter your name: ")
            if scanner.Scan() {
                name := scanner.Text() // 获取整行，包括空格
                fmt.Printf("Hello, %s!\n", name)
            }
            // 检查扫描过程中是否出错
            if err := scanner.Err(); err != nil {
                fmt.Fprintln(os.Stderr, "reading standard input:", err)
            }
        }
        ```

### 3. 使用 `bufio.Reader`

提供更底层的控制，适合需要精确控制读取过程的场景。

*   **`bufio.NewReader(os.Stdin)`**
    *   **常用方法**：
        *   **`ReadString(delim byte)`**：读取直到遇到指定的分隔符（如 `'\n'`）。
        *   **`ReadLine()`**：低级行读取，返回字节切片和是否行过长的标志。
        *   **`ReadBytes(delim byte)`**：读取直到分隔符，返回字节切片。
    *   **特点**：
        *   比 `Scanner` 更灵活，但代码稍复杂。
        *   返回的是 `[]byte`，通常需要转换为 `string`。
    *   **示例**：
        ```go
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter text: ")
        text, _ := reader.ReadString('\n')
        // text 包含换行符，通常需要去除
        text = strings.TrimSpace(text)
        ```

### 总结与选择建议

| 方法 | 适用场景 | 优点 | 缺点 |
| :--- | :--- | :--- | :--- |
| `fmt.Scan` / `Scanf` | 读取简单的、以空格分隔的格式化数据 | 代码简洁，类似C的`scanf` | 无法处理空格，错误处理不友好 |
| `bufio.Scanner` | **按行读取**，处理文本输入 | **高效、安全、推荐**，API 简洁 | 默认按行，需额外处理分割 |
| `bufio.Reader` | 需要精确控制读取过程 | 灵活性最高 | 代码相对复杂，需手动处理 |

**最佳实践**：
*   对于大多数情况，特别是需要读取整行（可能包含空格）的输入，**强烈推荐使用 `bufio.Scanner`**。
*   对于简单的、格式固定的输入（如两个用空格分开的数字），可以使用 `fmt.Scan`。
*   当需要处理特殊分隔符或进行低级I/O操作时，选择 `bufio.Reader`。