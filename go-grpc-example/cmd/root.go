package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd表示在没有任何子命令的情况下的基本命令
// 比如直接 go run main.go
// 输出：
/*
$ go run main.go
Run the gRPC hello-world server

Usage:
  grpc [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  server      Run the gRPC hello-world server

Flags:
  -h, --help   help for grpc

Use "grpc [command] --help" for more information about a command.
*/
var rootCmd = &cobra.Command{
	// Command的用法，Use是一个行用法消息
	Use: "grpc",
	// Short是help命令输出中显示的简短描述
	Short: "Run the gRPC hello-world server",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
