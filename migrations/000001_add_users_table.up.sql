CREATE TABLE "users" (
	"id" integer not null primary key autoincrement,
	"email" varchar not null,
	"password" varchar not null,
	"created_at" datetime null,
	"active" boolean not null check (active in (0, 1)) default 0
);
CREATE UNIQUE INDEX users_email_idx ON users(email);