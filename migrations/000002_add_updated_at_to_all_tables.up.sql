
-- Add updated_at to Users Table
ALTER TABLE users ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Add updated_at to Movies Table
ALTER TABLE movies ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Add updated_at to Votes Table
ALTER TABLE votes ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Add updated_at to Viewership Table
ALTER TABLE viewership ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;