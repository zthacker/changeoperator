CREATE DATABASE changesdb;

\c changesdb;

CREATE TABLE changes (
                         id serial PRIMARY KEY,
                         requester VARCHAR,
                         env VARCHAR,
                         type VARCHAR,
                         customerimpact VARCHAR,
                         description VARCHAR,
                         date DATE,
                         link VARCHAR,
                         linkback VARCHAR,
                         time TIMESTAMP
);