package main

// Application 真实的应用服务器
// 实现了server接口，提供实际的业务逻辑
type Application struct {
}

// handleRequest 处理具体的业务请求
func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Ok"
}
