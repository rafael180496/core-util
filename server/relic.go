package server

/**
including library based on https://github.com/jessie-codes/echo-relic/blob/master/echorelic.go
**/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

/*StRelicSend : Structure for sending role in new relic*/
type StRelicSend struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	LogType   string    `json:"logtype"`
	Entity    string    `json:"entity.name"`
	Module    string    `json:"module"`
}

const (
	LgInfo  = "info"
	LgSilly = "silly"
	LgTrace = "trace"
	LgDebug = "debug"
	LgWarn  = "warn"
	LgError = "error"
	LgFatal = "fatal"
	/*UrlRelic :  base url for logs*/
	UrlRelic = "https://log-api.newrelic.com/log/v1"
)

/*EchoRelic : structure to consume relic en go with echo*/
type EchoRelic struct {
	app        *newrelic.Application
	nameApp    string
	licenseKey string
}

/*Send : sending logs to new relic*/
func (e *EchoRelic) Send(tp, module, message string) {
	SendLogRelic(tp, module, message, e.nameApp, e.licenseKey)
}

/*NewEchoRelic : creating an instance in relic*/
func NewEchoRelic(appName, licenseKey string) (*EchoRelic, error) {
	app, err := newrelic.NewApplication(newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey))
	if err != nil {
		return nil, err
	}
	return &EchoRelic{
		app:        app,
		nameApp:    appName,
		licenseKey: licenseKey,
	}, nil
}

/*Transaction : service transactions*/
func (e *EchoRelic) Transaction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Request().Method + " " + c.Path()
		txn := e.app.StartTransaction(name)
		txn.AddAttribute("RealIP", c.RealIP())
		txn.AddAttribute("IsTLS", c.IsTLS())
		txn.AddAttribute("IsWebSocket", c.IsWebSocket())
		txn.AddAttribute("Query", c.QueryString())
		defer txn.End()
		next(c)
		return nil
	}
}

/*RequestNew : structure to send logs to a new relic application*/
func RequestNew(tp, module, message, appName, licenseKey string) (*http.Request, error) {
	body, err := json.Marshal(StRelicSend{
		Timestamp: time.Now(),
		Message:   message,
		LogType:   tp,
		Entity:    appName,
		Module:    module,
	})
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", UrlRelic, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Api-key", licenseKey)
	return request, nil
}

/*SendLogRelic : send a log to new relic*/
func SendLogRelic(tp, module, message, appName, licenseKey string) {
	client := &http.Client{}
	request, err := RequestNew(tp, module, message, appName, licenseKey)
	if err != nil {
		fmt.Printf("\n[SendLogRelicErr],%v", err.Error())
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("\n[SendLogRelicErr],%v", err.Error())
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("\n[SendLogRelicErr],%v", err.Error())
		return
	}
	if response.StatusCode != 200 && response.StatusCode != 202 {
		fmt.Printf("\n[SendLogRelic],body[%v]code[%v]\n", string(body), response.StatusCode)
	}
}
