package money_change

import (
	"reflect"
	"testing"
)

func TestGiveChange(t *testing.T) {
	tables := []struct {
		input        uint
		changeResult changeSet
	}{
		{input: 32, changeResult: changeSet{twenty: 1, ten: 1, two: 1}},
		{input: 0, changeResult: changeSet{}},
		{input: 331, changeResult: changeSet{fifty: 6, twenty: 1, ten: 1, one: 1}},
	}
	for _, table := range tables {
		changeResult := giveChange(table.input)
		if !reflect.DeepEqual(changeResult, table.changeResult) {
			t.Errorf("giveChange failed with input %d. Result is %#v, expected result is %#v", table.input, changeResult, table.changeResult)
		}
	}

}
