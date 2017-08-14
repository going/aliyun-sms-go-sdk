package main

import (
	"fmt"
	"os"

	"github.com/going/aliyun-sms-go-sdk/sms"
)

// modify it to yours
const (
	ENDPOINT  = "https://dysmsapi.aliyuncs.com/"
	ACCESSID  = "*******"
	ACCESSKEY = "***********"
)

func main() {
	sms.HttpDebugEnable = true
	c := sms.New(ACCESSID, ACCESSKEY)
	// send to one person
	e, err := c.SendOne("18616*****", "*****", "SMS_*******", `{"code":"1234"}`)
	if err != nil {
		fmt.Println("send sms failed", err, e.Error())
		os.Exit(0)
	}
	fmt.Println("send sms succeed", e.GetRequestId())
}
