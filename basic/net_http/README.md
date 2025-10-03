在 Go 的 `net/http` 包中，`ServeHTTP` 方法的调用链是一个精心设计的流程。它最终由 Go 的 HTTP 服务器在接收到客户端请求后自动触发。我们来一步步追踪这个调用过程，直到最底层。

### 核心调用链

`客户端请求` → `http.Server.Serve` → `serverHandler.ServeHTTP` → `mux.ServeHTTP` → **你的 Handler 的 `ServeHTTP`**

下面详细分解每一步：

---

### 1. 服务器启动 (`http.ListenAndServe`)

当你调用 `http.ListenAndServe(":8080", mux)` 时，底层会创建一个 `*http.Server` 实例，并调用其 `ListenAndServe` 方法。

```go
// 简化后的调用过程
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler} // handler 就是 mux
    return server.ListenAndServe()
}
```

### 2. 服务器开始监听和接受连接 (`Server.Serve`)

`ListenAndServe` 内部会监听端口，并在一个循环中不断接受新的 TCP 连接。

```go
// 伪代码：Server.ListenAndServe -> Server.Serve
func (srv *Server) Serve(l net.Listener) error {
    for {
        // 接受一个新连接
        rw, err := l.Accept()
        if err != nil {
            return err
        }
        // 为每个连接启动一个 goroutine 处理
        c := srv.newConn(rw)
        go c.serve(ctx)
    }
}
```

### 3. 为每个连接启动处理协程 (`conn.serve`)

每个连接 (`conn`) 都会在自己的 goroutine 中运行 `serve` 方法。这个方法会读取 HTTP 请求头，解析出 `*http.Request`，然后进入处理流程。

```go
// conn.serve 方法内部 (简化)
func (c *conn) serve(ctx context.Context) {
    for {
        // 读取请求
        w, err := c.readRequest(ctx)
        if err != nil {
            break
        }
        // 获取请求对象
        req := w.request
        
        // 关键步骤：创建 serverHandler 并调用其 ServeHTTP
        handler := c.server.Handler
        if handler == nil {
            handler = DefaultServeMux // 如果没传 handler，用默认的
        }
        serverHandler{c.server, handler}.ServeHTTP(w, req)
        
        // 处理下一个请求 (HTTP/1.1 Keep-Alive)
    }
}
```

### 4. `serverHandler.ServeHTTP`：进入处理管道

`serverHandler` 是一个内部类型，它包装了 `*Server` 和 `Handler`。它的 `ServeHTTP` 方法是请求处理的入口点。

```go
type serverHandler struct {
    server  *Server
    handler Handler
}

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
    // 这里可能会调用中间件，但最终会调用 sh.handler.ServeHTTP
    sh.handler.ServeHTTP(rw, req) // sh.handler 就是你的 mux
}
```

### 5. `mux.ServeHTTP`：路由分发

现在，控制权交到了你的 `*http.ServeMux`（或你传入的任何 `Handler`）的 `ServeHTTP` 方法。

```go
// http.ServeMux.ServeHTTP 的核心逻辑
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
    // 1. 根据 r.URL.Path 查找注册的 handler
    h, _ := mux.Handler(r)
    
    // 2. 调用找到的 handler 的 ServeHTTP 方法
    h.ServeHTTP(w, r) // h 是你注册的具体 handler，比如 http.HandlerFunc
}
```

### 6. 你的 Handler 的 `ServeHTTP` 被调用

最后，你注册的具体处理函数（例如通过 `mux.HandleFunc` 注册的）的 `ServeHTTP` 方法被执行。

```go
// 假设你注册了
mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
})

// 这个匿名函数被包装成了 http.HandlerFunc
// http.HandlerFunc 本身实现了 ServeHTTP
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r) // 调用你传入的函数
}
```

所以，你写的 `fmt.Fprintf(w, "Hello!")` 就在这里被执行。

---

### 总结：`ServeHTTP` 是何时被调用的？

**`ServeHTTP` 方法是在服务器成功解析完一个 HTTP 请求后，由 `conn.serve` 协程调用 `serverHandler.ServeHTTP` 时触发的。**

更精确的时间点是：

1.  **TCP 连接建立**：客户端与服务器的 TCP 连接完成。
2.  **HTTP 请求到达**：客户端发送了 HTTP 请求数据。
3.  **请求被解析**：服务器的 `conn.serve` 读取并解析出完整的 `*http.Request`。
4.  **进入处理流程**：`serverHandler.ServeHTTP` 被调用。
5.  **路由分发**：`mux.ServeHTTP` 被调用，根据 URL 查找目标 Handler。
6.  **最终执行**：**目标 Handler 的 `ServeHTTP` 方法被调用**，执行你的业务逻辑。

因此，`ServeHTTP` 是整个 HTTP 处理流程的**执行终点**，它由 Go 的 HTTP 服务器框架在每次有请求需要处理时自动调用。