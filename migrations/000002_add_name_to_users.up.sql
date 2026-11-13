ALTER TABLE users
    ADD COLUMN name TEXT;

UPDATE users
SET name = split_part(email, '@', 1);

ALTER TABLE users
    ALTER COLUMN name SET NOT NULL;
