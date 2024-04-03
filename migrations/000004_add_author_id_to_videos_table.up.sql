ALTER TABLE videos
ADD COLUMN author_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE videos
ADD CONSTRAINT fk_videos_users FOREIGN KEY (author_id) REFERENCES users (id);
