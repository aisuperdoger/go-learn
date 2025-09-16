package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

// 共同的业务实体
type Order struct {
	ID         string
	CustomerID string
	Amount     float64
	CreatedAt  time.Time
	Status     string
}

// 第三方支付API响应
type PaymentResponse struct {
	Success bool
	Message string
}

// ============= 版本1: 传统设计（难以测试，需要GoMonkey）=============

type LegacyOrderService struct {
	// 没有依赖注入，所有依赖都硬编码
}

func (s *LegacyOrderService) CreateOrder(customerID string, amount float64) (*Order, error) {
	// 1. 验证客户 - 直接调用HTTP API
	resp, err := http.Get(fmt.Sprintf("https://api.customer.com/validate/%s", customerID))
	if err != nil {
		return nil, fmt.Errorf("customer validation failed: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid customer")
	}
	
	// 2. 处理支付 - 直接调用支付API
	paymentResp, err := processPayment(customerID, amount)
	if err != nil {
		return nil, fmt.Errorf("payment failed: %w", err)
	}
	
	if !paymentResp.Success {
		return nil, fmt.Errorf("payment rejected: %s", paymentResp.Message)
	}
	
	// 3. 创建订单 - 直接使用系统时间
	order := &Order{
		ID:         generateOrderID(),
		CustomerID: customerID,
		Amount:     amount,
		CreatedAt:  time.Now(), // 硬编码时间
		Status:     "PAID",
	}
	
	// 4. 保存到数据库 - 直接创建数据库连接
	if err := saveOrderToDB(order); err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}
	
	return order, nil
}

// 第三方支付处理函数
func processPayment(customerID string, amount float64) (*PaymentResponse, error) {
	// 模拟调用外部支付API
	resp, err := http.Post("https://payment.api.com/charge", "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	// 简化处理
	return &PaymentResponse{Success: true, Message: "Payment successful"}, nil
}

// 数据库保存函数
func saveOrderToDB(order *Order) error {
	// 直接创建数据库连接
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/orders")
	if err != nil {
		return err
	}
	defer db.Close()
	
	_, err = db.Exec("INSERT INTO orders (id, customer_id, amount, created_at, status) VALUES (?, ?, ?, ?, ?)",
		order.ID, order.CustomerID, order.Amount, order.CreatedAt, order.Status)
	return err
}

func generateOrderID() string {
	return fmt.Sprintf("ORDER-%d", time.Now().UnixNano())
}

// ============= 版本2: 现代设计（易于测试，使用传统Mock）=============

// 定义接口
type CustomerValidator interface {
	ValidateCustomer(customerID string) error
}

type PaymentProcessor interface {
	ProcessPayment(customerID string, amount float64) (*PaymentResponse, error)
}

type OrderRepository interface {
	SaveOrder(order *Order) error
}

type TimeProvider interface {
	Now() time.Time
}

type IDGenerator interface {
	GenerateOrderID() string
}

// 现代订单服务 - 依赖注入
type ModernOrderService struct {
	validator   CustomerValidator
	payment     PaymentProcessor
	repository  OrderRepository
	timeProvider TimeProvider
	idGenerator IDGenerator
}

func NewModernOrderService(
	validator CustomerValidator,
	payment PaymentProcessor,
	repository OrderRepository,
	timeProvider TimeProvider,
	idGenerator IDGenerator,
) *ModernOrderService {
	return &ModernOrderService{
		validator:   validator,
		payment:     payment,
		repository:  repository,
		timeProvider: timeProvider,
		idGenerator: idGenerator,
	}
}

func (s *ModernOrderService) CreateOrder(customerID string, amount float64) (*Order, error) {
	// 1. 验证客户 - 通过接口
	if err := s.validator.ValidateCustomer(customerID); err != nil {
		return nil, fmt.Errorf("customer validation failed: %w", err)
	}
	
	// 2. 处理支付 - 通过接口
	paymentResp, err := s.payment.ProcessPayment(customerID, amount)
	if err != nil {
		return nil, fmt.Errorf("payment failed: %w", err)
	}
	
	if !paymentResp.Success {
		return nil, fmt.Errorf("payment rejected: %s", paymentResp.Message)
	}
	
	// 3. 创建订单 - 通过接口获取时间和ID
	order := &Order{
		ID:         s.idGenerator.GenerateOrderID(),
		CustomerID: customerID,
		Amount:     amount,
		CreatedAt:  s.timeProvider.Now(),
		Status:     "PAID",
	}
	
	// 4. 保存到数据库 - 通过接口
	if err := s.repository.SaveOrder(order); err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}
	
	return order, nil
}
