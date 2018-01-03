package tool

import (
	"bytes"
	rand4 "math/rand"
	"reflect"
)

//工具类

const(
	ReturnErrorParameNotNull = "参数不能为空"
	ReturnErrorTokenFailure = "Token 失效，请重新登录"
	ReturnErrorUserNonExistence = "用户不存在或被封禁"

	ReturnOperateFiled = "filed"
	ReturnOperateSuccess = "success"
)

type ResultData struct {
	Code    int64
	Message string
	Data    interface{}
}

func VerifyPassword(psssword string) bool {

	return false
}

func GetToken() string {
	rand1 := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	rand2 := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	rand3 := [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	buffer := bytes.Buffer{}
	for i := 0; i < 50; i++ {
		r := rand4.Intn(4)

		var s string
		if r == 0 {
			s = rand3[rand4.Intn(len(rand3))]
		} else if r > 0 && r <= 2 {
			s = rand2[rand4.Intn(len(rand2))]
		} else {
			s = rand1[rand4.Intn(len(rand1))]
		}

		buffer.WriteString(s)
	}

	return buffer.String()
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
