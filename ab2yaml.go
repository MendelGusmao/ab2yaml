package ab2yaml

import (
	"bufio"
	"bytes"
	"strings"

	"gopkg.in/yaml.v2"
)

func Convert(in []byte) ([]byte, error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(in))
	records := make([]map[string]string, 0)
	index := -1

	for scanner.Scan() {
		line := scanner.Text()
		size := len(line)

		if size > 0 && line[0] == '#' {
			continue
		}

		if size >= 3 && line[0] == '[' && line[size-1] == ']' {
			index++
			records = append(records, map[string]string{"_key": line[1 : size-1]})

			continue
		}

		if equal := strings.Index(line, "="); equal > -1 {
			records[index][line[:equal]] = line[equal+1:]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return yaml.Marshal(records)
}
