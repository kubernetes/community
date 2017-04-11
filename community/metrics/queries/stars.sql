  SELECT
  
  SUM(cnt) AS stars,
  From (
  SELECT
    actor.login,
    cnt,
    count(*) as contributor
  FROM (
    SELECT
      actor.login, COUNT(actor.login) AS cnt
    FROM
      ( TABLE_DATE_RANGE([githubarchive:day.], TIMESTAMP('2016-09-15'), TIMESTAMP('2016-09-15') ) )
    WHERE
      (repo.name IN ("kubernetes/kubernetes"))
      AND type = 'WatchEvent'
      AND NOT actor.login CONTAINS "-robot"
      AND NOT actor.login CONTAINS "-bot"
    GROUP BY
      actor.login) AS gh
  GROUP BY
    actor.login,
    cnt
)

