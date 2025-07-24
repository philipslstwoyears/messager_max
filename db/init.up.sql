CREATE TABLE IF NOT EXISTS users(
                                    id SERIAL PRIMARY KEY ,
                                    login TEXT NOT NULL UNIQUE,
                                    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS messages(
                                       id SERIAL PRIMARY KEY ,
                                       sender TEXT ,
                                       receiver TEXT,
                                       content TEXT NOT NULL,
                                       created_at     timestamptz     default timezone('Europe/Moscow'::text, now())
);