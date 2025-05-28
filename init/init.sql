-- init.sql
DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT FROM pg_database WHERE datname = 'tasker'
        ) THEN
            CREATE DATABASE tasker;
        END IF;
    END
$$;