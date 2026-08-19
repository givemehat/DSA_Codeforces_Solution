/* Write your PL/SQL query statement below */
select e.name AS Employee
from Employee e
Join Employee m
   On e.managerId=m.id
where e.salary>m.salary;