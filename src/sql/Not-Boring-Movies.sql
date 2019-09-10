select id, movie, description, rating from cinema where
description != "boring" group by id having mod(id,2) != 0  order by rating desc