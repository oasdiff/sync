package gcs

type InMemoryStore struct{}

func NewInMemoryStore() Client { return &InMemoryStore{} }

func (m *InMemoryStore) CreateFile(string, string, []byte) error { return nil }
