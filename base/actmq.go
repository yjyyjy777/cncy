package base

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"os/exec"
	"time"
)

func Activemq() {
	conn, err := stomp.Dial("tcp", "localhost:61613",
		stomp.ConnOpt.HeartBeat(0, 0),
		stomp.ConnOpt.Login("admin", "admin"))

	if err != nil {
		fmt.Println("Error connecting to ActiveMQ:", err)
		fmt.Println("activemq服务异常需要重新启动")
		cmd := exec.Command("service", "activemq", "restart")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))

		return
	}

	defer conn.Disconnect()

	// Subscribe to the queue
	sub, err := conn.Subscribe("/queue/test", stomp.AckAuto)
	if err != nil {
		fmt.Println("Error subscribing to queue:", err)
		return
	}
	defer sub.Unsubscribe()

	// Send a message to the queue
	conn.Send("/queue/test", "text/plain", []byte("读取测试数据，1"), nil)

	// Receive messages from the queue
	stomp.ConnOpt.HeartBeat(7200*time.Second, 7200*time.Second)

	msg := <-sub.C
	if msg.Err != nil {
		fmt.Println("Error receiving message:", msg.Err)
		return
	}
	fmt.Println("Received message:", string(msg.Body))
	if string(msg.Body) == "读取测试数据，1" {
		fmt.Println("activemq服务正常")
	} else {
		fmt.Println("activemq服务异常需要重新启动")
		cmd := exec.Command("service", "activemq", "restart")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
	}
}
