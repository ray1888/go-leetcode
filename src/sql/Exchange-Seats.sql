/*
1. need to count how many items in this table (solving the last odd problem)
2. If the table has even number of records or the id is the last one, check whether id is even or odd. If odd, add 1 as new id, otherwise minus 1 as new id
3. order by id
 */
SELECT IF(cnt % 2 = 1 AND id = cnt, id, IF(id % 2 = 1, id + 1, id - 1)) AS id, student FROM seat,
(SELECT COUNT(*) AS cnt FROM seat) AS t
ORDER BY id;
