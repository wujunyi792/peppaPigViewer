package dto

import (
	"fmt"
	"reflect"
	"strings"
)

type CourseListResp struct {
	TmpList []struct {
		Blyxrs             string `json:"blyxrs"`
		Blzyl              string `json:"blzyl"`
		Cxbj               string `json:"cxbj"`
		Date               string `json:"date"`
		DateDigit          string `json:"dateDigit"`
		DateDigitSeparator string `json:"dateDigitSeparator"`
		Day                string `json:"day"`
		Fxbj               string `json:"fxbj"`
		Jgpxzd             string `json:"jgpxzd"`
		JxbId              string `json:"jxb_id"`
		Jxbmc              string `json:"jxbmc"`
		Jxbzls             string `json:"jxbzls"`
		Kch                string `json:"kch"`
		KchId              string `json:"kch_id"`
		Kcmc               string `json:"kcmc"`
		Kcrow              string `json:"kcrow"`
		Kklxdm             string `json:"kklxdm"`
		Kzmc               string `json:"kzmc"`
		Listnav            string `json:"listnav"`
		LocaleKey          string `json:"localeKey"`
		Month              string `json:"month"`
		PageTotal          int    `json:"pageTotal"`
		Pageable           bool   `json:"pageable"`
		QueryModel         struct {
			CurrentPage   int           `json:"currentPage"`
			CurrentResult int           `json:"currentResult"`
			EntityOrField bool          `json:"entityOrField"`
			Limit         int           `json:"limit"`
			Offset        int           `json:"offset"`
			PageNo        int           `json:"pageNo"`
			PageSize      int           `json:"pageSize"`
			ShowCount     int           `json:"showCount"`
			Sorts         []interface{} `json:"sorts"`
			TotalCount    int           `json:"totalCount"`
			TotalPage     int           `json:"totalPage"`
			TotalResult   int           `json:"totalResult"`
		} `json:"queryModel"`
		Rangeable   bool   `json:"rangeable"`
		Sftj        string `json:"sftj"`
		TotalResult string `json:"totalResult"`
		UserModel   struct {
			Monitor    bool   `json:"monitor"`
			RoleCount  int    `json:"roleCount"`
			RoleKeys   string `json:"roleKeys"`
			RoleValues string `json:"roleValues"`
			Status     int    `json:"status"`
			Usable     bool   `json:"usable"`
		} `json:"userModel"`
		Xf         string `json:"xf"`
		Xxkbj      string `json:"xxkbj"`
		Year       string `json:"year"`
		Yxzrs      string `json:"yxzrs"`
		DetailList *CourseDetail
		HaveSet    bool // 是否可选（余量充足）
	} `json:"tmpList"`
	Sfxsjc string `json:"sfxsjc"`
}

type CourseDetail struct {
	Date               string `json:"date"`
	DateDigit          string `json:"dateDigit"`
	DateDigitSeparator string `json:"dateDigitSeparator"`
	Day                string `json:"day"`
	DoJxbId            string `json:"do_jxb_id"`
	Fxbj               string `json:"fxbj"`
	Jgpxzd             string `json:"jgpxzd"`
	Jsxx               string `json:"jsxx"`
	JxbId              string `json:"jxb_id"`
	Jxbrl              string `json:"jxbrl"`
	Jxdd               string `json:"jxdd"`
	Jxms               string `json:"jxms"`
	Kclbmc             string `json:"kclbmc"`
	Kcxzmc             string `json:"kcxzmc"`
	Kkxymc             string `json:"kkxymc"`
	Listnav            string `json:"listnav"`
	LocaleKey          string `json:"localeKey"`
	Month              string `json:"month"`
	PageTotal          int    `json:"pageTotal"`
	Pageable           bool   `json:"pageable"`
	QueryModel         struct {
		CurrentPage   int           `json:"currentPage"`
		CurrentResult int           `json:"currentResult"`
		EntityOrField bool          `json:"entityOrField"`
		Limit         int           `json:"limit"`
		Offset        int           `json:"offset"`
		PageNo        int           `json:"pageNo"`
		PageSize      int           `json:"pageSize"`
		ShowCount     int           `json:"showCount"`
		Sorts         []interface{} `json:"sorts"`
		TotalCount    int           `json:"totalCount"`
		TotalPage     int           `json:"totalPage"`
		TotalResult   int           `json:"totalResult"`
	} `json:"queryModel"`
	Rangeable   bool   `json:"rangeable"`
	Sksj        string `json:"sksj"`
	TotalResult string `json:"totalResult"`
	UserModel   struct {
		Monitor    bool   `json:"monitor"`
		RoleCount  int    `json:"roleCount"`
		RoleKeys   string `json:"roleKeys"`
		RoleValues string `json:"roleValues"`
		Status     int    `json:"status"`
		Usable     bool   `json:"usable"`
	} `json:"userModel"`
	XqhId string `json:"xqh_id"`
	Xqumc string `json:"xqumc"`
	Year  string `json:"year"`
	Yqmc  string `json:"yqmc"`
}

