package pag_test

import (
	"slices"
	"testing"

	pag "github.com/ndx-technologies/prometheus-aggregation-gateway"
)

func TestParseUserAgent(t *testing.T) {
	tests := []struct {
		s        string
		products []string
	}{
		{s: "", products: nil},
		{s: " , ,    ,", products: nil},
		{
			s:        "PriceTracker/202505.10.2217 (something blabla/1.2.3) CFNetwork/3826.500.111.2.2 Darwin/24.4.0",
			products: []string{"PriceTracker/202505.10.2217", "CFNetwork/3826.500.111.2.2", "Darwin/24.4.0"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			products := pag.ParseUserAgent(tc.s)
			if !slices.Equal(products, tc.products) {
				t.Error(tc.products, products)
			}
		})
	}
}
