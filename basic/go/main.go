package go


go func(name string, age int) {
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}("Alice", 25)  // 立即传参并调用