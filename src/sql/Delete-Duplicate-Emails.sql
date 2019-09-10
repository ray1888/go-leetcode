delete from Person  where id not in
(select Id from (
 select min(Id) as Id from Person group by Person.Email) as p
)