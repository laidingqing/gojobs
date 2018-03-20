package controller

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"
)

const (
	//InternalErrorMsg ..
	InternalErrorMsg = "Unable to complete request. Please try again later."
)

//OkResponse ok response
type OkResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

//WriteJSONResponse write json response
func WriteJSONResponse(w http.ResponseWriter, status int, value interface{}) {
	data, _ := json.Marshal(value)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

//GetIP ...
func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Unable to determine local IP address (non loopback). Exiting.")
}
