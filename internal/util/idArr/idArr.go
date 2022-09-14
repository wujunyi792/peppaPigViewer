package idArr

import (
	"fmt"
	"regexp"
)

func FindIDArr(body string, classNumber int) map[string][]string { //可以返回body中由正则匹配到的 onclick="queryCourse(this,'10','E0BDC4C7604BD44BE0536264A8C0B7EC','2020','0523')" 字符串
	returnTemp := make(map[string][]string)
	returnTemp["firstKklxdmArr"] = make([]string, classNumber)
	returnTemp["firstXkkzIdArr"] = make([]string, classNumber)
	returnTemp["firstNjdmIdArr"] = make([]string, classNumber)
	returnTemp["firstZyhIdArr"] = make([]string, classNumber)

	reg1 := regexp.MustCompile(`queryCourse\(this,(?s:(.*?))\)"`)
	reg2 := regexp.MustCompile(`'(?s:(.*?))'`)
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(body, -1)
	if len(result1) != 4 || len(result1[0]) != 2 {
		fmt.Println("re错误!") //TODO:错误处理
		return nil
	}
	for i := 0; i < classNumber; i++ {
		result2 := reg2.FindAllStringSubmatch(result1[i][1], -1)
		//fmt.Println(result1[i])
		if len(result2) != 4 || len(result2[0]) != 2 {
			fmt.Println("re错误!")
			return nil
		}
		returnTemp["firstKklxdmArr"][i] = result2[0][1]
		returnTemp["firstXkkzIdArr"][i] = result2[1][1]
		returnTemp["firstNjdmIdArr"][i] = result2[2][1]
		returnTemp["firstZyhIdArr"][i] = result2[3][1]
	}
	return returnTemp
}
