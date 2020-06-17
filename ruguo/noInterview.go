package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"personal/go_test/ruguo/utils"
	"strings"
	"text/template"
)

func main() {
	var (
		namesInput, tem string
		numInput        int
	)
	nameTemplate := "name1,name2,name3"
	fmt.Printf(`请输入未面试者的名字名字,格式参考：%s,中间用英文","区分`+"\n", nameTemplate)
	fmt.Scanln(&namesInput) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
	namesList := strings.Split(namesInput, ",")
	fmt.Println("请输入活动总人数:")
	fmt.Scanln(&numInput)
	// 模板定义
	tepl := `
------------{{.name}}------------>

哈喽，亲爱的{{.name}}。非常非常感谢你本次的报名，但是很遗憾这一次的活动要与你错过了。

放弃你，这个决定很艰难。我们本次的活动只挑选「性格阅历」各不相同的 {{.num}} 位参与者，这一次，我们综合「报名表填写内容」「报名时间」「报名次数」以及其他情况综合考虑决策。

但是请不要因为本次我们略显残忍的决定而放弃36h生活实验，或是放弃我们如果青年。我们还有例如「回声运动」「造作周末」「氧气戏剧社」等很多其他的优秀活动，都希望你能参加，完成本次没能相遇的遗憾。下一次的36h生活实验，欢迎你再来报名啊，如果你有空，一定要来啊～

请继续关注如果青年，并且欢迎随时勾搭小可爱


<------------{{.name}}------------ 
`
	// 解析模板
	tmpl, _ := template.New("test").Parse(tepl)
	buf := new(bytes.Buffer)
	for _, name := range namesList {
		tmpl.Execute(buf, map[string]interface{}{
			"name": name,
			"num":  numInput,
		})
	}
	writeString := buf.String()
	d1 := []byte(writeString)
	filename := "./未参加面试.text"
	err2 := ioutil.WriteFile(filename, d1, 0666)
	utils.Check(err2)

	// 提示用户信息
	outputPath, _ := os.Getwd()
	fmt.Printf("\n \n 运行结束啦(￣３￣)a，\n 1. 请到此位置 %s 未参加面试.text文件 \n 2. 请按下Enter键关闭该页面", outputPath)
	fmt.Scanln(&tem)
}
