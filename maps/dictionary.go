package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("can not add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("can not update word because it does not exist")
)

func (de DictionaryErr) Error() string {
	return string(de)
}

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]

	if ok {
		return ErrWordExists
	}
	d[key] = value
	return nil
}
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	if err == nil {
		d[key] = value
		return nil
	}
	return ErrWordDoesNotExist
}
