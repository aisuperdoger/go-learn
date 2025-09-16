# Go语言Mock测试最佳实践

```go
只有当函数有"外部依赖"时才需要接口：
// 典型的服务结构
type UserService struct {
    // 外部依赖 - 需要接口 (30%)
    repo   UserRepository
    cache  CacheService
    logger Logger
}

// 大部分方法包含纯函数逻辑 (70%)
func (s *UserService) ProcessUser(user *User) error {
    // 纯函数部分：不需要接口
    if !s.validateUserData(user) {        // 纯函数
        return errors.New("invalid user")
    }
    
    score := s.calculateUserScore(user)   // 纯函数
    user.Score = score
    
    // 外部依赖部分：需要接口
    return s.repo.Save(user)              // 需要Mock
}
```

## 核心原则

### 1. 面向接口编程
- 定义接口而不是直接依赖具体类型
- 接口应该小而专注（Interface Segregation Principle）
- 遵循"接受接口，返回结构体"的原则，具体含义是：
  - 这样可以接受任何实现了该接口的类型
  - 避免了接口返回值的类型断言问题

### 2. 依赖注入

**依赖注入（Dependency Injection, DI）** 是一种设计模式，它将对象的依赖关系从内部创建转移到外部注入，实现了控制反转（Inversion of Control, IoC）。

#### 2.1 通过构造函数注入依赖

构造函数注入是指在创建对象时，通过构造函数参数将所需的依赖传入，而不是在对象内部创建这些依赖。

**❌ 错误做法：内部创建依赖**
```go
// ❌ 违反依赖注入原则
type BadUserService struct {}

func (s *BadUserService) CreateUser(name, email string) error {
    // ❌ 在函数内部直接创建数据库连接
    db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/mydb")
    if err != nil {
        return err
    }
    defer db.Close()

    // ❌ 在函数内部直接创建日志器
    logger := log.New(os.Stdout, "USER: ", log.LstdFlags)

    // 业务逻辑...
    user := &User{Name: name, Email: email}
    _, err = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
    if err != nil {
        logger.Printf("Failed to save user: %v", err)
        return err
    }

    logger.Printf("User created: %s", user.Name)
    return nil
}

// 问题：
// 1. 无法测试 - 必须连接真实数据库
// 2. 配置硬编码 - 数据库连接字符串写死
// 3. 难以扩展 - 无法替换不同的数据库或日志实现
// 4. 违反单一职责 - 既要管理依赖创建，又要处理业务逻辑
```

**✅ 正确做法：构造函数注入**
```go
// ✅ 遵循依赖注入原则

// 1. 定义接口
type UserRepository interface {
    Save(user *User) error
    GetByID(id string) (*User, error)
}

type Logger interface {
    Info(message string)
    Error(message string, err error)
}

// 2. 服务结构体 - 依赖接口
type UserService struct {
    repo   UserRepository // 依赖接口，不是具体实现
    logger Logger         // 依赖接口，不是具体实现
}

// 3. 构造函数注入 - 核心！
func NewUserService(repo UserRepository, logger Logger) *UserService {
    return &UserService{
        repo:   repo,
        logger: logger,
    }
}

// 4. 业务方法 - 只关注业务逻辑
func (s *UserService) CreateUser(name, email string) error {
    // 验证输入（纯函数）
    if name == "" || email == "" {
        return errors.New("name and email are required")
    }

    // 创建用户对象
    user := &User{
        ID:    generateID(),
        Name:  name,
        Email: email,
    }

    // 使用注入的依赖
    if err := s.repo.Save(user); err != nil {
        s.logger.Error("Failed to save user", err)
        return err
    }

    s.logger.Info("User created successfully: " + user.Name)
    return nil
}
```

**构造函数注入的优势：**
1. **明确依赖**：从构造函数就能看出服务需要哪些依赖
2. **强制注入**：没有依赖就无法创建对象，避免空指针
3. **不可变性**：依赖在创建时确定，运行时不会改变
4. **易于测试**：可以轻松注入Mock实现

#### 2.2 避免在函数内部创建依赖

在函数内部直接创建依赖会导致紧耦合，使代码难以测试和维护。

**❌ 常见的错误模式**
```go
// ❌ 反模式1：在方法内部创建数据库连接
func (s *UserService) GetUser(id string) (*User, error) {
    // ❌ 每次调用都创建新连接
    db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/mydb")
    if err != nil {
        return nil, err
    }
    defer db.Close()
    // 查询逻辑...
}

// ❌ 反模式2：在方法内部创建HTTP客户端
func (s *PaymentService) ProcessPayment(amount float64) error {
    // ❌ 每次都创建新的HTTP客户端
    client := &http.Client{Timeout: 30 * time.Second}
    // 调用第三方API...
}

// ❌ 反模式3：在方法内部获取当前时间
func (s *UserService) CreateUser(name string) error {
    user := &User{
        Name:      name,
        CreatedAt: time.Now(), // ❌ 直接调用time.Now()
    }
    // ...
}
```

