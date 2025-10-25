--
-- PostgreSQL database dump
--

\restrict UzHGpX4kYYwcEa03dG3BMXlGJPtTJWthqOjpyMW7iLygTe7FCDYIem7mHR7Q54Z

-- Dumped from database version 18.0 (Ubuntu 18.0-1.pgdg25.10+3)
-- Dumped by pg_dump version 18.0 (Ubuntu 18.0-1.pgdg25.10+3)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    merchant_id bigint NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    price numeric(10,2) NOT NULL,
    stock bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id bigint NOT NULL,
    customer_id bigint NOT NULL,
    product_id bigint NOT NULL,
    quantity bigint NOT NULL,
    total_price numeric(10,2) NOT NULL,
    shipping_cost numeric(10,2) NOT NULL,
    discount numeric(10,2) NOT NULL,
    final_price numeric(10,2) NOT NULL,
    is_free_shipping boolean NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.transactions_id_seq OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(255) NOT NULL,
    role character varying(20) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, merchant_id, name, description, price, stock, created_at, updated_at) FROM stdin;
2	5	Mouse Wireless	Mouse wireless ergonomic	250000.00	50	2025-10-25 12:43:53.495969+07	2025-10-25 12:43:53.495969+07
3	5	Keyboard Mechanical	Keyboard mechanical RGB	800000.00	23	2025-10-25 12:44:03.533328+07	2025-10-25 13:03:19.488054+07
4	5	USB Cable	Kabel USB Type-C	50000.00	97	2025-10-25 12:44:13.490373+07	2025-10-25 13:04:40.987063+07
5	5	USB Cable	Kabel USB Type-C	50000.00	100	2025-10-25 13:07:15.590771+07	2025-10-25 13:07:15.590771+07
1	5	Laptop Gaming Pro	Laptop gaming spec tinggi	16000000.00	8	2025-10-25 12:39:19.198889+07	2025-10-25 13:11:24.481474+07
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (id, customer_id, product_id, quantity, total_price, shipping_cost, discount, final_price, is_free_shipping, created_at) FROM stdin;
1	10	3	2	1600000.00	0.00	160000.00	1440000.00	t	2025-10-25 13:03:19.485427+07
2	10	4	2	100000.00	0.00	10000.00	90000.00	t	2025-10-25 13:04:03.252742+07
3	10	4	1	50000.00	0.00	0.00	50000.00	t	2025-10-25 13:04:40.986057+07
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password, role, created_at, updated_at) FROM stdin;
1	John Doe	customer@example.com	$2a$10$H77NaUbzQUlNzTuTb.ARu.ayGpBhvQ4Cd5zaZfnNAzTTENMQK.ERO	customer	2025-10-25 12:29:17.468633+07	2025-10-25 12:29:17.468633+07
2	Customer Baru	customer2@example.com	$2a$10$nXgMBP04cVwUo5PrmcDPUOUBXs2cF1b/30jBO04n4Cfq9yZlW4JCa	customer	2025-10-25 12:30:37.73031+07	2025-10-25 12:30:37.73031+07
3	Customer Baru	customer3@example.com	$2a$10$kd9Py4CtC4VxVIRhDYi.SOYydIHTXUrAm1vZalE7hhKX6ygkLcPLO	customer	2025-10-25 12:32:55.718158+07	2025-10-25 12:32:55.718158+07
4	Customer Baru	customer4@example.com	password123	customer	2025-10-25 12:35:34.156664+07	2025-10-25 12:35:34.156664+07
5	Toko Serba Ada	merchant@example.com	password123	merchant	2025-10-25 12:36:59.312463+07	2025-10-25 12:36:59.312463+07
10	John Doe	customer10@example.com	password123	customer	2025-10-25 12:58:25.107628+07	2025-10-25 12:58:25.107628+07
\.


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 5, true);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 3, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 10, true);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: products fk_products_merchant; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_merchant FOREIGN KEY (merchant_id) REFERENCES public.users(id);


--
-- Name: transactions fk_transactions_customer; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_customer FOREIGN KEY (customer_id) REFERENCES public.users(id);


--
-- Name: transactions fk_transactions_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_product FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- PostgreSQL database dump complete
--

\unrestrict UzHGpX4kYYwcEa03dG3BMXlGJPtTJWthqOjpyMW7iLygTe7FCDYIem7mHR7Q54Z

