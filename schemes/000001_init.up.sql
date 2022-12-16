CREATE TABLE users (
    id serial NOT NULL unique,
    name varchar(255) NOT NULL,
    username varchar(255) NOT NULL unique,
    --// TODO статус пользователя
    -- status bool,
    password_hash varchar(255) NULL
);

CREATE TABLE essay (
    id serial NOT NULL unique,
    code_object uuid NOT NULL DEFAULT uuid_generate_v4(),
    title varchar(255) NOT NULL,
    address varchar(255) NULL,
    coordinates varchar(255) NULL,
    descriptions varchar(255) NULL 
)
