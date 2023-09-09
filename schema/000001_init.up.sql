CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id       serial                                           not null unique,
    user_id  int references users (id) on delete cascade      not null,
    lists_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255) not null unique,
    done        boolean      not null defalut false
);

CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255) not null unique
);

CREATE TABLE lists_items
(
    id      serial not null unique,
    user_id int reference users(id) on delete cascade not null,
    list_id int reference todo_lists(id) on delete cascade not null
);

