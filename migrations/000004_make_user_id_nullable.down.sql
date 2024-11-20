-- Revert the user_id column in the viewership table to NOT NULL
ALTER TABLE viewership
ALTER COLUMN user_id SET NOT NULL;

-- Re-add the foreign key constraint on user_id
ALTER TABLE viewership
ADD CONSTRAINT viewership_user_id_fkey
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;