package main

import (
	"errors"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ============= Mock实现 =============

type MockCustomerValidator struct {
	mock.Mock
}

func (m *MockCustomerValidator) ValidateCustomer(customerID string) error {
	args := m.Called(customerID)
	return args.Error(0)
}

type MockPaymentProcessor struct {
	mock.Mock
}

func (m *MockPaymentProcessor) ProcessPayment(customerID string, amount float64) (*PaymentResponse, error) {
	args := m.Called(customerID, amount)
	return args.Get(0).(*PaymentResponse), args.Error(1)
}

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) SaveOrder(order *Order) error {
	args := m.Called(order)
	return args.Error(0)
}

type MockTimeProvider struct {
	mock.Mock
}

func (m *MockTimeProvider) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

type MockIDGenerator struct {
	mock.Mock
}

func (m *MockIDGenerator) GenerateOrderID() string {
	args := m.Called()
	return args.String(0)
}

// ============= 使用传统Mock测试ModernOrderService =============

func TestModernOrderService_CreateOrder_Success_WithTraditionalMock(t *testing.T) {
	// 创建所有Mock
	mockValidator := new(MockCustomerValidator)
	mockPayment := new(MockPaymentProcessor)
	mockRepo := new(MockOrderRepository)
	mockTime := new(MockTimeProvider)
	mockIDGen := new(MockIDGenerator)
	
	// 设置期望行为
	fixedTime := time.Date(2023, 12, 25, 10, 0, 0, 0, time.UTC)
	
	mockValidator.On("ValidateCustomer", "CUSTOMER-123").Return(nil)
	mockPayment.On("ProcessPayment", "CUSTOMER-123", 100.50).Return(
		&PaymentResponse{Success: true, Message: "Payment successful"}, nil)
	mockTime.On("Now").Return(fixedTime)
	mockIDGen.On("GenerateOrderID").Return("ORDER-12345")
	mockRepo.On("SaveOrder", mock.MatchedBy(func(order *Order) bool {
		return order.ID == "ORDER-12345" &&
			order.CustomerID == "CUSTOMER-123" &&
			order.Amount == 100.50 &&
			order.CreatedAt.Equal(fixedTime) &&
			order.Status == "PAID"
	})).Return(nil)
	
	// 创建服务
	service := NewModernOrderService(mockValidator, mockPayment, mockRepo, mockTime, mockIDGen)
	
	// 执行测试
	order, err := service.CreateOrder("CUSTOMER-123", 100.50)
	
	// 验证结果
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, "ORDER-12345", order.ID)
	assert.Equal(t, "CUSTOMER-123", order.CustomerID)
	assert.Equal(t, 100.50, order.Amount)
	assert.Equal(t, fixedTime, order.CreatedAt)
	assert.Equal(t, "PAID", order.Status)
	
	// 验证所有Mock调用
	mockValidator.AssertExpectations(t)
	mockPayment.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockTime.AssertExpectations(t)
	mockIDGen.AssertExpectations(t)
}

func TestModernOrderService_CreateOrder_CustomerValidationFailed_WithTraditionalMock(t *testing.T) {
	// 创建Mock
	mockValidator := new(MockCustomerValidator)
	mockPayment := new(MockPaymentProcessor)
	mockRepo := new(MockOrderRepository)
	mockTime := new(MockTimeProvider)
	mockIDGen := new(MockIDGenerator)
	
	// 设置客户验证失败
	mockValidator.On("ValidateCustomer", "INVALID-CUSTOMER").Return(errors.New("customer not found"))
	
	// 创建服务
	service := NewModernOrderService(mockValidator, mockPayment, mockRepo, mockTime, mockIDGen)
	
	// 执行测试
	order, err := service.CreateOrder("INVALID-CUSTOMER", 100.50)
	
	// 验证错误处理
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Contains(t, err.Error(), "customer validation failed")
	
	// 验证只调用了验证器，其他服务没有被调用
	mockValidator.AssertExpectations(t)
	mockPayment.AssertNotCalled(t, "ProcessPayment")
	mockRepo.AssertNotCalled(t, "SaveOrder")
}

func TestModernOrderService_CreateOrder_PaymentFailed_WithTraditionalMock(t *testing.T) {
	// 创建Mock
	mockValidator := new(MockCustomerValidator)
	mockPayment := new(MockPaymentProcessor)
	mockRepo := new(MockOrderRepository)
	mockTime := new(MockTimeProvider)
	mockIDGen := new(MockIDGenerator)
	
	// 设置期望行为
	mockValidator.On("ValidateCustomer", "CUSTOMER-123").Return(nil)
	mockPayment.On("ProcessPayment", "CUSTOMER-123", 100.50).Return(
		&PaymentResponse{Success: false, Message: "Insufficient funds"}, nil)
	
	// 创建服务
	service := NewModernOrderService(mockValidator, mockPayment, mockRepo, mockTime, mockIDGen)
	
	// 执行测试
	order, err := service.CreateOrder("CUSTOMER-123", 100.50)
	
	// 验证错误处理
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Contains(t, err.Error(), "payment rejected")
	assert.Contains(t, err.Error(), "Insufficient funds")
	
	// 验证调用顺序：验证和支付被调用，但保存没有被调用
	mockValidator.AssertExpectations(t)
	mockPayment.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "SaveOrder")
}

func TestModernOrderService_CreateOrder_DatabaseError_WithTraditionalMock(t *testing.T) {
	// 创建Mock
	mockValidator := new(MockCustomerValidator)
	mockPayment := new(MockPaymentProcessor)
	mockRepo := new(MockOrderRepository)
	mockTime := new(MockTimeProvider)
	mockIDGen := new(MockIDGenerator)
	
	// 设置期望行为
	fixedTime := time.Date(2023, 12, 25, 10, 0, 0, 0, time.UTC)
	
	mockValidator.On("ValidateCustomer", "CUSTOMER-123").Return(nil)
	mockPayment.On("ProcessPayment", "CUSTOMER-123", 100.50).Return(
		&PaymentResponse{Success: true, Message: "Payment successful"}, nil)
	mockTime.On("Now").Return(fixedTime)
	mockIDGen.On("GenerateOrderID").Return("ORDER-12345")
	mockRepo.On("SaveOrder", mock.Anything).Return(errors.New("database connection failed"))
	
	// 创建服务
	service := NewModernOrderService(mockValidator, mockPayment, mockRepo, mockTime, mockIDGen)
	
	// 执行测试
	order, err := service.CreateOrder("CUSTOMER-123", 100.50)
	
	// 验证错误处理
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Contains(t, err.Error(), "failed to save order")
	
	// 验证所有步骤都被调用了
	mockValidator.AssertExpectations(t)
	mockPayment.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockTime.AssertExpectations(t)
	mockIDGen.AssertExpectations(t)
}
