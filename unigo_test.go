package unigo

import (
	"testing"
)

type data struct {
	in []byte
	out string
}

var testdata = []data {
	{ []byte(`\u044d\u0442\u043e \u0442\u0435\u0441\u0442 \u0441\u043e\u043e\u0431\u0449\u0435\u043d\u0438\u0435`),
	  "это тест сообщение"},
	{ []byte(`{"key":"\u044d\u0442\u043e \u0442\u0435\u0441\u0442 \u0441\u043e\u043e\u0431\u0449\u0435\u043d\u0438\u0435"}`),
	  `{"key":"это тест сообщение"}`},
	{ []byte(`{"key":"wasn\u0027t"}`),
	  `{"key":"wasn't"}`}}

// NOTE: this implicitly tests Encode, as well.
func TestEncodeToString(t *testing.T) {
	for _, v := range testdata {
		if EncodeToString(v.in) != v.out {
			t.Errorf("didn't get: %s", v.out)
		}
	}
}
