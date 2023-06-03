package ds

type InMemoryClient struct{}

func NewInMemoryClient() Client { return &InMemoryClient{} }

func (c *InMemoryClient) GetAll(kind Kind, dst interface{}) error { return nil }

func (c *InMemoryClient) Put(kind Kind, id string, src interface{}) error { return nil }
