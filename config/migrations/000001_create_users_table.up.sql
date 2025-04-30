CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username    varchar(40),
                       password   varchar(100),
                       role varchar(40)
);