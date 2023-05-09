CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS accounts (
	account_number TEXT PRIMARY KEY,
	user_id INTEGER NOT NULL,
    balance REAL NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS transfers (
    id INTEGER PRIMARY KEY,
    from_account TEXT NOT NULL,
    to_account TEXT NOT NULL,
    amount REAL NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (from_account) REFERENCES accounts(account_number),
    FOREIGN KEY (to_account) REFERENCES accounts(account_number)
);

CREATE INDEX users_name_idx ON users (name);

CREATE INDEX accounts_user_idx ON accounts (user_id);

CREATE INDEX transfers_from_account_idx ON transfers(from_account);

CREATE INDEX transfers_to_account_idx ON transfers(to_account);

CREATE INDEX transfers_from_to_idx ON transfers(from_account, to_account);

INSERT INTO users (name) VALUES ('Lise Ellison');

INSERT INTO users (name) VALUES ('Rosabel Hunnicutt ');

INSERT INTO users (name) VALUES ('Rick Beasley');

INSERT INTO users (name) VALUES ('Scott Evered');

