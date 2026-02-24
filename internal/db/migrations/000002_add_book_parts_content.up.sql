-- Add content column to book_parts table
ALTER TABLE book_parts ADD COLUMN content TEXT NOT NULL DEFAULT '';
