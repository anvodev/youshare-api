CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL
);
CREATE TABLE IF NOT EXISTS videos (
    id bigserial PRIMARY KEY,
    url text NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
ALTER TABLE videos
ADD COLUMN author_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE videos
ADD CONSTRAINT fk_videos_users FOREIGN KEY (author_id) REFERENCES users (id);

CREATE TABLE IF NOT EXISTS tokens (
    hash bytea PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    expiry timestamp(0) with time zone NOT NULL
);

-- Seed data
INSERT INTO users (name, email, password_hash) VALUES
('Alice', 'alice_test@example.com', 'pa55word');

INSERT INTO videos (url, title, description, author_id) VALUES
('https://www.youtube.com/watch?v=KLuTLF3x9sA', 'Norway 4K', 'Video description', 1);
