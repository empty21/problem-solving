Create table If Not Exists Products
(
    product_id  int,
    new_price   int,
    change_date date
);

with c1 as (select p.product_id, max(change_date) as mcd
            from Products p
            where p.change_date <= '2019-08-16'
            group by p.product_id),

     c2 as (select distinct product_id from Products)
select c2.product_id, if(p.change_date, p.change_date, 10) as price
from c2
         left join c1 l on c2.product_id = l.product_id
         left join Products p on l.product_id = p.product_id and l.mcd = p.change_date;
