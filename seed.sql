-- Create the 'entries' table
CREATE TABLE entries (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Insert data into the 'entries' table
INSERT INTO entries (id, name, created_at)
VALUES
    ('id1', 'entry_name1', '2024-01-16 00:00:00')
    ,('id2', 'entry_name2', '2024-01-16 00:00:00')
    ,('id3', 'entry_name3', '2024-01-16 00:00:00')
;
