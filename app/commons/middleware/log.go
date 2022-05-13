package middleware

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func AddMetadata() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			host, _, _ := net.SplitHostPort(req.RemoteAddr)
			transactionID := gettransactionID(req, res)
			c.Set("transactionID", transactionID)

			logger := zerolog.Ctx(req.Context())
			logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("uri", req.Method+" "+req.URL.String()).
					Str("host_ip", host).
					Str("transaction_id", transactionID).
					Str("user_agent", req.UserAgent()).
					Str("referer", req.Referer()).
					Str("remote_addr", req.RemoteAddr)
			})

			return next(c)
		}
	}
}

func SetLoggerInContext(logger zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log := logger.With().Logger()
			c.SetRequest(
				c.Request().WithContext(
					log.WithContext(c.Request().Context()),
				),
			)
			return next(c)
		}
	}
}

func gettransactionID(req *http.Request, res *echo.Response) string {
	transactionID := transactionIDOfCtx(req.Context())
	if transactionID == "" {
		transactionID = res.Header().Get("X-Transaction-Id")
		if transactionID == "" {
			transactionID = generateTID()
		}
	}
	return transactionID
}

func transactionIDOfCtx(ctx context.Context) string {
	transactionID, ok := ctx.Value("tid").(string)
	if !ok {
		transactionID = ""
	}
	return transactionID
}

func generateTID() string {
	id, err := uuid.New()
	if err != nil {
		panic(err)
	}
	return "starwars-api-" + fmt.Sprintf("%x", id)
}
