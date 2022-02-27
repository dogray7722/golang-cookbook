create table if not exists recipes (
   id   serial  primary key     not null,
   title        varchar(255)    not null,
   description  text,
   cooking_time text            not null,
   ingredients  varchar(255)[]  not null,
   instructions text            not null,
   date_created timestamptz     not null default now()
);