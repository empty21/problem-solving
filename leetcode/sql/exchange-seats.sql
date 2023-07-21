Create table If Not Exists Seat
(
    id      int,
    student varchar(255)
)

with c as (
    select count(*) as cnt from Seat
)
select if(s.id % 2 <> 0, if(s.id = c.cnt, s.id, s.id + 1), s.id - 1) id , s.student
from Seat s, c order by id
