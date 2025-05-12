package pag_test

import (
	"maps"
	"testing"

	pag "github.com/ndx-technologies/prometheus-aggregation-gateway"
	"github.com/ndx-technologies/prometheus-aggregation-gateway/language"
)

func TestParseAcceptLanguage(t *testing.T) {
	tests := []struct {
		s string
		w map[language.Language]float64
	}{
		{s: "", w: nil},
		{s: "*", w: nil},
		{s: "en", w: map[language.Language]float64{language.English: 1}},
		{s: "ja", w: map[language.Language]float64{language.Japanese: 1}},
		{s: "en-US,en;q=0.9", w: map[language.Language]float64{language.English: 1}},
		{s: "en;q=0.5", w: map[language.Language]float64{language.English: 0.5}},
		{s: "en, fr;q=0.8, es;q=0.7", w: map[language.Language]float64{language.English: 1, language.French: 0.8, language.Spanish: 0.7}},
		{s: "en;q=invalid", w: map[language.Language]float64{language.English: 0}},
		{s: "en-US", w: map[language.Language]float64{language.English: 1}},
		{s: "en, en;q=0.5", w: map[language.Language]float64{language.English: 1}},
		{s: " en , fr ; q=0.8 ", w: map[language.Language]float64{language.English: 1, language.French: 0.8}},
		{s: "de;q=0.9, en;q=0.8, fr;q=0.7", w: map[language.Language]float64{language.German: 0.9, language.English: 0.8, language.French: 0.7}},
		{s: "en, invalid, fr;q=0.5", w: map[language.Language]float64{language.English: 1, language.French: 0.5}},
		{s: ",en,,fr;q=0.5,", w: map[language.Language]float64{language.English: 1, language.French: 0.5}},
		{s: "en;q=, fr;q=0.5", w: map[language.Language]float64{language.English: 0, language.French: 0.5}},
		{s: "invalid", w: nil},
		{s: "invalid1, invalid2", w: nil},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			w := pag.ParseAcceptLanguage(tc.s)
			if !maps.Equal(w, tc.w) {
				t.Error(tc.w, w)
			}
		})
	}
}

func BenchmarkParseAcceptLanguage(b *testing.B) {
	tests := []struct {
		s string
	}{
		{s: "en-US,en;q=0.9"},
		{s: "en"},
	}
	for _, tc := range tests {
		b.Run(tc.s, func(b *testing.B) {
			for b.Loop() {
				pag.ParseAcceptLanguage(tc.s)
			}
		})
	}
}
