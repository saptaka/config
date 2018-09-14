package config

import (
	"os"
	"bufio"
	"log"
	"strings"

)

func SetConfiguration(config string) map[string]string {
	var splitconfig []string;
	mapConfig := make(map[string]string)
	for _, index := range readline(config) {
		if index != "" && strings.Index(index, "#")!=0 {
			strings.TrimSpace(index)
			splitconfig = strings.Split(index, "=")
			mapConfig[strings.TrimSpace(splitconfig[0])] = strings.TrimSpace(splitconfig[1])
		} else {
			continue
		}
	}
	return mapConfig
}

func readline(path string) []string {
	inFile, err := os.Open(path)
	defer inFile.Close()
	if err != nil {
		log.Println("File Not Found")
		return nil;
	} else {
		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)

		arr := make([]string, 0)
		for scanner.Scan() {
			arr = append(arr, scanner.Text())
		}
		return arr
	}

}

