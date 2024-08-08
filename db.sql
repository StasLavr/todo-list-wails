CREATE TABLE "users" (
	"id"	INTEGER NOT NULL,
	"name"	INTEGER NOT NULL,
	"password"	INTEGER,
	PRIMARY KEY("id")
);

CREATE TABLE "tasks" (
	"id"	INTEGER NOT NULL,
	"id_user"	INTEGER NOT NULL,
	"name"	TEXT,
	"date"	TEXT,
	"date_create"	TEXT,
	"done_task"	INTEGER,
	PRIMARY KEY("id")
)
