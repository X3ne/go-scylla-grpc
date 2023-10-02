CREATE KEYSPACE users WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy','DC1' : 3};

USE users;

CREATE TABLE users (
  id text,
  username text,
  password text,
  primary key((id, username))
);
