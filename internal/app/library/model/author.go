package model

// Author ...
type Author struct {
	ID   uint64
	Name string
}

// AuthorList ...
type AuthorList []*Author

// NewAuthor create new author model
func NewAuthor(id uint64, name string) (*Author, error) {
	m := &Author{
		ID:   id,
		Name: name,
	}
	if err := m.Validate(); err != nil {
		return nil, err
	}
	return m, nil
}

// Validate validatte category
func (m *Author) Validate() error {

	if m.Name == "" {
		return ErrInvalidModel
	}
	return nil
}
