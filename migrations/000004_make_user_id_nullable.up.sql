
-- Alter the user_id column in the viewership table to allow NULL values
ALTER TABLE viewership
ALTER COLUMN user_id DROP NOT NULL;

-- Drop the foreign key constraint on user_id
ALTER TABLE viewership DROP CONSTRAINT IF EXISTS viewership_user_id_fkey;