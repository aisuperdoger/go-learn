package main

// server 定义了服务的接口
// 代理和真实服务都必须实现这个接口
type server interface {
	handleRequest(string, string) (int, string)
}
