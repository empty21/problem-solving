Create table If Not Exists Logs (id int, num int)



with cte as (
    select ll.*, (select count(ll3.n2) from (select ll2.num n2 from Logs ll2 WHERE ll2.id >= ll.id limit 3) as ll3 where ll3.n2 = ll.num) as appearance from Logs ll
)
select distinct c.num as ConsecutiveNums from cte c WHERE c.appearance > 2
