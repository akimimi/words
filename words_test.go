package words

import "testing"

func TestCamelize(t *testing.T) {
	words := [][]string{
		{"shop", "Shop"},
		{"shop_account", "ShopAccount"},
		{"basic type", "Basic Type"},
		{"NothingHappens", "NothingHappens"},
	}

	for i := 0; i < len(words); i++ {
		actual := Camelize(words[i][0])
		if actual != words[i][1] {
			t.Error("Actual", actual, "Expected", words[i][1])
		}
	}
}

func TestCapitalize(t *testing.T) {
	words := [][]string{
		{"shop", "Shop"},
		{"shop_account", "Shop_account"},
		{"basic type", "Basic type"},
		{"NothingHappens", "NothingHappens"},
	}

	for i := 0; i < len(words); i++ {
		actual := Capitalize(words[i][0])
		if actual != words[i][1] {
			t.Error("Actual", actual, "Expected", words[i][1])
		}
	}
}

func TestLowercase(t *testing.T) {
	words := [][]string{
		{"Shop", "shop"},
		{"Shop_account", "shop_account"},
		{"Basic type", "basic type"},
	}

	for i := 0; i < len(words); i++ {
		actual := Lowercase(words[i][0])
		if actual != words[i][1] {
			t.Error("Actual", actual, "Expected", words[i][1])
		}
	}
}

func TestModelize(t *testing.T) {
	words := [][]string{
		{"Shop", "Shop"},
		{"shop", "Shop"},
		{"shops", "Shop"},
		{"shop_accounts", "ShopAccount"},
		{"Basic types", "BasicType"},
	}

	for i := 0; i < len(words); i++ {
		actual := Modelize(words[i][0])
		if actual != words[i][1] {
			t.Error("Actual", actual, "Expected", words[i][1])
		}
	}
}

func TestTableize(t *testing.T) {
	words := [][]string{
		{"Shop", "shops"},
		{"shop", "shops"},
		{"shops", "shops"},
		{"ShopAccount", "shop_accounts"},
		{"Basic types", "basic_types"},
	}

	for i := 0; i < len(words); i++ {
		actual := Tableize(words[i][0])
		if actual != words[i][1] {
			t.Error("Actual", actual, "Expected", words[i][1])
		}
	}
}

func TestWords(t *testing.T) {
	s := "GoodWork"
	w := words(s)
	if len(w) != 2 {
		t.Error()
	}
	if w[0] != "Good" || w[1] != "Work" {
		t.Error()
	}
	s = "  Good_work _ok_  "
	w = words(s)
	if len(w) != 3 {
		t.Error()
	}
	if w[0] != "Good" || w[1] != "work" || w[2] != "ok" {
		t.Error()
	}
	w = Words(s)
	if len(w) != 3 {
		t.Error()
	}
	if w[0] != "Good" || w[1] != "work" || w[2] != "ok" {
		t.Error()
	}
}
