/*
这个是对于用golang驱动树莓派点亮LED小灯的程序。在这个程序里面用到了github.com/stianeikeland/go-rpio/v4的开源项目。
go-rpio可以非常容易的驱动GPIO
*/
// 在开始驱动树莓派前，需要先安装golang

// 生成go.mod文件
//执行 go mod init [pwm-blinky] pwm-blinky 你自己的包名
package main

import (
    "fmt"
    "github.com/stianeikeland/go-rpio/v4"
    "os"
    "time"
)

var (
        led = rpio.Pin(17) // 定义BCM17脚为LED输出
        button = rpio.Pin(6) // 定义BCM6脚为LED输出
)

func main(){
      //开mem 在rpio并接收错误
  if err :=rpio.Open();err !=nil{
    fmt.Println(err)
    os.Exit(1)
  }
  defer rpio.Close()
  // 设定输出模式
  led.Output()
  // 设定button为输入模式
  button.Input()
  //下面的部分为在一个循环里点亮led
  //for x:=0;x<20;x++{
   // led.Toggle()
   // time.Sleep(time.Second/5)
  //}
  // 在死循环中判断开关的控制
  for{
       res :=button.Read()
       if res == 0 {
          led.High() // 点亮led，设置为高电平
          fmt.Println(“button on ”)
          time.Sleep(5*time.Second)// 休眠5 秒， 这里的意思是点亮5秒
	      }
	     led.Low()
          fmt.Println("button status:",res)
	  
  }
}
