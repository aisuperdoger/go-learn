package main

// Nginx 代理类，实现了server接口
// 提供访问控制、限流等功能
type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

// newNginxServer 创建新的Nginx代理实例
func newNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

// handleRequest 处理客户端请求
// 先进行限流检查，然后转发给真实的应用服务器
func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

// checkRateLimiting 检查请求是否超过限流阈值
func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
