package expect

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, wanted any, got any) {
	if !reflect.DeepEqual(wanted, got) {
		t.Fatalf("\n\n\twanted: %v\n\tgot: %v\n\n", wanted, got)
	}
}
