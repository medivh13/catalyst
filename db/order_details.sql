CREATE TABLE catalyst.order_details (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    order_id bigint NOT NULL,
    product_id bigint NOT NULL,
    quantity int NOT NULL,
    CONSTRAINT fk_order
      FOREIGN KEY(order_id) 
	  REFERENCES catalyst.orders(id),
    CONSTRAINT fk_product
      FOREIGN KEY(product_id) 
	  REFERENCES catalyst.products(id)
);

CREATE SEQUENCE catalyst.order_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE catalyst.order_details_id_seq OWNED BY catalyst.order_details.id;

ALTER TABLE ONLY catalyst.order_details ALTER COLUMN id SET DEFAULT nextval('catalyst.order_details_id_seq'::regclass);

ALTER TABLE ONLY catalyst.order_details
    ADD CONSTRAINT order_details_pkey PRIMARY KEY (id);


CREATE INDEX order_detailsindex ON catalyst.order_details USING btree (deleted_at, id);

