CREATE TABLE users (
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);


CREATE TABLE customer (
	id uuid NOT NULL,
	external_id text NULL,
	bar_code numeric(13) NOT NULL,
	activated bool NOT NULL,
	birthday timestamptz NULL,
	created timestamptz NOT NULL,
	first_name text NULL,
	last_name text NULL,
	second_name text NULL,
	update_at timestamptz NOT NULL,
	sex text NULL,
	referrer_code text NULL,
	CONSTRAINT rest_customer_pkey PRIMARY KEY (id)
);