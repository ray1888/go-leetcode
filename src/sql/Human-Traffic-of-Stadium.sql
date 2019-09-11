select distinct t1.id, t1.visit_date , t1.people from stadium t1, stadium t2, stadium t3
where t1.people > 100
      and t2.people > 100
      and t3.people > 100
      and
      (
          (
                t2.id - t1.id = 1
                and t3.id - t2.id = 1
          )
          or (
                t1.id - t2.id = 1
                and t2.id - t3.id = 1
          )
          or (
                t1.id - t2.id = 1
                and t3.id - t1.id = 1
          )
      )
order by t1.id