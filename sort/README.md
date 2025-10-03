在 Go 语言中，`sort.Ints(nums)` 是对整数切片进行排序的便捷方法之一。除了 `sort.Ints`，Go 的 `sort` 包还提供了多种排序方法，适用于不同类型和场景。以下是主要的排序方法：

---

### 1. **预定义类型的排序函数**

Go 为常见类型提供了专门的排序函数：

- `sort.Ints([]int)`：对整数切片排序（升序）
- `sort.Float64s([]float64)`：对 float64 切片排序
- `sort.Strings([]string)`：对字符串切片排序

示例：

```go
nums := []int{3, 1, 4, 1, 5}
sort.Ints(nums) // [1 1 3 4 5]

names := []string{"Bob", "Alice", "Charlie"}
sort.Strings(names) // [Alice Bob Charlie]
```

---

### 2. **自定义排序：`sort.Slice()`（推荐）**

Go 1.8+ 引入了 `sort.Slice`，可以对任意切片进行自定义排序，非常方便。

```go
people := []struct {
    Name string
    Age  int
}{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
}

// 按年龄升序排序
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
```

---

### 3. **使用 `sort.Sort()` 实现 `Interface` 接口**

如果需要更复杂的排序逻辑，可以实现 `sort.Interface` 接口（`Len`, `Less`, `Swap`）。

`sort.Interface` 接口定义如下：
```
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

各方法说明：
- `Len() int`：返回集合长度
- `Less(i, j int) bool`：比较两个元素，决定排序顺序（返回true表示i<j，即升序）
- `Swap(i, j int)`：交换两个元素的位置

完整示例：
```go
type Person struct {
    Name string
    Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }
    
    sort.Sort(ByAge(people))
    // 结果: Bob(25), Alice(30), Charlie(35)
}
```

还可以实现多重排序条件：
```go
type ByAgeThenName []Person

func (a ByAgeThenName) Len() int { return len(a) }
func (a ByAgeThenName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAgeThenName) Less(i, j int) bool {
    if a[i].Age != a[j].Age {
        return a[i].Age < a[j].Age  // 年龄升序
    }
    return a[i].Name < a[j].Name    // 年龄相同时按姓名升序
}
```

---

### 4. **降序排序**

- 使用 `sort.Reverse()` 包装：
  ```go
  sort.Sort(sort.Reverse(sort.IntSlice(nums)))
  ```
- 或在 `sort.Slice` 中反转比较逻辑：
  ```go
  sort.Slice(nums, func(i, j int) bool {
      return nums[i] > nums[j] // 降序
  })
  ```

---

### 5. **稳定排序：`sort.Stable()`**

当需要保持相等元素的原始顺序时，使用 `sort.Stable()`：

```go
sort.Stable(ByAge(people))
```

---

### 总结

| 方法 | 适用场景 |
|------|----------|
| `sort.Ints`, `sort.Strings` 等 | 基本类型切片 |
| `sort.Slice()` | 任意切片，自定义比较逻辑（最常用） |
| `sort.Sort()` + 接口实现 | 复杂排序逻辑或需复用排序规则 |
| `sort.Stable()` | 需要稳定排序时 |
| `sort.Reverse()` | 降序排列 |

**推荐**：对于大多数场景，使用 `sort.Slice` 是最简洁高效的方式。


# 数组可以使用这些函数吗
在 Go 语言中，标准库的 sort 包函数（如 sort.Ints、sort.Slice 等）不能直接用于数组，但可以通过切片（slice）的方式间接使用。
```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // 定义一个数组
    nums := [5]int{3, 1, 4, 1, 5}

    // 将数组转为切片，然后排序
    sort.Ints(nums[:])

    fmt.Println(nums) // 输出: [1 1 3 4 5]
}
```

