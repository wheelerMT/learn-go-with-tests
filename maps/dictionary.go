package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists within the dictionary")
	ErrWordDoesNotExist = DictionaryErr("word does not exist within the dictionary")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, found := d[word]
	if found {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, found := d[word]
	if !found {
		return ErrWordDoesNotExist
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, found := d[word]
	if !found {
		return ErrWordDoesNotExist
	}
	delete(d, word)
	return nil
}
