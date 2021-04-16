package gravity

import (
	"errors"
	"log"
)

func Engine() error {
	msg_rss := make(chan int)
	// err := getRss()
	go getRss(msg_rss)
	sign := <-msg_rss

	if sign == 1 {
		// amazing _nop_ pass
		log.Printf("获取RSS成功")
	}
	if 1 == 1 {
		return nil
	} else {
		err := errors.New("重力无法启动, 秩序无法建立.")
		// log.Fatal(err)
		log.Panic(err)
		return err
	}
}
