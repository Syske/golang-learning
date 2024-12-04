package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type NmData struct {
	Uuid              string  `json:"uuid"`
	ActivityId        string  `josn:"activityId"`
	TotalScore        float32 `json:"totalScore"`
	IdentityCard      string  `json:"identityCard"`
	ReportUrl         string  `json:"reportUrl"`
	ReportUrlJson     string  `json:"reportUrlJson"`
	DimScore          string  `json:"dimScore"`
	StartLevel        int     `json:"startLevel"`
	LevelExplain      string  `json:"levelExplain"`
	TestSite          string  `json:"testSite"`
	ReportUrlShowJson string  `json:"reportUrlShowJson"`
}

func ParseData(jsonStr string, numData *NmData) {
	// 将JSON字符串转换为struct对象
	err := json.Unmarshal([]byte(jsonStr), numData)
	if err != nil {
		panic(err)
	}
}

func main() {
	// JSON字符串
	jsonStr := `{ "name": "Alice", "age": 25 }`

	var person Person // 定义Person结构体变量

	err := json.Unmarshal([]byte(jsonStr), &person) // 将JSON字符串转换为struct对象
	if err != nil {
		panic(err)
	}

	fmt.Println("Name: ", person.Name)
	fmt.Println("Age: ", person.Age)

	personStr, _ := json.Marshal(person)

	fmt.Println(string(personStr))

	numStr := `{"uuid":"1904203977436827648","activityId":"n34979566936918016m4","totalScore":71.0,"identityCard":"","reportUrl":"http://test.com/ns-napm-web/joinNapm/downPdf.do?key=testKey","reportUrlJson":"{\"MBTI行为风格测验\":\"http://test.com/ns-napm-web/joinNapm/downPdf.do?key=testKey&reportTemplateId=n24647165937977344\",\"MBTI（精简版）\":\"http://test.com/ns-napm-web/joinNapm/downPdf.do?key=testKey&reportTemplateId=123213L\"}","dimScore":"{\"外向总分\":86,\"感觉总分\":54,\"判断\":15,\"外向-0\":1,\"感觉-1\":0,\"外向-1\":1,\"感觉\":14,\"感觉-0\":0,\"决策方式\":71,\"感知方式\":54,\"外向\":17,\"精力指向\":86,\"判断总分\":73,\"生活方式\":73,\"思维-1\":2,\"思维-0\":2,\"判断-0\":1,\"判断-1\":1,\"思维总分\":71,\"思维\":15}","dimScoreVo":"{\"dimensionScoreList\":[{\"dimensionScoreList\":[{\"dimensionScoreList\":[],\"id\":\"n24432170935781376\",\"name\":\"判断\",\"score\":15},{\"dimensionScoreList\":[],\"id\":\"n24432170935781377\",\"name\":\"判断-0\",\"score\":1},{\"dimensionScoreList\":[],\"id\":\"n24432170935781378\",\"name\":\"判断-1\",\"score\":1}],\"id\":\"n24432172465522688\",\"name\":\"判断总分\",\"score\":73}],\"id\":\"n24423194099189761\",\"name\":\"生活方式\",\"score\":73},{\"dimensionScoreList\":[{\"dimensionScoreList\":[{\"dimensionScoreList\":[],\"id\":\"n24432170933684226\",\"name\":\"感觉\",\"score\":14},{\"dimensionScoreList\":[],\"id\":\"n24432170933684227\",\"name\":\"感觉-0\",\"score\":0},{\"dimensionScoreList\":[],\"id\":\"n24432170933684228\",\"name\":\"感觉-1\",\"score\":0}],\"id\":\"n24432172454905856\",\"name\":\"感觉总分\",\"score\":54}],\"id\":\"n24423192229316609\",\"name\":\"感知方式\",\"score\":54},{\"dimensionScoreList\":[{\"dimensionScoreList\":[{\"dimensionScoreList\":[],\"id\":\"n24432170935781379\",\"name\":\"思维\",\"score\":15},{\"dimensionScoreList\":[],\"id\":\"n24432170935781380\",\"name\":\"思维-0\",\"score\":2},{\"dimensionScoreList\":[],\"id\":\"n24432170937747456\",\"name\":\"思维-1\",\"score\":2}],\"id\":\"n24432172467619840\",\"name\":\"思维总分\",\"score\":71}],\"id\":\"n24423195405453313\",\"name\":\"决策方式\",\"score\":71},{\"dimensionScoreList\":[{\"dimensionScoreList\":[{\"dimensionScoreList\":[],\"id\":\"n24432170931718144\",\"name\":\"外向\",\"score\":17},{\"dimensionScoreList\":[],\"id\":\"n24432170933684224\",\"name\":\"外向-0\",\"score\":1},{\"dimensionScoreList\":[],\"id\":\"n24432170933684225\",\"name\":\"外向-1\",\"score\":1}],\"id\":\"n24432172461459456\",\"name\":\"外向总分\",\"score\":86}],\"id\":\"n24423190718580737\",\"name\":\"精力指向\",\"score\":86}","startLevel":5,"levelExplain":"合格","testSite":"1610550872771792915","reportUrlShowJson":"{\"MBTI行为风格测验\":\"http://test.com/pdfjs/web/viewer.html?url=http://test.com/ns-napm-web/joinNapm/downPdf.do?key=testKey&reportTemplateId=n24647165937977344\",\"MBTI（精简版）\":\"http://test.com/pdfjs/web/viewer.html?url=http://test.com/ns-napm-web/joinNapm/downPdf.do?key=testKey&reportTemplateId=123213L\"}"}`
	var numData NmData
	ParseData(numStr, &numData)
	fmt.Println(numData.ReportUrlJson)
	fmt.Println(numData.ReportUrl)
	fmt.Println(numData.ReportUrlShowJson)

}
