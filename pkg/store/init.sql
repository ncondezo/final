CREATE DATABASE IF NOT EXISTS `dental_clinic`;
USE `dental_clinic`;

CREATE TABLE users
(
    id       VARCHAR(100) NOT NULL,
    name     VARCHAR(25)  NOT NULL,
    surname  VARCHAR(25)  NOT NULL,
    email    VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    CONSTRAINT users_id
        PRIMARY KEY (id),
    CONSTRAINT users_email
        UNIQUE (email)
);
