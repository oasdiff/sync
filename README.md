# sync

Sync is a service to manage the APIs that you depend on. It is similar to a package manager, such as [npm](https://www.npmjs.com/), but for APIs

### How does sync help me?
- Notifying you of API breaking changes
- Providing a single place to manage API dependencies

### How to use sync?
To use Sync, first create a tenant with a name:
```
curl -d '{"tenant": "test"}' https://sync.oasdiff.com/tenants
```
You will get a response with your tenant ID, that looks like this:
```
{"tenant-id":"2a849d6a-2221-41d7-bbc5-a9509582345"}
```

Now, for each OpenAPI spec that you depend on, create a webhook:
```
curl -d '{"callback": "https://api.example.com/webhooks", "spec": "https://some-service.com/balloons"}' https://sync.oasdiff.com/tenants/{tenant-id}/webhooks
```
You are all set. In case of a breaking API change, the sync service will notify you using the provided callback URL with the breaking errors. For example:
```
{
    "breaking-changes": [
        {
            "id": "response-success-status-removed",
            "text": "removed the success response with the status \u001b[1m'200'\u001b[0m",
            "level": 0,
            "operation": "GET",
            "operationId": "GetSecurityScore",
            "path": "/api/{domain}/{project}/badges/security-score",
            "source": "https://some-service.com/balloons"
        },
        {
            "id": "response-success-status-removed",
            "text": "removed the success response with the status \u001b[1m'201'\u001b[0m",
            "level": 0,
            "operation": "GET",
            "operationId": "GetSecurityScore",
            "path": "/api/{domain}/{project}/badges/security-score",
            "source": "https://some-service.com/balloons"
        },
        {
            "id": "request-parameter-removed",
            "text": "deleted the \u001b[1m'cookie'\u001b[0m request parameter \u001b[1m'test'\u001b[0m",
            "level": 1,
            "operation": "GET",
            "operationId": "GetSecurityScore",
            "path": "/api/{domain}/{project}/badges/security-score",
            "source": "https://some-service.com/balloons"
        },
        {
            "id": "request-parameter-removed",
            "text": "deleted the \u001b[1m'header'\u001b[0m request parameter \u001b[1m'user'\u001b[0m",
            "level": 1,
            "operation": "GET",
            "operationId": "GetSecurityScore",
            "path": "/api/{domain}/{project}/badges/security-score",
            "source": "https://some-service.com/balloons"
        },
        {
            "id": "request-parameter-removed",
            "text": "deleted the \u001b[1m'query'\u001b[0m request parameter \u001b[1m'filter'\u001b[0m",
            "level": 1,
            "operation": "GET",
            "operationId": "GetSecurityScore",
            "path": "/api/{domain}/{project}/badges/security-score",
            "source": "https://some-service.com/balloons"
        }
    ]
}
```
Note, provided callback service should response with HTTP 200 OK or 201 Created
