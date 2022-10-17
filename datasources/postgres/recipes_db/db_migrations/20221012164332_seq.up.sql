CREATE TABLE IF NOT EXISTS "recipes" (
   "id"   serial  primary key     NOT NULL,
   "title"        varchar(255)    NOT NULL,
   "description"  text,
   "cooking_time" text            NOT NULL,
   "ingredients"  text[]          NOT NULL DEFAULT '{}'::text[],
   "instructions" text            NOT NULL,
   "date_created" timestamptz     NOT NULL DEFAULT (now())
);