CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar NOT NULL,
  "created_at" timestamp,
  "admin" bool
);

CREATE TABLE "characters" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "class_id" integer,
  "strength" integer NOT NULL,
  "dexterity" integer NOT NULL,
  "constitution" integer NOT NULL,
  "intellect" integer NOT NULL,
  "wisdom" integer NOT NULL,
  "charisma" integer NOT NULL,
  "health" integer NOT NULL,
  "experience" integer NOT NULL,
  "level" integer NOT NULL,
  "pos_x" integer,
  "pos_y" integer
);

CREATE TABLE "classes" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "hit_dice" integer NOT NULL,
  "hit_sides" integer NOT NULL
);

CREATE TABLE "weapons" (
  "id" integer PRIMARY KEY,
  "character_id" integer,
  "title" varchar,
  "atk_dice" integer NOT NULL,
  "atk_sides" integer NOT NULL
);

CREATE TABLE "monsters" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "hit_dice" integer NOT NULL,
  "hit_sides" integer NOT NULL,
  "atk_dice" integer NOT NULL,
  "atk_sides" integer NOT NULL
);

CREATE TABLE "rooms" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "description" varchar,
  "pos_x" integer,
  "pos_y" integer
);

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "characters" ("user_id");

ALTER TABLE "classes" ADD FOREIGN KEY ("id") REFERENCES "characters" ("class_id");

ALTER TABLE "characters" ADD FOREIGN KEY ("id") REFERENCES "weapons" ("character_id");
