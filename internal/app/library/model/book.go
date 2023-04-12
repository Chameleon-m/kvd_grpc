package model

// Book ...
type Book struct {
	ID   int64
	Name string
}

// BookList ...
type BookList []*Author

// NewBook create new Book model
func NewBook(id int64, name string) (*Book, error) {
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
