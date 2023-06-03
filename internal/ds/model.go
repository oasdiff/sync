package ds

type Tenant struct {
	Id      string `json:"id" datastore:"id"`
	Name    string `json:"name" datastore:"name"`
	Created int64  `json:"created" datastore:"created"`
}

type Webhook struct {
	TenantId string `json:"tenant_id" datastore:"tenant_id"`
	Url      string `json:"url" datastore:"url"`
	Spec     string `json:"spec" datastore:"spec"`
	Created  int64  `json:"created" datastore:"created"`
}
