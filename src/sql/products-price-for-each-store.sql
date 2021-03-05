select product_id,
sum(case when store = 'store1' then price else null end )as store1,
sum(case when store = 'store2' then price else null end )as store2,
sum(case when store = 'store3' then price else null end )as store3
from products
group by product_id

