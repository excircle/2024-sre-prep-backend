CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    domain VARCHAR(75) NOT NULL,
    description TEXT
);

CREATE INDEX idx_category_domain ON category(domain);

CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    description TEXT,
    category_id INT NOT NULL,
    FOREIGN KEY (category_id)
        REFERENCES category(id)
        ON DELETE CASCADE
);
