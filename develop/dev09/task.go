package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

# Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func getLinks(body io.Reader) map[string]struct{} {
	links := make(map[string]struct{})

	tokenizer := html.NewTokenizer(body)
	for {
		token := tokenizer.Next()

		switch token {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links[attr.Val] = struct{}{}
					}
				}
			}

		}
	}
}

var visited map[string]struct{}
var mutex sync.Mutex

func checkURL(url string) bool {
	mutex.Lock()
	_, ok := visited[url]
	if !ok {
		visited[url] = struct{}{}
	}
	mutex.Unlock()
	return !ok
}

func dumpResponse(resp *http.Response, filename string) (bool, error) {
	dump, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	mime := http.DetectContentType(dump)
	if strings.HasPrefix(mime, "text/") && !strings.HasSuffix(filename, ".html") {
		filename = filepath.Join(filename, "/", filepath.Base(filename)+".html")
	}
	path, err := os.Getwd()
	if err != nil {
		return false, err
	}
	resppath := filepath.Join(path, filename)
	dirpath := filepath.Dir(resppath)
	err = os.MkdirAll(dirpath, 0644)
	if err != nil {
		return false, err
	}
	file, err := os.Create(resppath)
	if err != nil {
		return false, err
	}
	defer file.Close()
	return strings.HasPrefix(mime, "text/"), os.WriteFile(filename, dump, 0644)
}

func recursiveDownload(root string, dir string, curDepth int, maxDepth int) error {
	if curDepth >= maxDepth {
		return nil
	}
	pathname, err := url.JoinPath(root, dir)
	if err != nil {
		return err
	}
	resp, err := http.Get(pathname)
	if err != nil {
		log.Fatal(err)
	}
	host, err := url.Parse(pathname)
	if err != nil {
		log.Fatal(err)
	}
	html, err := dumpResponse(resp, host.Host+host.Path)

	if err != nil {
		log.Fatal(err)
	}

	resp, err = http.Get(pathname)
	if err != nil {
		log.Fatal(err)
	}

	if html {
		for next := range getLinks(resp.Body) {
			if !strings.HasPrefix(next, "http") && !strings.Contains(next, "#") && !strings.Contains(next, "?") {
				fmt.Println(next)
				if checkURL(next) {
					err = recursiveDownload(root, next, curDepth+1, maxDepth)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func main() {
	visited = make(map[string]struct{})
	err := recursiveDownload("https://go.dev/", "", 0, 5)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
