package main

import "github.com/wanmei002/docker-repository-cli/cmd"

func main() {
	cmd.Init()
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
