package gopacket

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/lflxp/showme/utils"
)

// func Run(in string) {
// 	fmt.Println(in)
// 	tmp := strings.Split(in, " ")
// 	if len(tmp) != 3 {
// 		fmt.Println("命令长度不为三，退出")
// 		return
// 	}

// 	// ok := true
// 	// // 获取退出信号
// 	// c := make(chan os.Signal, 1)
// 	// signal.Notify(c, os.Interrupt, os.Kill)
// 	// num := 0
// 	// go func() {
// 	// 	for {
// 	// 		num++
// 	// 		select {
// 	// 		case s := <-c:
// 	// 			fmt.Printf("\n\033[1;4;31m%s:罒灬罒:小伙子走了哟！\033[0m\n", s)
// 	// 			ok = false
// 	// 			break
// 	// 			os.Exit(1)
// 	// 		}
// 	// 		if !ok {
// 	// 			break
// 	// 		}
// 	// 	}
// 	// }()
// 	utils.WatchDogEasy(tmp[2])
// }

func Run(in string) {
	fmt.Println(in)
	tmp := strings.Split(in, " ")
	if len(tmp) != 3 {
		fmt.Println("命令长度不为三，退出")
		return
	}

	ok := true
	// 获取退出信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	bigChan := make(chan interface{}, 1000)
	// defer close(bigChan)

	go utils.WatchDog(bigChan, tmp[2])

	num := 0
	for {
		num++
		select {
		case s := <-c:
			fmt.Printf("\n\033[1;4;31m%s:罒灬罒:小伙子走了哟！\033[0m\n", s)
			ok = false
			close(bigChan)
			break
		case data, ok := <-bigChan:
			if !ok {
				fmt.Println("is ok?")
				return
			}
			switch data.(type) {
			case *utils.Data:
				json, err := json.Marshal(data.(*utils.Data))
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println(string(json))
			}
		}

		if !ok {
			break
		}
		fmt.Println("num", num)
	}
}