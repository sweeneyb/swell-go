package swell

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const swellUrl string = "https://api.swell.store"

type Client struct {
	HostUrl    string
	HTTPClient *http.Client
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Store  string
	Secret string
}

type Error struct {
	Errors map[string]map[string]interface{} `json:"errors"`
}

func NewClient() (*Client, error) {
	c := Client{
		HostUrl: swellUrl,
		HTTPClient: &http.Client{
			Timeout: time.Second * 2, // Timeout after 2 seconds
		},
		Auth: AuthStruct{
			Store:  os.Getenv("swell_store"),
			Secret: os.Getenv("swell_secret"),
		},
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	req.SetBasicAuth(os.Getenv("swell_store"), os.Getenv("swell_secret"))
	req.Header.Set("User-Agent", "go-libray")

	res, getErr := c.HTTPClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%v", res.StatusCode)
	}

	if res.Body != nil {
		defer res.Body.Close()
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		return body, nil
	}
	return nil, nil
}
