package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInPathFiles(word, path string) []FindInfo {
	var resultFindInfo []FindInfo

	fileList, err := GetFileList(path)
	if err != nil {
		log.Warnln("Can not find the file. err : ", err)
		return resultFindInfo
	}

	ch := make(chan FindInfo)
	cnt := len(fileList)
	recvCnt := 0

	for _, filename := range fileList {
		log.Println(filename)
		//resultFindInfo = append(resultFindInfo, FindWordInFile(word, filename))
		go FindWordInFile(word, filename, ch)
	}

	for findInfo := range ch {
		resultFindInfo = append(resultFindInfo, findInfo)
		recvCnt++
		if cnt == recvCnt {
			break
		}
	}

	return resultFindInfo
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename: filename, lines: []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Cannot open the given file: ", file)
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
	return
}

func PrintAllFiles(word string, files []string) []FindInfo {
	var resultInfo []FindInfo

	for _, path := range files {
		resultInfo = append(resultInfo, FindWordInPathFiles(word, path)...)
	}
	return resultInfo
}

func main() {
	if len(os.Args) < 3 {
		log.Errorln("Need more than 2 arguments. i.e,) go-find word filepath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	log.Infoln("You are looking for:", word)
	result := PrintAllFiles(word, files)

	for _, info := range result {
		fmt.Println(info.filename)
		fmt.Println("------------------")
		for _, lineInfo := range info.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("------------------")
	}
}
