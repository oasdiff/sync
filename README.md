# sync

Sync is a service designed to help developers stay ahead of API breaking changes and maintain application reliability. It enables you to register the APIs you rely on and receive notifications whenever a breaking change occurs.

### Key Benefits of Sync
- Proactive Notification System: Sync proactively alerts you to API breaking changes, allowing you to take timely action and prevent disruptions to your applications.
- Centralized API Dependency Management: Sync provides a single platform to manage your API dependencies, simplifying the monitoring process and ensuring that you stay informed about changes across all your APIs.

### How it works?
Signing up for Sync is simple and straightforward. Just provide us with the URLs of the APIs you rely on, and we'll take care of the rest. Our tool will continuously monitor these OpenAPIs for changes and notify you via Slack whenever a breaking change is detected.

The Slack message includes:
- Webhook Name: The name of the webhook triggering the notification.
- Changelog Link: A hyperlink to the detailed changelog.
- Summary of Breaking Changes: A brief summary of the detected breaking changes, including counts of errors and warnings.

### How to use sync?
1. Create a Tenant: Sign up for Sync and create a tenant with a name, your email address and a Slack channel URL. This tenant serves as the central hub for managing your API dependencies.
2. Define Webhooks for Each API: For each OpenAPI specification you depend on, create a webhook.
3. Receive Breaking Change Notifications. Example:
```
OpenAI Changelog
179 breaking changes: 2 error, 177 warning
```
Explore a detailed breakdown of changes in this comprehensive HTML changelog.

#### Create a Tenant
To use Sync, first create a tenant with a name:
```
curl -d '{
    "tenant": "my-company",
    "email": "james@my-company.com",
    "callback": "https://api.my-company.com/webhooks",
    "slack_channel": "https://hooks.slack.com/services/TLDF14G/AG123/abcd"
}' https://sync.oasdiff.com/tenants
```
You will get a response with your tenant ID, that looks like this:
```
{"id": "2ahh9d6a-2221-41d7-bbc5-a950958345"}
```

#### Define Webhooks for Each API
Now, for each OpenAPI spec that you depend on, create a webhook:
```
curl -d '{
    "webhook_name": "OpenAI",
    "owner": "openai",
    "repo": "openai-openapi",
    "path": "openapi.yaml",
    "branch": "master",
    "spec": "https://github.com/openai/openai-openapi/raw/e145786e70bf5fc1bc73c7cd19884f445d52c383/openapi.yaml"
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

Note, currently only GitHub is supported.

You are all set :)

Tip: to discover a vast collection of OpenAPI specifications, consider exploring this open-source [catalog](https://apis.guru/).
