SET search_path TO shop;

-- Напишите запросы на вставку, редактирование и удаление пользователей и продуктов.
INSERT INTO users (name, email, password)
VALUES ('user_1', 'user_1@example.com', 'password1'),
       ('user_2', 'user_2@example.com', 'password2'),
       ('user_3', 'user_3@example.com', 'password3'),
       ('user_4', 'user_4@example.com', 'password4'),
       ('user_5', 'user_5@example.com', 'password5');

INSERT INTO products (name, price)
VALUES ('product_1', 11.11),
       ('product_2', 22.22),
       ('product_3', 33.33),
       ('product_4', 44.44),
       ('product_5', 55.55);

UPDATE users
SET name = 'user_3_updated'
WHERE id = 3;

UPDATE products
SET name  = 'product_5_updated',
    price = 77.77
WHERE id = 5;

DELETE
FROM users
WHERE id = 4;

DELETE
FROM products
WHERE id = 1;

-- Напишите запрос на сохранение и удаление заказов
INSERT INTO orders (user_id, order_date, total_amount)
VALUES (1, '2023-07-01', 33.33),
       (2, '2023-07-02', 55.55),
       (3, '2023-07-03', 77.77),
       (4, '2023-07-04', 99.99),
       (5, '2023-07-05', 66.66);

INSERT INTO orders_products (order_id, product_id)
VALUES (1, 1),
       (1, 2),
       (2, 2),
       (2, 3),
       (3, 3),
       (3, 4),
       (4, 4),
       (4, 5),
       (5, 5),
       (5, 1);

DELETE
FROM orders
WHERE id = 1;

-- Напишите запрос на выборку пользователей и выборку товаров
SELECT id, name, email
FROM users
WHERE email = 'user_1@example.com';

SELECT id, name, price
FROM products
WHERE price < 40;

-- Напишите запрос на выборку заказов по пользователю
SELECT u.name         as user_name,
       u.email        as user_email,
       o.id           as order_nubmer,
       o.order_date   as order_date,
       o.total_amount as amout
FROM users u
         LEFT JOIN orders o
                   ON o.user_id = u.id
WHERE u.email = 'user_1@example.com';

-- Напишите запрос на выборку статистики по пользователю (общая сумма заказов/средняя цена товара)
SELECT u.name       AS user_name,
       u.email      AS user_email,
       ot.amount    AS total_orders_amount,
       pt.avg_price AS avg_products_price
FROM users u
         LEFT JOIN (SELECT user_id, SUM(total_amount) AS amount
                    FROM orders
                    GROUP BY user_id) ot ON u.id = ot.user_id
         LEFT JOIN (SELECT o.user_id, AVG(p.price) AS avg_price
                    FROM orders o
                             INNER JOIN orders_products op ON o.id = op.order_id
                             INNER JOIN products p ON op.product_id = p.id
                    GROUP BY o.user_id) pt ON u.id = pt.user_id
WHERE u.email = 'user_1@example.com';