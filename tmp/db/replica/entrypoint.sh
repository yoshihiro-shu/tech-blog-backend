#!/bin/bash

set -e

echo "psql -h $DB_PRIMARY_HOST -U $DB_PRIMARY_USER -d $DB_PRIMARY_NAME -p $DB_PRIMARY_PORT"


while ! psql -h $DB_PRIMARY_HOST -U $DB_PRIMARY_USER -d $DB_PRIMARY_NAME -p $DB_PRIMARY_PORT -c "select 'it is running';" 2>&1 ; do \
	sleep 1s ; \
done

# load backup from primary instance
pg_basebackup -h primary -p 5432 -D $PGDATA -S replication_slot_slave1 --progress -X stream -U replicator -Fp -R || :

# start postgres
bash /usr/local/bin/docker-entrypoint.sh -c 'config_file=/etc/postgresql/postgresql.conf' -c 'hba_file=/etc/postgresql/pg_hba.conf'
