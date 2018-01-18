package main

import (
	"simulateLock/talk"
)

func main(){

	talk.NewSocketTalk()

	select {}
}