**✅ 正确的做法**
```go
// ✅ 正确：通过接口注入所有依赖
type UserService struct {
    repo         UserRepository
    httpClient   HTTPClient
    logger       Logger
    timeProvider TimeProvider  // 连时间都通过接口注入
}

type HTTPClient interface {
    Post(url, contentType string, body io.Reader) (*http.Response, error)
}

type TimeProvider interface {
    Now() time.Time
}

func NewUserService(
    repo UserRepository,
    httpClient HTTPClient,
    logger Logger,
    timeProvider TimeProvider,
) *UserService {
    return &UserService{
        repo:         repo,
        httpClient:   httpClient,
        logger:       logger,
        timeProvider: timeProvider,
    }
}

func (s *UserService) CreateUser(name string) error {
    user := &User{
        Name:      name,
        CreatedAt: s.timeProvider.Now(), // ✅ 通过接口获取时间
    }

    s.logger.Info("Creating user: " + name) // ✅ 使用注入的日志器
    return s.repo.Save(user) // ✅ 使用注入的仓储
}
```

#### 2.3 使用工厂模式或依赖注入容器

对于复杂的应用，手动管理所有依赖关系会变得困难。这时可以使用工厂模式或依赖注入容器来集中管理。

**工厂模式示例**
```go
// ============= 工厂模式示例 =============

// 配置结构
type Config struct {
    DatabaseURL string
    RedisURL    string
    LogLevel    string
}

// 服务工厂
type ServiceFactory struct {
    config *Config
    db     *sql.DB
    redis  *redis.Client
    logger Logger
}

func NewServiceFactory(config *Config) (*ServiceFactory, error) {
    // 创建数据库连接
    db, err := sql.Open("mysql", config.DatabaseURL)
    if err != nil {
        return nil, err
    }

    // 创建Redis连接
    redisClient := redis.NewClient(&redis.Options{
        Addr: config.RedisURL,
    })

    // 创建日志器
    var logger Logger
    switch config.LogLevel {
    case "debug":
        logger = &DebugLogger{}
    case "info":
        logger = &InfoLogger{}
    default:
        logger = &StandardLogger{}
    }

    return &ServiceFactory{
        config: config,
        db:     db,
        redis:  redisClient,
        logger: logger,
    }, nil
}

// 工厂方法：创建UserService
func (f *ServiceFactory) CreateUserService() *UserService {
    repo := &MySQLUserRepository{db: f.db}
    cache := &RedisCache{client: f.redis}

    return NewUserService(repo, cache, f.logger)
}

// 工厂方法：创建OrderService
func (f *ServiceFactory) CreateOrderService() *OrderService {
    repo := &MySQLOrderRepository{db: f.db}
    userService := f.CreateUserService() // 复用其他服务

    return NewOrderService(repo, userService, f.logger)
}

// 使用工厂
func main() {
    config := &Config{
        DatabaseURL: "mysql://user:pass@localhost:3306/mydb",
        RedisURL:    "localhost:6379",
        LogLevel:    "info",
    }

    factory, err := NewServiceFactory(config)
    if err != nil {
        log.Fatal(err)
    }

    // 通过工厂创建服务
    userService := factory.CreateUserService()
    orderService := factory.CreateOrderService()

    // 使用服务...
}
```

**简单的依赖注入容器**
```go
// ============= 简单DI容器示例 =============

type Container struct {
    services  map[string]interface{}
    factories map[string]func() interface{}
}

func NewContainer() *Container {
    return &Container{
        services:  make(map[string]interface{}),
        factories: make(map[string]func() interface{}),
    }
}

// 注册单例服务
func (c *Container) RegisterSingleton(name string, service interface{}) {
    c.services[name] = service
}

// 注册工厂函数
func (c *Container) RegisterFactory(name string, factory func() interface{}) {
    c.factories[name] = factory
}

// 获取服务
func (c *Container) Get(name string) interface{} {
    // 先查找单例
    if service, exists := c.services[name]; exists {
        return service
    }

    // 再查找工厂
    if factory, exists := c.factories[name]; exists {
        return factory()
    }

    return nil
}

// 使用示例
func setupContainer() *Container {
    container := NewContainer()

    // 注册配置
    config := &Config{DatabaseURL: "mysql://..."}
    container.RegisterSingleton("config", config)

    // 注册数据库连接
    container.RegisterSingleton("db", func() *sql.DB {
        db, _ := sql.Open("mysql", config.DatabaseURL)
        return db
    }())

    // 注册日志器
    container.RegisterSingleton("logger", &StandardLogger{})

    // 注册UserRepository工厂
    container.RegisterFactory("userRepo", func() interface{} {
        db := container.Get("db").(*sql.DB)
        return &MySQLUserRepository{db: db}
    })

    // 注册UserService工厂
    container.RegisterFactory("userService", func() interface{} {
        repo := container.Get("userRepo").(UserRepository)
        logger := container.Get("logger").(Logger)
        return NewUserService(repo, logger)
    })

    return container
}

func main() {
    container := setupContainer()

    // 从容器获取服务
    userService := container.Get("userService").(*UserService)

    // 使用服务
    err := userService.CreateUser("Alice", "alice@example.com")
    if err != nil {
        log.Fatal(err)
    }
}
```

