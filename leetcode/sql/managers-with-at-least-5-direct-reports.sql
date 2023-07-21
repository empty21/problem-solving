Create table If Not Exists Employee (id int, name varchar(255), department varchar(255), managerId int)

with c1 as (
    select e.managerId, sum(1) c from Employee e group by e.managerId
) select e.name from Employee e left join c1 on e.id = c1.managerId where c1.c > 4
