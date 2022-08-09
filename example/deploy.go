package main

import (
	"log"

	_ "github.com/gongzhxu/packagedeploy/deploy" //手动添加部署依赖
)

func main() {
	log.Println("test deploy")
}
