package maps_01

type DictionaryErr string

const (
	ErrAlreadyExist   = DictionaryErr("Key already added")
	ErrNotFound       = DictionaryErr("Key not found in dictionary")
	ErrNotFoundUpdate = DictionaryErr("No key found to update")
	ErrNotFoundDelet  = DictionaryErr("The key to be deleted was not found")
)

func (e DictionaryErr) Error() string {
	return string(e)
}
