package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GetCurrentDirectory() string {
	dir, _ := os.Executable()
	exPath := filepath.Dir(dir)
	return exPath
}
func WriteFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)
		_, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
		}
	}
}
func ReadFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return []byte("")
	}
	defer f.Close()
	buf := make([]byte, 1024*1024)
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF {
		fmt.Println("err1 = ", err1)
		return []byte("")
	}
	return buf[:n]
}
func ReadFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		//遇到'\n'结束读取，但是'\n'也读取进来
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err =", err)
		}
		fmt.Println(string(buf))
	}
}

func LogPath() string {
	dir, _ := os.Getwd()
	Out := dir + "/" + ".Out.log"
	CreateFile(Out)
	return Out
}
func CreateFile(file_name string) {
	_, err := os.Create(file_name)
	if err != nil {
		fmt.Println(err)
	}
}
