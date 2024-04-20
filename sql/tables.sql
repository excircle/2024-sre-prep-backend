CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    category VARCHAR(75) NOT NULL,
    description TEXT
);

CREATE INDEX idx_categories ON categories(category);

CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    description TEXT,
    category_id INT NOT NULL,
    FOREIGN KEY (category_id)
        REFERENCES categories(id)
        ON DELETE CASCADE
);
