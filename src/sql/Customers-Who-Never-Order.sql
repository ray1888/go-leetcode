select Name as Customers from Customers as c left join
Orders as o on c.Id = o.CustomerId
where o.CustomerId is null