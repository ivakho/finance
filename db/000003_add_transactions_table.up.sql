create table
    if not exists transactions (
        id serial primary key,
        category_id integer not null references category (id) on delete restrict,
        type text not null,
        amount integer not null,
        created_at timestamp not null,
        updated_at timestamp
    );