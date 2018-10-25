package http

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
)

// GetJSONBytes queries an HTTP endpoint and returns the JSON response
// payload as a byte array
func GetJSONBytes(url string) ([]byte, error) {
	bytes, err := GetBytes(url)
	if err != nil {
		return nil, err
	}
	if !gjson.ValidBytes(bytes) {
		return nil, fmt.Errorf("Received invalid JSON data from: %s", url)
	}

	return bytes, nil
}

type ContentRequest struct {
	Content json.RawMessage `json:"content"`
}

func PostJSON(url string, body []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.Header.SetMethod("POST")
	req.Header.Add("Accept", "application/json")

	req.SetRequestURI(url)
	req.SetBody(body)

	err := fasthttp.DoTimeout(req, res, 30*time.Second)
	if err != nil {
		return nil, err
	}

	bytes := res.Body()

	if !gjson.ValidBytes(bytes) {
		return nil, fmt.Errorf("Received invalid JSON data from: %s", url)
	}

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
	return bytes, nil
}

// GetBytes queries an HTTP endpoint and returns the raw []byte response
func GetBytes(url string) ([]byte, error) {
	statusCode, body, err := fasthttp.GetTimeout(nil, url, 30*time.Second)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("Received unsuccessful HTTP status code \"%d\" from: %s", statusCode, url)
	}
	return body, nil
}

// JoinStrings ... concatenates an arbitrary amount of provided strings
func JoinStrings(strs ...string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}

// BuildURL constructs a URL string
func BuildURL(protocol string, host string, port string, route string) string {
	var builder strings.Builder
	builder.WriteString(protocol)
	builder.WriteString("://")
	builder.WriteString(host)
	builder.WriteString(":")
	builder.WriteString(port)
	builder.WriteString(route)
	return builder.String()
}

// BuildNetworkGatewayEndpoint builds a URL for an endpoint at the Network Gateway
func BuildNetworkGatewayEndpoint(route string) string {
	protocol := viper.GetString("NetworkGatewayProtocol")
	host := viper.GetString("NetworkGatewayHostname")
	port := viper.GetString("NetworkGatewayPort")
	return BuildURL(protocol, host, port, route)
}
