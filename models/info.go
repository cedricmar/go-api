package models

import (
	"runtime"
	"time"
)

var (
	name        = "toto"
	description = "something..."
)

// ServiceInfo service
type ServiceInfo struct {
	Name        string    `json:"name"`
	Desc        string    `json:"desc"`
	OS          string    `json:"os"`
	Proc        int       `json:"proc"`
	CurrentTime time.Time `json:"current_time"`
}

// GetServiceInfo getter
func GetServiceInfo() *ServiceInfo {
	return &ServiceInfo{
		Name:        name,
		Desc:        description,
		OS:          getOS(),
		Proc:        runtime.NumCPU(),
		CurrentTime: time.Now(),
	}
}

func getOS() string {
	return runtime.GOOS
}
