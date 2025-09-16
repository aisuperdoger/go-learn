package main

import (
	"database/sql"
	"net/http"
	"testing"
	"time"
	
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

// ============= 使用GoMonkey测试LegacyOrderService =============

func TestLegacyOrderService_CreateOrder_Success_WithGoMonkey(t *testing.T) {
	service := &LegacyOrderService{}
	
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	
	// Mock HTTP客户验证请求
	patches.ApplyFunc(http.Get, func(url string) (*http.Response, error) {
		// 模拟成功的客户验证响应
		return &http.Response{StatusCode: 200}, nil
	})
	
	// Mock支付处理
	patches.ApplyFunc(processPayment, func(customerID string, amount float64) (*PaymentResponse, error) {
		return &PaymentResponse{Success: true, Message: "Payment successful"}, nil
	})
	
	// Mock时间
	fixedTime := time.Date(2023, 12, 25, 10, 0, 0, 0, time.UTC)
	patches.ApplyFunc(time.Now, func() time.Time {
		return fixedTime
	})
	
	// Mock ID生成
	patches.ApplyFunc(generateOrderID, func() string {
		return "ORDER-12345"
	})
	
	// Mock数据库保存
	patches.ApplyFunc(saveOrderToDB, func(order *Order) error {
		return nil // 模拟成功保存
	})
	
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
}

func TestLegacyOrderService_CreateOrder_CustomerValidationFailed_WithGoMonkey(t *testing.T) {
	service := &LegacyOrderService{}
	
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	
	// Mock HTTP客户验证请求 - 返回错误状态
	patches.ApplyFunc(http.Get, func(url string) (*http.Response, error) {
		return &http.Response{StatusCode: 404}, nil // 客户不存在
	})
	
	// 执行测试
	order, err := service.CreateOrder("INVALID-CUSTOMER", 100.50)
	
	// 验证错误处理
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Contains(t, err.Error(), "invalid customer")
}

func TestLegacyOrderService_CreateOrder_PaymentFailed_WithGoMonkey(t *testing.T) {
	service := &LegacyOrderService{}
	
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	
	// Mock成功的客户验证
	patches.ApplyFunc(http.Get, func(url string) (*http.Response, error) {
		return &http.Response{StatusCode: 200}, nil
	})
	
	// Mock支付失败
	patches.ApplyFunc(processPayment, func(customerID string, amount float64) (*PaymentResponse, error) {
		return &PaymentResponse{Success: false, Message: "Insufficient funds"}, nil
	})
	
	// 执行测试
	order, err := service.CreateOrder("CUSTOMER-123", 100.50)
	
	// 验证错误处理
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Contains(t, err.Error(), "payment rejected")
	assert.Contains(t, err.Error(), "Insufficient funds")
}

func TestLegacyOrderService_CreateOrder_DatabaseError_WithGoMonkey(t *testing.T) {
	service := &LegacyOrderService{}
	
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	
	// Mock成功的前置条件
	patches.ApplyFunc(http.Get, func(url string) (*http.Response, error) {
		return &http.Response{StatusCode: 200}, nil
	})
	
	patches.ApplyFunc(processPayment, func(customerID string, amount float64) (*PaymentResponse, error) {
		return &PaymentResponse{Success: true, Message: "Payment successful"}, nil
	})
	
	patches.ApplyFunc(time.Now, func() time.Time {
		return time.Date(2023, 12, 25, 10, 0, 0, 0, time.UTC)
	})
	
	patches.ApplyFunc(generateOrderID, func() string {
		return "ORDER-12345"
	})
	
	// Mock数据库错误
	patches.ApplyFunc(saveOrderToDB, func(order *Order) error {
		return sql.ErrConnDone // 模拟数据库连接错误
	})
	
	// 执行测试
	order, err := service.CreateOrder("CUSTOMER-123", 100.50)
	
	// 验证错误处理
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Contains(t, err.Error(), "failed to save order")
}
