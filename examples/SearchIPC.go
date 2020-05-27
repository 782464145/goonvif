package main

import (
	"log"
	"strconv"

	"github.com/782464145/goonvif"
)

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

	getOnLineCamera(info)
	return
}
