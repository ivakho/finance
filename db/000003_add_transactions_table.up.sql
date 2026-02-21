create table if not exists transaction 
(
    id serial primary key,
    category_id integer not null references category(id) on delete restrict,
    amount float not null check (amount > 0),
    created_at timestamp not null,
    updated_at timestamp
);
