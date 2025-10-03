package main

import (
	"errors"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ============= 6. Mock实现（测试环境） =============

// Mock UserRepository - 实现相同的接口
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id string) (*User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// Mock CacheService
type MockCacheService struct {
	mock.Mock
}

func (m *MockCacheService) Set(key string, value interface{}) error {
	args := m.Called(key, value)
	return args.Error(0)
}

func (m *MockCacheService) Get(key string) (interface{}, error) {
	args := m.Called(key)
	return args.Get(0), args.Error(1)
}

// Mock Logger
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(message string) {
	m.Called(message)
}

func (m *MockLogger) Error(message string, err error) {
	m.Called(message, err)
}

// ============= 7. 测试用例 =============

func TestUserService_ProcessUser_Success(t *testing.T) {
	// 创建Mock对象
	mockRepo := new(MockUserRepository)
	mockCache := new(MockCacheService)
	mockLogger := new(MockLogger)
	
	// 设置Mock期望行为
	mockRepo.On("Save", mock.MatchedBy(func(user *User) bool {
		return user.Name == "Alice" && user.Email == "alice@example.com"
	})).Return(nil)  // 模拟保存成功
	
	mockLogger.On("Info", mock.AnythingOfType("string"))
	
	// 创建服务 - 注入Mock实现
	service := NewUserService(mockRepo, mockCache, mockLogger)
	
	// 准备测试数据
	user := &User{
		ID:    "user-123",
		Name:  "Alice",
		Email: "alice@example.com",
	}
	
	// 执行测试
	err := service.ProcessUser(user)
	
	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, 100, user.Score) // 验证纯函数计算结果
	
	// 验证Mock调用
	mockRepo.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}

func TestUserService_ProcessUser_SaveError(t *testing.T) {
	// 创建Mock对象
	mockRepo := new(MockUserRepository)
	mockCache := new(MockCacheService)
	mockLogger := new(MockLogger)
	
	// 设置Mock期望行为 - 模拟保存失败
	saveError := errors.New("database connection failed")
	mockRepo.On("Save", mock.Anything).Return(saveError)
	mockLogger.On("Error", "Failed to save user", saveError)
	
	// 创建服务
	service := NewUserService(mockRepo, mockCache, mockLogger)
	
	// 准备测试数据
	user := &User{
		ID:    "user-123",
		Name:  "Alice",
		Email: "alice@example.com",
	}
	
	// 执行测试
	err := service.ProcessUser(user)
	
	// 验证错误处理
	assert.Error(t, err)
	assert.Equal(t, saveError, err)
	
	// 验证Mock调用
	mockRepo.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}

func TestUserService_ProcessUser_InvalidUser(t *testing.T) {
	// 创建Mock对象（但不会被调用）
	mockRepo := new(MockUserRepository)
	mockCache := new(MockCacheService)
	mockLogger := new(MockLogger)
	
	// 创建服务
	service := NewUserService(mockRepo, mockCache, mockLogger)
	
	// 准备无效的测试数据
	user := &User{
		ID:    "user-123",
		Name:  "", // 无效：空名称
		Email: "alice@example.com",
	}
	
	// 执行测试
	err := service.ProcessUser(user)
	
	// 验证错误处理
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid user")
	
	// 验证没有调用外部依赖（因为验证失败了）
	mockRepo.AssertNotCalled(t, "Save")
	mockLogger.AssertNotCalled(t, "Info")
	mockLogger.AssertNotCalled(t, "Error")
}

// ============= 8. 对比：如果没有接口会怎样？ =============

// ❌ 错误设计：直接依赖具体实现
type BadUserService struct {
	repo   *MySQLUserRepository  // 具体类型，不是接口
	cache  *RedisCache          // 具体类型，不是接口
	logger *StandardLogger      // 具体类型，不是接口
}

// 这样设计的问题：
// 1. 无法在测试中替换为Mock
// 2. 无法在不同环境使用不同实现
// 3. 代码紧耦合，难以维护

// func TestBadUserService(t *testing.T) {
//     // ❌ 无法Mock，必须使用真实的数据库、缓存、日志
//     repo := &MySQLUserRepository{...}  // 需要真实数据库连接
//     cache := &RedisCache{...}          // 需要真实Redis连接
//     logger := &StandardLogger{}        // 会产生真实日志输出
//     
//     service := &BadUserService{repo, cache, logger}
//     // 测试会很慢，不稳定，难以控制
// }
