-- +goose Up
--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0 (Ubuntu 16.0-1.pgdg22.04+1)
-- Dumped by pg_dump version 16.0 (Ubuntu 16.0-1.pgdg22.04+1)

SET client_encoding = 'UTF8';

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public ON CONFLICT DO NOTHING;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;


ALTER TABLE public.bank_services OWNER TO postgres;

--
-- Name: banks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.banks (
    id UUID NOT NULL DEFAULT uuid_generate_v1(),
    salepointname text NOT NULL,
    address text NOT NULL,
    status text NOT NULL,
    rko boolean text,
    officetype text NOT NULL,
    salepointformat text NOT NULL,
    suoavailability boolean NOT NULL,
    hasramp boolean,
    latitude numeric NOT NULL,
    longitude numeric NOT NULL,
    metrostation text,
    distance int NOT NULL,
    kep boolean,
    mybranch boolean NOT NULL,
    CONSTRAINT id_tbl PRIMARY KEY (id)
);


ALTER TABLE public.banks OWNER TO postgres;

--
-- Name: clients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.clients (
    id UUID NOT NULL DEFAULT uuid_generate_v1(),
    latitude numeric,
    longitude numeric,
    CONSTRAINT id_tbl PRIMARY KEY (id)
);


ALTER TABLE public.clients OWNER TO postgres;

--
-- Name: reviews; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reviews (
    id UUID NOT NULL DEFAULT uuid_generate_v1(),
    bank_id integer,
    content text,
    CONSTRAINT id_tbl PRIMARY KEY (id)
    CONSTRAINT fk_bank
      FOREIGN KEY(bank_id) 
	  REFERENCES banks(bank_id)
);


ALTER TABLE public.reviews OWNER TO postgres;

--
-- Name: services; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.services (
    id UUID NOT NULL DEFAULT uuid_generate_v1(),
    description text,
    CONSTRAINT id_tbl PRIMARY KEY (id)
);


ALTER TABLE public.services OWNER TO postgres;

--
-- PostgreSQL database dump complete
--

-- +goose Down
