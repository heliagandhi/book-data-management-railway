-- +migrate Up

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP DEFAULT NULL,
    modified_by VARCHAR(50)
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP DEFAULT NULL,
    modified_by VARCHAR(50)
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    description TEXT,
    image_url VARCHAR(255),
    release_year INT CHECK (release_year > 0),
    price DECIMAL(12, 2) NOT NULL DEFAULT 0,
    total_page INT CHECK (total_page > 0),
    thickness VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP NULL,
    modified_by VARCHAR(50),
    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE user_sessions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    token TEXT NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_sessions_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down

DROP TABLE user_sessions;
DROP TABLE books;
DROP TABLE categories;
DROP TABLE users;