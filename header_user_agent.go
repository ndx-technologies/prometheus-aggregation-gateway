package pag

import "strings"

func ParseUserAgent(s string) (products []string) {
	if s == "" {
		return nil
	}

	for part := range strings.SplitSeq(s, " ") {
		part = strings.TrimSpace(part)
		if len(part) < 2 {
			continue
		}

		// skip details
		if part[0] == '(' || part[len(part)-1] == ')' {
			continue
		}

		products = append(products, part)
	}

	return products
}
