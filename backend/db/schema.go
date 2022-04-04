package db

var schema = `
CREATE TABLE user (
    id       VARCHAR(16)  PRIMARY KEY,
	username VARCHAR(255) UNIQUE NOT NULL,
	email    VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	salt     VARCHAR(255) NOT NULL
);
`
