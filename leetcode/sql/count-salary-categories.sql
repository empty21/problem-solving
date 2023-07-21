Create table If Not Exists Accounts
(
    account_id int,
    income     int
)

with c1 as (select sum(if(a.income < 20000, 1, 0))                 as low,
                   sum(if(a.income BETWEEN 20000 AND 50000, 1, 0)) as average,
                   sum(if(a.income > 50000, 1, 0))                 as high
            from Accounts a)
select 'High Salary' category, high accounts_count
from c1

union all
select 'Low Salary' category, low accounts_count
from c1

union all
select 'Average Salary' category, average accounts_count
from c1;

