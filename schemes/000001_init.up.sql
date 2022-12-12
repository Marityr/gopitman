CREATE TABLE users (
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

-- CREATE TABLE category {
--     id not null unique,
--     status bool,
--     title varchar(255) not null,
--     descriptions text
-- }

-- CREATE TABLE product {
--     id not null unique,
--     title varchar(255) not null,
--     category_id int
--     CONSTRAINT "table3_category_4ud7av49y_foreign" FOREIGN KEY ("category_id") REFERENCES "category" ("id")
-- }