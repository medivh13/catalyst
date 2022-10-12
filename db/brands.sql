CREATE TABLE catalyst.brands (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(50) NOT NULL
);

CREATE SEQUENCE catalyst.brands_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE catalyst.brands_id_seq OWNED BY catalyst.brands.id;

ALTER TABLE ONLY catalyst.brands ALTER COLUMN id SET DEFAULT nextval('catalyst.brands_id_seq'::regclass);

ALTER TABLE ONLY catalyst.brands
    ADD CONSTRAINT brands_pkey PRIMARY KEY (id);


CREATE INDEX brandsindex ON catalyst.brands USING btree (deleted_at, id);


INSERT INTO catalyst.brands (name)
VALUES ('brand1');
INSERT INTO catalyst.brands (name)
VALUES ('brand2');
