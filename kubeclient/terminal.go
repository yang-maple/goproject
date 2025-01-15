package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// 初始化k8s客户端
func initialClientSet(path string) (*kubernetes.Clientset, *rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		log.Fatal(err)
	}

	ClientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	return ClientSet, config, err
}

func initialWS(c *gin.Context) (*websocket.Conn, error) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//将http协议提升为ws
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ws, err
}

// 为remotecommand.StreamOptions提供方法
type streamHandler struct {
	ws          *websocket.Conn                 //ws
	inputMsg    chan []byte                     //客户端输入数据
	resizeEvent chan remotecommand.TerminalSize //窗口调整事件
}

// 获取调整窗口事件
func (handler *streamHandler) Next() *remotecommand.TerminalSize {
	resize := <-handler.resizeEvent
	return &resize
}

// 从ws获取客户端输入的数据
func (handler *streamHandler) Read(p []byte) (size int, err error) {
	data, ok := <-handler.inputMsg
	if !ok {
		return 0, errors.New("I/O data reading failed")
	}
	copy(p, data)
	return len(data), nil
}

// 将标准输出、错误写入ws（客户端）
func (handler *streamHandler) Write(p []byte) (int, error) {
	// 处理非utf8字符
	if !utf8.Valid(p) {
		bufStr := string(p)
		buf := make([]rune, 0, len(bufStr))
		for _, r := range bufStr {
			if r == utf8.RuneError {
				buf = append(buf, []rune("@")...)
			} else {
				buf = append(buf, r)
			}
		}
		p = []byte(string(buf))
	}
	err := handler.ws.WriteMessage(websocket.TextMessage, p)
	return len(p), err
}

// 将字符串转换为int类型
func ToInt(str string) int {
	data, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

// 处理ws输入数据
func executeTask(handler *streamHandler) {
	for {
		_, msg, err := handler.ws.ReadMessage()
		if err != nil {
			return
		}
		//心跳检测
		if string(msg) == "ping" {
			continue
		}
		//调整窗口宽高
		if strings.Contains(string(msg), "resize") {
			resizeSlice := strings.Split(string(msg), ":")
			rows, _ := strconv.Atoi(resizeSlice[1])
			cols, _ := strconv.Atoi(resizeSlice[2])
			handler.resizeEvent <- remotecommand.TerminalSize{
				Width:  uint16(cols),
				Height: uint16(rows),
			}
			continue
		}
		handler.inputMsg <- msg
	}
}

func podTerminal(c *gin.Context) {
	podName := c.Query("podName")
	namespace := c.Query("namespace")
	containerName := c.Query("containerName")
	cols := c.Query("cols")
	rows := c.Query("rows")

	ClientSet, config, err := initialClientSet("./config/config")
	if err != nil {
		return
	}

	//初始化请求体
	req := ClientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).        //podName
		Namespace(namespace). //namespace
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: containerName, //containerName
			Command:   []string{"bash"},
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true, // 启用终端
		}, scheme.ParameterCodec)

	// http转SPDY,添加X-Stream-Protocol-Version等相关header并发送请求
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		log.Println(err)
		return
	}

	ws, err := initialWS(c)
	defer func() {
		ws.Close()
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	handler := &streamHandler{
		ws:          ws,
		inputMsg:    make(chan []byte, 1024),
		resizeEvent: make(chan remotecommand.TerminalSize, 1),
	}
	//将初次获取的窗口cols、rows写入channel
	handler.resizeEvent <- remotecommand.TerminalSize{Width: uint16(ToInt(cols)), Height: uint16(ToInt(rows))}

	//获取ws输入数据
	go executeTask(handler)

	if err := exec.Stream(remotecommand.StreamOptions{
		Stdin:             handler,
		Stdout:            handler,
		Stderr:            handler,
		Tty:               true,
		TerminalSizeQueue: handler,
	}); err != nil {
		ws.Close()
		return
	}
	//定时超时时间
	stopTimer := time.NewTimer(time.Minute * 30)
	for {
		select {
		case <-stopTimer.C:
			return
		}
	}
}

func main() {
	// 创建一个客户端配置
	r := gin.Default()
	r.GET("", podTerminal)
	r.Run(":8081")
}
