package osc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Url      string
	Response *http.Response
}

func OpenClient(url string) (client *Client, error error) {
	client = new(Client)
	client.Url = url
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
	client.Response, error = http.Get(client.Url + "/osc/info")
	if error != nil {
		return
	}
	body, error := ioutil.ReadAll(client.Response.Body)
	if error != nil {
		return
	}
	info = new(Info)
	json.Unmarshal(body, &info)
	return
}

type CameraState struct {
	SessionId      string
	BatteryLevel   float32
	StorageChanged bool
}

type State struct {
	Fingerprint string
	State       CameraState
}

func (client *Client) State() (state *State, error error) {
	client.Response, error = http.PostForm(client.Url+"/osc/state", nil)
	if error != nil {
		return
	}
	body, error := ioutil.ReadAll(client.Response.Body)
	if error != nil {
		return
	}
	state = new(State)
	json.Unmarshal(body, &state)
	return
}