type CourseChosenResp struct {
	Bdzcbj             string `json:"bdzcbj"`
	Cxbj               string `json:"cxbj"`
	Date               string `json:"date"`
	DateDigit          string `json:"dateDigit"`
	DateDigitSeparator string `json:"dateDigitSeparator"`
	Day                string `json:"day"`
	DoJxbId            string `json:"do_jxb_id"`
	IsInxksj           string `json:"isInxksj"`
	Jgpxzd             string `json:"jgpxzd"`
	Jsxx               string `json:"jsxx"`
	JxbId              string `json:"jxb_id"`
	Jxbmc              string `json:"jxbmc"`
	Jxbrs              string `json:"jxbrs"`
	Jxbzls             string `json:"jxbzls"`
	Jxdd               string `json:"jxdd"`
	Kch                string `json:"kch"`
	KchId              string `json:"kch_id"`
	Kcmc               string `json:"kcmc"`
	Kklxdm             string `json:"kklxdm"`
	Kklxmc             string `json:"kklxmc"`
	Kklxpx             string `json:"kklxpx"`
	Krrl               string `json:"krrl"`
	Listnav            string `json:"listnav"`
	LocaleKey          string `json:"localeKey"`
	Month              string `json:"month"`
	PageTotal          int    `json:"pageTotal"`
	Pageable           bool   `json:"pageable"`
	QueryModel         struct {
		CurrentPage   int           `json:"currentPage"`
		CurrentResult int           `json:"currentResult"`
		EntityOrField bool          `json:"entityOrField"`
		Limit         int           `json:"limit"`
		Offset        int           `json:"offset"`
		PageNo        int           `json:"pageNo"`
		PageSize      int           `json:"pageSize"`
		ShowCount     int           `json:"showCount"`
		Sorts         []interface{} `json:"sorts"`
		TotalCount    int           `json:"totalCount"`
		TotalPage     int           `json:"totalPage"`
		TotalResult   int           `json:"totalResult"`
	} `json:"queryModel"`
	Qz          string `json:"qz"`
	Rangeable   bool   `json:"rangeable"`
	Rlkz        string `json:"rlkz"`
	Rlzlkz      string `json:"rlzlkz"`
	Rwlx        string `json:"rwlx"`
	Sfktk       string `json:"sfktk"`
	Sfxkbj      string `json:"sfxkbj"`
	Sksj        string `json:"sksj"`
	Sxbj        string `json:"sxbj,omitempty"`
	TKchId      string `json:"t_kch_id"`
	Tktjrs      string `json:"tktjrs"`
	TotalResult string `json:"totalResult"`
	UserModel   struct {
		Monitor    bool   `json:"monitor"`
		RoleCount  int    `json:"roleCount"`
		RoleKeys   string `json:"roleKeys"`
		RoleValues string `json:"roleValues"`
		Status     int    `json:"status"`
		Usable     bool   `json:"usable"`
	} `json:"userModel"`
	Xf     string `json:"xf"`
	Xkgz   string `json:"xkgz"`
	XkkzId string `json:"xkkz_id"`
	Xxkbj  string `json:"xxkbj"`
	Year   string `json:"year"`
	Yxzrs  string `json:"yxzrs"`
	Zckz   string `json:"zckz"`
	Zixf   string `json:"zixf"`
	Zntgpk string `json:"zntgpk"`
	Zy     string `json:"zy"`
}

