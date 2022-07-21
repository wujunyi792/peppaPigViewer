package core

import (
	"strconv"
	"time"
)

var (
	JwHost         = "newjw.hdu.edu.cn"
	JwBase         = "https://newjw.hdu.edu.cn/jwglxt"
	JwLoginFailUri = "/jwglxt/xtgl/login_slogin.html"

	JwApiCourseChosen = "/xsxk/zzxkyzb_cxZzxkYzbChoosedDisplay.html?"
	JwDisplayPage     = "/xsxk/zzxkyzb_cxZzxkYzbDisplay.html?"
	JwApiCourseList   = "/xsxk/zzxkyzb_cxZzxkYzbPartDisplay.html?"
	JwApiCourseDetail = "/xsxk/zzxkyzbjk_cxJxbWithKchZzxkYzb.html?"
	JwApiChooseCourse = "/xsxk/zzxkyzbjk_xkBcZyZzxkYzb.html?"
)

// JwApiHome 主页api
type JwApiHome string

func (JwApiHome) GetString() string {
	return "/xtgl/index_initMenu.html?jsdm=xs&_t=" + strconv.Itoa(int(time.Now().UnixMilli()))
}
