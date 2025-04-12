package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/petermattis/goid"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println(time.Now().Format(time.DateTime))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		simpleReverseProxyTestDefer(c)

		c.JSONP(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
}

func simpleReverseProxyTestDefer(c *gin.Context) {

	defer func() {
		// time.Sleep(time.Second * 3)
		go func() {
			fmt.Println("Over Log...", goid.Get())
		}()
	}()

	// 定义目标服务地址
	target, _ := url.Parse("https://www.baidu.com")

	proxyStart(target, c.Writer, c.Request)
}

func proxyStart(remote *url.URL, rw http.ResponseWriter, re *http.Request) {

	// 创建反向代理实例
	proxy := httputil.NewSingleHostReverseProxy(remote)

	//proxy.ModifyResponse = func(resp *http.Response) error {
	//	reader := bufio.NewReader(resp.Body)
	//	bytes, err := io.ReadAll(reader)
	//	if err != nil {
	//		log.Fatal(err)
	//		return err
	//	}
	//	respStr := string(bytes)
	//	length := len(respStr)
	//	if length > 50 {
	//		fmt.Printf("结果返回，结果前50字符: %s, 长度: %d \n", respStr[0:50], length)
	//	} else {
	//		fmt.Printf("结果返回，结果字符: %s, 长度: %d \n", respStr, length)
	//	}
	//
	//	return nil
	//}

	// 设置响应拦截逻辑
	proxy.ModifyResponse = func(resp *http.Response) error {
		// 记录状态码
		log.Printf("响应状态码: %d", resp.StatusCode)

		// 读取响应体
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// 必须重置 Body 以便后续读取
		resp.Body = io.NopCloser(bytes.NewReader(body))

		// 记录响应体长度
		log.Printf("响应体长度: %d", len(body))

		// 修改内容（示例：替换字符串）
		modifiedBody := bytes.ReplaceAll(body, []byte("旧文本"), []byte("新文本"))

		// 重置 Body 并更新 Content-Length
		resp.Body = io.NopCloser(bytes.NewReader(modifiedBody))
		resp.ContentLength = int64(len(modifiedBody)) // 关键步骤
		resp.Header.Set("Content-Length", strconv.Itoa(len(modifiedBody)))

		// 移除分块编码标记（如果存在）
		resp.Header.Del("Transfer-Encoding")

		return nil
	}

	// 设置错误处理
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("代理错误: %v", err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("502 - 代理服务异常"))
	}

	// 启动代理服务器
	proxy.ServeHTTP(rw, re)
	// log.Fatal(http.ListenAndServe(":8080", proxy))
}

// 简单反向代理(单目标)
func simpleReverseProxy() {
	// 定义目标服务地址
	target, _ := url.Parse("http://backend-server:8080")

	// 创建反向代理实例
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 启动代理服务器
	log.Fatal(http.ListenAndServe(":8000", proxy))
}

// 路由配置表
var routeMap = map[string]string{
	"/api/":    "http://api-server:3000",
	"/static/": "http://static-server:4000",
}

// 动态路由(多目标)
func routerReverseProxy() {
	// 自定义请求处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 查找匹配的路由
		var targetURL *url.URL
		for pathPrefix, target := range routeMap {
			if strings.HasPrefix(r.URL.Path, pathPrefix) {
				targetURL, _ = url.Parse(target)
				break
			}
		}

		// 未找到路由返回404
		if targetURL == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// 创建动态代理
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// 修改请求头（示例）
		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host
			req.URL.Path = strings.TrimPrefix(req.URL.Path, "/api")
			req.Header.Set("X-Proxy", "Go-Reverse-Proxy")
		}

		// 执行代理请求
		proxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
