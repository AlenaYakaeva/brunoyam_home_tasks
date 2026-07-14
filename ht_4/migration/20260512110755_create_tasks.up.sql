CREATE TYPE task_status AS ENUM ('0', '1', '2');

CREATE TABLE tasks(
    tid      UUID PRIMARY KEY  DEFAULT gen_random_uuid(),
    uid      UUID NOT NULL ,
    title    TEXT NOT NULL,
    description TEXT,
    status   task_status DEFAULT '0' NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted boolean DEFAULT false,

    CONSTRAINT fk_user FOREIGN KEY (uid) REFERENCES users(uid) ON DELETE CASCADE
)