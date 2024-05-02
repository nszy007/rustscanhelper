package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputPath := flag.String("i", "", "解析rustscan -g输出结果的文件路径")
	outputPath := flag.String("o", "", "输出结果文件路径")

	flag.Parse()

	if *inputPath == "" {
		fmt.Println("输入文件路径不能为空！")
		return
	}
	content, err := os.ReadFile(*inputPath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	rawData := string(content)
	data := make(map[string][]string)
	lines := strings.Split(rawData, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			parts := strings.Split(line, " -> ")
			ip := parts[0]
			var ports []string
			portsStr := parts[1][1:]
			portsStr = portsStr[:len(portsStr)-1]
			ports = strings.Split(portsStr, ",")
			data[ip] = ports
		}
	}

	var ipPortList []string
	for ip, ports := range data {
		for _, port := range ports {
			ipPortList = append(ipPortList, ip+":"+port)
		}
	}

	result := strings.Join(ipPortList, "\n")

	// 如果指定了输出文件，则写入文件，否则打印到控制台
	if *outputPath != "" {
		outputFile, err := os.Create(*outputPath)
		if err != nil {
			fmt.Println("创建输出文件失败:", err)
			return
		}
		defer outputFile.Close()

		outputFile.WriteString(result)
	} else {
		fmt.Println(result)
	}
}
