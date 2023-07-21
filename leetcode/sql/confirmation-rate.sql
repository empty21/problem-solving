Create table If Not Exists Signups
(
    user_id    int,
    time_stamp datetime
);
Create table If Not Exists Confirmations
(
    user_id    int,
    time_stamp datetime,
    action     ENUM ('confirmed','timeout')
);

with c1 as (select c.user_id, SUM(IF(action = 'timeout', 1, 0)) ct, SUM(1) as cc from Confirmations c group by c.user_id)
select s.user_id, round(coalesce(c1.ct, 0) / IF(coalesce(c1.cc, 0) = 0, 1, c1.cc), 2) as confirmation_rate
from Signups s
         LEFT JOIN c1 on s.user_id = c1.user_id
