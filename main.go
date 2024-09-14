package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 指定虚拟环境中的 Python 解释器路径
	pythonPath := "./venv/bin/python"
	scriptPath := "./run.py"

	// 创建命令
	cmd := exec.Command(pythonPath, scriptPath, "-h")

	// 运行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// 打印输出
	fmt.Printf("Output: %s\n", output)
}
