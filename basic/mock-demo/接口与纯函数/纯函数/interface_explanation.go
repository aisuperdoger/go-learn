// 本文件是接口与纯函数的说明文件

package main

import (
	"errors"
	"fmt"
)

// ============= 1. 定义接口（抽象） =============

// UserRepository 接口 - 这就是"需要接口"的意思
type UserRepository interface {
	Save(user *User) error
	GetByID(id string) (*User, error)
	Delete(id string) error
}

type CacheService interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type Logger interface {
	Info(message string)
	Error(message string, err error)
}

// ============= 2. 业务实体 =============

type User struct {
	ID    string
	Name  string
	Email string
	Score int
}

// ============= 3. 服务层（依赖接口） =============

type UserService struct {
	// 这些字段都是接口类型，不是具体实现
	repo   UserRepository  // 接口，不是具体的数据库实现
	cache  CacheService    // 接口，不是具体的Redis实现
	logger Logger          // 接口，不是具体的日志实现
}

// 构造函数：接受接口参数
func NewUserService(repo UserRepository, cache CacheService, logger Logger) *UserService {
	return &UserService{
		repo:   repo,
		cache:  cache,
		logger: logger,
	}
}

// 业务方法：混合纯函数和外部依赖
func (s *UserService) ProcessUser(user *User) error {
	// 纯函数部分：不需要接口，直接调用
	if !s.validateUserData(user) {        // 纯函数
		return errors.New("invalid user")
	}
	
	score := s.calculateUserScore(user)   // 纯函数
	user.Score = score
	
	// 外部依赖部分：通过接口调用
	// 这里的 s.repo.Save() 调用的是接口方法，不是具体实现
	if err := s.repo.Save(user); err != nil {
		s.logger.Error("Failed to save user", err)
		return err
	}
	
	s.logger.Info(fmt.Sprintf("User %s processed successfully", user.ID))
	return nil
}

// 纯函数：不需要接口，直接实现
func (s *UserService) validateUserData(user *User) bool {
	if user == nil {
		return false
	}
	if user.Name == "" || user.Email == "" {
		return false
	}
	return true
}

// 纯函数：不需要接口，直接实现
func (s *UserService) calculateUserScore(user *User) int {
	// 简化的评分逻辑
	score := len(user.Name) * 10
	if user.Email != "" {
		score += 50
	}
	return score
}

// ============= 4. 具体实现（生产环境） =============

// MySQL实现 - 实现UserRepository接口
type MySQLUserRepository struct {
	connectionString string
}

func (r *MySQLUserRepository) Save(user *User) error {
	// 实际的数据库保存逻辑
	fmt.Printf("Saving user %s to MySQL database\n", user.ID)
	return nil
}

func (r *MySQLUserRepository) GetByID(id string) (*User, error) {
	// 实际的数据库查询逻辑
	fmt.Printf("Getting user %s from MySQL database\n", id)
	return &User{ID: id, Name: "John", Email: "john@example.com"}, nil
}

func (r *MySQLUserRepository) Delete(id string) error {
	// 实际的数据库删除逻辑
	fmt.Printf("Deleting user %s from MySQL database\n", id)
	return nil
}

// Redis缓存实现 - 实现CacheService接口
type RedisCache struct {
	host string
	port int
}

func (c *RedisCache) Set(key string, value interface{}) error {
	fmt.Printf("Setting cache key %s in Redis\n", key)
	return nil
}

func (c *RedisCache) Get(key string) (interface{}, error) {
	fmt.Printf("Getting cache key %s from Redis\n", key)
	return nil, nil
}

// 标准日志实现 - 实现Logger接口
type StandardLogger struct{}

func (l *StandardLogger) Info(message string) {
	fmt.Printf("[INFO] %s\n", message)
}

func (l *StandardLogger) Error(message string, err error) {
	fmt.Printf("[ERROR] %s: %v\n", message, err)
}

// ============= 5. 生产环境使用 =============

func main() {
	// 创建具体实现
	repo := &MySQLUserRepository{connectionString: "mysql://..."}
	cache := &RedisCache{host: "localhost", port: 6379}
	logger := &StandardLogger{}
	
	// 注入到服务中（依赖注入）
	userService := NewUserService(repo, cache, logger)
	
	// 使用服务
	user := &User{
		ID:    "user-123",
		Name:  "Alice",
		Email: "alice@example.com",
	}
	
	err := userService.ProcessUser(user)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
