CREATE TABLE catalyst.orders (
    id bigint NOT NULL,
    order_code character varying(50) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    total float DEFAULT 0 NOT NULL
);

CREATE SEQUENCE catalyst.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE catalyst.orders_id_seq OWNED BY catalyst.orders.id;

ALTER TABLE ONLY catalyst.orders ALTER COLUMN id SET DEFAULT nextval('catalyst.orders_id_seq'::regclass);

ALTER TABLE ONLY catalyst.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


CREATE INDEX ordersindex ON catalyst.orders USING btree (deleted_at, id, order_code);
