CREATE TABLE IF NOT EXISTS owners (
    owner_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    removed BOOLEAN NOT NULL DEFAULT false,
    username TEXT NOT NULL UNIQUE,
    folder TEXT NOT NULL,
    UNIQUE (username, folder)
);

CREATE TABLE IF NOT EXISTS files (
    file_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    removed BOOLEAN NOT NULL DEFAULT false,
    virtual_name TEXT UNIQUE DEFAULT '',
    state INT NOT NULL DEFAULT 100,
    hash_sum TEXT DEFAULT '', 
    file_name TEXT NOT NULL, 
    size_file INT DEFAULT 0, 
    owner_id INT NOT NULL, 
    mod_time INT NOT NULL, 
    client TEXT NOT NULL, 
    UNIQUE (hash_sum, file_name, owner_id, virtual_name),
    FOREIGN KEY (owner_id) REFERENCES owners (owner_id)
);