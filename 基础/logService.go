package main

import (
	"errors"
	"fmt"
	"os"
)

//日志对外接口
type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	writeList []LogWriter
}

func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writeList = append(l.writeList, writer)
}

func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writeList {
		writer.Write(data)
	}
}

func NewLogger() *Logger {
	return &Logger{}
}

//文件写入器
type fileWriter struct {
	file *os.File
}

func (f *fileWriter) SetFile(filename string) (err error) {
	if f.file != nil {
		f.file.Close()
	}

	f.file, err = os.Create(filename)

	return err
}

func (f *fileWriter) Write(data interface{}) error {
	if f.file == nil {
		return errors.New("file not created")
	}
	str := fmt.Sprintf("%v\n", data)

	_, err := f.file.Write([]byte(str))
	return err
}

func newFileWriter() *fileWriter {
	return &fileWriter{}
}

//命令行写入器
type consoleWriter struct{}

func (f *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)

	_, err := os.Stdout.Write([]byte(str))
	return err
}

func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}

//使用日志
func createLogger() *Logger {
	l := NewLogger()
	cw := newConsoleWriter()
	l.RegisterWriter(cw)
	fw := newFileWriter()
	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}
	l.RegisterWriter(fw)
	return l
}
func main() {
	l := createLogger()
	l.Log("hello1")
}
