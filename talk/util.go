package talk

import (
	"crypto/md5"
	"encoding/hex"
	//"io"
	//"encoding/base64"
	//"time"
	//"strconv"
	//"math/rand"
	//"crypto/rand"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

var (
	num1 string = "43d6b5314f408991b9d61474f8301281"
	num2 string = "79c1e29cd1c42ed01ec8d7e5bbde83f0"
	num3 string = "8a70e3eba04d9e4ade286928ac40270c"
	num4 string = "fe36199881182d0c6369255a8ecf4b9a"
	num5 string = "e0ccca2c7662be88df05631c4f30b345"
	num6 string = "c760e45d1cd477599697daec17f60ca3"
	num7 string = "7dacdfba88884383752680d9f6323d64"
	num8 string = "b9f95af1f8435e06b51e2c61e122889e"
)

func CreateDeviceId(num int) string {

	switch num {
	case 1:
		return num1
	case 2:
		return num2
	case 3:
		return num3
	case 4:
		return num4
	case 5:
		return num5
	case 6:
		return num6
	case 7:
		return num7
	case 8:
		return num8
	default:
		return num8
	}
	//
	//return
	//
	//list := make(map[string]string)
	//
	//
	//
	//
	//
	//
	//b := make([]byte, 48)
	//if _, err := io.ReadFull(rand.Reader, b); err != nil {
	//	return ""
	//}
	//deviceId := GetMd5String(base64.URLEncoding.EncodeToString(b))
	//
	//return deviceId
}

//func GetRandNum(index int) string{
//
//	var str string
//
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//
//	for i:=0; i < index; i++ {
//
//		num := strconv.Itoa(r.Intn(10))
//
//		str += num
//	}
//
//	return str
//}