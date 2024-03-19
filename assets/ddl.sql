create database ecommerce;
create extension if not exists 
"uuid-ossp";

create type userrole as enum ('admin','costumer');

create table users (
    id uuid default uuid_generate_v4() primary key,
    username varchar(50) unique,
    userrole userrole,
    address text,
    email varchar(70),
    password text,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);
create table category (
    id uuid DEFAULT uuid_generate_v4() primary key,
    category_name VARCHAR(50),
    created_at timestamp default current_timestamp,
    updated_at timestamp
);
create table images (
    id uuid DEFAULT uuid_generate_v4() primary key,
    image VARCHAR(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp
);
create table products (
    id uuid default uuid_generate_v4() primary key,
    product_name VARCHAR(80),
    description text,
    price int,
    stock_quantity INT,
    category_id UUID,
    image_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    Foreign Key (category_id) REFERENCES category(id),
    Foreign Key (image_id) REFERENCES images(id)
);

create table orderTable (
    id uuid DEFAULT uuid_generate_v4() primary key,
    user_id uuid,
    order_date date,
    amount int,
    created_at timestamp default current_timestamp,
    updated_at timestamp default now(),
    Foreign Key (user_id) REFERENCES users(id)
);
create table orderDetails (
    id uuid DEFAULT uuid_generate_v4() primary key,
    order_id uuid,
    product_id uuid,
    quantity int,
    total_amount int,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    Foreign Key (order_id) REFERENCES orderTable(id),
    Foreign Key (product_id) REFERENCES products(id)
);

insert into orderDetails (order_id,product_id, quantity) VALUES ('b91f035a-1bcb-403d-98b2-4b37fc99bca3','948acd64-01d8-4ce5-9ab7-3a28b3c7f609', 5);
select * from orderDetails;

select ot.user_id,ot.amount - ( p.price * od.quantity)  as total
from orderDetails od 
join products  p on od.product_id = p.id
join orderTable ot on od.order_id = ot.id
WHERE ot.id = 'b91f035a-1bcb-403d-98b2-4b37fc99bca3';



