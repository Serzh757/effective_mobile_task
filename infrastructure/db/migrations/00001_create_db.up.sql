SELECT 'CREATE DATABASE musicdb'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'musicdb');