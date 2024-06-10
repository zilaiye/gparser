package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	f, err := os.Open("./autogen/structs")
	if err != nil {
		log.Fatalln("read file err ", err)
		return
	}
	defer f.Close()
	createDir("./autogen/gen")
	reader := bufio.NewReader(f)
	lineNum := 0
	wg := sync.WaitGroup{}
	for {
		lineNum++
		line, err := reader.ReadString(byte('\n'))
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		wg.Add(1)
		go func() {
			writeFile("./autogen/gen", line)
			wg.Done()
		}()
		fmt.Printf(" line no = %d , line = %s \n", lineNum, line)
	}
	wg.Wait()
}

func createDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
		os.Chmod(path, 0777)
	}
}

func writeFile(dir, structName string) {
	s := dir + "/" + UnCapitalize(structName) + ".go"
	outputFile, err := os.OpenFile(s, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer outputFile.Close()
	if err != nil {
		log.Fatalln("打开文件失败", s)
		return
	}
	builder := strings.Builder{}
	builder.WriteString("package base\n")
	builder.WriteString("type ")
	builder.WriteString(structName)
	builder.WriteString(" struct {} \n")

	outputFile.WriteString(builder.String())
}

// unCapitalize 字符首字母小写
func UnCapitalize(str string) string {
	size := len(str)
	if size <= 0 {
		return str
	}
	vv := []rune(str)
	result := make([]rune, size)
	for i := 0; i < size; i++ {
		if i == 0 && vv[0] >= 65 && vv[0] <= 90 {
			result[i] = vv[i] + 32
		} else {
			result[i] = vv[i]
		}
	}
	return string(result)
}
