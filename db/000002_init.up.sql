create table if not exists category 
(
    id serial primary key,
    name text not null,
    created_at timestamp not null,
    updated_at timestamp
)