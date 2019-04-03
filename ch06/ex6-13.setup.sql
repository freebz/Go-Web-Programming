-- Listing 6.13  Setting up two related tables

drop table posts casdate if exists;
drop table comments if exists;

create table posts (
  id     serial primary key,
  content text,
  author  varchar(255)
);

create table comments (
  id     serial primary key,
  content text,
  author  varchar(255),
  post_id integer references posts(id)
);
