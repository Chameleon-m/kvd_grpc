package model

// Book ...
type Book struct {
	ID   uint64
	Name string
}

// BookList ...
type BookList []*Book

// NewBook create new Book model
func NewBook(id uint64, name string) (*Book, error) {
	m := &Book{
		ID:   id,
		Name: name,
	}
	if err := m.Validate(); err != nil {
		return nil, err
	}
	return m, nil
}

// Validate validatte category
func (m *Book) Validate() error {

	if m.Name == "" {
		return ErrInvalidModel
	}
	return nil
}
