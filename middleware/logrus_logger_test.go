package middleware

import (
	"bytes"
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLogrusLogger(t *testing.T) {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{}
	logger.Level = logrus.DebugLevel
	buf := new(bytes.Buffer)
	logger.Out = buf

	e := echo.New()
	req := test.NewRequest(echo.GET, "https://github.com/o1egl/echox", nil)
	req.Header().Add(echo.HeaderXRealIP, "127.0.0.1")
	req.Header().Set("Referer", "https://github.com/")
	req.Header().Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:47.0) Gecko/20100101 Firefox/47.0")

	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)

	h := LogrusLogger(logger)(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})

	// Status 2xx
	h(c)
	assert.Contains(t, buf.String(), "method=GET path=\"/o1egl/echox\" referer=\"https://github.com/\" remote_ip=127.0.0.1 rx_bytes=0 status=200 tx_bytes=4 uri=\"https://github.com/o1egl/echox\" user_agent=\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:47.0) Gecko/20100101 Firefox/47.0\"")
	buf.Reset()

	// Status 3xx
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	h = LogrusLogger(logger)(func(c echo.Context) error {
		return c.String(http.StatusTemporaryRedirect, "test")
	})
	h(c)
	assert.Contains(t, buf.String(), "method=GET path=\"/o1egl/echox\" referer=\"https://github.com/\" remote_ip=127.0.0.1 rx_bytes=0 status=307 tx_bytes=4 uri=\"https://github.com/o1egl/echox\" user_agent=\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:47.0) Gecko/20100101 Firefox/47.0\"")
	buf.Reset()

	// Status 4xx
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	h = LogrusLogger(logger)(func(c echo.Context) error {
		return c.String(http.StatusNotFound, "test")
	})
	h(c)
	assert.Contains(t, buf.String(), "method=GET path=\"/o1egl/echox\" referer=\"https://github.com/\" remote_ip=127.0.0.1 rx_bytes=0 status=404 tx_bytes=4 uri=\"https://github.com/o1egl/echox\" user_agent=\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:47.0) Gecko/20100101 Firefox/47.0\"")
	buf.Reset()

	// Status 5xx with empty path
	rec = test.NewResponseRecorder()
	c = e.NewContext(req, rec)
	h = LogrusLogger(logger)(func(c echo.Context) error {
		return errors.New("error")
	})
	h(c)
	assert.Contains(t, buf.String(), "method=GET path=\"/o1egl/echox\" referer=\"https://github.com/\" remote_ip=127.0.0.1 rx_bytes=0 status=500 tx_bytes=21 uri=\"https://github.com/o1egl/echox\" user_agent=\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:47.0) Gecko/20100101 Firefox/47.0\"")
	buf.Reset()
}
