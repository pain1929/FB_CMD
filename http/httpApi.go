package http

import (
	"encoding/json"
	"log"
	"net/http"
	"phoenixbuilder/fastbuilder/environment"
)

// SendCmdRequest 定义与 C++ 结构体相对应的 Go 结构体
type SendCmdRequest struct {
	Cmd string `json:"cmd"` // JSON 中的 "cmd" 映射到结构体的 Cmd 字段
}

// 处理 /api/sendCmd 的请求
func sendCmdHandler(w http.ResponseWriter, r *http.Request, env *environment.PBEnvironment) {
	// 检查请求方法是否为 POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 解析 JSON 数据
	var req SendCmdRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 打印接收到的命令
	log.Printf("Received command: %s", req.Cmd)

	// 调用 SendCommand 方法
	err := env.GameInterface.SendCommand(req.Cmd)
	var response map[string]interface{}

	if err != nil {
		// 失败响应
		response = map[string]interface{}{
			"status": 1,
			"error":  err.Error(),
		}
	} else {
		// 成功响应
		response = map[string]interface{}{
			"status": 0,
		}
	}

	// 返回 JSON 响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func InitHttpApi(env *environment.PBEnvironment) {
	// sendCmd 路由
	http.HandleFunc("/api/sendCmd", func(w http.ResponseWriter, r *http.Request) {
		sendCmdHandler(w, r, env)
	})

	// 启动 HTTP 服务器
	port := ":1319"
	log.Printf("httpApiServer is running on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
