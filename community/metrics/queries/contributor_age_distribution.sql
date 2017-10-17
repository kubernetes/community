  SELECT actor.login, DATEDIFF(MAX(created_at), MIN(created_at)) AS DAYS_INVOLVED
  FROM [githubarchive:year.2015], [githubarchive:year.2016]  
  WHERE repo.name = "kubernetes/kubernetes"
    AND type = "PullRequestEvent"
    AND NOT actor.login CONTAINS "-robot"
    AND NOT actor.login CONTAINS "-bot"
  GROUP BY actor.login
  HAVING DAYS_INVOLVED > 1
  LIMIT 100
