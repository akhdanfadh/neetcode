-- Write your query below
select c.customer_id, c.customer_name
from customers c
where
    exists (select 1 from orders o where o.customer_id = c.customer_id and o.product_name = 'A')
    and exists (select 1 from orders o where o.customer_id = c.customer_id and o.product_name = 'B')
    and not exists (select 1 from orders o where o.customer_id = c.customer_id and o.product_name = 'C')
order by customer_name;
