package node

import (
	"errors"
	"github.com/pquerna/otp/totp"
	log "github.com/sirupsen/logrus"
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/openapi"
	"github.com/xueqianLu/twitter-bee/rapid"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Node struct {
	api         *openapi.OpenAPI
	service     *handleService
	getBalancer rapid.ServiceBalancer
	available   bool
	quit        chan struct{}
}

func newTransport(conf *config.Config) *http.Transport {
	proxyURL, err := url.Parse(conf.Proxy)
	if len(conf.Proxy) > 0 && err == nil {
		return &http.Transport{
			Proxy:               http.ProxyURL(proxyURL),
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
			DisableCompression:  true,
			DialContext: (&net.Dialer{
				Timeout:   20 * time.Second,
				KeepAlive: 20 * time.Second,
			}).DialContext,
		}
	} else {
		return &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
			DisableCompression:  true,
			DialContext: (&net.Dialer{
				Timeout:   20 * time.Second,
				KeepAlive: 20 * time.Second,
			}).DialContext,
		}
	}
}

func generateTOTP(secret string) (string, error) {
	return totp.GenerateCode(secret, time.Now())
}

func NewNode(conf *config.Config) (*Node, error) {
	n := new(Node)
	keylist := getKeyList(conf.KeyList)
	if len(keylist) == 0 {
		return nil, errors.New("no key found")
	}
	httpClient := &http.Client{
		Timeout:   20 * time.Second,
		Transport: newTransport(conf),
	}
	n.getBalancer = rapid.GetAllServices(keylist, httpClient)
	n.service = newService(n, conf)

	n.available = true
	api := openapi.NewOpenAPI(conf, n.service)
	n.api = api
	n.quit = make(chan struct{})
	return n, nil
}

func (n *Node) Start() error {
	// start openapi server.
	if err := n.api.Run(); err != nil {
		log.WithError(err).Error("start openapi server failed")
		return err
	}

	return nil
}

func (n *Node) Stop() {
	close(n.quit)
}

func (n *Node) Balancer() rapid.ServiceBalancer {
	return n.getBalancer
}
