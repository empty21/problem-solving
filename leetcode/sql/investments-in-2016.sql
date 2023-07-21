Create Table If Not Exists Insurance
(
    pid      int,
    tiv_2015 float,
    tiv_2016 float,
    lat      float,
    lon      float
)

with c1 as (select i.tiv_2016, count(1) over (partition by i.tiv_2015) cy, count(1) over (partition by i.lat, i.lon) cl
            from Insurance i)
select round(sum(c1.tiv_2016), 2) as tiv_2016
from c1
where cy > 1
  and cl = 1;
