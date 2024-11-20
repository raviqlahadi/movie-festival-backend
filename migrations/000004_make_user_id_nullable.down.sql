-- Alter the user_id column in the viewership table to allow NULL values
ALTER TABLE viewership
ALTER COLUMN user_id DROP NOT NULL;