#!/bin/bash
sleep 5

export PGPASSWORD=1234

echo "PostgreSQL is ready. Running the init.sql script..."
psql -h db -U user_db -d ozon -f /root/bash/init.sql