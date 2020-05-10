CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS links CASCADE;
DROP TABLE IF EXISTS files CASCADE;

CREATE TABLE links (
    created_at TIMESTAMP DEFAULT NOW(),
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    folderpath TEXT NOT NULL,
    cookie uuid NOT NULL,
    accesspassword TEXT,
    expirationdate TIMESTAMP NOT NULL   
);

CREATE TABLE files (
    created_at TIMESTAMP DEFAULT NOW(),
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    originalname VARCHAR(300) NOT NULL,
    storedname VARCHAR(300) NOT NULL,
    cookie uuid NOT NULL,
    linkid uuid references links(id)
);