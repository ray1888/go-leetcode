select t.Request_at As Day,
    round(sum(if(status != 'completed', 1, 0)) / sum(1), 2) AS 'Cancellation Rate'   from Trips t
left join Users u on t.Client_Id  = u.Users_id
where u.banned = "No" and t.Request_at BETWEEN "2013-10-01" AND "2013-10-03"
group by t.Request_at
