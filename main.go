package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var PORT = ":8000"

type ReqParams struct {
	url  string
	host string
}

func main() {
	http.HandleFunc("/", showVersion)

	fmt.Println("Server running bind port", PORT)
	http.ListenAndServe(PORT, nil)
}

func showVersion(w http.ResponseWriter, req *http.Request) {
	// 返回hello
	fmt.Fprintf(w, "Server Version: 0.01\n\n\n\n")
	fmt.Fprintf(w, "Your Req Header is:\n\n")

	info := fmt.Sprintln("URL", req.URL, "HOST", req.Host, "Method", req.Method, "RequestURL", req.RequestURI, "RawQuery", req.URL.RawQuery)
	fmt.Fprintln(w, info)

	len_body := req.ContentLength
	body := make([]byte, len_body)
	req.Body.Read(body)
	// reqParams := ReqParams{req.URL, req.Host}
	fmt.Fprintln(w, info, string(body)) //, reqParams)
	writeLog(info, "./task_server.log")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func writeLog(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "")
	buf := []byte(content)
	fd.Write(buf)
}
