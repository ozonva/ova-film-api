-- +goose Up
-- +goose StatementBegin
CREATE table movies (
    id serial primary key,
    user_id bigint,
    title text,
    year text
);

INSERT INTO movies (user_id, title, year)
VALUES
    (2, 'Titanic', '1997'),
    (2, 'Dirty Dances', '1989'),
    (3, 'Avatar', '2009');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table movies;
-- +goose StatementEnd
