Create table If Not Exists RequestAccepted
(
    requester_id int  not null,
    accepter_id  int  null,
    accept_date  date null
);

select f.fuck as id, sum(1) as num
from (select r1.requester_id as fuck
      from RequestAccepted r1
      union all
      select r2.accepter_id as fuck
      from RequestAccepted r2) as f
group by f.fuck
order by num desc
limit 1
