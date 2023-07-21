create table station
(
    id     int,
    city   varchar(21),
    state  varchar(3),
    lat_n  long,
    long_w long
);

with name_length as (select s.city, length(s.city) as len
                     from station s),
     min_record as (select city, len
                    from name_length n
                    order by n.len, n.city
                    limit 1),
     max_record as (select city, len
                    from name_length n
                    order by n.len desc, n.city
                    limit 1)
select *
from min_record
union
select *
from max_record;
