package list

import (
	"fmt"
	"strings"
)

func ReadList(content string) (map[string]string, error) {
	idToName := make(map[string]string)

	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		idAndName := strings.Split(line, "\t")

		if len(idAndName) != 2 {
			return idToName, fmt.Errorf(
				"Kegg list should contains ID and Name.\nError in %s", line,
			)
		}

		idToName[idAndName[0]] = idAndName[1]
	}

	return idToName, nil
}
