#!/bin/bash

# set -e
psql -U "$POSTGRES_USER" <<-EOSQL
	CREATE USER docker SUPERUSER;
	CREATE DATABASE docker;
    ALTER USER docker WITH PASSWORD 'docker';
	GRANT ALL ON schema public TO docker;
EOSQL

psql -U "docker" <<-EOSQL
create table batches (
    batch_id int generated always as identity not null,
    username    text unique not null,
    folder      text unique not null,
    client      text unique not null,
    path        text not null,
    hash        text not null,
    mod_time    int  not null,
    part        int  not null,
    count_parts int  not null,
    part_size   int  not null,
    byte_offset int  not null,
    size_file   int  not null,
    constraint PK_batches_batch_id primary key(batch_id)
);

INSERT INTO batches (
		username, folder, client, path, hash, mod_time, 
		part, count_parts, part_size, byte_offset, size_file 
	) 
VALUES ('username', 'folder', 'client', 'path', 'hash', 12345, 
	   1, 2, 3, 4, 5)
RETURNING batch_id;
EOSQL