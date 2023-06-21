# sync

Sync is a service to manage the APIs that you depend on. It is similar to a package manager, such as [npm](https://www.npmjs.com/), but for APIs

### How does sync help me?
- Notifying you of API breaking changes
- Providing a single place to manage API dependencies

### How to use sync?
To use Sync, first create a tenant with a name:
```
curl -d '{"tenant": "my-company", "email": "james@my-company.com", "callback": "https://api.my-company.com/webhooks", "slack_channel": "https://hooks.slack.com/services/TLDF14G/AG123/abcd"}' https://sync.oasdiff.com/tenants
```
You will get a response with your tenant ID, that looks like this:
```
{"tenant-id": "2ahh9d6a-2221-41d7-bbc5-a950958345"}
```

Now, for each OpenAPI spec that you depend on, create a webhook:
```
curl -d '{"spec": "https://some-service.com/balloons"}' https://sync.oasdiff.com/tenants/{tenant-id}/webhooks
```
Notes:
1. You can specify a `callback` an endpoint to be called on breaking changes, or `slack_channel` or both.
2. [in development] You will be able to specify an `event_type` that can be one of the follow: `diff` or `breaking-changes` (default) or `changelog`.

You are all set. In case of a breaking API change, the sync service will notify you using the provided callback URL with the breaking errors. For example:
```
{
    "breaking-changes": [
        {
            "id": "response-success-status-removed",
            "text": "removed the success response with the status '200'",
            "level": 0,
            "operation": "GET",
            "operationId": "GetSecurityScore",
            "path": "/api/{domain}/{project}/badges/security-score",
            "source": "https://some-service.com/balloons"
        },
        {
            "id": "request-parameter-removed",
            "text": "deleted the 'cookie' request parameter 'test'",
            "level": 1,
            "operation": "GET",
            "operationId": "GetSecurityScore",
            "path": "/api/{domain}/{project}/badges/security-score",
            "source": "https://some-service.com/balloons"
        },
        {
            "id": "request-parameter-removed",
            "text": "deleted the 'header' request parameter 'user'",
            "level": 1,
            "operation": "GET",
            "operationId": "GetSecurityScore",
            "path": "/api/{domain}/{project}/badges/security-score",
            "source": "https://some-service.com/balloons"
        },
        {
            "id": "request-parameter-removed",
            "text": "deleted the 'query' request parameter 'filter'",
            "level": 1,
            "operation": "GET",
            "operationId": "GetSecurityScore",
            "path": "/api/{domain}/{project}/badges/security-score",
            "source": "https://some-service.com/balloons"
        }
    ]
}
```
Note, provided callback service should response with HTTP '200 OK' or '201 Created'

