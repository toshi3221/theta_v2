package osc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	url      string
	response *http.Response
}

func OpenClient(url string) (client *Client, error error) {
	client = new(Client)
	client.url = url
	return
}

type Endpoints struct {
	HttpPort         int
	HttpUpdatesPort  int
	HttpsPort        int
	HttpsUpdatesPort int
}

type Info struct {
	Manufacturer    string
	Model           string
	SerialNumber    string
	FirmwareVersion string
	SupportUrl      string
	Gps             bool
	Gyro            bool
	Uptime          int
	Api             []string
	Endpoints       Endpoints
}

func (client *Client) Info() (info *Info, error error) {
	response, error := http.Get(client.url + "/osc/info")
	if error != nil {
		return
	}
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return
	}
	info = new(Info)
	json.Unmarshal(body, &info)
	return
}
