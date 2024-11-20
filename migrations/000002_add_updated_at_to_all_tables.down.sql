-- Remove updated_at from Users Table
ALTER TABLE users DROP COLUMN updated_at;

-- Remove updated_at from Movies Table
ALTER TABLE movies DROP COLUMN updated_at;

-- Remove updated_at from Votes Table
ALTER TABLE votes DROP COLUMN updated_at;

-- Remove updated_at from Viewership Table
ALTER TABLE viewership DROP COLUMN updated_at;