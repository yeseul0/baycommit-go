CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    github_email VARCHAR NOT NULL,
    wallet_address VARCHAR NOT NULL
);

CREATE TABLE studies (
    id SERIAL PRIMARY KEY,
    proxy_address VARCHAR NOT NULL,
    study_start_time BIGINT NOT NULL,
    study_end_time BIGINT NOT NULL
);

CREATE TABLE user_studies (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    study_id INT NOT NULL REFERENCES studies(id),
    repo_url VARCHAR NOT NULL
);

CREATE TABLE study_sessions (
    id SERIAL PRIMARY KEY,
    study_id INT NOT NULL REFERENCES studies(id),
    study_date VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    started_at BIGINT,
    closed_at BIGINT,
    blockchain_tx_hash VARCHAR,
    study_midnight_utc BIGINT NOT NULL
);

CREATE TABLE commit_records (
    id SERIAL PRIMARY KEY,
    study_session_id INT NOT NULL REFERENCES study_sessions(id),
    user_id INT NOT NULL REFERENCES users(id),
    commit_timestamp BIGINT NOT NULL,
    commit_id VARCHAR NOT NULL,
    commit_message VARCHAR
);

-- 테스트 데이터
INSERT INTO users (github_email, wallet_address) VALUES
    ('test@example.com', '0x1234567890abcdef1234567890abcdef12345678');

INSERT INTO studies (proxy_address, study_start_time, study_end_time) VALUES
    ('0xabcdef1234567890abcdef1234567890abcdef12', 0, 86400);

INSERT INTO user_studies (user_id, study_id, repo_url) VALUES
    (1, 1, 'https://github.com/yeseul0/rust-study');
