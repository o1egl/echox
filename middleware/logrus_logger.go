package middleware

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"net"
	"strconv"
	"time"
)

func LogrusLogger(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()

			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			ra := req.RemoteAddress()
			if ip := req.Header().Get(echo.HeaderXRealIP); ip != "" {
				ra = ip
			} else if ip = req.Header().Get(echo.HeaderXForwardedFor); ip != "" {
				ra = ip
			} else {
				ra, _, _ = net.SplitHostPort(ra)
			}

			path := req.URL().Path()
			if path == "" {
				path = "/"
			}
			status := res.Status()
			latency := stop.Sub(start).Nanoseconds() / 1000
			latency_human := stop.Sub(start).String()
			rx_bytes := req.Header().Get(echo.HeaderContentLength)
			if rx_bytes == "" {
				rx_bytes = "0"
			}
			tx_bytes := strconv.FormatInt(res.Size(), 10)

			entry := logger.WithFields(logrus.Fields{
				"host":          req.Host(),
				"uri":           req.URI(),
				"method":        req.Method(),
				"path":          path,
				"remote_ip":     ra,
				"referer":       req.Referer(),
				"user_agent":    req.UserAgent(),
				"status":        status,
				"latency":       latency,
				"latency_human": latency_human,
				"rx_bytes":      rx_bytes,
				"tx_bytes":      tx_bytes,
			})

			switch {
			case status >= 500:
				entry.Error()
			default:
				entry.Info()
			}

			return nil
		}
	}
}
