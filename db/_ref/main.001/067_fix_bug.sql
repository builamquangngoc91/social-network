ALTER TABLE variant ALTER COLUMN etop_status DROP NOT NULL;

ALTER TABLE order_line ALTER COLUMN status DROP NOT NULL;
ALTER TABLE order_line ALTER COLUMN wholesale_price_0 DROP NOT NULL;
ALTER TABLE order_line ALTER COLUMN wholesale_price DROP NOT NULL;