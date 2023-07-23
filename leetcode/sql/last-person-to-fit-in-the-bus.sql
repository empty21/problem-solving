Create table If Not Exists Queue
(
    person_id   int,
    person_name varchar(30),
    weight      int,
    turn        int
);

with c1 as (select q.*, sum(q.weight) over (order by turn) as sw
            from Queue q)
select c1.person_name
from c1
where c1.sw <= 1000 order by c1.turn desc limit 1
