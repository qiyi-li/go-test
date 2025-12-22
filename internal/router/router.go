package router

import (
	"go-test/internal/handler" // 记得引入 handler
	"net/http"
)

// NewRouter 不需要返回什么，因为它直接操作 http 的默认路由
// 或者我们可以让它返回一个 *http.ServeMux (更高级的写法)，但现在先保持简单
func RegisterRoutes() {
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/user", handler.GetUserHandler)
	http.HandleFunc("/user/update", handler.UpdateUserHandler)
	http.HandleFunc("/user/delete", handler.DeleteUserHandler)
}
