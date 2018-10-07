create user solar_user with password 'solar_pass';

create database solar_db;

grant all privileges on database solar_db to solar_user;

CREATE TABLE public.vacancy
(
  id         serial PRIMARY KEY NOT NULL,
  name       text               NOT NULL,
  salary     int                NOT NULL,
  experience text               NOT NULL,
  place      text               NOT NULL
);
CREATE UNIQUE INDEX vacancy_id_uindex
  ON public.vacancy (id);


create function get_all()
  returns json
language plpgsql
as $$
begin

  return (select coalesce(to_json(array_agg(t)), '[]') from (select * from vacancy order by name) t);

end;

$$;

alter function get_all()
  owner to solar_user;