package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	// 定义文件名和大小
	fileName := "data.bin"
	fileSize := 1 * 1024 * 1024 // 1MB

	// 创建一个字节数组来存储随机数据
	data := make([]byte, fileSize)

	// 使用 crypto/rand 生成随机数据
	_, err := rand.Read(data)
	if err != nil {
		fmt.Println("Error generating random data:", err)
		return
	}

	// 打开文件进行写入
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 将随机数据写入文件
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("1MB random binary file created: %s\n", fileName)
}