type FindClassReq struct {
	FilterList []string `form:"filter_list"`

	Rwlx         string `form:"rwlx"`
	Xkly         string `form:"xkly"`
	Bklx_id      string `form:"bklx_id"`
	Sfkkjyxdxnxq string `form:"sfkkjyxdxnxq"`
	Xqh_id       string `form:"xqh_id"`
	Jg_id        string `form:"jg_id"`
	Njdm_id_1    string `form:"njdm_id_1"`
	Zyh_id_1     string `form:"zyh_id_1"`
	Zyh_id       string `form:"zyh_id"`
	Zyfx_id      string `form:"zyfx_id"`
	Njdm_id      string `form:"njdm_id"`
	Bh_id        string `form:"bh_id"`
	Xbm          string `form:"xbm"`
	Xslbdm       string `form:"xslbdm"`
	Mzm          string `form:"mzm"`
	Xz           string `form:"xz"`
	Ccdm         string `form:"ccdm"`
	Xsbj         string `form:"xsbj"`
	Sfkknj       string `form:"sfkknj"`
	Sfkkzy       string `form:"sfkkzy"`
	Kzybkxy      string `form:"kzybkxy"`
	Sfznkx       string `form:"sfznkx"`
	Zdkxms       string `form:"zdkxms"`
	Sfkxq        string `form:"sfkxq"`
	Sfkcfx       string `form:"sfkcfx"`
	Kkbk         string `form:"kkbk"`
	Kkbkdj       string `form:"kkbkdj"`
	Sfkgbcx      string `form:"sfkgbcx"`
	Sfrxtgkcxd   string `form:"sfrxtgkcxd"`
	Tykczgxdcs   string `form:"tykczgxdcs"`
	Xkxnm        string `form:"xkxnm"`
	Xkxqm        string `form:"xkxqm"`
	Kklxdm       string `form:"kklxdm"`
	Bbhzxjxb     string `form:"bbhzxjxb"`
	Rlkz         string `form:"rlkz"`
	Xkzgbj       string `form:"xkzgbj"`
	Kspage       string `form:"kspage"`
	Jspage       string `form:"jspage"`
	Jxbzb        string `form:"jxbzb"`
}

type GetDisplayReq struct {
	Xkkz_id string `form:"xkkz_id"`
	Xszxzt  string `form:"xszxzt"`
	Kspage  string `form:"kspage"`
	Jspage  string `form:"jspage"`
}

type GetCourseDetailReq struct {
	FilterList []string `form:"filter_list"`

	Rwlx         string `form:"rwlx"`
	Xkly         string `form:"xkly"`
	BklxId       string `form:"bklx_id"`
	Sfkkjyxdxnxq string `form:"sfkkjyxdxnxq"`
	XqhId        string `form:"xqh_id"`
	JgId         string `form:"jg_id"`
	ZyhId        string `form:"zyh_id"`
	ZyfxId       string `form:"zyfx_id"`
	NjdmId       string `form:"njdm_id"`
	BhId         string `form:"bh_id"`
	Xbm          string `form:"xbm"`
	Xslbdm       string `form:"xslbdm"`
	Mzm          string `form:"mzm"`
	Xz           string `form:"xz"`
	Bbhzxjxb     string `form:"bbhzxjxb"`
	Ccdm         string `form:"ccdm"`
	Xsbj         string `form:"xsbj"`
	Sfkknj       string `form:"sfkknj"`
	Sfkkzy       string `form:"sfkkzy"`
	Kzybkxy      string `form:"kzybkxy"`
	Sfznkx       string `form:"sfznkx"`
	Zdkxms       string `form:"zdkxms"`
	Sfkxq        string `form:"sfkxq"`
	Sfkcfx       string `form:"sfkcfx"`
	Kkbk         string `form:"kkbk"`
	Kkbkdj       string `form:"kkbkdj"`
	Xkxnm        string `form:"xkxnm"`
	Xkxqm        string `form:"xkxqm"`
	Xkxskcgskg   string `form:"xkxskcgskg"`
	Rlkz         string `form:"rlkz"`
	Kklxdm       string `form:"kklxdm"`
	KchId        string `form:"kch_id"`
	Jxbzcxskg    string `form:"jxbzcxskg"`
	XkkzId       string `form:"xkkz_id"`
	Cxbj         string `form:"cxbj"`
	Fxbj         string `form:"fxbj"`
}

