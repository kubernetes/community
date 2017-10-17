select nth(
  51, quantiles(
    datediff( REPLACE(JSON_EXTRACT(payload, '$.pull_request.merged_at'), '"', ''), REPLACE(JSON_EXTRACT(payload, '$.pull_request.created_at'), '"', '') )
    , 101)
    // 50th quantile is approximately equivalent to median value, and is measured in days thanks to `datediff`
)
from [githubarchive:month.201611]
where repo.name = 'kubernetes/kubernetes'
and type = 'PullRequestEvent'
and JSON_EXTRACT(payload, '$.pull_request.merged') = 'true'
