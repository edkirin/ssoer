-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	uuid uuid NOT NULL,
	"password" varchar(200) NOT NULL,
	first_name varchar(50) NOT NULL,
	last_name varchar(50) NOT NULL,
	email varchar(250) NOT NULL,
	is_active bool NOT NULL,
	deleted bool NOT NULL,
	date_joined timestamptz NOT NULL,
	last_login timestamptz NULL,
	CONSTRAINT unique_email UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_uuid_key UNIQUE (uuid)
);