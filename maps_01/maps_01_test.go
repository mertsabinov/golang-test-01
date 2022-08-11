package maps_01

import (
	"testing"
)

// test dictionary
var dictionary = Dictionary{
	"test": "just test",
}

func chackError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err.Error())
	}
}

func chack(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got = %s want = %s", got, want)
	}
}

func TestDictionary_Search(t *testing.T) {
	got, err := dictionary.Search("test")
	chackError(t, err)
	want := "just test"
	chack(t, got, want)
}

func TestDictionary_Add(t *testing.T) {
	err := dictionary.Add("just", "just value")
	chackError(t, err)
	got, err := dictionary.Search("just")
	chackError(t, err)
	want := "just value"
	chack(t, got, want)
}

func TestDictionary_Update(t *testing.T) {
	err := dictionary.Update("test", "just")
	chackError(t, err)
	got, err := dictionary.Search("test")
	want := "just"
	chack(t, got, want)
}

func TestDictionary_Delet(t *testing.T) {
	err := dictionary.Delet("test")
	chackError(t, err)
	_, err = dictionary.Search("test")
	want := ErrNotFound
	chack(t, err.Error(), want.Error())
}
