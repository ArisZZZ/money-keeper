-- 给 users 表添加 phone 字段
ALTER TABLE users
ADD COLUMN phone VARCHAR(255) NOT NULL UNIQUE;