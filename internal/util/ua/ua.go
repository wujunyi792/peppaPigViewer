package ua

import browser "github.com/EDDYCJY/fake-useragent"

func GetUA() string {
	return browser.Chrome()
}
