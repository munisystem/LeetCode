SELECT CASE WHEN COUNT(*) > 0 THEN t.Salary ELSE null END as SecondHighestSalary FROM (SELECT Salary FROM Employee GROUP BY Salary ORDER BY Salary DESC LIMIT 1 OFFSET 1) AS t;
