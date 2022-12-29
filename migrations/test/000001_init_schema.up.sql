CREATE TABLE IF NOT EXISTS users (
    user_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username TEXT NOT NULL UNIQUE,
    removed BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS folders (
    folder_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    folder TEXT NOT NULL,
    removed BOOLEAN NOT NULL DEFAULT false,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE IF NOT EXISTS files (
    file_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    client TEXT NOT NULL,
    removed BOOLEAN NOT NULL DEFAULT false,
    file_name TEXT NOT NULL,
    mod_time INT NOT NULL,
    virtual_name TEXT DEFAULT '',
    size_file INT DEFAULT 0,
    hash_sum TEXT DEFAULT '',
    state INT NOT NULL DEFAULT 100,
    folder_id INT NOT NULL,
    FOREIGN KEY (folder_id) REFERENCES folders (folder_id)
);

-- -- Username 1
INSERT INTO
    users (username, removed)
VALUES
    ('Albert_Robinson', false);

-- Folder 1
INSERT INTO
    folders (folder, removed, user_id)
VALUES
    ('folder1', false, 1);

INSERT INTO
    files (
        client,
        file_name,
        mod_time,
        size_file,
        hash_sum,
        removed,
        folder_id,
        virtual_name,
        state
    )
VALUES
    (
        'IPhone',
        'Photos/family.jpg',
        1672320628,
        3675,
        'd05958620e57c91568afe45f8dc90269',
        false,
        1,
        '',
        100
    ),
    (
        'IPhone',
        'Photos/schooL/high/photo1.jpg',
        1672320627,
        45668,
        'd05958sdfgt20e57c91568afehdc90266',
        false,
        1,
        'AA235g5ea-22',
        200
    ),
    (
        'MAC',
        'Photos/schooL/high/photo2.jpg',
        1672320626,
        4566867,
        'ftgrh6aaaaa0e57c91568afehdc90266',
        false,
        1,
        'Ghyiii-4__',
        300
    ),
    (
        'IPhone',
        'Photos/schooL/high/photo3.jpg',
        1672320625,
        123457,
        'ftgrh6aaaеониfehdc90266',
        true,
        1,
        'ADIGU$$$',
        300
    ),
    (
        'IPhone',
        'Photos/games/nfs.exe',
        1672320624,
        343665756,
        '4tg4yng48uh6657h65967bh7',
        false,
        1,
        'K^HH%%',
        300
    );

INSERT INTO
    files (
        client,
        file_name,
        mod_time,
        removed,
        folder_id,
        state
    )
VALUES
    (
        'IPhone',
        'Photos/',
        1672320623,
        false,
        1,
        300
    ),
    (
        'IPhone',
        'Photos/schooL/',
        1672320622,
        false,
        1,
        300
    ),
    (
        'MAC',
        'Photos/schooL/high/',
        1672320621,
        false,
        1,
        300
    ),
    (
        'IPhone',
        'Photos/schooL/middle/',
        1672320620,
        true,
        1,
        300
    ),
    (
        'MAC',
        'Photos/games/',
        1672320619,
        false,
        1,
        300
    );

-- -- Username 2
INSERT INTO
    users (username, removed)
VALUES
    ('Joseph_Johnson', false);

-- Folder 1
INSERT INTO
    folders (folder, removed, user_id)
VALUES
    ('myFolder', false, 2);

INSERT INTO
    files (
        client,
        file_name,
        mod_time,
        size_file,
        hash_sum,
        removed,
        folder_id,
        virtual_name,
        state
    )
VALUES
    (
        'Xiaomi',
        'notes/1.txt',
        1672320618,
        54565675,
        'd05958620e57c91568afe45f8dc90269',
        false,
        2,
        'VITU::::',
        200
    ),
    (
        'Windows',
        'myLife/notes/1.txt',
        1672320617,
        45668,
        'd05958sdfgt20e57c91568afehdc90266',
        false,
        2,
        'HEGUUUPE**__E',
        300
    );

INSERT INTO
    files (
        client,
        file_name,
        mod_time,
        removed,
        folder_id,
        state
    )
VALUES
    (
        'Win',
        'notes/',
        1672320616,
        false,
        1,
        300
    ),
    (
        'MAC',
        'myLife/',
        1672320615,
        false,
        1,
        300
    ),
    (
        'IPhone',
        'Photos/schooL/middle/',
        1672320614,
        true,
        1,
        300
    ),
    (
        'MAC',
        'myLife/notes/',
        1672320613,
        false,
        1,
        300
    );