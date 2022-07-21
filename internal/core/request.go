package core

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

// newRequestClient 初始化请求客户端
func (u *User) newRequestClient() *resty.Client {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"User-Agent": u.config.ua,
		"Cookie":     u.getCookie(),
	})
	//client.SetDebug(true)
	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		if u.Error() != nil {
			return u.Error()
		}
		log.Printf("【Req】 %s", request.URL)

		err := u.getRequestRate().Wait(context.Background())
		return err
	})
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		if r.StatusCode() != 200 {
			log.Printf("Request get a status code %d when %s %s", r.StatusCode(), r.Request.Method, r.Request.URL)
		}
		return nil
	})
	client.SetRedirectPolicy(resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
		if req.URL.Path == JwLoginFailUri {
			u.e = errors.New("登陆失效")
			return http.ErrUseLastResponse
		}
		return nil
	}))
	return client
}
