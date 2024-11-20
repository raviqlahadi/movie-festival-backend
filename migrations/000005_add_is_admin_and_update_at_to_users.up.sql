-- Add is_admin and updated_at columns to users table
ALTER TABLE users
ADD COLUMN is_admin BOOLEAN DEFAULT FALSE;