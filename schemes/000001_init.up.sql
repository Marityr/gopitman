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


CREATE TABLE customermeta (
	id serial4 NOT NULL,
	meta_tag text NULL,
	value text NULL,
	customer_id uuid NOT NULL,
	CONSTRAINT idx_customermeta_value UNIQUE (value),
	CONSTRAINT customermeta_meta_tag_value_73ff1b99_uniq UNIQUE (meta_tag, value),
	CONSTRAINT customermeta_pkey PRIMARY KEY (id),
	CONSTRAINT customermeta_customer_id_85c1d53a_fk_customer_id FOREIGN KEY (customer_id) REFERENCES customer(id) DEFERRABLE INITIALLY DEFERRED
);
CREATE INDEX customermeta_customer_id_85c1d53a ON customermeta USING btree (customer_id);


CREATE TABLE customerreferral (
	referral_id uuid NOT NULL,
	referrer_id uuid NOT NULL,
	CONSTRAINT customerreferral_referral_id_aa82c024_pk PRIMARY KEY (referral_id),
	CONSTRAINT customerreferral_referral_id_aa82c024_uniq UNIQUE (referral_id),
	CONSTRAINT customerreferral_referrer_id_referral_id_88f8ae1a_uniq UNIQUE (referrer_id, referral_id),
	CONSTRAINT customerreferral_referral_id_aa82c024_fk_customer_id FOREIGN KEY (referral_id) REFERENCES customer(id) DEFERRABLE INITIALLY DEFERRED,
	CONSTRAINT customerreferral_referrer_id_3f4874ee_fk_customer_id FOREIGN KEY (referrer_id) REFERENCES customer(id) DEFERRABLE INITIALLY DEFERRED
);
CREATE INDEX customerreferral_referrer_id_3f4874ee ON customerreferral USING btree (referrer_id);


CREATE TABLE transactionsystemname (
	id int4 NOT NULL,
	value varchar(64) NOT NULL,
	CONSTRAINT transactiontype_pkey PRIMARY KEY (id)
);


CREATE TABLE transaction (
	id serial4 NOT NULL,
	customer_id uuid NOT NULL,
	kind varchar(20) NOT NULL,
	name_system_id int4 NOT NULL,
	status int4 NOT NULL,
	"timestamp" timestamptz NOT NULL,
	"cost" numeric(64, 2) NOT NULL,
	bonuses numeric(64, 2) NOT NULL,
	external_source_id varchar(128) NULL,
	source_transaction_id int4 NULL,
	date_update date NULL,
	CONSTRAINT transaction_pkey PRIMARY KEY (id),
	CONSTRAINT transaction_customer_id_588d99cd_fk_rest_customer_id FOREIGN KEY (customer_id) REFERENCES customer(id) DEFERRABLE INITIALLY DEFERRED,
	CONSTRAINT transaction_name_system_id_208dbf73_fk_rest_tran FOREIGN KEY (name_system_id) REFERENCES transactionsystemname(id) DEFERRABLE INITIALLY DEFERRED,
	CONSTRAINT transaction_source_transaction_i_0f5438ca_fk_rest_tran FOREIGN KEY (source_transaction_id) REFERENCES transaction(id) DEFERRABLE INITIALLY DEFERRED
);
CREATE INDEX transaction_customer_id_588d99cd ON transaction USING btree (customer_id);
CREATE INDEX transaction_name_system_id_208dbf73 ON transaction USING btree (name_system_id);
CREATE INDEX transaction_source_transaction_id_0f5438ca ON transaction USING btree (source_transaction_id);
-- external_source_id код чека, номер заказа из ритейла
-- source_transaction_id идентификатор покупки

CREATE TABLE bonuses (
	id serial4 NOT NULL,
	cost_step numeric(64, 2) NOT NULL,
	bonuses numeric(64, 2) NOT NULL,
	percentage_bonuses numeric(5, 2) NOT NULL,
	direction int4 NOT NULL,
	"first" bool NOT NULL,
	CONSTRAINT bonuses_pkey PRIMARY KEY (id)
);