type ChooseClassPrvReq struct {
	JxbIds string `form:"jxb_ids"`
	KchId  string `form:"kch_id"`
	Kcmc   string `form:"kcmc"`
	Rwlx   string `form:"rwlx"`
	Rlkz   string `form:"rlkz"`
	Rlzlkz string `form:"rlzlkz"` //?field有
	Sxbj   string `form:"sxbj"`   //?if(rlkz=="1" || rlzlkz=="1"){sxbj = "1";}else{ sxbj = "0";}
	Xxkbj  string `form:"xxkbj"`  //?有先行课
	Qz     string `form:"qz"`     //?权重
	Cxbj   string `form:"cxbj"`
	XkkzId string `form:"xkkz_id"`
	NjdmId string `form:"njdm_id"`
	ZyhId  string `form:"zyh_id"`
	Kklxdm string `form:"kklxdm"`
	Xklc   string `form:"xklc"` //?选课轮数 flied有
	Xkxnm  string `form:"xkxnm"`
	Xkxqm  string `form:"xkxqm"`
}

type GetChosenCourseReq struct {
	JgId   string `form:"jg_id"`
	ZyhId  string `form:"zyh_id"`
	NjdmId string `form:"njdm_id"`
	ZyfxId string `form:"zyfx_id"`
	BhId   string `form:"bh_id"`
	Xz     string `form:"xz"`
	Ccdm   string `form:"ccdm"`
	Xkxnm  string `form:"xkxnm"`
	Xkxqm  string `form:"xkxqm"`
	Xkly   string `form:"xkly"`
}

func (c *GetChosenCourseReq) MakeForm() string {
	var builder strings.Builder
	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(*c)
	for k := 0; k < t.NumField(); k++ {
		builder.WriteString(fmt.Sprintf("%s=%v", t.Field(k).Tag.Get("form"), v.Field(k).Interface()))
		if k != t.NumField()-1 {
			builder.WriteString("&")
		}
	}
	return builder.String()
}

func MakeGetChosenClassReq(field map[string]string) *GetChosenCourseReq {
	var req GetChosenCourseReq
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(&req)
	v = v.Elem()
	for k := 0; k < t.NumField(); k++ {
		if v.Field(k).Type().Kind() != reflect.String {
			continue
		}
		v.Field(k).SetString(field[t.Field(k).Tag.Get("form")])
	}
	req.JgId = field["jg_id_1"]
	return &req
}

func (c *FindClassReq) MakeForm() string {
	var builder strings.Builder
	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(*c)
	for k := 0; k < t.NumField(); k++ {
		switch v.Field(k).Type().Kind() {
		case reflect.Slice:
			{
				count := v.Field(k).Len() // Len() 函数

				for i := 0; i < count; i++ {
					child := v.Field(k).Index(i) // Index() 函数
					s := child.String()
					builder.WriteString(fmt.Sprintf("%s[%d]=%s", t.Field(k).Tag.Get("form"), i, s))
					if i != count-1 {
						builder.WriteString("&")
					}
				}
			}
		default:
			builder.WriteString(fmt.Sprintf("%s=%v", t.Field(k).Tag.Get("form"), v.Field(k).Interface()))
		}
		if k != t.NumField()-1 {
			builder.WriteString("&")
		}
	}
	return builder.String()
}

