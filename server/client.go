package server

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	utl "github.com/rafael180496/core-util/utility"
)

/*BasicAuth : autentificacion basica de un client http */
func BasicAuth(username, password string) string {
	auth := fmt.Sprintf("%s:%s", username, password)
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

/*TokenHeaders : autentificacion por medio de token bearer*/
func TokenHeaders(token string) map[string]string {
	Headers := make(map[string]string)
	Headers["Content-Type"] = "application/json"
	if utl.IsNilStr(token) {
		Headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	}
	return Headers
}

/*NewNetClient : crea una instancia de http client*/
func NewNetClient(timeout int) *http.Client {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(timeout) * time.Second,
		}).Dial,
		TLSHandshakeTimeout: time.Duration(timeout) * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * time.Duration(timeout),
		Transport: netTransport,
	}
	return netClient
}

/*SendRequest : envia una peticion de una api rest*/
func (c *Client) SendByte(req Request) ([]byte, int, error) {
	resp, err := c.Send(req)
	if err != nil {
		return nil, 0, err
	}
	return utl.StrtoByte(resp.Body), resp.StatusCode, nil
}

/*AddQueryParameters : agrega los parametros para un param querie*/
func AddQueryParameters(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}

/*BuildRequestObject : Contruccion de objeto http*/
func BuildRequestObject(request Request) (*http.Request, error) {
	if len(request.QueryParams) != 0 {
		request.BaseURL = AddQueryParameters(request.BaseURL, request.QueryParams)
	}
	req, err := http.NewRequest(string(request.Method), request.BaseURL, bytes.NewBuffer(request.Body))
	if err != nil {
		return req, err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	_, exists := req.Header["Content-Type"]
	if len(request.Body) > 0 && !exists {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, err
}

/*MakeRequest :  obtiene la respuesta de una peticion*/
func MakeRequest(req *http.Request) (*http.Response, error) {
	return DefaultClient.HTTPClient.Do(req)
}

/*BuildResponse : construccion de una respuesta*/
func BuildResponse(res *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	response := Response{
		StatusCode: res.StatusCode,
		Body:       string(body),
		Headers:    res.Header,
	}
	return &response, err
}

/*Send : envia una peticion a un servicio*/
func Send(request Request) (*Response, error) {
	return DefaultClient.Send(request)
}

/*MakeRequest : envia la peticion de un servicio con un cliente por defecto*/
func (c *Client) MakeRequest(req *http.Request) (*http.Response, error) {
	return c.HTTPClient.Do(req)
}

/*Send : envia una peticion a un servicio y regresa la respuesta*/
func (c *Client) Send(request Request) (*Response, error) {
	req, err := BuildRequestObject(request)
	if err != nil {
		return nil, err
	}
	res, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	return BuildResponse(res)
}

/*NewClient : crea una instancia client http*/
func NewClient(timeout int) *Client {
	return &Client{
		HTTPClient: NewNetClient(timeout),
	}
}
