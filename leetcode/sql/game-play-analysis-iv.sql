Create table If Not Exists Activity
(
    player_id    int,
    device_id    int,
    event_date   date,
    games_played int
);

with c1 as (select a.player_id, min(event_date) as fld
            from Activity a
            group by player_id),
     c2 as (select a.player_id
            from Activity a
                     left join c1 on a.player_id = c1.player_id
            where date_sub(a.event_date, INTERVAL 1 DAY) = c1.fld),
     c1c as (select count(c1.player_id) as cnt
             from c1),
     c2c as (select count(c2.player_id) as cnt
             from c2)
select round(c2c.cnt / c1c.cnt, 2) as fraction
from c1c,
     c2c
