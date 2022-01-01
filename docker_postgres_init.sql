CREATE TABLE public.connection
(
    l_user bigint NOT NULL,
    r_user bigint NOT NULL,
    status "char" NOT NULL,
    CONSTRAINT l_user_fkey FOREIGN KEY (l_user)
        REFERENCES public.user (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT r_user_fkey FOREIGN KEY (r_user)
        REFERENCES public.user (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

CREATE TABLE public.course
(
    course_id bigint NOT NULL,
    department_id bigint NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT course_pkey PRIMARY KEY (course_id),
    CONSTRAINT department_id_fkey FOREIGN KEY (department_id)
        REFERENCES public.department (department_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

CREATE TABLE public.department
(
    department_id bigint NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT department_pkey PRIMARY KEY (department_id)
)

CREATE TABLE public.hostel
(
    hostel_id bigint NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT hostel_pkey PRIMARY KEY (hostel_id)
)

CREATE TABLE public.por
(
    place_id bigint NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT por_pkey PRIMARY KEY (place_id)
)

CREATE TABLE public.user
(
    user_id bigint NOT NULL,
    first_name text COLLATE pg_catalog."default",
    last_name text COLLATE pg_catalog."default",
    dob date,
    email_id text COLLATE pg_catalog."default",
    contact_no text COLLATE pg_catalog."default",
    skills text COLLATE pg_catalog."default",
    year_of_admission text COLLATE pg_catalog."default",
    year_of_completion text COLLATE pg_catalog."default",
    department_id integer NOT NULL,
    course_id integer NOT NULL,
    curr_sem integer,
    bio text COLLATE pg_catalog."default",
    por_id bigint NOT NULL,
    hostel_id bigint NOT NULL,
    resume bytea,
    status boolean NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (user_id),
    CONSTRAINT course_id_fkey FOREIGN KEY (course_id)
        REFERENCES public.course (course_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT department_id_fkey FOREIGN KEY (department_id)
        REFERENCES public.department (department_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT hostel_id_fkey FOREIGN KEY (hostel_id)
        REFERENCES public.hostel (hostel_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT por_id_fkey FOREIGN KEY (por_id)
        REFERENCES public.por (place_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)