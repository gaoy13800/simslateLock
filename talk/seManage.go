package talk

import (
	"sync"
)

type SessionManage struct {

	SeList map[int64]InterFaceSession

	SeSync sync.RWMutex

	sesIDAcc    int64
}



func (this *SessionManage) Add(sess InterFaceSession){

	//this.SeSync.Lock()
	//
	//defer this.SeSync.Unlock()


	id := this.sesIDAcc

	id++


	socketSes := sess.(*DialSession)
	socketSes.id = id

	this.SeList[id] = sess
}



func newSessionManager() *SessionManage {
	return &SessionManage{
		SeList: make(map[int64]InterFaceSession),
	}
}
