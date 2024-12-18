package pag

import (
	"errors"
	"io"
	"sort"
	"strconv"
	"strings"
)

var (
	ErrEmptyName           = errors.New("empty name")
	ErrEmptyLabel          = errors.New("empty label")
	ErrLabelValueNotString = errors.New("label value is not string")
	ErrInvalidLabel        = errors.New("invalid label")
)

type MetricType string

func (s MetricType) String() string { return string(s) }

const (
	Counter   MetricType = "counter"
	Histogram MetricType = "histogram"
)

type MetricConfig struct {
	Help    string     `json:"help" yaml:"help"`
	Type    MetricType `json:"type" yaml:"type"`
	Buckets []string   `json:"buckets" yaml:"buckets"`
}

func (s MetricConfig) LabelValues() map[string]map[string]bool {
	if s.Type == Histogram && len(s.Buckets) > 0 {
		vs := make(map[string]bool, len(s.Buckets))
		for _, b := range s.Buckets {
			vs[b] = true
		}
		return map[string]map[string]bool{"le": vs}
	}
	return nil
}

func EncodeLabels(labels map[string]string) string {
	if len(labels) == 0 {
		return ""
	}

	vs := make([]string, 0, len(labels))

	for k, v := range labels {
		vs = append(vs, k+"="+`"`+v+`"`)
	}

	sort.Strings(vs)

	return "{" + strings.Join(vs, ",") + "}"
}

func ParseMetric(s string) (name string, labels map[string]string, err error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return "", nil, ErrEmptyName
	}

	i := strings.IndexByte(s, '{')
	j := strings.LastIndexByte(s, '}')

	if (i == -1) != (j == -1) {
		return "", nil, ErrInvalidLabel
	}

	if i == -1 && j == -1 {
		return s, nil, nil
	}

	if i == 0 {
		return "", nil, ErrInvalidLabel
	}

	if i >= j {
		return "", nil, ErrInvalidLabel
	}

	if j != len(s)-1 {
		return "", nil, ErrInvalidLabel
	}

	name = s[:i]

	for _, kv := range strings.Split(s[i+1:j], ",") {
		p := strings.SplitN(kv, "=", 2)

		if len(p) != 2 {
			return "", nil, ErrInvalidLabel
		}

		label := strings.TrimSpace(p[0])
		value := strings.TrimSpace(p[1])

		if len(label) == 0 || len(value) == 0 {
			return "", nil, ErrEmptyLabel
		}

		if len(value) <= 2 || value[0] != '"' || value[len(value)-1] != '"' {
			return "", nil, ErrLabelValueNotString
		}

		if labels == nil {
			labels = make(map[string]string)
		}
		labels[label] = value[1 : len(p[1])-1]
	}

	return name, labels, nil
}

func PrintMetric(w io.Writer, prefix string, name string, config MetricConfig, values map[string]map[string]float64) (count int) {
	names := []string{name}

	if config.Type == Histogram {
		names = []string{name + "_bucket", name + "_sum", name + "_count"}
	}

	for _, m := range names {
		for l, v := range values[m] {
			if count == 0 {
				w.Write([]byte("# HELP " + prefix + name + " " + config.Help + "\n"))
				w.Write([]byte("# TYPE " + prefix + name + " " + config.Type.String() + "\n"))
			}

			w.Write([]byte(prefix + m))
			w.Write([]byte(l))
			w.Write([]byte(" "))
			w.Write([]byte(strconv.FormatFloat(v, 'f', -1, 64)))

			w.Write([]byte("\n"))
			count++
		}
	}

	if count > 0 {
		w.Write([]byte("\n\n"))
	}

	return count
}

var histSuffixes = [...]string{"_count", "_sum", "_bucket"}

func StripHistSuffix(s string) string {
	for _, suffix := range histSuffixes {
		if strings.HasSuffix(s, suffix) {
			return s[:len(s)-len(suffix)]
		}
	}
	return s
}
