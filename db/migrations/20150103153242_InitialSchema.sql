
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user_account (
  id SERIAL NOT NULL PRIMARY KEY,
  nickname VARCHAR,
  email VARCHAR NOT NULL UNIQUE,
  username VARCHAR NOT NULL UNIQUE,
  password VARCHAR NOT NULL,
  gender VARCHAR,
  avatar_filename VARCHAR,
  inviter_id INTEGER REFERENCES user_account(id) ON DELETE RESTRICT,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_login_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  is_archived BOOLEAN NOT NULL DEFAULT 'f',
  locale VARCHAR NOT NULL DEFAULT 'en'
);

CREATE TABLE verification_token (
  user_account_id INTEGER REFERENCES user_account(id) ON DELETE RESTRICT,
  code VARCHAR NOT NULL PRIMARY KEY,
  validation_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '1 hour',
  is_valid BOOLEAN NOT NULL DEFAULT 't'
);

CREATE TABLE user_login_oath (
  user_account_id INTEGER REFERENCES user_account(id) ON DELETE RESTRICT,
  service VARCHAR NOT NULL,
  identity VARCHAR NOT NULL,
  access_token VARCHAR NOT NULL,
  PRIMARY KEY (user_account_id, service)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_login_oath;
DROP TABLE verification_token;
DROP TABLE user_account;
