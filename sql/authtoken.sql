create table authtokens (
    token_id bigint UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    token varchar(255),
    user_id BIGINT UNSIGNED,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(user_id)
);