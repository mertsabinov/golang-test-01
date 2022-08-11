package maps_01

type Dictionary map[string]string

func (d Dictionary) Search(value string) (string, error) {
	definition, ok := d[value]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(value string, definition string) error {
	_, err := d.Search(value)
	switch err {
	case nil:
		return ErrAlreadyExist
	case ErrNotFound:
		d[value] = definition
		break
	}
	return nil
}

func (d Dictionary) Update(value string, definition string) error {
	_, err := d.Search(value)
	switch err {
	case ErrNotFound:
		return ErrNotFoundUpdate
	case nil:
		d[value] = definition
		break
	}
	return nil
}

func (d Dictionary) Delet(value string) error {
	_, err := d.Search(value)
	switch err {
	case ErrNotFound:
		return ErrNotFoundDelet
	case nil:
		delete(d, value)
		break
	}
	return nil
}
