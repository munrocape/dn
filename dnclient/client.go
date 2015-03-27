package dnclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	BaseUrl       string
	StoriesSuffix string
	ClientId      string
}

func NewClient() *Client {
	client_id := os.Getenv("DN_CLIENT_ID")
	if client_id == "" {
		fmt.Printf("ERROR: Could not find Designer News Client Id\n")
		os.Exit(1)
	}
	var c = Client{
		BaseUrl:       "https://api-news.layervault.com/api/v1/",
		StoriesSuffix: "stories",
		ClientId:      "?client_id=" + client_id,
	}
	return &c
}

func (c *Client) Get(url string, params url.Values) ([]byte, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	// make the request
	req.Close = true
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (c *Client) GetStories() (Stories, error) {
	url := c.BaseUrl + c.StoriesSuffix + c.ClientId
	rep, err := c.Get(url, nil)
	var stories Stories
	if err != nil {
		return stories, err
	}
	err = json.Unmarshal(rep, &stories)
	return stories, err
}

// func main() {
// 	c := NewClient()
// 	s, _ := c.GetStories()
// 	for _, element := range s.Stories {
// 		fmt.Printf("%d %s %s %s %d\n", element.VoteCount, element.Title, element.Url, element.SiteUrl, len(element.Comments))
// 	}
// }
