package csvtojson

import (
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	fr, err := os.Open("examples/month.csv")
	if err != nil {
		t.Error(err)
	}
	_, err = Convert(fr)
	if err != nil {
		t.Error(err)
	}
}
