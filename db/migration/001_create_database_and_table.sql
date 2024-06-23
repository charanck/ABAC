CREATE TABLE IF NOT EXISTS policy(
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    name TEXT NOT NULL,
    action TEXT NOT NULL,
    rego_policy TEXT NOT NULL,
    description TEXT,
    updated DATE,
    deleted DATE,
    created DATE
);

CREATE TABLE IF NOT EXISTS policy_input(
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    policy_id TEXT NOT NULL,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    required BOOLEAN NOT NULL,
    UNIQUE(policy_id, name)
);

CREATE TABLE IF NOT EXISTS resource (
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    name TEXT NOT NULL,
    owner_id TEXT NOT NULL,
    policy_id TEXT NOT NULL,
    description TEXT,
    updated DATE,
    deleted DATE,
    created DATE,
    FOREIGN KEY(policy_id) REFERENCES policy(id)
);

CREATE TABLE IF NOT EXISTS user(
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    name TEXT NOT NULL,
    external_id TEXT,
    updated DATE,
    deleted DATE,
    created DATE
);

CREATE TABLE IF NOT EXISTS role(
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    name TEXT NOT NULL,
    updated DATE,
    deleted DATE,
    created DATE
);

CREATE TABLE IF NOT EXISTS role_user_mapping(
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    role_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    updated DATE,
    deleted DATE,
    created DATE,
    FOREIGN KEY(role_id) REFERENCES role(id),
    FOREIGN KEY(user_id) REFERENCES user(id),
    UNIQUE(role_id, user_id)
);

CREATE TABLE IF NOT EXISTS attribute(
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    name TEXT NOT NULL,
    string_value TEXT,
    integer_value INTEGER,
    float_value FLOAT,
    bool_value BOOLEAN,
    date_value DATE,
    type TEXT NOT NULL,
    updated DATE,
    deleted DATE,
    created DATE
);

CREATE TABLE IF NOT EXISTS role_attribute_mapping(
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    role_id TEXT NOT NULL,
    attribute_id TEXT NOT NULL,
    FOREIGN KEY(role_id) REFERENCES role(id),
    FOREIGN KEY(attribute_id) REFERENCES attribute(id)
);

CREATE TABLE IF NOT EXISTS resource_attribute_mapping(
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    resource_id TEXT NOT NULL,
    attribute_id TEXT NOT NULL,
    FOREIGN KEY(resource_id) REFERENCES resource(id),
    FOREIGN KEY(attribute_id) REFERENCES attribute(id)
);

CREATE TABLE IF NOT EXISTS user_attribute_mapping(
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    user_id TEXT NOT NULL,
    attribute_id TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES user(id),
    FOREIGN KEY(attribute_id) REFERENCES attribute(id)
);



