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
  id_lxd int default null references lxds (id),
  host_port int not null,
  container_port int not null,
  error_message text default null
);

create table metrics (
  id serial primary key,
  id_lxd int not null references lxds (id) unique,
  cpu_usage numeric not null default 0.0,
  memory_usage numeric not null default 0,
  counter int not null default 1
);
