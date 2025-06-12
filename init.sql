CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);

INSERT INTO users (id, name, email, password) VALUES (
    '1',
    'admin',
    'admin@example.com',
    '$2a$10$eXMnTxTVesfC0rUIPR6yxuRvJ2hjWX32BNr9cJl0l7RL1i5/UG5MO'  

);

