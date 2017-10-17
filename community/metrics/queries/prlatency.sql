//counts number of PRs merged with respect to month 
select left(string(created_at),7), count (*)
from [githubarchive:month.201611],[githubarchive:month.201612],[githubarchive:month.201701]
where repo.name = 'kubernetes/kubernetes'
and type = 'PullRequestEvent'
and JSON_EXTRACT(payload, '$.pull_request.merged') = 'true'
group by 1
