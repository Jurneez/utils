package file

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
create file and append write
content:the content of needing write into file
path: new file's path
newline: 加入数据后是否换行
*/
func CreateAndAppend(content interface{}, path string, newline bool) (err error) {
	b, err := json.MarshalIndent(content, "", "")
	if err != nil {
		return err
	}
	// 去掉 string(b)字符串中带有的 双引号
	b1, strconverr := strconv.Unquote(string(b))
	if strconverr != nil {
		return strconverr
	}

	if !isVaildPath(path) {
		return errors.New("path error!")
	}
	//创建一个新文件
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(string(b1))
	write.WriteString("\n")
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	return nil
}

func isVaildPath(path string) bool {
	if len(path) == 0 {
		return false
	}

	splitPaths := strings.Split(path, ".")
	if len(splitPaths) != 2 {
		return false
	}

	return true
}
