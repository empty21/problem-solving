Create table If Not Exists Customer
(
    customer_id int,
    product_key int
);
Create table Product
(
    product_key int
);

with c1 as (select count(distinct p.product_key) cnt from Product p),
     c2 as (select customer_id cid, count(distinct product_key) cnt
            from Customer c
            group by customer_id)
select c2.cid customer_id from c2, c1 where c2.cnt = c1.cnt;
