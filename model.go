package main

import (
	"fmt"
	"sort"
)

type Object map[string]interface{}

func (o Object) SortedKeys() []string {
	columns := make([]string, len(o), len(o))
	i := 0
	for column, _ := range o {
		columns[i] = column
		i++
	}
	sort.Strings(columns) // sorting could be done differently from alphabeticaly
	return columns
}

func (o Object) String() string {
	columns := o.SortedKeys()
	values := make([]string, len(o), len(o))

	for column, value := range o {
		for k, c := range columns {
			if c == column {
				values[k] = fmt.Sprintf("%v", value) // this can cause problems
				break
			}
		}
	}
	return ToCSVLine(values)
}

func ToCSVLine(values []string) string {
	out := ""
	for k, value := range values {
		out = out + value
		if k != len(values)-1 {
			out = out + separator
		}
	}
	return out
}
