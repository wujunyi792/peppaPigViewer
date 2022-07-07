package main

import hdu_cas_helper "github.com/wujunyi792/hdu-cas-helper"

var cookie *hdu_cas_helper.NewJW

func login() *hdu_cas_helper.NewJW {
	ticket := hdu_cas_helper.CasPasswordLogin("20081131", "")
	return hdu_cas_helper.NewJWLogin(ticket)
}

func getCookie() string {
	if cookie == nil {
		cookie = login()
	}
	return cookie.GetCookie()
}
