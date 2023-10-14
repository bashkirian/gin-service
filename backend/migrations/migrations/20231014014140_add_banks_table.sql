-- +goose Up

GRANT ALL ON schema public TO postgres;

CREATE SCHEMA IF NOT EXISTS public;

ALTER SCHEMA public OWNER TO pg_database_owner;

CREATE TABLE public.banks (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    salepointname text NOT NULL,
    address text NOT NULL,
    status text NOT NULL,
    rko boolean,
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

CREATE TABLE public.clients (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    latitude numeric,
    longitude numeric,
    CONSTRAINT idc_tbl PRIMARY KEY (id)
);


ALTER TABLE public.clients OWNER TO postgres;

CREATE TABLE public.reviews (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    bank_id UUID,
    content text,
    CONSTRAINT idr_tbl PRIMARY KEY (id),
    CONSTRAINT fk_bank
      FOREIGN KEY(bank_id) 
	  REFERENCES public.banks(id)
);


ALTER TABLE public.reviews OWNER TO postgres;

CREATE TABLE public.services (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    description text,
    CONSTRAINT idss_tbl PRIMARY KEY (id)
);


ALTER TABLE public.services OWNER TO postgres;


CREATE TABLE public.banks_services (
    bank_id UUID REFERENCES public.banks(id),
    services_id UUID REFERENCES public.services(id),
    CONSTRAINT idbs_tbl PRIMARY KEY (bank_id, services_id)
);

-- +goose Down
DROP TABLE public.banks_services;
DROP TABLE public.reviews;
DROP TABLE public.banks;
DROP TABLE public.services;
DROP TABLE public.clients;