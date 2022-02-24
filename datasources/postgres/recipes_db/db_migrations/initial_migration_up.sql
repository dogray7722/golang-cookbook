create table if not exists recipes (
   id   serial  primary key     not null,
   name         varchar(255)    not null,
   description  text,
   instructions text            not null,
   status       varchar(128)    not null,
   date_created timestamptz     not null default now()
);

create table if not exists ingredients (
    id  serial      primary key     not null,
    serving_size    varchar(255)    not null,
    item            varchar(255)    not null,
    date_created    timestamptz     not null default now()
);

create table if not exists recipes_to_ingredients (
    id  serial      primary key not null,
    recipe_id       int not null,
    ingredient_id   int not null
)