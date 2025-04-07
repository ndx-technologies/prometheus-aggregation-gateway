package pag

import (
	"bufio"
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

	for kv := range strings.SplitSeq(s[i+1:j], ",") {
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

func PrintMetric(w io.Writer, prefix string, name string, config MetricConfig, values map[string]map[string]float64) {
	if config.Type == Histogram {
		if len(values[name+"_bucket"]) == 0 {
			return
		}
	} else {
		if len(values[name]) == 0 {
			return
		}
	}

	b := bufio.NewWriterSize(w, (len(prefix)+len(name)+10)*3+len(values[name])*8)

	b.WriteString("# HELP " + prefix + name + " " + config.Help + "\n")
	b.WriteString("# TYPE " + prefix + name + " " + config.Type.String() + "\n")

	if config.Type == Histogram {
		printMetric(b, prefix, name, "_bucket", values)
		printMetric(b, prefix, name, "_sum", values)
		printMetric(b, prefix, name, "_count", values)
	} else {
		printMetric(b, prefix, name, "", values)
	}

	b.WriteRune('\n')
	b.WriteRune('\n')

	b.Flush()
}

func printMetric(b *bufio.Writer, prefix, name, suffix string, values map[string]map[string]float64) {
	for l, v := range values[name+suffix] {
		b.WriteString(prefix)
		b.WriteString(name)
		b.WriteString(suffix)
		b.WriteString(l)
		b.WriteRune(' ')
		b.Write(strconv.AppendFloat(b.AvailableBuffer(), v, 'f', -1, 64))
		b.WriteRune('\n')
	}
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
