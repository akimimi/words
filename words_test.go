package words

import "testing"

func TestCamelize(t *testing.T) {
  words := [][]string{
		[]string{"shop", "Shop"},
		[]string{"shop_account", "ShopAccount"},
		[]string{"basic type", "Basic Type"},
		[]string{"NothingHappens", "NothingHappens"},
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
		[]string{"shop", "Shop"},
		[]string{"shop_account", "Shop_account"},
		[]string{"basic type", "Basic type"},
		[]string{"NothingHappens", "NothingHappens"},
	}

	for i := 0; i < len(words); i++ {
		actual := Capitalize(words[i][0])
		if actual != words[i][1] {
			t.Error("Actual", actual, "Expected", words[i][1])
		}
	}
}
