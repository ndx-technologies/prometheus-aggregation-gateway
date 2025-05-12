package pag

import (
	"strconv"
	"strings"

	"github.com/ndx-technologies/prometheus-aggregation-gateway/language"
)

func ParseAcceptLanguage(s string) map[language.Language]float64 {
	if s == "" || s == "*" {
		return nil
	}

	weight := make(map[language.Language]float64)

	for e := range strings.SplitSeq(s, ",") {
		parts := strings.Split(e, ";")
		if len(parts) == 0 {
			continue
		}

		// language
		lparts := strings.Split(strings.TrimSpace(parts[0]), "-")
		if len(lparts) == 0 {
			continue
		}

		var lang language.Language
		if err := lang.UnmarshalText([]byte(lparts[0])); err != nil {
			continue
		}
		if lang == language.Unknown {
			continue
		}

		// weight
		var w float64 = 1

		if len(parts) > 1 {
			wparts := strings.Split(strings.TrimSpace(parts[1]), "=")
			if len(wparts) == 2 {
				w, _ = strconv.ParseFloat(wparts[1], 64)
			}
		}

		if wold, ok := weight[lang]; !ok || wold < w {
			weight[lang] = w
		}
	}

	return weight
}
