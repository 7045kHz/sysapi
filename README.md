
# SysAPI

## Environment Variables

Create an .env file in the same directory as the API Binary. Add the following variables to it.

```bash
$ cat .env
PG_CONNECT="your posgress connection string"
SECRET="your JWT Token secret"
```

```go
go env GO111MODULE=off
```

## Postgress DB Scheme

## Create the table with the following sql statement

```sql

create table users (
  id serial primary key,
  email text not null unique,
  password text not null,
  level int,
  token text
);

drop table systems;
create  table systems (
  id serial primary key,
  hostname text not null unique,
  domain text not null,
  uuid text not null unique,
  device_type text not null,
  device_sn text not null unique,
  model text not null unique,
  vendor text not null,
  environment text,
  scan_timestamp timestamp  
);

drop table cmd_sysctl_a;
create  table cmd_sysctl_a (
  id serial primary key,
  machine_id text not null unique,
  setting text not null unique,
  value text,
  scan_timestamp timestamp  
);

drop table procfs_mounts;
create  table procfs_mounts (
  id serial primary key,
  machine_id text not null unique,
  mount_point text not null unique,
  mount_source text not null unique,
  fs_type text not null unique,
  fs_mntops text not null unique,
  fs_freq text not null unique,
  fs_passno text not null unique,
  scan_timestamp timestamp  
);

drop table procfs_meminfo;
create  table procfs_meminfo (
  id serial primary key,
  machine_id text not null unique,
  setting text not null unique,
  unit int,
  unit_symbol text not null unique,
  scan_timestamp timestamp  
);

drop table procfs_swaps;
create  table procfs_swaps (
  id serial primary key,
  machine_id text not null unique,
  swap_device text not null unique,
  swap_type text not null  ,
  size int   not null ,
  used int   not null ,
  priority int not null  ,
  scan_timestamp timestamp  
);
```

## Example insert

```json
{"systems":[{"hostname":"deskctop1","domain":"example.com","uuid":"XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX","device_type":"Desktop","device_sn":"SNABC1234","model":"XXYYZZ","vendor":"HP","environment":"Production"}]}
```
