CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE app_user (
  user_id uuid DEFAULT uuid_generate_v4() not null,
  primary key (user_id)
);