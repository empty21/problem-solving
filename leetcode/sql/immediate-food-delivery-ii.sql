Create table If Not Exists Delivery
(
    delivery_id                 int,
    customer_id                 int,
    order_date                  date,
    customer_pref_delivery_date date
);

with ranked_orders as (select d.*, rank() over (partition by d.customer_id order by d.order_date) rn from Delivery d),
     first_orders as (select d.customer_id, order_date, customer_pref_delivery_date
                      from ranked_orders d
                      where d.rn = 1),
     c1 as (select sum(if(fo.order_date = fo.customer_pref_delivery_date, 1, 0)) cnt1, count(1) as cnt2
            from first_orders fo)
select round(c1.cnt1 * 100 / c1.cnt2, 2) immediate_percentage
from c1
