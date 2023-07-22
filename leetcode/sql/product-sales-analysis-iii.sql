Create table If Not Exists Sales (sale_id int, product_id int, year int, quantity int, price int);
Create table If Not Exists Product (product_id int, product_name varchar(10));

with first_year as (select s.product_id, s.year, s.quantity, s.price, min(s.year) over (partition by s.product_id) fy from Sales s)
select f.product_id, f.year first_year, f.quantity, f.price from first_year f where f.year = f.fy


