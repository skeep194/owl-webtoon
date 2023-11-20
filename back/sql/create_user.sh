#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER owl PASSWORD '$OWL_DB_PASSWORD';
	CREATE DATABASE owl_webtoon;
	ALTER DATABASE owl_webtoon OWNER TO owl;
	GRANT ALL PRIVILEGES ON DATABASE owl_webtoon TO owl;
EOSQL