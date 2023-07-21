Create table If Not Exists Scores (id int, score DECIMAL(3,2))

SELECT s.score, DENSE_RANK() over (ORDER BY s.score DESC) AS `rank` FROM Scores s
