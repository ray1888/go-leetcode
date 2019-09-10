select id from Weather w1, Weather w2 where
w1.Temperature > w2.Temperature and
TO_DAYS(w1.RecordDate) - TO_DAYS(w2.RecordDate) =1