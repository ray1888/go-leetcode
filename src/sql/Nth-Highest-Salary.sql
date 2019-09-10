CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      SELECT MAX(Salary) FROM Employee E1
      WHERE N - 1 =
      (SELECT COUNT(DISTINCT(E2.Salary)) FROM Employee E2
      WHERE E2.Salary > E1.Salary)
  );
END