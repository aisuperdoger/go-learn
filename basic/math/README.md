在 Go 语言中，常见的数学运算函数主要位于 `math` 包中。以下是一些常用的数学运算函数：

### 1. 绝对值 (Absolute Value)
*   **`math.Abs(x float64) float64`**: 返回 `x` 的绝对值。
    ```go
    package main

    import (
        "fmt"
        "math"
    )

    func main() {
        fmt.Println(math.Abs(-5.5)) // 输出: 5.5
        fmt.Println(math.Abs(3.0))  // 输出: 3
    }
    ```

### 2. 幂运算 (Power)
*   **`math.Pow(x, y float64) float64`**: 返回 `x` 的 `y` 次幂 (`x^y`)。
    ```go
    fmt.Println(math.Pow(2, 3))  // 输出: 8 (2的3次方)
    fmt.Println(math.Pow(4, 0.5)) // 输出: 2 (4的平方根)
    ```

### 3. 平方根 (Square Root)
*   **`math.Sqrt(x float64) float64`**: 返回 `x` 的平方根。如果 `x` 是负数，返回 `NaN`。
    ```go
    fmt.Println(math.Sqrt(16)) // 输出: 4
    fmt.Println(math.Sqrt(2))  // 输出: 1.4142135623730951
    ```


### 7. 四舍五入与取整
*   **`math.Round(x float64) float64`**: 四舍五入到最接近的整数。
*   **`math.Floor(x float64) float64`**: 向下取整（返回小于或等于 `x` 的最大整数）。
*   **`math.Ceil(x float64) float64`**: 向上取整（返回大于或等于 `x` 的最小整数）。
    ```go
    fmt.Println(math.Round(3.7))  // 输出: 4
    fmt.Println(math.Round(3.2))  // 输出: 3
    fmt.Println(math.Floor(3.7))  // 输出: 3
    fmt.Println(math.Ceil(3.2))   // 输出: 4
    ```

### 8. 最大值与最小值
*   **`math.Max(x, y float64) float64`**: 返回 `x` 和 `y` 中的最大值。
*   **`math.Min(x, y float64) float64`**: 返回 `x` 和 `y` 中的最小值。
    ```go
    fmt.Println(math.Max(10, 5)) // 输出: 10
    fmt.Println(math.Min(10, 5)) // 输出: 5
    ```

### int最大值
推荐方法：使用 math.MaxInt 和 math.MinInt (Go 1.17+)
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    maxInt := math.MaxInt
    minInt := math.MinInt

    fmt.Printf("int 的最大值: %d\n", maxInt) // 在 64 位系统上输出: 9223372036854775807
    fmt.Printf("int 的最小值: %d\n", minInt) // 在 64 位系统上输出: -9223372036854775808
}
```

### 注意事项
*   `math` 包中的大多数函数都操作 `float64` 类型。如果你有 `int` 或其他类型的变量，需要先进行类型转换。
*   进行浮点数运算时，要注意精度问题。
*   记得在代码中导入 `math` 包：`import "math"`。