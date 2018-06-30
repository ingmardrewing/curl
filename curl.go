package curl

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func NewCurl() Curl {
	c := new(curl)
	c.method = "GET"
	c.userAgent = "curl"
	c.timeout = time.Duration(10 * time.Second)
	return c
}

type Curl interface {
	SetUserAgent(userAgent string)
	SetUrl(url string)
	SetMethod(method string)
	SetTimeout(timeout time.Duration)
	Execute() int
	Data() string
}

type curl struct {
	url       string
	userAgent string
	method    string
	data      string
	timeout   time.Duration
}

// Set user agent
func (c *curl) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

// Sets the the timout for the http request in seconds
func (c *curl) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// Set url
func (c *curl) SetUrl(url string) {
	c.url = url
}

// Set http method (GET,POST,etc.)
func (c *curl) SetMethod(method string) {
	c.method = method
}

// Actually executes the http request,
// stores the returned data (if any)
// and returns the http status code
func (c *curl) Execute() int {
	if len(c.url) == 0 {
		log.Fatalln("No url given")
	}
	client := http.Client{
		Timeout: c.timeout,
	}

	req, err := http.NewRequest(c.method, c.url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	c.data = string(responseData)
	return resp.StatusCode
}

// Return retrieved data
func (c *curl) Data() string {
	return c.data
}
