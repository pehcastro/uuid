package uuid

import "testing"

func TestNullUUIDZeroValueIsInvalid(t *testing.T) {
	var nu NullUUID
	if nu.Valid || nu.UUID != Nil {
		t.Errorf("zero NullUUID = %+v, want Valid false and Nil", nu)
	}
}
