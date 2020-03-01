DROP DATABASE IF EXISTS ms;
CREATE DATABASE ms;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    u_id SERIAL PRIMARY KEY,
    u_pretty_id CHAR(36) NOT NULL DEFAULT '',
    u_name VARCHAR(255) NOT NULL DEFAULT '',
    u_password VARCHAR(255) NOT NULL DEFAULT '',
    u_create_time BIGINT NOT NULL DEFAULT 0,
    u_update_time BIGINT NOT NULL DEFAULT 0,
    u_delete_time BIGINT NOT NULL DEFAULT 0,
    UNIQUE(u_name)
);