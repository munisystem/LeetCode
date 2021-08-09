# Write your MySQL query statement below
SELECT e1.Name as Employee
FROM Employee as e1
LEFT JOIN Employee as e2 ON e1.ManagerId = e2.Id
WHERE e1.Salary > e2.Salary
