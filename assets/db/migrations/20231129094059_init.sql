-- migrate:up
create table users (
  id integer,
  shortURL varchar(255),
  longURL varchar(255)
);

-- migrate:down
drop table users;
