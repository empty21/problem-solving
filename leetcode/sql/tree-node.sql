Create table If Not Exists Tree
(
    id   int,
    p_id int
);

select t.id, if(t.p_id, if(count(t2.id) > 0, 'Inner', 'Leaf'), 'Root') type
from Tree t
         left join Tree t2 on t.id = t2.p_id
group by t.id;
