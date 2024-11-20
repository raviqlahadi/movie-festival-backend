-- Revert the user_id column in the viewership table to NOT NULL
ALTER TABLE viewership
ALTER COLUMN user_id SET NOT NULL;