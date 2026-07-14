CREATE TABLE users(
    
	uid      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name     TEXT NOT NULL,
	email    TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted boolean DEFAULT false
)