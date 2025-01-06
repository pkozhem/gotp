package envparser

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Env struct {
	Path string
	CommentChar string
	items map[string]string
}

func (e *Env) GetEnvMembers() map[string]string {
	if e.items == nil {
		e.items = make(map[string]string)
		e.parse()
	}
	return e.items
}

func (e *Env) parse() {
	envS := e.readEnv()
	lines := strings.Split(envS, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if !e.lineStartsWithCommentChar(line) {
			eqIdx := strings.Index(line, "=")
			if eqIdx < 0 {
				errMsg := fmt.Sprintf("No '=' operator in .env file. Line %v", i+1)
				panic(errMsg)
			}
			e.items[line[:eqIdx]] = line[eqIdx+1:]
		}
	}
}

func (e *Env) readEnv() string {
	envFile := e.openEnvFile()
	s := e.readEnvFileToString(envFile)
	envFile.Close()
	return strings.TrimSpace(s)
}

func (e *Env) openEnvFile() *os.File {
	file, err := os.OpenFile(e.Path, os.O_RDONLY, 0644)
	if err != nil {
        panic(err)
    }
	return file
}

func (e *Env) readEnvFileToString(envFile *os.File) string {
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, envFile)
	s := string(buf.Bytes())
	return s
}

func (e *Env) lineStartsWithCommentChar(line string) bool {
	return string(line[0]) == e.CommentChar
}
