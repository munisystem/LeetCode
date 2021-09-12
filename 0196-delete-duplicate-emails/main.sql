# Write your MySQL query statement below
DELETE FROM Person WHERE Id NOT IN (SELECT mid FROM (SELECT MIN(Id) as mid FROM Person GROUP BY Email) temp);
