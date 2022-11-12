# SQL查漏补缺
## 1.min函数的用法
> sql中的min函数返回一组中的最小值。`MIN(expression)`

使用示例：
查找员工最低薪水：
```sql
SELECT
    employee_id,
    first_name,
    last_name,
    salary
FROM
    employees
WHERE
    salary = (
        SELECT
            MIN(salary)
        FROM
            employees
    );
```

与group by混用：
查找一个部门薪水最低的员工
```sql
SELECT 
    d.department_id,
    department_name,
    MIN(salary)
FROM 
    employees e
INNER JOIN departments d ON d.department_id = e.department_id
GROUP BY
    d.department_id
```
