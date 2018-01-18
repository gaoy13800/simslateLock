package talk

import (
	"sync"
	"errors"
)

type IClients interface {

	Add(terminalId string, sess InterFaceSession)

	Remove(terminalId string)

	GetSessionByTerminalId(terminalId string) (InterFaceSession, error)
}

type clientManage struct {

	terminalList map[string]InterFaceSession

	syncTex sync.RWMutex

}

func (this *clientManage) Add(terminalId string, sess InterFaceSession) {

	//fmt.Println(terminalId, sess)

	this.syncTex.Lock()
	defer this.syncTex.Unlock()

	this.terminalList[terminalId] = sess
}

func (this *clientManage) Remove(terminalId string) {
	this.syncTex.Lock()
	defer this.syncTex.Unlock()

	delete(this.terminalList, terminalId)
}

func (this *clientManage) GetSessionByTerminalId(terminalId string) (InterFaceSession, error) {
	if v, ok := this.terminalList[terminalId]; ok {

		return v, nil
	}

	return nil, errors.New("not exits")
}

func NewTerminalManage() *clientManage {

	return &clientManage{
		terminalList: make(map[string]InterFaceSession),
	}

}
