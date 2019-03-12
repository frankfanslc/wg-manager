package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// API is a utility for communicating with the
type API struct {
	Username string
	Password string
	BaseURL  string
	Client   *http.Client
}

// WireguardPeerList is a list of Wireguard peers
type WireguardPeerList []WireguardPeer

// WireguardPeer is a wireguard peer
type WireguardPeer struct {
	IPLeastsig int    `json:"ip_leastsig"`
	Ports      []int  `json:"ports"`
	Pubkey     string `json:"pubkey"`
}

// GetWireguardPeers fetches a list of wireguard peers from the API and returns it
func (a *API) GetWireguardPeers() (WireguardPeerList, error) {
	req, err := http.NewRequest("GET", a.BaseURL+"/wg/active-pubkeys/", nil)
	if err != nil {
		return WireguardPeerList{}, err
	}
	req.SetBasicAuth(a.Username, a.Password)

	response, err := a.Client.Do(req)
	if err != nil {
		return WireguardPeerList{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return WireguardPeerList{}, err
	}

	var decodedResponse WireguardPeerList
	err = json.Unmarshal(body, &decodedResponse)
	if err != nil {
		return WireguardPeerList{}, fmt.Errorf("error decoding %s, %s", body, err.Error())
	}

	return decodedResponse, nil

}
