package talk

import (
	"net"
	"log"
	"simulateLock/conf"
)


type InterFaceSession interface {
	Send (msg string)

	ID() int64

	Close()

	Recv()

	SetGuid(guid string)

	GetGuid() string

	SetCstu(cstu int)

	GetCstu() int

	GetTerminalType() string

	SendFlag(carNum string)

}

type DialSession struct {

	id int64

	guid string

	conn net.Conn

	Cstu int

	plam *MsgPlamt

	ClientType string
}


func (this *DialSession) Send(msg string){

	this.conn.Write([]byte(msg))
}


func (this *DialSession) SetGuid(guid string){

	this.guid = guid
}

func (this *DialSession) GetGuid() string{
	return this.guid
}


func (this *DialSession) ID () int64{
	return this.id
}

func (this *DialSession) Close() {
	this.conn.Close()
}

func (this *DialSession) Recv(){

	for  {
		byt := make([]byte, 1024)

		index, err := this.conn.Read(byt)

		if err != nil {

			log.Println("连接中断！！！" , err)

			this.Close()
			break
		}


		msg := string(byt[0:index])

		this.plam.GetMessage(this, msg)
	}

}


func (this *DialSession) SetCstu(cstu int){
	this.Cstu = cstu
}

func (this *DialSession) GetCstu() int{
	return this.Cstu
}

func (this *DialSession) SendFlag(carNum string){
	this.plam.SendMsg(this, this.guid+"flag"+carNum)
}


func (this *DialSession) GetTerminalType() string{
	return this.ClientType
}


func NewSession(conn net.Conn) *DialSession{

	self := &DialSession{
		conn:conn,
		plam:NewMsg(),
		Cstu:3,
	}

	self.ClientType = conf.Conf["client_type"]

	return self
}
