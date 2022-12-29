CREATE TABLE IF NOT EXISTS users (
    user_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

    username TEXT NOT NULL UNIQUE,
    removed BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS folders (
    folder_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    
    folder TEXT NOT NULL,
    removed BOOLEAN NOT NULL,

    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE IF NOT EXISTS files (
    file_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

    client TEXT NOT NULL,
    removed BOOLEAN NOT NULL,
    file_name TEXT NOT NULL,
    mod_time INT NOT NULL,
    virtual_name TEXT DEFAULT '',
    size_file INT DEFAULT 0,
    hash_sum TEXT DEFAULT '', 
    state INT NOT NULL DEFAULT 1,

    folder_id INT NOT NULL,
    FOREIGN KEY (folder_id) REFERENCES folders (folder_id),
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

-- Value 1

INSERT INTO
    users (username, removed)
VALUES
    ('u1', false);
INSERT INTO
    folders (folder, removed, user_id)
VALUES
    ('f1', false, 1);

INSERT INTO
    files (
        client,
        file_name,
        mod_time,
        size_file,
        hash_sum,
        removed,
        folder_id,
        user_id,
        state
    )
VALUES
    (
        'c1',
        'C:\1\abc.exe',
        12345,
        2048,
        'ghy67678',
        false,
        1,
        1,
        1
    );

-- Value 2

INSERT INTO
    users (username, removed)
VALUES
    ('u2', false);
INSERT INTO
    folders (folder, removed, user_id)
VALUES
    ('f2', false, 2);

INSERT INTO
    files (
        client,
        file_name,
        mod_time,
        size_file,
        hash_sum,
        virtual_name,
        removed,
        folder_id,
        user_id,
        state
    )
VALUES
    (
        'c1',
        'path',
        12345,
        2048,
        'ghy67678',
        'virtualname1',
        false,
        2,
        2,
        1
    );
