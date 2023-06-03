package gcs

type InMemoryStore struct{}

func NewInMemoryStore() Client { return &InMemoryStore{} }

func (m *InMemoryStore) SaveSpec(tenantId string, name string) error { return nil }
