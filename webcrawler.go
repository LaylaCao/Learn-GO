package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

/*
构造get请求函数
*/
func httpGet(url string) string{

	resp, err := http.Get(url)

	if err != nil {

		fmt.Println(err)

		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		fmt.Println(err)

		return ""

	}

	return string(body)
}

/*
构造正则匹配函数
*/

func httpRegex(str string) []string {

	//正则表达式，有点菜，只会(.*?)
	regex := "<span class="+"\""+"img-hash"+"\""+">(.*?)</span>"

	reg := regexp.MustCompile(regex)

	dataS := reg.FindAllSubmatch([]byte(str), -1)

	results := make([]string,0)

	for _,v := range dataS {

		results = append(results,string(v[1]))
	}

	return results

}
/*
构造保存图片函数
本函数是在互联网上搜到的
*/

func getImg(url string) (n int64, err error) {

	path := strings.Split(url, "/")

	var name string

	if len(path) > 1 {

		name = path[len(path)-1]
	}

	fmt.Println(name)

	out, err := os.Create(name)   //创建文件

	defer out.Close()

	pix := httpGet(url) //获取图片

	pic := []byte(pix)

	n, err = io.Copy(out, bytes.NewReader(pic))

	return

}
/*
主方法
*/

func main() {

	err := os.MkdirAll("image", os.ModePerm)   //创建image目录

	if err != nil {

		fmt.Println(err)

	}

	os.Chdir("./image")   //修改工作目录

	str := httpGet("http://jiandan.net/ooxx")

	reg := httpRegex(str)

	results := make([]string,0)

	for _, v := range reg{

		decodeBytes, err := base64.StdEncoding.DecodeString(v)  //base64解码

		if err != nil {

			fmt.Println(err)

		}

		results = append(results,"http:" + string(decodeBytes))

	}

	//遍历url

	for _, url := range results {

		getImg(url)

		fmt.Println("-------------------------")

	}



}

