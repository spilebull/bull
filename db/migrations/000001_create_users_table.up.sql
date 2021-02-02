ALTER DATABASE template SET timezone TO 'Asia/Tokyo';

create table if not exists  users
(
  id              INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  first_name      TEXT NOT NULL CHECK(length(first_name) >= 1 AND length(first_name) <= 50),
  last_name       TEXT NOT NULL CHECK(length(last_name) >= 1 AND length(last_name) <= 50),
  email           TEXT NOT NULL UNIQUE,
  phone_number    TEXT NOT NULL UNIQUE CHECK(length(phone_number) >= 10 AND length(phone_number) <= 15),
  created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);
