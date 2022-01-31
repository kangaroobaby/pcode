package runtime

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"main/core"
	"strconv"
	"strings"
)

type Func struct {
	content     string
	sourceCodes []string
	eip         int
}

var (
	EOF = errors.New("EOF")
)

func NewFunc(content string) *Func {
	return &Func{
		content:     content,
		sourceCodes: make([]string, 0, 100),
	}
}

func (f *Func) Initialize() error {
	r := bufio.NewReader(strings.NewReader(f.content))
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// 去掉前面空格、TAB键
		for {
			if line[0] != ' ' && line[0] != '\t' {
				break
			}
			line = line[1:]
		}

		// 忽略注释、空行
		if line[0] == ';' || line[0] == '\n' {
			continue
		}

		f.sourceCodes = append(f.sourceCodes, line)
	}
	return nil
}

func (f *Func) Next() (core.Command, error) {
	var sourceCode string
	for {
		if f.eip >= len(f.sourceCodes) {
			return core.Command{}, EOF
		}

		sourceCode = f.sourceCodes[f.eip]
		f.eip++

		// 跳过标签
		line := sourceCode
		for {
			if line[0] == ' ' || line[0] == ':' || line[0] == '\n' {
				break
			}
			line = line[1:]
		}
		if line[0] == ' ' || line[0] == '\n' {
			break
		}
	}

	return f.parseCode(sourceCode)
}

func (f *Func) Seek(eip int) error {
	f.eip = eip
	return nil
}

func (f *Func) parseCode(code string) (core.Command, error) {
	var operate string
	var values = make([]interface{}, 0, 5)

	line := code
	for {
		if line[0] == ' ' || line[0] == '\n' {
			break
		}
		operate += line[:1]
		line = line[1:]
	}

	for {
		var value string

		// 去掉前面空格
		for {
			if line[0] != ' ' {
				break
			}
			line = line[1:]
		}

		if line[0] == ';' || line[0] == '\n' {
			break
		}

		// 引号情况
		if line[0] == '"' {
			line = line[1:]

			for {
				if line[0] == '"' || line[0] == ';' || line[0] == '\n' {
					break
				}
				value += line[:1]
				line = line[1:]
			}

			if line[0] == '"' {
				line = line[1:]
			}

		} else {
			for {
				if line[0] == ' ' || line[0] == ',' || line[0] == ';' || line[0] == '\n' {
					break
				}
				value += line[:1]
				line = line[1:]
			}
			if line[0] == ',' {
				line = line[1:]
			}
		}

		i, err := strconv.ParseInt(value, 10, 32)
		if err == nil {
			values = append(values, int(i))
		} else {
			values = append(values, value)
		}

	}

	return core.Command{f.eip, operate, values}, nil
}

func (f *Func) Goto(labelName string) error {
	for i, sourceCode := range f.sourceCodes {
		if strings.HasPrefix(sourceCode, labelName+":") {
			f.eip = i
			return nil
		}
	}
	return fmt.Errorf("can't find label: %s", labelName)
}
