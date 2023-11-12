# sync

Sync is a service to manage the APIs that you depend on. It is similar to a package manager, such as [npm](https://www.npmjs.com/), but for APIs

### How does sync help me?
- Notifying you of API breaking changes
- Providing a single place to manage API dependencies

### How to use sync?
To use Sync, first create a tenant with a name:
```
curl -d '{
    "tenant": "my-company", \
    "email": "james@my-company.com", \
    "callback": "https://api.my-company.com/webhooks", \
    "slack_channel": "https://hooks.slack.com/services/TLDF14G/AG123/abcd"\
}' https://sync.oasdiff.com/tenants
```
You will get a response with your tenant ID, that looks like this:
```
{"id": "2ahh9d6a-2221-41d7-bbc5-a950958345"}
```

Now, for each OpenAPI spec that you depend on, create a webhook:
```
curl -d '{
    "webhook_name": "OpenAI", \
    "owner": "openai", \
    "repo": "openai-openapi", \
    "path": "openapi.yaml", \
    "branch": "master", \
    "spec": "https://raw.githubusercontent.com/openai/openai-openapi/master/openapi.yaml" \
}' https://sync.oasdiff.com/tenants/{tenant-id}/webhooks
```
You will get a response with created webhook ID, that looks like this:
```
{"id": "2ahh9d6a-3344-41d7-bbc5-a950958345"}
```

Schema input parameters:
1. `webhook_name` a unique webhook name
2. `owner` GitHub repo's owner
3. `repo` GitHub repo
4. `path` path to OpenAPI revision file
5. `branch` branch which the revision OpenAPI file (usually main or master)
6. `spec` full URL to base OpenAPI file
7. `slack_channel` will be called on breaking changes.

Notes:
1. Currently only GitHub is supported.
3. [in development] You will be able to specify an `event_type` that can be one of the follow: `diff` or `breaking-changes` (default) or `changelog`.

You are all set :)

Our service will pull new commit. In case of a new commit's breaking API change, the sync service will notify the provided slack channel.

Providing a *slack channel* for callback, it will notify with the API breaking-changes in a [markdown](https://en.wikipedia.org/wiki/Markdown) format. 

In case of providing a *callback* URL it will send a JSON like the follow:
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
Note, provided callback service should response with HTTP `200 OK` or `201 Created`.
