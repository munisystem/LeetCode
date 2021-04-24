SELECT p.FirstName, p.LastName, a.City, a.State FROM Person as p LEFT OUTER JOIN Address as a ON p.PersonId = a.PersonId;
