/*Script to create the database and table in PostgreSQL*/

CREATE DATABASE products_db;

\c products_db;

CREATE TABLE "products" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR,
    "description" VARCHAR,
    "price" DECIMAL,
    "quantity" INTEGER
);
