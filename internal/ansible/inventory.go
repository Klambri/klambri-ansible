package ansible

import "strings"

func createInventory(hosts []string) string {
	var builder strings.Builder

	for _, host := range hosts {
		builder.WriteString(host)
		builder.WriteByte(',')
	}

	return builder.String()
}
