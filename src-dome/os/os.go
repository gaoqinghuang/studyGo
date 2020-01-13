package main

import (
	"fmt"
	"os"
)

func main() {
	shellPath := "/home/xx/test.sh"
	argv := make([]string, 1)
	attr := new(os.ProcAttr)
	newProcess, err := os.StartProcess(shellPath, argv, attr) //运行脚本
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Process PID", newProcess.Pid)
	processState, err := newProcess.Wait() //等待命令执行完
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("processState PID:", processState.Pid()) //获取PID
	fmt.Println("ProcessExit:", processState.Exited())   //获取进程是否退出

	//file, _ := os.OpenFile("xiaogao.txt", os.O_APPEND, 0)
	//defer file.Close()
	//
	//data := make([]byte, 1024)
	//n, _ := file.Read(data)
	//fmt.Println(file.WriteString("\n"))
	//fmt.Println(string(data[:n]))

}