**第三方DI框架（推荐）**

对于复杂项目，可以使用成熟的DI框架：

```go
// 使用 github.com/google/wire 示例

//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// 提供者集合
var SuperSet = wire.NewSet(
    // 基础设施
    NewDatabase,
    NewRedisClient,
    NewLogger,

    // 仓储层
    NewUserRepository,
    NewOrderRepository,

    // 服务层
    NewUserService,
    NewOrderService,

    // 控制器层
    NewUserController,
)

// Wire生成的函数
func InitializeUserController() (*UserController, error) {
    wire.Build(SuperSet)
    return &UserController{}, nil
}
```

#### 依赖注入的核心价值

**1. 可测试性**
```go
// 生产环境
func main() {
    repo := &MySQLUserRepository{db: prodDB}
    service := NewUserService(repo, &ProdLogger{})
}

// 测试环境
func TestUserService(t *testing.T) {
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo, &MockLogger{})
    // 完全隔离的单元测试
}
```

**2. 灵活性**
```go
// 可以轻松切换实现
func createUserService(env string) *UserService {
    var repo UserRepository
    switch env {
    case "prod":
        repo = &MySQLUserRepository{}
    case "test":
        repo = &MockUserRepository{}
    case "dev":
        repo = &MemoryUserRepository{}
    }
    return NewUserService(repo, logger)
}
```

**3. 可维护性**
- 依赖关系清晰明确
- 单一职责原则
- 易于重构和扩展

**4. 配置集中化**
- 所有依赖在一个地方配置
- 便于管理不同环境的配置
- 减少重复代码

### 3. 可测试性设计
- 将外部依赖抽象为接口
- 避免全局变量和单例模式
- 使用纯函数尽可能多

## Mock工具选择指南

### 1. GoMock (推荐 - 大型项目)
```bash
go install github.com/golang/mock/mockgen@latest
mockgen -source=interface.go -destination=mocks/mock.go
```
**优势：** 类型安全、性能好、Google维护
**适用：** 大型项目、复杂接口、团队协作

### 2. 手动Mock (推荐 - 简单项目)
```go
type MockClient struct {
    GetFunc func(url string) (*http.Response, error)
}

func (m *MockClient) Get(url string) (*http.Response, error) {
    return m.GetFunc(url)
}
```
**优势：** 简单直接、无额外依赖、完全控制
**适用：** 简单接口、小型项目、学习阶段

### 3. testify/mock (可选 - 中型项目)
```go
type MockDB struct {
    mock.Mock
}

func (m *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
    arguments := m.Called(query, args)
    return arguments.Get(0).(*sql.Rows), arguments.Error(1)
}
```
**优势：** 功能丰富、API友好、灵活断言
**缺点：** 运行时类型检查、性能略低

### 4. 函数式注入 (特定场景)
```go
type Service struct {
    timeProvider func() time.Time
    logger       func(string)
}
```
**适用：** 简单函数依赖、时间/随机数等

## 最佳实践

1. **接口定义在使用方**，而不是实现方
2. **接口应该尽可能小**，遵循单一职责原则
3. **使用依赖注入**，避免在代码中直接创建依赖
4. **为外部依赖创建接口**（数据库、HTTP客户端、文件系统等）
5. **测试中验证行为**，而不仅仅是状态
6. **使用表驱动测试**处理多种场景

## 常见模式

### Repository模式
```go
type UserRepository interface {
    GetByID(id int) (*User, error)
    Save(user *User) error
}
```

### Service模式
```go
type UserService struct {
    repo UserRepository
    logger Logger
}
```

### 适配器模式
```go
type HTTPClientAdapter struct {
    client *http.Client
}

func (a *HTTPClientAdapter) Get(url string) (*http.Response, error) {
    return a.client.Get(url)
}
```

## 运行测试

```bash
cd mock-demo
go mod tidy
go test -v ./...
```
