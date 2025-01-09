CREATE TABLE IF NOT EXISTS user_profile(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nickname TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    gender TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    age INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS session(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    token UNIQUE NOT NULL,
    expire_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user_profile(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS message(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sender INTEGER NOT NULL,
    receiver INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender) REFERENCES user_profile(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver) REFERENCES user_profile(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS post(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user_profile(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS comment(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user_profile(id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS post_react(
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    type INTEGER NOT NULL,
    PRIMARY KEY (post_id, user_id),
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user_profile(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS comment_react(
    comment_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    type INTEGER NOT NULL,
    PRIMARY KEY (comment_id, user_id),
    FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user_profile(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS category(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    category_name TEXT UNIQUE NOT NULL
);
CREATE TABLE IF NOT EXISTS post_category(
    post_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    PRIMARY KEY(post_id, category_id),
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS token_bucket(
    ip TEXT NOT NULL,
    route TEXT NOT NULL,
    tokens INTEGER NOT NULL,
    maxTokens INTEGER NOT NULL,
    refill INTEGER NOT NULL,
    lastRefill DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (ip, route)
);
INSERT
    OR IGNORE INTO category (category_name)
VALUES ('tech'),
    ('sports'),
    ('games')