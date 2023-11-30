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

CREATE TABLE IF NOT EXISTS dentists
(
    id       VARCHAR(100) NOT NULL,
    name     VARCHAR(25)  NOT NULL,
    surname  VARCHAR(25)  NOT NULL,
    registry VARCHAR(10)  NOT NULL,
    CONSTRAINT dentists_id
        PRIMARY KEY (id),
    CONSTRAINT dentists_registry
        UNIQUE (registry)
);

CREATE TABLE IF NOT EXISTS patient
(
    id            VARCHAR(100) NOT NULL,
    name          VARCHAR(25)  NOT NULL,
    lastname      VARCHAR(25)  NOT NULL,
    address       VARCHAR(250) NOT NULL,
    dni           VARCHAR(10)  NOT NULL,
    dateup        DATE         NOT NULL,
    CONSTRAINT patient_id
        PRIMARY KEY (id),
    CONSTRAINT patient_dni
        UNIQUE (dni)
);
