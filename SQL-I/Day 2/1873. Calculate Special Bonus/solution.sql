SELECT
    employee_id,
    IF(mod(employee_id,2) = 0 OR LEFT(name, 1) = 'M', 0, salary) as bonus
FROM
    Employees
ORDER BY
    employee_id