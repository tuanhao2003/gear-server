CREATE TABLE
    users (
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        hash_password TEXT NOT NULL,
        is_active BOOLEAN DEFAULT TRUE,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
    );

CREATE TABLE
    profiles (
        id TEXT PRIMARY KEY,
        full_name TEXT NOT NULL,
        phone_number TEXT NOT NULL UNIQUE,
        address TEXT DEFAULT NULL,
        avatar_url TEXT DEFAULT NULL,
        user_id TEXT NOT NULL UNIQUE,
        is_active BOOLEAN DEFAULT TRUE,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW (),
        CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
    );