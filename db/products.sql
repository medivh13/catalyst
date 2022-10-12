CREATE TABLE catalyst.products (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(50) NOT NULL,
    price float DEFAULT 0 NOT NULL,
    brand_id bigint NOT NULL,
    CONSTRAINT fk_brand
      FOREIGN KEY(brand_id) 
	  REFERENCES catalyst.brands(id)
);

CREATE SEQUENCE catalyst.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE catalyst.products_id_seq OWNED BY catalyst.products.id;

ALTER TABLE ONLY catalyst.products ALTER COLUMN id SET DEFAULT nextval('catalyst.products_id_seq'::regclass);

ALTER TABLE ONLY catalyst.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


CREATE INDEX productsindex ON catalyst.products USING btree (deleted_at, id, brand_id);


INSERT INTO catalyst.products (name, price, brand_id)
VALUES ('p1', 1000, 1);
INSERT INTO catalyst.products (name, price, brand_id)
VALUES ('p2', 1000, 2);
