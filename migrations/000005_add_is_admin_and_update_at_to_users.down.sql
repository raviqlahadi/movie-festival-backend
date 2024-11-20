a-- Remove is_admin and updated_at columns from users table
ALTER TABLE users
DROP COLUMN is_admin;