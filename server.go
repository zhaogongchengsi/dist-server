package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	var staticDir string
	var port int
	flag.StringVar(&staticDir, "dir", "./dist", "需要代理的文件夹")
	flag.IntVar(&port, "port", 8080, "端口")
	flag.Parse()

	fileServer := http.FileServer(http.Dir(staticDir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		urlStr := path.Join(staticDir, r.URL.Path)

		if p := filepath.Ext(r.URL.Path); p == "" {
			if r.URL.Path == "/" {
				urlStr = path.Join(staticDir, "index")
			}

			urlStr = urlStr + ".html"

			info, err := os.Stat(urlStr)

			if err != nil || !info.Mode().IsRegular() {
				http.Error(w, fmt.Sprintf("文件无后缀且不是html文件 %s:,%v", r.URL.Path, err), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "text/html")

			data, err := os.ReadFile(urlStr)

			if err != nil {
				http.Error(w, fmt.Sprintf("Failed to read file: %v", err), http.StatusInternalServerError)
				return
			}

			_, err = w.Write(data)

			if err != nil {
				http.Error(w, fmt.Sprintf("写入错误错误, %v", err), http.StatusInternalServerError)
				return
			}

			return
		}

		if _, err := os.Stat(urlStr); err == nil {
			fileServer.ServeHTTP(w, r)
			return
		}

	})

	portStr := fmt.Sprintf(":%v", port)
	// 启动 HTTP 服务器

	log.Printf("Static file proxy server running at http://localhost%s", portStr)

	if err := http.ListenAndServe(portStr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
