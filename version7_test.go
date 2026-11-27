package uuid

import "testing"

func TestNewV7VersionAndVariant(t *testing.T) {
	u, err := NewV7()
	if err != nil {
		t.Fatalf("NewV7: %v", err)
	}
	if u.Version() != 7 || u.Variant() != RFC4122 {
		t.Errorf("NewV7() = %s: version %d, variant %v, want 7 and %v", u, u.Version(), u.Variant(), RFC4122)
	}
}
