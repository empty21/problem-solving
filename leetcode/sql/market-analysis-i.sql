Create table If Not Exists Users
(
    user_id        int,
    join_date      date,
    favorite_brand varchar(10)
);
Create table If Not Exists Orders
(
    order_id   int,
    order_date date,
    item_id    int,
    buyer_id   int,
    seller_id  int
);
Create table If Not Exists Items
(
    item_id    int,
    item_brand varchar(10)
);

select u.user_id as buyer_id, u.join_date, coalesce(count(o.order_id), 0) as orders_in_2019
from Users u left join Orders o on u.user_id = o.buyer_id and YEAR(o.order_date) = 2019 group by u.user_id, u.join_date;
