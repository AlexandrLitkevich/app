CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    encryptor_password varchar not null
);