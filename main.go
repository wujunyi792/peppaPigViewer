package main

import (
	"fmt"
	"log"
	"net/url"
	"time"
)

func main() {
	entrance := getEntranceList()

	reqUrl := url.URL{
		Scheme: "https",
		Host:   JWHOST,
		Path:   "/jwglxt",
	}
	flag := false
	lid := ""
	for _, list := range entrance.list {
		if list.name == "自主选课" {
			reqUrl.Path += list.uri
			query := make(url.Values)
			query.Add("gnmkdm", list.id)
			query.Add("layout", "default")
			query.Add("su", getUserID())
			reqUrl.RawQuery = query.Encode()
			flag = true
			lid = list.id
			break
		}
	}
	if !flag {
		log.Fatalln("look like not in correct time")
	}
	p := getCoursePage(reqUrl.String())
	fields := getInputField(p, nil)

	baseQuery := make(url.Values)
	baseQuery.Add("gnmkdm", lid)
	baseQuery.Add("su", getUserID())

	p = getDisplayPage(JWBASE+"/xsxk/zzxkyzb_cxZzxkYzbDisplay.html?"+baseQuery.Encode(), MakeGetDisplayReq(fields))
	fields = getInputField(p, fields)

	chosenCourse := getChosenCourse(baseQuery.Encode(), MakeGetChosenClassReq(fields))

	findClassBaseField := MakeFindClassReq(fields)
	findClassBaseField.FilterList = getTarget()
	list := getCourseList(JWBASE+"/xsxk/zzxkyzb_cxZzxkYzbPartDisplay.html?"+baseQuery.Encode(), findClassBaseField)
	getClassDetailField := MakeGetClassDetailReq(fields)
	for i := 0; i < len(list.TmpList); {
		getClassDetailField.KchId = list.TmpList[i].KchId
		getClassDetailField.FilterList = getTarget()

		time.Sleep(3 * time.Second)
		details := getCourseDetail(JWBASE+"/xsxk/zzxkyzbjk_cxJxbWithKchZzxkYzb.html?"+baseQuery.Encode(), getClassDetailField)
		if *details == nil {
			id := list.TmpList[i].KchId
			for j := 0; j < len(list.TmpList); j++ {
				if list.TmpList[j].KchId == id {
					i++
				}
			}
			continue
		}
		for index, detail := range *details {
			for j := 0; j < len(list.TmpList); j++ {
				if list.TmpList[j].JxbId == detail.JxbId {
					list.TmpList[j].DetailList = &(*details)[index]
					i++
					break
				}
			}
		}
	}

	for i := 0; i < len(list.TmpList) && list.TmpList[i].DetailList == nil; i++ {
		fmt.Printf("%s 课程号 %s 班级号 %s    总容量 %s 已选 %s\n",
			list.TmpList[i].Kcmc,
			list.TmpList[i].Kch,
			list.TmpList[i].Jxbmc,
			(*list.TmpList[i].DetailList).Jxbrl,
			list.TmpList[i].Yxzrs)

		chose := false
		for j := 0; j < len(*chosenCourse); j++ {
			if (*chosenCourse)[j].JxbId == list.TmpList[i].Jxbmc {
				chose = true
				break
			}
		}
		if chose {
			continue
		}
		if (*list.TmpList[i].DetailList).Jxbrl > list.TmpList[i].Yxzrs {
			time.Sleep(3 * time.Second)

			prvChooseReq := MakeChooseClassPrvReq(fields)
			prvChooseReq.JxbIds = (*list.TmpList[i].DetailList).DoJxbId
			prvChooseReq.KchId = list.TmpList[i].Kch
			prvChooseReq.Cxbj = list.TmpList[i].Cxbj
			prvChooseReq.Xxkbj = list.TmpList[i].Xxkbj
			prvChooseReq.Kcmc = fmt.Sprintf("(%s)%s - %s 学分", list.TmpList[i].Kch, list.TmpList[i].Kcmc, list.TmpList[i].Xf)
			err := prvChooseCourse(JWBASE+"/xsxk/zzxkyzbjk_xkBcZyZzxkYzb.html?"+baseQuery.Encode(), prvChooseReq)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(list.TmpList[i].Kch, " 选课成功")
			chosenCourse = getChosenCourse(baseQuery.Encode(), MakeGetChosenClassReq(fields))
		}
	}

}
