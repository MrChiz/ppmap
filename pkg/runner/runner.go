package runner

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/MrChiz/ppmap/pkg/cmd"
)

// Check params
func Pmap() {
	var wg sync.WaitGroup
	if cmd.List != "" && cmd.Url == "" {
		list := readFile(cmd.List)
		wg.Add(len(list))
		for _, link := range list {
			if cmd.Output != "" {
				go queryEnum(link, "?", true, &wg)
				go queryEnum(link, "#", true, &wg)
			} else {
				log.Fatalf("[%s] Output file not import!", red("Error"))
			}
		}
	} else if cmd.Output != "" || cmd.Url != "" {
		wg.Add(2)
		go queryEnum(cmd.Url, "?", true, &wg)
		go queryEnum(cmd.Url, "#", true, &wg)
	}
	wg.Wait()
}

// Read file
func readFile(path string) []string {
	//open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("[%s] %s\n", red("Error"), err)
	}
	scanner := bufio.NewScanner(file)
	var list []string
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	file.Close()
	return list
}

// Save data
func Save(name, data string) {
	file, _ := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file.WriteString(data)
	defer file.Close()
}
