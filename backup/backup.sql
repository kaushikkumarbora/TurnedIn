CREATE SEQUENCE por_por_id_seq START 1;
CREATE SEQUENCE user_user_id_seq START 1;

-- Table: public.por

-- DROP TABLE public.por;

CREATE TABLE public.por
(
    por_id bigint NOT NULL DEFAULT nextval('por_por_id_seq'::regclass),
    name text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT por_pkey PRIMARY KEY (por_id)
)

TABLESPACE pg_default;

ALTER TABLE public.por
    OWNER to postgres;

-- Table: public.user

-- DROP TABLE public."user";

CREATE TABLE public."user"
(
    user_id bigint NOT NULL DEFAULT nextval('user_user_id_seq'::regclass),
    first_name text COLLATE pg_catalog."default",
    last_name text COLLATE pg_catalog."default",
    dob date,
    email_id text COLLATE pg_catalog."default",
    contact_no text COLLATE pg_catalog."default",
    skills text COLLATE pg_catalog."default",
    year_of_admission text COLLATE pg_catalog."default",
    year_of_completion text COLLATE pg_catalog."default",
    semester smallint,
    bio text COLLATE pg_catalog."default",
    por_id bigint,
    education text COLLATE pg_catalog."default",
    department text COLLATE pg_catalog."default",
    course text COLLATE pg_catalog."default",
    hostel text COLLATE pg_catalog."default",
    resume bytea,
    status boolean,
    username text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    internship text COLLATE pg_catalog."default",
    CONSTRAINT user_pkey PRIMARY KEY (user_id),
    CONSTRAINT user_username_key UNIQUE (username),
    CONSTRAINT por_id_fkey FOREIGN KEY (por_id)
        REFERENCES public.por (por_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE public."user"
    OWNER to postgres;

-- Table: public.connection

-- DROP TABLE public.connection;

CREATE TABLE public.connection
(
    l_user_id bigint NOT NULL,
    r_user_id bigint NOT NULL,
    status character(1) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT l_user_id FOREIGN KEY (l_user_id)
        REFERENCES public."user" (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT r_user_id FOREIGN KEY (r_user_id)
        REFERENCES public."user" (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.connection
    OWNER to postgres;

INSERT INTO por (name) VALUES ('Charaideo'), ('Dhemaji'), ('Dibrugarh'), ('Golaghat'), ('Jorhat'), ('Lakhimpur'), ('Majuli Sivasagar'), ('Tinsukia');

INSERT INTO "user" (first_name, last_name, dob, email_id, contact_no, skills, year_of_admission, year_of_completion, semester, bio, por_id, education, department, course, hostel, status, username, password, internship)
VALUES ('Kaushik', 'Kumar Bora', '17-01-2001', 'kaushikkumarbora@gmail.com', '9999999999', 'C, C++, Golang', '2018', '2022', 7, 'Hey', (select por_id from por where name='Jorhat'), 'B.Tech, M.Tech, PhD', 'CSE', 'B.Tech', 'CMH', true, 'kaushik', 'kaushik123', 'blablabla'),
('Aasurjya', 'Bikash Handique', '17-01-2001', 'ahandique8@gmail.com', '9999999999', 'C, C++, Golang', '2018', '2022', 7, 'Hey', (select por_id from por where name='Dibrugarh'), 'B.Tech, M.Tech, PhD', 'CSE', 'B.Tech', 'CMH', true, 'shivangshu', 'shivangshu123', 'blablabla'),
('Mithuraj', 'Borgohain', '13-02-1997', 'mithuraj@gmail.com', '9999999999', 'C, C++, Golang', '2018', '2022', 7, 'Hey', (select por_id from por where name='Dhemaji'), 'B.Tech, M.Tech, PhD', 'FET', 'B.Tech', 'CMH', true, 'mithuraj', 'mithuraj123', 'blablabla');

INSERT INTO connection (l_user_id, r_user_id, status)
VALUES (1,2,'p'),
(1,3,'a');