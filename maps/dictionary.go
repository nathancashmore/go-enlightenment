/*
Simple Dictionary providing the ability to Search, Add, Update or Delete items.

Provides examples of the following:
 Create maps
 Search for items in maps
 Add new items to maps
 Update items in maps
 Delete items from a map
 Learned more about errors
 How to create errors that are constants
 Writing error wrappers
*/
package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound           = DictionaryErr("could not find the word you were looking for")
	ErrAddAlreadyExists   = DictionaryErr("cannot add word because it already exists")
	ErrUpdateDoesNotExist = DictionaryErr("cannot find the word to update")
)

func (d Dictionary) Search(searchTerm string) (string, error) {
	result, ok := d[searchTerm]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

func (d Dictionary) Add(word string, description string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = description
		return nil
	case nil:
		return ErrAddAlreadyExists
	default:
		return err
	}
}

func (d Dictionary) Update(word string, description string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = description
		return nil
	case ErrNotFound:
		return ErrUpdateDoesNotExist
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
