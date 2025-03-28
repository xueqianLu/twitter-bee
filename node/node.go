package node

import (
	"errors"
	twitterscraper "github.com/imperatrona/twitter-scraper"
	"github.com/pquerna/otp/totp"
	log "github.com/sirupsen/logrus"
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/openapi"
	"github.com/xueqianLu/twitter-bee/types"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Node struct {
	api       *openapi.OpenAPI
	spider    *twitterscraper.Scraper
	account   types.TwitterAccount
	RateLimit *rate.Limiter
	service   *handleService
	available bool
	quit      chan struct{}
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
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
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
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		}
	}
}

func generateTOTP(secret string) (string, error) {
	return totp.GenerateCode(secret, time.Now())
}

func NewNode(conf *config.Config, user string) (*Node, error) {
	n := new(Node)
	userlib := getUserLib(conf.UserLib)
	if acc, ok := userlib[user]; ok {
		n.account = acc
	} else {
		return nil, errors.New("user not found")
	}
	n.available = false
	n.service = newService(n, conf)
	n.spider = twitterscraper.NewWithTransport(newTransport(conf))
	api := openapi.NewOpenAPI(conf, n.service)
	n.api = api
	// set a rate limit to 40 requests per 15 minutes.

	n.RateLimit = rate.NewLimiter(rate.Every(15*time.Minute/40), 40)
	n.quit = make(chan struct{})
	return n, nil
}

func (n *Node) Start() error {
	go n.loop()
	// start openapi server.
	if err := n.api.Run(); err != nil {
		log.WithError(err).Error("start openapi server failed")
		return err
	}

	return nil
}

func (n *Node) checkHealth() bool {
	_, err := n.spider.GetProfile("bitcoin")
	if err != nil {
		return false
	}
	return true
}

func (n *Node) login() error {
	var v = n.account
	var err error
	if v.F2A != "" {
		code, _ := generateTOTP(v.F2A)
		err = n.spider.AutoLogin(v.Username, v.Password, v.Email, code)
	} else {
		err = n.spider.AutoLogin(v.Username, v.Password, v.Email)
	}
	if err != nil {
		log.WithError(err).Error("login failed")
	}
	if v.Token != "" && v.CSRFToken != "" {
		log.WithField("user", v.Username).Info("login with token")
		err = n.bindToken(v.Token, v.CSRFToken)
	}
	return err
}

func (n *Node) bindToken(token, csrf string) error {
	n.spider.SetAuthToken(twitterscraper.AuthToken{
		Token:     token,
		CSRFToken: csrf,
	})
	if n.spider.IsLoggedIn() {
		return nil
	} else {
		return errors.New("login failed")
	}
}

func (n *Node) loop() error {
	keepAlive := time.NewTicker(5 * time.Minute)
	defer keepAlive.Stop()
	login := time.NewTicker(time.Second)
	defer login.Stop()
	for {
		select {
		case <-n.quit:
			return nil
		case <-keepAlive.C:
			n.available = n.checkHealth()
			log.WithField("available", n.available).Info("keep alive")
		case <-login.C:
			if !n.available {
				if err := n.login(); err != nil {
					log.WithError(err).Error("login failed")
					login.Reset(time.Minute)
				} else {
					log.WithField("user", n.account.Username).Info("login success")
					n.available = true
					login.Reset(time.Hour)
					n.RateLimit.SetLimit(rate.Every(15 * time.Minute / 40))
				}
			}
		}
	}
}

func (n *Node) Stop() {
	close(n.quit)
}
