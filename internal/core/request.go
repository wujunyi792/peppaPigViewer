package core

import (
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"net/http"
)

var (
	ERROR_NEED_RELOGIN = errors.New("需要重新登陆")
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

		<-u.getRequestTicket().C
		return nil
	})
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		if r.StatusCode() != 200 {
			if r.StatusCode() == 302 {
				log.Printf("Request 302 get when %s %s", r.Request.Method, r.Request.URL)
				return ERROR_NEED_RELOGIN
			}
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
