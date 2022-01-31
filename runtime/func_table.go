package runtime

import (
	"bufio"
	"fmt"
	"io"
	"main/core"
	"strings"
)

type FuncTable struct {
	content string
	funcs   map[string]*Func
}

func NewFuncTable(content string) *FuncTable {
	return &FuncTable{
		content: content,
		funcs:   make(map[string]*Func),
	}
}

func (f *FuncTable) Initialize() error {
	r := bufio.NewReader(strings.NewReader(f.content))
	return f.extractFunc(r, "")
}

func (f *FuncTable) extractFunc(r *bufio.Reader, funcName string) error {
	var content string
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// 去掉前面空格，tab键
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

		// 函数体
		if strings.HasPrefix(line, "FUNC") {
			// 找到'@'
			for {
				if line[0] == '\n' {
					return fmt.Errorf("invalid `FUNC`: %s", line)
				}
				if line[0] == '@' {
					break
				}
				line = line[1:]
			}

			line = line[1:]

			// 计算函数名
			var name string
			for {
				if line[0] == '\n' {
					return fmt.Errorf("invalid `FUNC`: %s", line)
				}
				if line[0] == ':' {
					break
				}
				name += line[:1]
				line = line[1:]
			}

			err = f.extractFunc(r, name)
			if err != nil {
				return err
			}

			continue
		}

		// 函数体
		if strings.HasPrefix(line, "ENDFUNC") {
			break
		}

		// 添加新代码行
		content += line
	}

	if content == "" {
		return nil
	}

	newFunc := NewFunc(content)
	err := newFunc.Initialize()
	f.funcs[funcName] = newFunc

	return err
}

func (f *FuncTable) SearchFunc(funcName string) (core.Func, bool) {
	_func, ok := f.funcs[funcName]
	return _func, ok
}
