package main

import (
	"bufio"
	"os"
)

func openFile(filePath string) *os.File {
	// 打开文件，file叫文件对象/文件指针/文件句柄
	file, err := os.Open(filePath)
	// 及时关闭，否则会有内存泄漏
	defer file.Close()
	if err != nil {
		panic(err)
	}
	return file
}

/*
func readFile(file *os.File)  {
	// 方法一：带缓冲的读取文件
	// 创建一个*reader，是带缓冲的，默认缓冲区4096
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString("\n")
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	fmt.Println("read file over")

	// 方法二：一次性读取文件，适合读取小文件
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("read file data:", string(data))
}
*/

func writeFile() {
	filePath := "test.txt"
	// 打开文件：若为打开新文件，则用os.O_WRONLY|os.O_CREATE；
	// 若为打开已存在的文件，且要覆盖原有内容，则用os.O_WRONLY|os.O_TRUNC;
	// 若为追加的方式，则用os.O_WRONLY|os.O_APPEND;
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	// 及时关闭file句柄，防止内存泄漏
	defer file.Close()

	if err != nil {
		panic(err)
	}
	// 创建 *writer, 使用带缓冲的 *writer
	writer := bufio.NewWriter(file)
	//写入内容
	_, err = writer.WriteString("hello go")
	if err != nil {
		panic(err)
	}
	// 注意：因为使用带缓冲的 *writer，其实时先写入缓冲buf，因此需要flush写入文件
	_ = writer.Flush()
}

func main() {
	//file := openFile("test.txt")
	//readFile(file)

	//写文件
	writeFile()
}
