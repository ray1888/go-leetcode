select max(salary) as SecondHighestSalary from Employee where salary < (select max(distinct(salary)) from Employee) as p
