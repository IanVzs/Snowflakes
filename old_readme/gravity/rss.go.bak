package gravity

import (
	"errors"
	"log"

	"encoding/json"

	"github.com/IanVzs/Snowflakes/helpers"
)

type RssListItem struct {
	Url         string   `json:"url"`
	Description string   `json:"description"`
	Srtdesc     string   `json:"srtdesc"`
	Articles    []string `json:"articles"`
}

func getFromUrl() error {
	// 获取RSS列表文章
	if 1 == 1 {
		log.Printf("获取Rss")
		return nil
	} else {
		err := errors.New("获取失败")
		// log.Fatal(err)
		// log.Panic(err) // 连带系统退出Emmmmm...不好
		// 之后用`zap`吧, demo中测试感觉还行.
		log.Println(err)
		return err
	}
}

func getRssSum() int {
	// 获取Rss总数, 获取失败时返回0
	// TODO
	return 100
}

func save2DB(listData []RssListItem) error {
	// 循环列表,从中获取url,拉取数据,保存入库
	for i := 0; i < len(listData); i++ {
		data := listData[i]
		str_data, _ := json.Marshal(data)
		log.Printf("正在请求写入: %s", string(str_data))
		_, err := helpers.Post("http://127.0.0.1:5481/rss/save_details", "json", str_data)
		if err != nil {
			log.Println("save2DB|faild: ", data.Description)
			return err
		} else {
			log.Println("save2DB|sucess: ", data.Description)
		}
	}
	return nil
}

func getRss(msg chan int) {
	// 获取RSS列表文章
	sign := 0
	log.Printf("获取Rss")
	num_all := getRssSum()
	if num_all > 100 {
		loop_num := num_all / 100
		for i := 0; i < loop_num; i++ {
			listData, err := getRssList(i*100, i*100+100)
			if err == nil {
				err = save2DB(listData)
				if err == nil {
					sign = 1
				}
			}
		}
	} else {
		listData, err := getRssList(0, 100)
		if err == nil {
			err = save2DB(listData)
			if err == nil {
				sign = 1
			}
		}
	}
	msg <- sign
}

func getRssList(skip int, limit int) ([]RssListItem, error) {
	// 获取RSS列表
	var data []RssListItem
	log.Printf("获取Rss列表")
	body, err := helpers.Get("http://127.0.0.1:5481/list_title?skip=0&limit=100")
	log.Println("请问有运行到这里吗？")
	if err == nil {
		err = json.Unmarshal(body, &data)
		if err != nil {
			err := errors.New("解析请求数据失败")
			log.Println(err)
			return data, err
		}
		return data, nil
	} else {
		err := errors.New("获取失败")
		// log.Fatal(err)
		// log.Panic(err) // 连带系统退出Emmmmm...不好
		// 之后用`zap`吧, demo中测试感觉还行.
		log.Println(err)
		return data, err
	}
}
