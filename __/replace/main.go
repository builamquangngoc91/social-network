package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "https://scontent.fmaa10-1.fna.fbcdn.net/v/t15.5256-10/183608934_1136279103557797_34050213524739064_n.jpg?_nc_cat=107\u0026ccb=1-3\u0026_nc_sid=f2c4d5\u0026_nc_ohc=XLzmkgqExxIAX9deACL\u0026_nc_ht=scontent.fmaa10-1.fna\u0026oh=2891940c8a7f7e6a5f11ba45a22a15ae\u0026oe=60C2F2B5"

	fmt.Println(string(bytes.ReplaceAll([]byte(str), []byte("\u0026"), []byte("&"))))
}
