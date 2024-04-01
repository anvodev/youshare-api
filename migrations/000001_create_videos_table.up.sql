CREATE TABLE IF NOT EXISTS videos (
    id bigserial PRIMARY KEY,
    url text NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
