## API call extension points latency SLIs details

### Definition

| Status | SLI |
| --- | --- |
| WIP | Admission latency for each admission plugin type, measured as 99th percentile over last 5 minutes |
| WIP | Webhook call latency for each webhook type, measured as 99th percentile over last 5 minutes
| WIP | Initializer latency for each initializer, measured as 99th percentile over last 5 minutes |

### User stories
- As an administrator, if API calls are slow, I would like to know if this is
because slow extension points (admission plugins, webhooks, initializers) and
if so which ones are responsible for it.
