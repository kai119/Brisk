package object

import "testing"

func TestStringDictionaryKey(t *testing.T) {
	hello1 := &String{Value: "hello world"}
	hello2 := &String{Value: "hello world"}
	diff1 := &String{Value: "this is a test"}
	diff2 := &String{Value: "this is a test"}

	if hello1.DictionaryKey() != hello2.DictionaryKey() {
		t.Errorf("strings with same content have different dictionary keys")
	}

	if diff1.DictionaryKey() != diff2.DictionaryKey() {
		t.Errorf("strings with same content have different dictionary keys")
	}

	if hello1.DictionaryKey() == diff1.DictionaryKey() {
		t.Errorf("strings with different content have the same dictionary keys")
	}
}
