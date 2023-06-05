package ds

import "reflect"

type InMemoryClient struct{}

func NewInMemoryClient() Client { return &InMemoryClient{} }

func (c *InMemoryClient) Get(kind Kind, id string, dst interface{}) error {

	if kind == KindTenant {
		reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(Tenant{
			Id:   id,
			Name: "test-123",
		}))
	}

	return nil
}

func (c *InMemoryClient) GetAll(kind Kind, dst interface{}) error { return nil }

func (c *InMemoryClient) Put(kind Kind, id string, src interface{}) error { return nil }
