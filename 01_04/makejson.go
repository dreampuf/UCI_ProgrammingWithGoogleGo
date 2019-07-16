package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	m := map[string]string{}
	var name, address, line string
	var err error

	r := bufio.NewReader(os.Stdin)
	if line, err = r.ReadString('\n'); err != nil {
		return
	}
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}
	if spacePos := strings.LastIndex(line, " "); spacePos != -1 {
		name = line[:spacePos]
		address = line[spacePos+1:]
		m["name"] = name
		m["address"] = address
	}
	data, _ := json.Marshal(m)
	fmt.Printf("%s\n", string(data))
}
