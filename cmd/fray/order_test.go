package main

import (
	"math/big"
	"testing"

	"github.com/fraybet/cli/core"
)

func TestPriceToMicro(t *testing.T) {
	cases := []struct {
		in   string
		want int64
		bad  bool
	}{
		{"0.5", 500_000, false},
		{"0.55", 550_000, false},
		{"1", 1_000_000, false},
		{"0.000001", 1, false},
		{"0.1234567", 123_456, false}, // truncates beyond 6dp
		{"0", 0, true},                // not allowed (<= 0)
		{"1.5", 0, true},              // > 1
		{"-0.5", 0, true},
		{"abc", 0, true},
	}
	for _, c := range cases {
		got, err := priceToMicro(c.in)
		if c.bad {
			if err == nil {
				t.Errorf("priceToMicro(%q): want error", c.in)
			}
			continue
		}
		if err != nil {
			t.Errorf("priceToMicro(%q): %v", c.in, err)
			continue
		}
		if got.Int64() != c.want {
			t.Errorf("priceToMicro(%q) = %d, want %d", c.in, got.Int64(), c.want)
		}
	}
}

func TestOrderAmounts(t *testing.T) {
	size := big.NewInt(100_000_000) // 100 tokens (6dp)
	price := big.NewInt(550_000)    // 0.55

	mk, tk := orderAmounts(core.OrderBuy, size, price)
	if mk.Int64() != 55_000_000 || tk.Int64() != 100_000_000 {
		t.Errorf("buy: maker/taker = %s/%s, want 55000000/100000000", mk, tk)
	}
	mk, tk = orderAmounts(core.OrderSell, size, price)
	if mk.Int64() != 100_000_000 || tk.Int64() != 55_000_000 {
		t.Errorf("sell: maker/taker = %s/%s, want 100000000/55000000", mk, tk)
	}
}

func TestPrice1e18ToDecimal(t *testing.T) {
	cases := map[string]string{
		"500000000000000000":  "0.5",
		"1000000000000000000": "1",
		"550000000000000000":  "0.55",
	}
	for in, want := range cases {
		if got := price1e18ToDecimal(in); got != want {
			t.Errorf("price1e18ToDecimal(%s) = %s, want %s", in, got, want)
		}
	}
}
