Create table If Not Exists Transactions
(
    id         int,
    country    varchar(4),
    state      enum ('approved', 'declined'),
    amount     int,
    trans_date date
)

with c1 as (select t.*, month(t.trans_date) m, year(t.trans_date) y from Transactions t)
select concat(c1.y, '-', lpad(c1.m, 2, '0'))                      as `month`,
       c1.country,
       count(1)                                     as trans_count,
       sum(if(c1.state = 'approved', 1, 0))         as approved_count,
       sum(c1.amount)                               as trans_total_amount,
       sum(if(c1.state = 'approved', c1.amount, 0)) as approved_total_amount
from c1
group by c1.y, c1.m, c1.country
