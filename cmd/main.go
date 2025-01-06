package main

import (
	"fmt"
	"gotp/pkg/envparser"
	"gotp/pkg/tcpchat/tcpserver"
	"os"
)

func main() {
	envs := getEnvs()
	protocol := envs["PROTOCOL"]
	host := envs["HOST"]
	port := envs["PORT"]
	chat.StartTCPChat(protocol, host, port)
}

func getEnvs () map[string]string {
	wd, _ := os.Getwd()
	pathToEnv := fmt.Sprintf("%s/.env", wd)
	env := envparser.Env{Path: pathToEnv, CommentChar: "#"}
	return env.GetEnvMembers()
}
