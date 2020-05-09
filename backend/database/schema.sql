CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE files (
    created_at TIMESTAMP DEFAULT NOW(),
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    extension VARCHAR(30) NOT NULL,
    originalname VARCHAR(300) NOT NULL,
    storedname VARCHAR(300) NOT NULL,
    cookie uuid NOT NULL
);

CREATE TABLE links (
    created_at TIMESTAMP DEFAULT NOW(),
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    foldername TEXT NOT NULL,
    accesspassword TEXT NOT NULL,
    expirationdate TIMESTAMP NOT NULL,
);

CREATE TABLE file_link (
    linkid uuid NOT NULL,
    fileid uuid NOT NULL,
    PRIMARY KEY (linkid, fileid),
    FOREIGN KEY (linkid) REFERENCES links(id) ON DELETE CASCADE,
    FOREIGN KEY (fileid) REFERENCES files(id)
);