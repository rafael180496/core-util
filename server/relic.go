package server

/**
including library based on https://github.com/jessie-codes/echo-relic/blob/master/echorelic.go
**/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

/*StRelicSend : Structure for sending role in new relic*/
type StRelicSend struct {
	Timestamp time.Time   `json:"timestamp"`
	Message   string      `json:"message"`
	LogType   string      `json:"logtype"`
	Entity    string      `json:"entity.name"`
	Project   string      `json:"project"`
	Scope     string      `json:"scope"`
	Payload   interface{} `json:"payload"`
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
func (e *EchoRelic) Send(tp, scope, message, project string, payload interface{}) {
	SendLogRelic(tp, scope, message, e.nameApp, project, e.licenseKey, payload)
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
func RequestNew(tp, scope, message, appName, project, licenseKey string, payload interface{}) (*http.Request, error) {
	body, err := json.Marshal(StRelicSend{
		Timestamp: time.Now(),
		Message:   message,
		LogType:   tp,
		Entity:    appName,
		Scope:     scope,
		Project:   project,
		Payload:   payload,
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
func SendLogRelic(tp, scope, message, appName, project, licenseKey string, payload interface{}) {
	client := &http.Client{}
	request, err := RequestNew(tp, scope, message, appName, project, licenseKey, payload)
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
}
