DROP TABLE IF EXISTS lxds, lxcs, metrics;

create table lxds (
  id serial primary key,
  name text default null unique,
  ip_address text not null unique
);

create table lxcs (
  id serial primary key,
  name text not null unique,
  ip_address text default null unique,
  image text not null,
  status text not null default 'pending',
  id_lxd int default null references lxds (id)
);

create table metrics (
  id serial primary key,
  id_lxd int not null references lxds (id),
  cpu_usage numeric not null default 0.0,
  memory_usage numeric not null default 0
);
