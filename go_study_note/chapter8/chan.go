package main

import "time"

func main() {
	// 创建通道。因为仅是通知，数据并没有实际意义
	exit := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		println("goroutine done.")
		// 关闭通道，发出信号
		close(exit)
	}()

	println("main ...")
	// 如通道关闭，立即接触阻塞
	<-exit
	println("main exit...")
}
