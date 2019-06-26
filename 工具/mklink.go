package main

import (
	"fmt"
	"os"
)

var oldproject = [...]string{
	"proto",
	"core",
}

func main() {
	//随手写的一个建立软连接的
	project := [...]string{
		"grpc-services-antitrash",
		"grpc-services-circle",
		"grpc-services-data",
		"grpc-services-diaries",
		"grpc-services-friends",
		"grpc-services-gravidity",
		"grpc-services-mp",
		"grpc-services-user",
		"http-services-user",
	}

	var a int
	oldPathmap := map[string]os.FileInfo{}
	for _, v := range oldproject {
		oldpath := "d:\\goVm\\src\\grpc-services-" + v
		oldPathmap[v], _ = os.Stat(oldpath)
	}

	for _, value := range project {
		for _, v := range oldproject {
			path := "d:\\goVm\\src\\" + value + "\\vendor\\gitlab.meiyou.com\\meiyou-services\\grpc-services-" + v + ".git"
			oldpath := "d:\\goVm\\src\\grpc-services-" + v
			//文件是否存在
			pathFile, err := os.Stat(path)
			//不存在
			if err != nil {
				err = os.Symlink(oldpath, path)
				if err != nil {
					for {
						fmt.Println(err)
						_, _ = fmt.Scanln(&a)
						if a == 1 {
							return
						}
					}
				}
				continue
			}
			//存在
			//判断是否建立了软链接
			//建立了
			if os.SameFile(oldPathmap[v], pathFile) {
				continue
			}

			_ = os.RemoveAll(path)
			err = os.Symlink(oldpath, path)
			if err != nil {
				for {
					fmt.Println(err)
					_, _ = fmt.Scanln(&a)
					if a == 1 {
						return
					}
				}
			}
		}
	}
}
