package utils

import "fmt"

//goland:noinspection GoUnusedFunction
func niy(method string, params ...any) {
	str := ""
	for _, param := range params {
		if len(str) != 0 {
			str += ", "
		}
		str += fmt.Sprintf("'%v'", param)
	}
	panic(method + "(" + str + ")")
}
