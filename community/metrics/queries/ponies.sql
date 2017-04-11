 SELECT
    actor.login,
    actors.company,
    cnt,
    count(*) as contributor
  FROM (
    SELECT
      actor.login, COUNT(actor.login) AS cnt
    FROM
      //small date rage, for testing:
      //( TABLE_DATE_RANGE([githubarchive:day.], TIMESTAMP('2016-09-15'), TIMESTAMP('2016-09-15') ) )
      // k8s 1.2 timeframe
      //( TABLE_DATE_RANGE([githubarchive:day.], TIMESTAMP('2015-11-09'), TIMESTAMP('2016-03-16') ) )
      // k8s 1.3 timeframe
      //( TABLE_DATE_RANGE([githubarchive:day.], TIMESTAMP('2016-03-17'), TIMESTAMP('2016-07-06') ) )
      // k8s 1.4 timeframe
      //( TABLE_DATE_RANGE([githubarchive:day.], TIMESTAMP('2016-07-07'), TIMESTAMP('2016-09-26') ) )
      // k8s 1.5 timeframe
         ( TABLE_DATE_RANGE([githubarchive:day.], TIMESTAMP('2016-09-26'), TIMESTAMP('2016-12-12') ) )
      // k8s 1.6 timeframe
      // ( TABLE_DATE_RANGE([githubarchive:day.], TIMESTAMP('2016-12-13'), TIMESTAMP('2017-3-22') ) )
    WHERE
      (repo.name IN ("kubernetes/kubernetes"))
    // GH Issue filed:
    //  AND type = 'IssuesEvent'
    // GH Issue comment:
    //  AND type = 'IssueCommentEvent'
    // PRs filed:
    AND type = 'PullRequestEvent'
    // PR reviews:
    // AND type = 'PullRequestReviewEvent'
      AND NOT actor.login CONTAINS "-robot"
      AND NOT actor.login CONTAINS "-bot"
    GROUP BY
      actor.login) AS gh
  JOIN (
    SELECT
      login,
      company
    FROM
      [coreos-community-development:community.actors]) AS actors
  ON
    gh.actor.login = actors.login
  WHERE
    actors.company IS NOT NULL
  GROUP BY
    actor.login,
    actors.company,
    cnt
  ORDER BY cnt DESC
  limit 100