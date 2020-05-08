CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE files (
    created_at TIMESTAMP NOT NULL,
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    originalname VARCHAR(300) NOT NULL,
    storedname VARCHAR(300) NOT NULL,
    cookiename uuid NOT NULL
);

CREATE TABLE links (
    created_at TIMESTAMP NOT NULL,
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    foldername TEXT NOT NULL,
    accesspassword TEXT NOT NULL,
    expirationdate TIMESTAMP NOT NULL,
    fileid uuid NOT NULL,
    FOREIGN KEY (fileid) REFERENCES files (id) ON DELETE CASCADE
);