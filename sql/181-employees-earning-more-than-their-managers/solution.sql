SELECT e2.name AS Employee
    FROM Employee e1 
JOIN Employee e2 on e1.id = e2.managerId 
    WHERE e2.salary > e1.salary