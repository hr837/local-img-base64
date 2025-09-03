package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// 处理根路由请求
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 只响应根路径请求
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// 读取测试页面文件
	htmlPath := filepath.Join(".", "test", "index.html")
	data, err := os.ReadFile(htmlPath)
	if err != nil {
		http.Error(w, "Test page not found", http.StatusNotFound)
		return
	}

	// 设置响应头并返回HTML内容
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// 处理/getImageFile路由
func getImageFileHandler(w http.ResponseWriter, r *http.Request) {
	// 获取查询参数file
	filePath := r.URL.Query().Get("file")
	if filePath == "" {
		http.Error(w, "Missing 'file' query parameter", http.StatusBadRequest)
		return
	}

	// 获取文件扩展名
	ext := filepath.Ext(filePath)

	// 检查是否为支持的图片类型
	acceptExts := ".jpg,.jpeg,.png,.gif,.bmp,.webp,.svg,.avif"

	// 获取索引getImageFileHandler
	idx := strings.Index(acceptExts, ext)
	if idx == -1 {
		http.Error(w, "File is not an image", http.StatusBadRequest)
		return
	}

	// 检查文件是否存在并读取
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
		}
		return
	}

	// 转换为base64并添加固定头部
	base64Str := base64.StdEncoding.EncodeToString(data)
	fullBase64 := fmt.Sprintf("data:image/%s;base64,%s", ext[1:], base64Str)

	// 设置响应头并返回结果
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fullBase64))
}

// CORS中间件，处理跨域请求
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置允许跨域的源，*表示允许所有源
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许的请求方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		// 允许的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 预检请求的缓存时间
		w.Header().Set("Access-Control-Max-Age", "86400")

		// 处理预检请求(OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 调用下一个处理器
		next.ServeHTTP(w, r)
	})
}

func main() {
	// 使用CORS中间件包装处理器
	http.Handle("/", corsMiddleware(http.HandlerFunc(rootHandler)))
	http.Handle("/getImageFile", corsMiddleware(http.HandlerFunc(getImageFileHandler)))

	// 修改端口为21333
	fmt.Println("Server starting on :21333...")
	fmt.Println("Visit http://localhost:21333 in your browser")
	err := http.ListenAndServe(":21333", nil)
	if err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}
