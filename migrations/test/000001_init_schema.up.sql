CREATE TABLE IF NOT EXISTS owners (
    owner_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    removed BOOLEAN NOT NULL DEFAULT false,
    username TEXT NOT NULL,
    folder TEXT NOT NULL,
    UNIQUE (username, folder)
);

CREATE TABLE IF NOT EXISTS files (
    file_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    removed BOOLEAN NOT NULL DEFAULT false,
    state INT NOT NULL DEFAULT 100,
    virtual_name TEXT DEFAULT '',
    hash_sum TEXT DEFAULT '',
    file_name TEXT NOT NULL,
    size_file INT DEFAULT 0,
    owner_id INT NOT NULL,
    mod_time INT NOT NULL,
    client TEXT NOT NULL,
    UNIQUE (file_name, owner_id),
    FOREIGN KEY (owner_id) REFERENCES owners (owner_id)
);