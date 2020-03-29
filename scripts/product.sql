DROP TABLE public.product;

CREATE TABLE public.product
(
    id serial PRIMARY KEY,
    name VARCHAR(100),
    quantity integer,
    price double precision
)
