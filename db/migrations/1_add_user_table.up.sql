CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE app_user (
  user_id uuid DEFAULT uuid_generate_v4() not null,
  primary key (user_id)
);

CREATE TABLE user_account_token (
  user_id uuid not null,
  private_token varchar,
  item_id varchar,
  PRIMARY KEY (user_id, item_id),
  FOREIGN KEY (user_id) REFERENCES app_user (user_id)
);