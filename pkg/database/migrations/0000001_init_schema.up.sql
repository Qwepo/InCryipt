CREATE TYPE type_message AS ENUM (
    'photo',
    'string'
);
CREATE TYPE type_chat AS ENUM (
    'group',
    'private'
);

CREATE TABLE users
(
    id bigserial PRIMARY KEY,
    username varchar(100) NOT NULL,
    email varchar(50) NOT NULL,
    password varchar(150) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE sessions 
(
    user_id bigint NOT NULL,
    token varchar(255) NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE chats (

    id bigserial PRIMARY KEY,
    name varchar NOT NULL,
    type type_chat NOT NULL,
    created_at timestamp NOT NULL
);
CREATE TABLE chats_users 
(
    id bigserial PRIMARY KEY,
    chat_id bigint NOT NULL, 
    user_id bigint NOT NULL,

    FOREIGN KEY (chat_id) REFERENCES chats (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);


CREATE TABLE message 
(
    id bigserial PRIMARY KEY,
    chat_id bigint NOT NULL,
    sender bigint NOT NULL,
    type type_message NOT NULL, 
    content_text text NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,

    FOREIGN KEY (chat_id) REFERENCES chats (id)
);

CREATE TABLE wallet 
(
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    public_key varchar NOT NULL,
    private_key varchar NOT NULL,
    balance bigint NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id)
);