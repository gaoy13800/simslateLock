package talk

import (
	"net/http"
	"log"
	"github.com/admpub/mahonia"
	"fmt"
)

var Clients IClients

func ListenHttp(clients IClients){

	Clients = clients

	http.HandleFunc("/flag", sendFlag)
	err := http.ListenAndServe(":10001", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GetGuids(w http.ResponseWriter, req *http.Request){

}

func sendFlag(w http.ResponseWriter, req *http.Request){

	req.ParseForm()

	//guid := req.Form["guid"][0]
	carNum := req.Form["carNum"][0]
	guid := req.Form["guid"][0]

	sess, _ := Clients.GetSessionByTerminalId(guid)

	serializeFlag := mahonia.NewDecoder("utf8").ConvertString(carNum)


	enc:=mahonia.NewEncoder("gbk").ConvertString(serializeFlag)


	fmt.Println(enc)

	sess.SendFlag(enc)
}
