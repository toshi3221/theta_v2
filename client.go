package osc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
	defer client.Response.Body.Close()
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
	defer client.Response.Body.Close()
	if error != nil {
		return
	}
	state = new(State)
	json.Unmarshal(body, &state)
	return
}

type CheckForUpdatesResponse struct {
	StateFingerprint string
	ThrottleTimeout  int
}

func (client *Client) CheckForUpdates(stateFingerprint string, waitTimeout *int) (response *CheckForUpdatesResponse, error error) {
	request_json := `{"stateFingerprint": "` + stateFingerprint + `"`
	if waitTimeout != nil {
		request_json = request_json + `,` + strconv.Itoa(*waitTimeout)
	}
	request_json = request_json + `}`
	client.Response, error = http.Post(client.Url+"/osc/checkForUpdates", "application/json", strings.NewReader(request_json))
	if error != nil {
		return
	}
	body, error := ioutil.ReadAll(client.Response.Body)
	defer client.Response.Body.Close()
	response = new(CheckForUpdatesResponse)
	json.Unmarshal(body, &response)
	return
}

type CommandExecError struct {
	Code    string
	Message string
}
type CommandExecProgress struct {
	Completion float32
}
type CommandExecResponse struct {
	Name     string
	State    string
	Id       string
	Results  interface{}
	Error    CommandExecError
	Progress CommandExecProgress
}

func (client *Client) CommandExecute(name string, parameters interface{}, results interface{}) (response *CommandExecResponse, error error) {
	parameters_json, _ := json.Marshal(parameters)
	request_json := `{"name": "` + name + `", "parameters": ` + string(parameters_json) + `}`
	client.Response, error = http.Post(client.Url+"/osc/commands/execute", "application/json", strings.NewReader(request_json))
	if error != nil {
		return
	}
	body, error := ioutil.ReadAll(client.Response.Body)
	defer client.Response.Body.Close()
	response = new(CommandExecResponse)
	response.Results = results
	json.Unmarshal(body, &response)
	return
}

func (client *Client) CommandStatus(id string, results interface{}) (response *CommandExecResponse, error error) {
	request_json := `{"id": "` + id + `"}`
	client.Response, error = http.Post(client.Url+"/osc/commands/status", "application/json", strings.NewReader(request_json))
	if error != nil {
		return
	}
	body, error := ioutil.ReadAll(client.Response.Body)
	defer client.Response.Body.Close()
	response = new(CommandExecResponse)
	response.Results = results
	json.Unmarshal(body, &response)
	return
}
