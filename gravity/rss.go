package gravity

import (
	"errors"
	"log"
)

func getRss(msg chan int) {
	// 获取RSS列表文章
	sign := 0
	if 1 == 1 {
		log.Printf("获取Rss")
		sign = 1
	} else {
		err := errors.New("获取失败")
		log.Println(err)
		sign = 0
	}
	msg <- sign

}

func getOther() error {
	// 获取RSS列表文章
	if 1 == 1 {
		log.Printf("获取Rss")
		return nil
	} else {
		err := errors.New("获取失败")
		// log.Fatal(err)
		// log.Panic(err) // 连带系统退出Emmmmm...不好
		log.Println(err)
		return err
	}
}
