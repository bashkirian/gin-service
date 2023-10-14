-- +goose Up
CREATE SCHEMA IF NOT EXISTS bank;

GRANT ALL ON schema bank TO root;

ALTER SCHEMA bank OWNER TO pg_database_owner;

CREATE TABLE bank.banks (
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


ALTER TABLE bank.banks OWNER TO root;

CREATE TABLE bank.clients (
                                id UUID NOT NULL DEFAULT gen_random_uuid(),
                                latitude numeric,
                                longitude numeric,
                                CONSTRAINT idc_tbl PRIMARY KEY (id)
);


ALTER TABLE bank.clients OWNER TO root;

CREATE TABLE bank.reviews (
                                id UUID NOT NULL DEFAULT gen_random_uuid(),
                                bank_id UUID,
                                content text,
                                CONSTRAINT idr_tbl PRIMARY KEY (id),
                                CONSTRAINT fk_bank
                                    FOREIGN KEY(bank_id)
                                        REFERENCES bank.banks(id)
);


ALTER TABLE bank.reviews OWNER TO root;

CREATE TABLE bank.services (
                                 id UUID NOT NULL DEFAULT gen_random_uuid(),
                                 description text,
                                 CONSTRAINT idss_tbl PRIMARY KEY (id)
);


ALTER TABLE bank.services OWNER TO root;


CREATE TABLE bank.banks_services (
                                       bank_id UUID REFERENCES bank.banks(id),
                                       services_id UUID REFERENCES bank.services(id),
                                       CONSTRAINT idbs_tbl PRIMARY KEY (bank_id, services_id)
);

-- +goose Down
DROP TABLE bank.banks_services;
DROP TABLE bank.reviews;
DROP TABLE bank.banks;
DROP TABLE bank.services;
DROP TABLE bank.clients;
DROP SCHEMA bank;
