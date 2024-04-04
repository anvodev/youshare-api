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

DROP TABLE IF EXISTS "schema_migrations";
CREATE TABLE "public"."schema_migrations" (
    "version" bigint NOT NULL,
    "dirty" boolean NOT NULL,
    CONSTRAINT "schema_migrations_pkey" PRIMARY KEY ("version")
) WITH (oids = false);

INSERT INTO "schema_migrations" ("version", "dirty") VALUES
(4,	'f');