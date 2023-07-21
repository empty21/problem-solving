Create table If Not Exists Employee
(
    id           int,
    name         varchar(255),
    salary       int,
    departmentId int
);
Create table If Not Exists Department
(
    id   int,
    name varchar(255)
);

with c1 as (select e.*, dense_rank() over (PARTITION BY departmentId order by salary desc) as r from Employee e)
select d.name as Department, c1.name as Employee, c1.salary as Salary
from c1
         left join Department d on c1.departmentId = d.id
where c1.r = 1
