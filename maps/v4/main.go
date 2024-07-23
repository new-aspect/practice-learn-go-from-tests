package main

const ErrWordExist = DictionaryErr("error: already have exist word")
const ErrNotFound = DictionaryErr("could not find the word you are look for")

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Add(key, content string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = content
	case nil:
		return ErrWordExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Search(key string) (string, error) {
	content, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return content, nil
}
