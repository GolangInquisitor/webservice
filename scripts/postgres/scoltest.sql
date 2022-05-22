--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: fio_create(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.fio_create() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
	BEGIN
	-- Ro
	NEW.fio=concat_ws(' ',NEW.SURNAME,NEW.NAME,NEW.MIDLENAME);
 
	RETURN new;
END$$;


ALTER FUNCTION public.fio_create() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    id bigint NOT NULL,
    product uuid NOT NULL,
    user_id uuid NOT NULL,
    price numeric NOT NULL,
    description text NOT NULL
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: price_inc; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.price_inc
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.price_inc OWNER TO postgres;

--
-- Name: price; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.price (
    id bigint DEFAULT nextval('public.price_inc'::regclass) NOT NULL,
    product uuid NOT NULL,
    currency character varying NOT NULL,
    amount numeric NOT NULL
);


ALTER TABLE public.price OWNER TO postgres;

--
-- Name: product; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product (
    uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    description text NOT NULL,
    left_in_stock bigint NOT NULL
);


ALTER TABLE public.product OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255),
    surname character varying(255),
    midlename character varying(255),
    fio character varying(255),
    gender character varying(255),
    age integer
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (uuid, id, product, user_id, price, description) FROM stdin;
1ace0632-1c40-47cb-bf73-c56f9240cd20	1	9ff79fc3-4fd0-4d66-8d15-4833946e4537	bfc4cff1-396e-4ae9-9317-46ddf2cd010e	1011.22	Конь копченый
b3acfe3e-7307-4422-8c1a-8d5307cc39cc	1	5050fd2d-5694-4bd9-b4d6-9ec8fc4b6620	bfc4cff1-396e-4ae9-9317-46ddf2cd010e	13011.22	Брус деревянный
82ac32c6-b4ab-4f54-b8c9-baba380563ce	1	29d61665-f810-45aa-8ab3-ddfa5cc8524e	bfc4cff1-396e-4ae9-9317-46ddf2cd010e	10.33	Чашка 
08f5fcf6-d101-4b81-8e5c-8411092e6a7c	1	9ff79fc3-4fd0-4d66-8d15-4833946e4537	cfc4eaae-5153-498e-95bd-bcc726d14060	1011.22	Конь копченый
b5a32298-c680-4b3b-ac12-9dbf2deb9f45	1	9ff79fc3-4fd0-4d66-8d15-4833946e4537	cfc4eaae-5153-498e-95bd-bcc726d14060	1011.22	Конь копченый
7c3376e2-58ca-4078-b4a9-1123c477c8e2	1	206be398-0a2d-45af-97ee-e133e4da7eb7	cfc4eaae-5153-498e-95bd-bcc726d14060	3011.22	Клавиатура
\.


--
-- Data for Name: price; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.price (id, product, currency, amount) FROM stdin;
20	29d61665-f810-45aa-8ab3-ddfa5cc8524e	BTC	10.33
21	9ff79fc3-4fd0-4d66-8d15-4833946e4537	RUB	1011.22
22	206be398-0a2d-45af-97ee-e133e4da7eb7	RUB	3011.22
23	5050fd2d-5694-4bd9-b4d6-9ec8fc4b6620	RUB	13011.22
24	ddd00cc4-fa55-4d6f-ba0b-2cb361808ff2	RUB	13011.22
25	27a5d402-d698-46bf-ab1f-c5b9693a5cde	RUB	13011.22
26	32c64710-0576-42e7-900b-925df9221063	EUR	1.22
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product (uuid, description, left_in_stock) FROM stdin;
ddd00cc4-fa55-4d6f-ba0b-2cb361808ff2	Ролики	111
27a5d402-d698-46bf-ab1f-c5b9693a5cde	Стакан	1
32c64710-0576-42e7-900b-925df9221063	Краска	1
5050fd2d-5694-4bd9-b4d6-9ec8fc4b6620	Брус деревянный	2221
29d61665-f810-45aa-8ab3-ddfa5cc8524e	Чашка 	40
9ff79fc3-4fd0-4d66-8d15-4833946e4537	Конь копченый	219
206be398-0a2d-45af-97ee-e133e4da7eb7	Клавиатура	221
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (uuid, name, surname, midlename, fio, gender, age) FROM stdin;
bfc4cff1-396e-4ae9-9317-46ddf2cd010e	Иван	Васильев	Петрович	Васильев Иван Петрович	male	23
cfc4eaae-5153-498e-95bd-bcc726d14060	Алла	Иванова	Семновна	Иванова Алла Семновна	female	30
288c8739-a90d-41e1-bd32-d57a787e541e	Ростислав	Каренин	Витольдович	Каренин Ростислав Витольдович	male	60
\.


--
-- Name: price_inc; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.price_inc', 26, true);


--
-- Name: orders order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT order_pkey PRIMARY KEY (uuid);


--
-- Name: price price_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.price
    ADD CONSTRAINT price_pkey PRIMARY KEY (id);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (uuid);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (uuid);


--
-- Name: users fio_before_ins; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER fio_before_ins BEFORE INSERT OR UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.fio_create();


--
-- Name: orders product_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT product_fk FOREIGN KEY (product) REFERENCES public.product(uuid) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: price product_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.price
    ADD CONSTRAINT product_fk FOREIGN KEY (product) REFERENCES public.product(uuid) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: orders user_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES public.users(uuid) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

