package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/yakovlevdmv/goonvif"
	"github.com/yakovlevdmv/goonvif/Device"
	"github.com/yakovlevdmv/goonvif/xsd/onvif"
	"github.com/yakovlevdmv/gosoap"
)

const (
	login    = "admin"
	password = "Supervisor"
)

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

type CameraInfo struct {
	DeviceId string `json:"deviceId"`
	Name     string `json:"name"`
	Ip       string `json:"ip"`
	Port     int    `json:port`
}

func getOnLineCamera(camInfo []CameraInfo) ([]string, error) {
	var onLineCamera []string

	for i := 0; i < len(camInfo); i++ {
		dev, err := goonvif.NewDevice(camInfo[i].Ip + ":" + strconv.Itoa(camInfo[i].Port))
		if err != nil {
			log.Println(err)
		} else {
			log.Println(camInfo[i].DeviceId, "is onLine : ", dev)
			onLineCamera = append(onLineCamera, camInfo[i].DeviceId)
		}
	}
	log.Println("end", onLineCamera)
	return onLineCamera, nil
}

func main() {

	info := []CameraInfo{
		CameraInfo{
			DeviceId: "a-b-c-d-1",
			Name:     "camera1",
			Ip:       "10.1.68.15",
			Port:     8000,
		},
		CameraInfo{
			DeviceId: "a-b-c-d-2",
			Name:     "camera2",
			Ip:       "10.1.68.16",
			Port:     8001,
		},
		CameraInfo{
			DeviceId: "a-b-c-d-3",
			Name:     "camera3",
			Ip:       "10.1.68.17",
			Port:     8002,
		},
	}

	// info := make([]CameraInfo, 3)
	// info[0].DeviceId = "a-1";
	getOnLineCamera(info)
	return

	//Getting an camera instance
	dev, err := goonvif.NewDevice("10.1.68.15:8000")
	if err != nil {
		panic(err)
	}

	//Authorization
	dev.Authenticate(login, password)

	//Preparing commands
	systemDateAndTyme := Device.GetSystemDateAndTime{}
	getCapabilities := Device.GetCapabilities{Category: "All"}
	createUser := Device.CreateUsers{User: onvif.User{
		Username:  "TestUser",
		Password:  "TestPassword",
		UserLevel: "User",
	},
	}

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
	createUserResponse, err := dev.CallMethod(createUser)
	if err != nil {
		log.Println(err)
	} else {
		/*
			You could use https://github.com/yakovlevdmv/gosoap for pretty printing response
		*/
		fmt.Println(gosoap.SoapMessage(readResponse(createUserResponse)).StringIndent())
	}

}
