# GoMonkey vs 传统Mock - 相同功能对比分析

## 🎯 测试的相同功能

两种方式都测试了**订单创建服务**的相同业务逻辑：
1. 验证客户有效性
2. 处理支付
3. 生成订单ID和时间戳
4. 保存订单到数据库
5. 处理各种错误场景

## 📊 详细对比

### **代码结构对比**

| 方面 | GoMonkey方式 | 传统Mock方式 |
|------|-------------|-------------|
| **服务代码** | `LegacyOrderService` - 硬编码依赖 | `ModernOrderService` - 依赖注入 |
| **测试代码行数** | ~80行 | ~120行 |
| **Mock设置** | 函数级别替换 | 接口级别替换 |
| **依赖管理** | 无需修改原代码 | 需要重构为依赖注入 |

### **测试能力对比**

#### ✅ **GoMonkey的优势**

1. **无需重构原代码**
   ```go
   // 原代码可以保持不变
   func (s *LegacyOrderService) CreateOrder(customerID string, amount float64) (*Order, error) {
       resp, err := http.Get(fmt.Sprintf("https://api.customer.com/validate/%s", customerID))
       // ... 直接调用外部依赖
   }
   ```

2. **可以Mock任何函数**
   ```go
   patches.ApplyFunc(http.Get, mockHttpGet)
   patches.ApplyFunc(time.Now, mockTimeNow)
   patches.ApplyFunc(processPayment, mockPayment)
   ```

3. **快速添加测试覆盖**
   - 对遗留代码立即可测试
   - 不需要大规模重构

#### ✅ **传统Mock的优势**

1. **更清晰的依赖关系**
   ```go
   service := NewModernOrderService(
       mockValidator,    // 明确的依赖
       mockPayment,      // 类型安全
       mockRepo,         // 编译时检查
       mockTime,
       mockIDGen,
   )
   ```

2. **更好的测试验证**
   ```go
   // 可以精确验证调用参数和次数
   mockValidator.AssertExpectations(t)
   mockPayment.AssertNotCalled(t, "ProcessPayment") // 验证未调用
   ```

3. **类型安全**
   - 编译时发现接口变化
   - IDE支持更好
   - 重构更安全

### **错误处理对比**

#### **GoMonkey方式**
```go
func TestLegacyOrderService_CreateOrder_PaymentFailed_WithGoMonkey(t *testing.T) {
    patches.ApplyFunc(processPayment, func(customerID string, amount float64) (*PaymentResponse, error) {
        return &PaymentResponse{Success: false, Message: "Insufficient funds"}, nil
    })
    
    // 测试执行...
    // 🚨 问题：无法验证是否真的调用了processPayment
}
```

#### **传统Mock方式**
```go
func TestModernOrderService_CreateOrder_PaymentFailed_WithTraditionalMock(t *testing.T) {
    mockPayment.On("ProcessPayment", "CUSTOMER-123", 100.50).Return(
        &PaymentResponse{Success: false, Message: "Insufficient funds"}, nil)
    
    // 测试执行...
    
    // ✅ 优势：可以验证精确的调用
    mockValidator.AssertExpectations(t)  // 验证被调用
    mockPayment.AssertExpectations(t)    // 验证被调用
    mockRepo.AssertNotCalled(t, "SaveOrder") // 验证未被调用
}
```

## 🎯 **实际项目中的选择建议**

### **选择GoMonkey的场景：**

1. **遗留系统改造**
   ```go
   // 已有的遗留代码，重构成本太高
   func ProcessLegacyOrder() {
       db := sql.Open("mysql", hardcodedConnectionString)
       resp := http.Get(hardcodedAPIURL)
       // ... 大量硬编码逻辑
   }
   ```

2. **第三方库集成测试**
   ```go
   // 无法控制的第三方库调用
   result := thirdPartyLib.ComplexOperation()
   ```

3. **系统级函数测试**
   ```go
   // 需要控制系统时间、环境变量等
   now := time.Now()
   env := os.Getenv("CONFIG")
   ```

### **选择传统Mock的场景：**

1. **新项目开发**
   ```go
   // 从设计阶段就考虑可测试性
   type OrderService struct {
       validator CustomerValidator
       payment   PaymentProcessor
   }
   ```

2. **可重构的代码**
   ```go
   // 可以改造为依赖注入的代码
   func (s *Service) Process(deps Dependencies) error {
       return s.processor.Process(deps.Data)
   }
   ```

3. **团队协作项目**
   - 需要明确的接口契约
   - 重视代码质量和可维护性

## 🏆 **最佳实践建议**

### **混合使用策略：**

1. **新功能** → 使用依赖注入 + 传统Mock
2. **遗留代码** → 使用GoMonkey快速增加覆盖率
3. **逐步重构** → 将GoMonkey测试的代码逐步改造为依赖注入

### **团队规范：**

```go
// ✅ 推荐：新代码使用依赖注入
type NewService struct {
    deps ServiceDependencies
}

// ⚠️ 可接受：遗留代码使用GoMonkey
func TestLegacyCode(t *testing.T) {
    // TODO: 重构为依赖注入后移除GoMonkey
    patches := gomonkey.NewPatches()
    defer patches.Reset()
    // ...
}
```

## 📈 **性能和维护性对比**

| 指标 | GoMonkey | 传统Mock |
|------|----------|----------|
| **测试执行速度** | 较慢（运行时替换） | 快（编译时确定） |
| **内存使用** | 较高 | 较低 |
| **并发安全** | 需要注意 | 天然安全 |
| **调试难度** | 困难 | 容易 |
| **重构友好** | 差 | 好 |
| **IDE支持** | 一般 | 优秀 |

## 🎯 **结论**

**相同的功能测试，两种方式各有优劣：**

- **GoMonkey**：适合快速给遗留代码添加测试，但长期维护性差
- **传统Mock**：需要前期设计投入，但提供更好的代码质量和可维护性

**最佳策略是根据具体情况选择，并制定明确的使用规范。**
