# Write your MySQL query statement below
SELECT c.Name AS Customers FROM Customers AS c LEFT OUTER JOIN Orders AS O ON c.Id = o.CustomerId WHERE o.id IS NULL;