package server

/**
including library based on https://github.com/jessie-codes/echo-relic/blob/master/echorelic.go
**/
import (
	echo "github.com/labstack/echo/v4"
	newrelic "github.com/newrelic/go-agent"
)

/*EchoRelic : structure to consume relic en go with echo*/
type EchoRelic struct {
	app newrelic.Application
}

/*NewEchoRelic : creating an instance in relic*/
func NewEchoRelic(appName, licenseKey string) (*EchoRelic, error) {
	config := newrelic.NewConfig(appName, licenseKey)
	app, err := newrelic.NewApplication(config)
	if err != nil {
		return nil, err
	}
	return &EchoRelic{
		app: app,
	}, nil
}

/*Transaction : service transactions*/
func (e *EchoRelic) Transaction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Request().Method + " " + c.Path()
		txn := e.app.StartTransaction(name, c.Response().Writer, c.Request())
		txn.AddAttribute("RealIP", c.RealIP())
		txn.AddAttribute("IsTLS", c.IsTLS())
		txn.AddAttribute("IsWebSocket", c.IsWebSocket())
		txn.AddAttribute("Query", c.QueryString())
		defer txn.End()
		next(c)
		return nil
	}
}
