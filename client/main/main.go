package main

import (
	"fmt"
	"go_code/chatPro/client/processes"
	"os"
)

var userName string
var userPwd string

func main() {
	
	var key int // 标识输入	

	// var loop bool = true // 菜单显示标志位

	for true {
		fmt.Println("------------welcome to chat room------------")
		fmt.Println("------------1 login-------------------------")
		fmt.Println("------------2 new user register-------------")
		fmt.Println("------------3 logout------------------------")
		fmt.Println("------------please choose(1-3)--------------")
		fmt.Scanf("%d\n", &key)
		switch key {
			case 1:
				fmt.Println("------------login chat room------------")
				fmt.Println("------------please input username")
				fmt.Scanf("%s\n", &userName)
				fmt.Println("------------please input password")
				fmt.Scanf("%s\n", &userPwd)
				userProcess := &processes.UserProcess{}
				userProcess.Login(userName, userPwd)
				// loop = false
			case 2:
				fmt.Println("------------register new user------------")
				// loop = false
			case 3:
				fmt.Println("l------------ogout the system------------")
				os.Exit(0)
			default:
				fmt.Println("wrong input, Please re-input number(1-3)")
		}
	}

	// if 1== key {
	// 	fmt.Println("------------please input username")
	// 	fmt.Scanf("%s\n", &userName)
	// 	fmt.Println("------------please input password")
	// 	fmt.Scanf("%s\n", &userPwd)
	// 	userProcess := &processes.UserProcess{}
	// 	userProcess.Login(userName, userPwd)
	// } else if 2 == key {
	// 	fmt.Println("------------register task")
	// }
}
