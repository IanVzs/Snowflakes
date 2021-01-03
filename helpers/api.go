package helpers

import (
    "log"
    "bytes"
    "errors"
    "net/http"
    "io/ioutil"
)
   
func Get(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        // log.Panic("api.Get Panic", err)
        // TODO log Panic 都会结束进程...emmmm那么应该怎么记录错误日志还不终端运行嘞?
        log.Println("api.Get Panic", err)
        return []byte(``), err
    }
    defer resp.Body.Close()
    
    log.Printf("Response status: %s", resp.Status)
    body, _ := ioutil.ReadAll(resp.Body)

    return body, nil
}

func Post(url string, method string, data []byte) ([]byte, error) {
    body := []byte(``)
    reader := bytes.NewReader(data)
    request, err := http.NewRequest("POST", url, reader)
    if err != nil {
        log.Println(err.Error())
        return body, err
    } else {
        if method == "json" {
            request.Header.Set("Content-Type", "application/json;charset=UTF-8")
        } else if method == "data" {
            request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
        } else {}
        client := http.Client{}
        resp, err := client.Do(request)
        if err != nil {
            log.Println(err.Error())
        } else {
            body, err = ioutil.ReadAll(resp.Body)
            log.Println("post|status", resp.Status)
            if resp.StatusCode != 200 {
                log.Println("post|status_code: ", resp.StatusCode)
                log.Println("post|err: ", body)
                err = errors.New("请求错误")
                return body, err            
            }
            
        }
    }
    return body, err
}