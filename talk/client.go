package talk

import (
	"net"
	"log"
	"time"
	"fmt"
	"simulateLock/conf"
	//"strconv"
)



type SocketTalk struct {

	Conn *net.TCPConn

	msgClient *MsgPlamt

	*SessionManage

	terminal *clientManage
}


type josn map[string]string



func (this *SocketTalk)NewClient(){

	go ListenHttp(this.terminal)

	//num, _ := strconv.Atoi(conf.Conf["server_addr"])

	for i :=1; i < 9; i++ {

		deviceId := CreateDeviceId(i)

		fmt.Println("deviceId:", deviceId)

		go this.Dial(deviceId)
	}


}



func (this *SocketTalk) Dial(guid string){

	//tcpAddr, err := net.ResolveTCPAddr("tcp", "192.168.5.49:6000")
	//tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:44444")

	tcpAddr, err := net.ResolveTCPAddr("tcp", conf.Conf["server_addr"])

	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.DialTCP("tcp", nil,  tcpAddr)

	this.Conn = conn

	if err != nil {
		log.Println("net Dial TCP error", err)
		return
	}
	session := NewSession(conn)

	this.terminal.Add(guid, session)

	session.SetGuid(guid)
	//this.SessionManage.Add(session)


	//新的通讯协议不再发送ida
	//this.msgClient.IdaSend(session, guid)

	go this.MessagePipe(guid, conn, session)

	go func() {
		tick := time.Tick(time.Second * 15)

		for  {
			select {
			case <- tick:
				this.msgClient.TickIdb(session, guid)
			}
		}
	}()

}


func (this *SocketTalk)MessagePipe(guid string, conn *net.TCPConn, sess InterFaceSession){
	sess.Recv()
}


func NewSocketTalk() *SocketTalk{
	self := &SocketTalk{
		msgClient:NewMsg(),
		terminal:NewTerminalManage(),
	}

	self.NewClient()

	return self
}