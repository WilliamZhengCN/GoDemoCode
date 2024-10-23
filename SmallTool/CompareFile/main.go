package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

var (
	fileMaps sync.Map
	wg       sync.WaitGroup
)

//a
func main() {
	var dir string
	flag.StringVar(&dir, "path", "F:\\jisuxt", "scanFile")
	flag.Parse()
	if len(dir) == 0 {
		return
	}

	folders := LoadFolder(dir)
	wg.Add(len(folders))
	for i := 0; i < len(folders); i++ {
		go LoadFile(folders[i])
	}

	wg.Wait()
	fileMaps.Range(GetDuplicateFile)
}

func LoadFolder(path string) []string {
	folders := make([]string, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Print("read file error. err:" + err.Error())
	}
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, path+"\\"+file.Name())
		}
	}
	return folders
}

func LoadFile(path string) {
	defer wg.Done()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Print("read file error err:" + err.Error())
	}
	for _, file := range files {
		if paths, ok := fileMaps.Load(file.Name()); ok {
			p := paths.(string) + ";" + path + "\\" + file.Name()
			fileMaps.Store(file.Name(), p)
		} else {
			fileMaps.Store(file.Name(), path+"\\"+file.Name())
		}
	}
}

func GetDuplicateFile(key, value interface{}) bool {
	if strings.Contains(value.(string), ";") {
		fmt.Printf("key: %s, value: %s", key, value)
	}
	return true
}
