CREATE TABLE users (
    id TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE todo (
    id TEXT PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL
    updated_at TIMESTAMP NOT NULL,
    todonote TEXT NOT NULL,
    user_id TEXT REFERENCES users(id)
);