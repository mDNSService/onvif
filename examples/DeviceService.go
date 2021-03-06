package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	goonvif "github.com/mDNSService/onvif"
	"github.com/mDNSService/onvif/device"
	//"github.com/mDNSService/onvif/gosoap"
	//"github.com/mDNSService/onvif/xsd/onvif"
)

const (
	login    = "admin"
	password = ""
)

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	//Getting an camera instance
	dev, err := goonvif.NewDevice("192.168.123.211:8899")
	if err != nil {
		panic(err)
	}
	//Authorization
	dev.Authenticate(login, password)
	log.Println(dev.GetServices())
	//Preparing commands
	systemDateAndTyme := device.GetSystemDateAndTime{}
	getCapabilities := device.GetCapabilities{Category: "All"}
	//createUser := device.CreateUsers{User: onvif.User{
	//	Username:  "TestUser",
	//	Password:  "TestPassword",
	//	UserLevel: "User",
	//},
	//}

	//Commands execution
	systemDateAndTymeResponse, err := dev.CallMethod(systemDateAndTyme)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(readResponse(systemDateAndTymeResponse))
	}
	getCapabilitiesResponse, err := dev.CallMethod(getCapabilities)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(readResponse(getCapabilitiesResponse))
	}
	//createUserResponse, err := dev.CallMethod(createUser)
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	/*
	//		You could use https://github.com/mDNSService/onvif/gosoap for pretty printing response
	//	*/
	//	fmt.Println(gosoap.SoapMessage(readResponse(createUserResponse)).StringIndent())
	//}

}
