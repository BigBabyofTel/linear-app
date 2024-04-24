CREATE TABLE IF NOT EXISTS issues(
    id bigserial PRIMARY KEY,
    created_at timestamp(0) NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    description text,
    status text NOT NULL DEFAULT 'backlog',
    priority text,
    due_date date,
    version integer NOT NULL DEFAULT 1 
);

CREATE TABLE IF NOT EXISTS workspaces(
    id bigserial PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    image text,
    slug text NOT NULL UNIQUE,
    version integer NOT NULL DEFAULT 1
);



CREATE INDEX idx_slug on workspaces(slug);

CREATE TABLE IF NOT EXISTS permissions(
    id bigserial PRIMARY KEY,
    code text NOT NULL
);


CREATE TABLE IF NOT EXISTS user_workspace_issues(
    creator_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    issue_id bigint NOT NULL REFERENCES issues(id) ON DELETE CASCADE,
    user_id bigint REFERENCES users(id) ON DELETE SET NULL,
    workspace_id bigint REFERENCES workspaces(id) ON DELETE SET NULL,
    PRIMARY KEY (issue_id, creator_id),
    UNIQUE(issue_id, user_id, workspace_id, creator_id)
);

CREATE TABLE IF NOT EXISTS requests(
    user_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    workspace_id bigint NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE
);

INSERT INTO permissions (code)
VALUES ('workspace:admin'), ('workspace:moderator'), ('workspace:member');


CREATE TABLE IF NOT EXISTS notes(
    id bigserial PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT NOW(),
    content text NOT NULL,
    creator_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS issue_comments(
    id bigserial PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT NOW(),
    content text NOT NULL,
    creator_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    issue_id bigint NOT NULL REFERENCES issues(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comment_reactions(
    id bigserial PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT NOW(),
    reaction text NOT NULL,
    creator_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    comment_id bigint NOT NULL REFERENCES issue_comments(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_workspace_permissions(
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    workspace_id bigint NOT NULL REFERENCES workspaces ON DELETE CASCADE,
    permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE,
    PRIMARY KEY (workspace_id, user_id, permission_id)
);