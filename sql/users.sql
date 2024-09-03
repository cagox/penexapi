create table users (
    user_id bigint UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    email varchar(255),
    password_hash varchar(255),
    is_admin bool,
    is_verified bool,
    role text
);