CREATE DATABASE main WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';
ALTER DATABASE main OWNER TO postgres;

\connect main

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TYPE public.transaction_status AS ENUM (
    'reserved',
    'approved'
);
ALTER TYPE public.transaction_status OWNER TO postgres;

CREATE TABLE public.transaction (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    amount integer NOT NULL,
    order_id uuid NOT NULL,
    service_id uuid NOT NULL,
    status public.transaction_status NOT NULL,
    user_id uuid NOT NULL
);
ALTER TABLE public.transaction OWNER TO postgres;

CREATE TABLE public."user" (
    id uuid NOT NULL,
    amount integer DEFAULT 0 NOT NULL
);
ALTER TABLE public."user" OWNER TO postgres;


ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_order_id_uniq UNIQUE (order_id);

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);
    
CREATE INDEX fki_fk_user_id_transaction ON public.transaction USING btree (user_id);

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT fk_transaction_user_id FOREIGN KEY (user_id) REFERENCES public."user"(id) NOT VALID;

