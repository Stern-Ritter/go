CREATE SCHEMA shop AUTHORIZATION admin;

CREATE TABLE IF NOT EXISTS shop.users
(
    id       BIGSERIAL PRIMARY KEY,
    name     VARCHAR(255) NOT NULL DEFAULT '',
    email    VARCHAR(255) NOT NULL,
    password TEXT         NOT NULL,
    CONSTRAINT uq_users_email_idx UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS shop.orders
(
    id           BIGSERIAL PRIMARY KEY,
    user_id      BIGINT         NOT NULL,
    order_date   DATE           NOT NULL DEFAULT CURRENT_DATE,
    total_amount NUMERIC(10, 2) NOT NULL,
    CONSTRAINT fk_orders_to_users FOREIGN KEY (user_id) REFERENCES shop.users (id),
    CONSTRAINT ck_orders_total_amount CHECK (total_amount > 0)
);

CREATE TABLE IF NOT EXISTS shop.products
(
    id    BIGSERIAL PRIMARY KEY,
    name  varchar(255)   NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    CONSTRAINT uq_products_name_idx UNIQUE (name),
    CONSTRAINT ck_products_price CHECK (price > 0)
);

CREATE TABLE IF NOT EXISTS shop.orders_products
(
    order_id   BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    PRIMARY KEY (order_id, product_id),
    CONSTRAINT fk_orders_products_to_orders FOREIGN KEY (order_id) REFERENCES shop.orders (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_orders_products_to_products FOREIGN KEY (product_id) REFERENCES shop.products (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_products_price ON shop.products (price);