func (c *GetCourseDetailReq) MakeForm() string {
	var builder strings.Builder
	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(*c)
	for k := 0; k < t.NumField(); k++ {
		switch v.Field(k).Type().Kind() {
		case reflect.Slice:
			{
				count := v.Field(k).Len() // Len() 函数

				for i := 0; i < count; i++ {
					child := v.Field(k).Index(i) // Index() 函数
					s := child.String()
					builder.WriteString(fmt.Sprintf("%s[%d]=%s", t.Field(k).Tag.Get("form"), i, s))
					if i != count-1 {
						builder.WriteString("&")
					}
				}
			}
		default:
			builder.WriteString(fmt.Sprintf("%s=%v", t.Field(k).Tag.Get("form"), v.Field(k).Interface()))
		}
		if k != t.NumField()-1 {
			builder.WriteString("&")
		}
	}
	return builder.String()
}

func MakeFindClassReq(field map[string]string) *FindClassReq { //获取基本参数
	var req FindClassReq
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(&req)
	v = v.Elem()
	for k := 0; k < t.NumField(); k++ {
		if v.Field(k).Type().Kind() != reflect.String {
			continue
		}
		v.Field(k).SetString(field[t.Field(k).Tag.Get("form")])
	}
	req.Kspage = "1"
	req.Jspage = "10"
	req.Kklxdm = field["firstKklxdm"] //修改
	//req.Kklxdm = "10"
	req.Jg_id = field["jg_id_1"]
	return &req
}

func MakeChooseClassPrvReq(field map[string]string) *ChooseClassPrvReq {
	var req ChooseClassPrvReq
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(&req)
	v = v.Elem()
	for k := 0; k < t.NumField(); k++ {
		if v.Field(k).Type().Kind() != reflect.String {
			continue
		}
		v.Field(k).SetString(field[t.Field(k).Tag.Get("form")])
	}
	req.Kklxdm = field["firstKklxdm"] //修改
	//req.Kklxdm = "10"
	req.XkkzId = field["firstXkkzId"]
	req.Qz = "0"
	if req.Rlzlkz == "1" || req.Rlkz == "1" {
		req.Sxbj = "1"
	} else {
		req.Sxbj = "0"
	}
	return &req
}

func MakeGetClassDetailReq(field map[string]string) *GetCourseDetailReq {
	var req GetCourseDetailReq
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(&req)
	v = v.Elem()
	for k := 0; k < t.NumField(); k++ {
		if v.Field(k).Type().Kind() != reflect.String {
			continue
		}
		v.Field(k).SetString(field[t.Field(k).Tag.Get("form")]) //根据预设的键值获取field中对应标签id的值
	}
	req.Kklxdm = field["firstKklxdm"] //修改
	//req.Kklxdm = "10"
	req.XkkzId = field["firstXkkzId"]
	req.JgId = field["jg_id_1"]
	req.Cxbj = "0"
	req.Fxbj = "0"
	return &req
}

func MakeGetDisplayReq(field map[string]string) *GetDisplayReq {
	var req GetDisplayReq
	req.Xkkz_id = field["firstXkkzId"]
	req.Xszxzt = field["xszxzt"]
	req.Kspage = "0"
	req.Jspage = "0"
	return &req
}

func (c *GetDisplayReq) MakeForm() string {
	var builder strings.Builder
	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(*c)
	for k := 0; k < t.NumField(); k++ {
		builder.WriteString(fmt.Sprintf("%s=%v", t.Field(k).Tag.Get("form"), v.Field(k).Interface()))
		if k != t.NumField()-1 {
			builder.WriteString("&")
		}
	}
	return builder.String()
}

func (c *ChooseClassPrvReq) MakeForm() string {
	var builder strings.Builder
	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(*c)
	for k := 0; k < t.NumField(); k++ {
		builder.WriteString(fmt.Sprintf("%s=%v", t.Field(k).Tag.Get("form"), v.Field(k).Interface()))
		if k != t.NumField()-1 {
			builder.WriteString("&")
		}
	}
	return builder.String()
}
