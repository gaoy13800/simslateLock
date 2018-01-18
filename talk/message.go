package talk

import (
	"net"
	"sync"
	"fmt"
	"strconv"
	"log"
)



type MsgPlamt struct {

	conn net.Conn

	syncTax sync.WaitGroup

	single *SingleSe

}


func (this *MsgPlamt) IdaSend(sess InterFaceSession, guid string){
	msg := "ida-"+ guid

	sess.Send(msg)
}


func (this *MsgPlamt) GetMessage(sess InterFaceSession, serverMsg string){

	fmt.Println("收到指令 ", serverMsg)

	if len(serverMsg) == 2 && serverMsg == "OK" {
		return
	}

	var dicate, guid string

	if len(serverMsg) >32 {
		dicate = serverMsg[32:]

		guid =  serverMsg[:32]
	}else {
		dicate = "stus"
		guid = sess.GetGuid()
	}



	switch dicate {
	case "stus":
		fmt.Println("接收电量stus指令, guid", guid)
		this.stusSend(sess)
		break
	case "cstu":
		fmt.Println("接收状态ctus指令, guid", guid)
		this.SendMsg(sess, "cstu-" + guid + "|" + this.single.GetCstu())
		break
	case "brut":
		fmt.Println("接收brut指令", guid)
		this.SendMsg(sess, "brut-" + guid)
		break
	case "over":
		fmt.Println("接收over指令", guid)
	case "open":
		fmt.Println("接收open指令", guid)
		sess.SetCstu(3)
		this.SendMsg(sess, guid+"open")
		break
	case "clse":
		fmt.Println("接收clse指令",guid)
		sess.SetCstu(4)
		this.SendMsg(sess, guid+"clse")
	default:
		fmt.Println("未知命令", serverMsg)
	}

}

func (this *MsgPlamt) SendMsg(sess InterFaceSession, msg string){

	log.Println("msg", msg)

	sess.Send(msg)
}

func (this *MsgPlamt) stusSend(sess InterFaceSession){
	msg := sess.GetGuid() + "stus80"

	sess.Send(msg)
}


func (this *MsgPlamt) TickIdb(sess InterFaceSession, guid string){

	var status string

	if sess.GetTerminalType() == "lock" {
		status = strconv.Itoa(sess.GetCstu())
	}else {
		status = "0"
	}

	this.SendMsg(sess, guid+"hblv" + status)
}


func NewMsg() *MsgPlamt{

	return &MsgPlamt{
		single:NewSingleRe(),
	}
}