insert into lxds (name, ip_address) values ('Cluster 1','127.0.0.1');
insert into lxds (name, ip_address) values ('Cluster 2','127.0.0.2');
insert into lxds (name, ip_address) values ('Cluster 3','127.0.0.3');

insert into lxcs (name, image, id_lxd) values ('Ruby','xenial64',1);
insert into lxcs (name, image, id_lxd) values ('GOPAY Scheduler','xenial64',1);
insert into lxcs (name, image, id_lxd) values ('GOPAY Backend','xenial64',2);


insert into metrics (id_lxd, cpu_usage, memory_usage, counter) values (2,10.00,4096,2);
insert into metrics (id_lxd, cpu_usage, memory_usage, counter) values (3,20.00,2048,2);
