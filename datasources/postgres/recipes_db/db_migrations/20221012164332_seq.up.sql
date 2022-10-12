create table if not exists recipes (
   id   serial  primary key     not null,
   title        varchar(255)    not null,
   description  text,
   cooking_time text            not null,
   ingredients  text[]          not null default '{}'::text[],
   instructions text            not null,
   date_created timestamptz     not null default now()
);