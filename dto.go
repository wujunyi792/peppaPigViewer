package main

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
