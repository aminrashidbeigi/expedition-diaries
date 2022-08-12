CREATE TABLE countries (
    id serial PRIMARY KEY,
    code varchar(2) NOT NULL UNIQUE,
    name text NOT NULL UNIQUE
);

CREATE TABLE resources (  
    id serial PRIMARY KEY,
    title text NOT NULL UNIQUE,
    link text NOT NULL,
    image text NOT NULL
);

CREATE TABLE travelers (
    id serial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    link text NOT NULL
);

CREATE TABLE travels (
    id serial PRIMARY KEY,
    title text NOT NULL UNIQUE,
    started_at varchar(4) NOT NULL,
    ended_at varchar(4) NOT NULL
);

CREATE TABLE travel_countries (
    travel_id int REFERENCES travels(id) ON DELETE CASCADE,
    country_id int REFERENCES countries(id) ON DELETE CASCADE,
    PRIMARY KEY(travel_id, country_id)
);

CREATE TABLE travel_travelers (
    travel_id int REFERENCES travels(id) ON DELETE CASCADE,
    traveler_id int REFERENCES travelers(id) ON DELETE CASCADE,
    PRIMARY KEY(travel_id, traveler_id)
);

CREATE TABLE travel_resources (
    travel_id int REFERENCES travels(id) ON DELETE CASCADE,
    resource_id int REFERENCES resources(id) ON DELETE CASCADE,
    PRIMARY KEY(travel_id, resource_id)
);