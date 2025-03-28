package openapi

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/types"
	"time"
)

type OpenAPI struct {
	conf    *config.Config
	backend types.ServiceBackend
}

func NewOpenAPI(conf *config.Config, backend types.ServiceBackend) *OpenAPI {
	return &OpenAPI{conf: conf, backend: backend}
}

func (s *OpenAPI) Run() error {
	log.WithField("address", s.conf.ServiceUrl).Info("openapi server start")
	return s.startHttp(s.conf.ServiceUrl)
}

func (s *OpenAPI) startHttp(address string) error {
	router := gin.Default()
	router.Use(cors())
	router.Use(ginLogrus())
	handler := apiHandler{conf: s.conf, backend: s.backend}
	v1 := router.Group("/tbapi/v2")
	{
		follower := v1.Group("/follower")
		follower.POST("/profile", handler.UserProfile)
		follower.POST("/list", handler.FollowerList)
	}
	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ch := make(chan error)
	go func() {
		err := router.Run(address)
		ch <- err
	}()
	time.Sleep(100 * time.Millisecond)
	select {
	case v := <-ch:
		return v
	default:
		return nil
	}
}

// gin use logrus
func ginLogrus() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.WithFields(log.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"query":  c.Request.URL.RawQuery,
			"ip":     c.ClientIP(),
		}).Info("request")
		c.Next()
	}
}

// enable cors
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
