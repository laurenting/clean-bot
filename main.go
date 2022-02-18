package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	fileType := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".mp4":  true,
	}
	u, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		_, _ = fmt.Scanln()
		return
	}
	rootDir := []string{
		fmt.Sprintf("%s\\Documents\\JDM", u.HomeDir),
		fmt.Sprintf("%s\\Documents\\WeChat Files", u.HomeDir),
	}
	count := 0
	for _, dir := range rootDir {
		fmt.Println(dir)
		root, err := os.Open(dir)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("目录不存在")
			_, _ = fmt.Scanln()
			return
		}
		err = filepath.Walk(root.Name(), func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				fmt.Printf("当前目录: %s \n", info.Name())
			}
			if fileType[strings.ToLower(filepath.Ext(info.Name()))] {
				count++
				fmt.Printf("媒体文件: %s \n",info.Name())
				// _ := os.Remove(path)
				return nil
			}
			return nil
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Printf("找到媒体文件共 %d 个", count)
	_, _ = fmt.Scanln()
}
