CREATE TABLE products (
                          id BIGSERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          price NUMERIC(12,2) NOT NULL DEFAULT 0
);