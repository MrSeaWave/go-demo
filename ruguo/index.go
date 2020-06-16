//使用一个文件单独运行的程序必须放在package main下面，否则在go run运行时会报错：”go run: cannot run non-main package“
package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	var (
		namesInput, addressInput, dateInput, tem string
	)
	nameTemplate := "name1,name2,name3"
	fmt.Printf(`请输入面试者名字,格式参考：%s,中间用英文","区分`+"\n", nameTemplate)
	fmt.Scanln(&namesInput) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
	namesList := strings.Split(namesInput, ",")
	fmt.Println("请输入活动地址:")
	fmt.Scanln(&addressInput)
	fmt.Println("请输入活动开始时间:")
	fmt.Scanln(&dateInput)
	// 模板定义
	tepl := `
亲爱的 {{.name}}：

  恭喜你成功通过面试，加入如果青年第11期36小时生活实验室。🎉🎉🎉🎉🎉🎉
 
  下面是活动的地点和时间。
  活动地点：{{.address}}
  
  活动时间：{{.time}}
 
  为了此次活动的顺利开展，我们需要再次向你确认是否能够全程参与。能够参与请回复“是”以及本人身份证号，并向以下二维码转账报名费350元。不能参与请回复“否”
 
  Ps.24小时内未回复视为自动放弃，以转账成功时间为准。
 
  感谢小伙伴对实验室的支持与热爱，期待周末同你们美妙相遇
 
可爱又迷人的如果青年
`
	// 解析模板
	tmpl, _ := template.New("test").Parse(tepl)
	for _, name := range namesList {
		fmt.Printf("----%s----> \n", name)
		// 使用data渲染模板，并将结果写入os.Stdout
		tmpl.Execute(os.Stdout, map[string]interface{}{
			"name":    name,
			"address": addressInput,
			"time":    dateInput,
		})
		fmt.Printf("\n")
		fmt.Printf("<----%s---- \n", name)
	}

	// 提示用户
	fmt.Printf("\n \n 运行结束啦(￣３￣)a，请按下Enter键关闭该页面")
	fmt.Scanln(&tem)
}
