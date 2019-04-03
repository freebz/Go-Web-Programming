-- Listing 6.5  Script that creates our database

create table posts (
  id     serial primary key,
  content text,
  author  varchar(255)
);
