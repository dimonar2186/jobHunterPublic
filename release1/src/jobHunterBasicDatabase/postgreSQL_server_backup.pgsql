--
-- PostgreSQL database cluster dump
--

-- Started on 2024-10-23 23:48:56

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE "jobHunterBasic";
ALTER ROLE "jobHunterBasic" WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:VMnnrdTWZ45oRDAS5DLPwg==$LNGYXxAHOEmeqJVPmnYNp3zJNeBYORvyGJ4FlIYI52M=:zhgkhylhN9FGKmuObUEE1j7shnHvD8sd1WsT2GPCMjg=';
COMMENT ON ROLE "jobHunterBasic" IS 'Regular user of jobHunterBasic database';
CREATE ROLE "jobHunterBasic_su";
ALTER ROLE "jobHunterBasic_su" WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:QmsASFQaHDHI2R6MbPwMfQ==$GXyW44PNTGCVzAmQm30RhAgjUf2RqwrzDXMAtB4fmDk=:cQOF5VgufG6cqFgVsgsBT3seB68ItuCiZYM15+gSEnE=';
COMMENT ON ROLE "jobHunterBasic_su" IS 'Super user of the jobHunterBasic database';
CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:e/1HlggzWAtWe7pSIknmsA==$Lm0eGG9rFdSy1PSNCMX23v8asZsXc6/edyd7GxIF5b8=:X0rE0xNllPDSCgL0t3o37n2RyzD4fqv0j77eDDGT1KE=';

--
-- User Configurations
--








--
-- Databases
--

--
-- Database "template1" dump
--

\connect template1

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

-- Started on 2024-10-23 23:48:56

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

-- Completed on 2024-10-23 23:48:56

--
-- PostgreSQL database dump complete
--

--
-- Database "jobHunterBasic" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

-- Started on 2024-10-23 23:48:56

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4958 (class 1262 OID 16400)
-- Name: jobHunterBasic; Type: DATABASE; Schema: -; Owner: jobHunterBasic_su
--

CREATE DATABASE "jobHunterBasic" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';


ALTER DATABASE "jobHunterBasic" OWNER TO "jobHunterBasic_su";

\connect "jobHunterBasic"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4959 (class 0 OID 0)
-- Dependencies: 4958
-- Name: DATABASE "jobHunterBasic"; Type: COMMENT; Schema: -; Owner: jobHunterBasic_su
--

COMMENT ON DATABASE "jobHunterBasic" IS 'Database for basic jobHunter''s data.';


--
-- TOC entry 5 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: jobHunterBasic_su
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO "jobHunterBasic_su";

--
-- TOC entry 895 (class 1247 OID 16514)
-- Name: application_stage_statuses_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.application_stage_statuses_enum AS ENUM (
    'active',
    'waiting',
    'blocked',
    'passed',
    'failed',
    'deleted'
);


ALTER TYPE public.application_stage_statuses_enum OWNER TO postgres;

--
-- TOC entry 862 (class 1247 OID 16504)
-- Name: job_searching_process_statuses_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.job_searching_process_statuses_enum AS ENUM (
    'active',
    'archived',
    'finished',
    'deleted'
);


ALTER TYPE public.job_searching_process_statuses_enum OWNER TO postgres;

--
-- TOC entry 856 (class 1247 OID 16482)
-- Name: job_types_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.job_types_enum AS ENUM (
    'full-time',
    'part_time',
    'contract'
);


ALTER TYPE public.job_types_enum OWNER TO postgres;

--
-- TOC entry 853 (class 1247 OID 16475)
-- Name: table_type_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.table_type_enum AS ENUM (
    'offers',
    'job_searching_processes',
    'vacancies',
    'employers'
);


ALTER TYPE public.table_type_enum OWNER TO postgres;

--
-- TOC entry 859 (class 1247 OID 16495)
-- Name: vacancy_statuses_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.vacancy_statuses_enum AS ENUM (
    'active',
    'failed',
    'success',
    'deleted'
);


ALTER TYPE public.vacancy_statuses_enum OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 227 (class 1259 OID 16556)
-- Name: application_stages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.application_stages (
    id uuid NOT NULL,
    vacancy uuid NOT NULL,
    index_number money,
    name character varying(250),
    status public.application_stage_statuses_enum NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone
);


ALTER TABLE public.application_stages OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16418)
-- Name: countries; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.countries (
    id uuid NOT NULL,
    name character varying(150) NOT NULL,
    code character varying(10) NOT NULL
);


ALTER TABLE public.countries OWNER TO "jobHunterBasic_su";

--
-- TOC entry 219 (class 1259 OID 16428)
-- Name: currencies; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.currencies (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    short_name character varying(10) NOT NULL,
    iso_code integer NOT NULL
);


ALTER TABLE public.currencies OWNER TO "jobHunterBasic_su";

--
-- TOC entry 220 (class 1259 OID 16450)
-- Name: employers; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.employers (
    id uuid NOT NULL,
    name character varying(250) NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE public.employers OWNER TO "jobHunterBasic_su";

--
-- TOC entry 222 (class 1259 OID 16462)
-- Name: hrcontacts; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.hrcontacts (
    id uuid NOT NULL,
    value character varying(250) NOT NULL,
    messenger uuid NOT NULL,
    preferred boolean NOT NULL,
    is_deleted boolean NOT NULL,
    hr_manager uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone
);


ALTER TABLE public.hrcontacts OWNER TO "jobHunterBasic_su";

--
-- TOC entry 221 (class 1259 OID 16455)
-- Name: hrmanagers; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.hrmanagers (
    id uuid NOT NULL,
    first_name character varying(250) NOT NULL,
    last_name character varying(250),
    second_name character varying(250),
    employer uuid NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE public.hrmanagers OWNER TO "jobHunterBasic_su";

--
-- TOC entry 216 (class 1259 OID 16406)
-- Name: job_searching_processes; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.job_searching_processes (
    id uuid NOT NULL,
    name character varying(1000) NOT NULL,
    min_mounthly_salary money,
    max_mounthly_salary money,
    "position" character varying(2000)[],
    "user" uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    is_deleted boolean,
    deleted_at timestamp with time zone,
    currency uuid,
    status public.job_searching_process_statuses_enum NOT NULL
);


ALTER TABLE public.job_searching_processes OWNER TO "jobHunterBasic_su";

--
-- TOC entry 224 (class 1259 OID 16489)
-- Name: links_to_job_types; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.links_to_job_types (
    id uuid NOT NULL,
    object_id uuid NOT NULL,
    table_type public.table_type_enum NOT NULL,
    job_type public.job_types_enum NOT NULL
);


ALTER TABLE public.links_to_job_types OWNER TO "jobHunterBasic_su";

--
-- TOC entry 225 (class 1259 OID 16527)
-- Name: locations; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.locations (
    id uuid NOT NULL,
    object_id uuid NOT NULL,
    table_type character varying(30) NOT NULL,
    country uuid NOT NULL,
    CONSTRAINT locations_table_type_check CHECK (((table_type)::text = ANY ((ARRAY['employers'::character varying, 'job_searching_processes'::character varying, 'vacancies'::character varying])::text[])))
);


ALTER TABLE public.locations OWNER TO "jobHunterBasic_su";

--
-- TOC entry 217 (class 1259 OID 16413)
-- Name: messengers; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.messengers (
    id uuid NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.messengers OWNER TO "jobHunterBasic_su";

--
-- TOC entry 226 (class 1259 OID 16539)
-- Name: offer; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.offer (
    id uuid NOT NULL,
    vacancy uuid NOT NULL,
    mounthly_salary money,
    currency uuid,
    contract_length integer,
    "position" character varying(250),
    commentary text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    job_type public.job_types_enum
);


ALTER TABLE public.offer OWNER TO "jobHunterBasic_su";

--
-- TOC entry 215 (class 1259 OID 16401)
-- Name: users; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.users (
    id uuid NOT NULL
);


ALTER TABLE public.users OWNER TO "jobHunterBasic_su";

--
-- TOC entry 223 (class 1259 OID 16467)
-- Name: vacancies; Type: TABLE; Schema: public; Owner: jobHunterBasic_su
--

CREATE TABLE public.vacancies (
    id uuid NOT NULL,
    job_searching_process uuid NOT NULL,
    min_mounthly_salary money,
    max_mounthly_salary money,
    currency uuid NOT NULL,
    contract_length integer,
    benefits text,
    responsibilities text,
    "position" character varying(250),
    commentary text,
    employer uuid NOT NULL,
    applied_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    status public.vacancy_statuses_enum NOT NULL
);


ALTER TABLE public.vacancies OWNER TO "jobHunterBasic_su";

--
-- TOC entry 228 (class 1259 OID 16591)
-- Name: vacancies_to_hrmanagers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vacancies_to_hrmanagers (
    id uuid NOT NULL,
    hrmanager uuid NOT NULL,
    vacancy uuid NOT NULL
);


ALTER TABLE public.vacancies_to_hrmanagers OWNER TO postgres;

--
-- TOC entry 4951 (class 0 OID 16556)
-- Dependencies: 227
-- Data for Name: application_stages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.application_stages (id, vacancy, index_number, name, status, created_at, updated_at) FROM stdin;
\.


--
-- TOC entry 4942 (class 0 OID 16418)
-- Dependencies: 218
-- Data for Name: countries; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.countries (id, name, code) FROM stdin;
\.


--
-- TOC entry 4943 (class 0 OID 16428)
-- Dependencies: 219
-- Data for Name: currencies; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.currencies (id, name, short_name, iso_code) FROM stdin;
\.


--
-- TOC entry 4944 (class 0 OID 16450)
-- Dependencies: 220
-- Data for Name: employers; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.employers (id, name, created_at) FROM stdin;
\.


--
-- TOC entry 4946 (class 0 OID 16462)
-- Dependencies: 222
-- Data for Name: hrcontacts; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.hrcontacts (id, value, messenger, preferred, is_deleted, hr_manager, created_at, updated_at) FROM stdin;
\.


--
-- TOC entry 4945 (class 0 OID 16455)
-- Dependencies: 221
-- Data for Name: hrmanagers; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.hrmanagers (id, first_name, last_name, second_name, employer, created_at) FROM stdin;
\.


--
-- TOC entry 4940 (class 0 OID 16406)
-- Dependencies: 216
-- Data for Name: job_searching_processes; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.job_searching_processes (id, name, min_mounthly_salary, max_mounthly_salary, "position", "user", created_at, updated_at, is_deleted, deleted_at, currency, status) FROM stdin;
\.


--
-- TOC entry 4948 (class 0 OID 16489)
-- Dependencies: 224
-- Data for Name: links_to_job_types; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.links_to_job_types (id, object_id, table_type, job_type) FROM stdin;
\.


--
-- TOC entry 4949 (class 0 OID 16527)
-- Dependencies: 225
-- Data for Name: locations; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.locations (id, object_id, table_type, country) FROM stdin;
\.


--
-- TOC entry 4941 (class 0 OID 16413)
-- Dependencies: 217
-- Data for Name: messengers; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.messengers (id, name) FROM stdin;
\.


--
-- TOC entry 4950 (class 0 OID 16539)
-- Dependencies: 226
-- Data for Name: offer; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.offer (id, vacancy, mounthly_salary, currency, contract_length, "position", commentary, created_at, updated_at, job_type) FROM stdin;
\.


--
-- TOC entry 4939 (class 0 OID 16401)
-- Dependencies: 215
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.users (id) FROM stdin;
\.


--
-- TOC entry 4947 (class 0 OID 16467)
-- Dependencies: 223
-- Data for Name: vacancies; Type: TABLE DATA; Schema: public; Owner: jobHunterBasic_su
--

COPY public.vacancies (id, job_searching_process, min_mounthly_salary, max_mounthly_salary, currency, contract_length, benefits, responsibilities, "position", commentary, employer, applied_at, created_at, updated_at, status) FROM stdin;
\.


--
-- TOC entry 4952 (class 0 OID 16591)
-- Dependencies: 228
-- Data for Name: vacancies_to_hrmanagers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.vacancies_to_hrmanagers (id, hrmanager, vacancy) FROM stdin;
\.


--
-- TOC entry 4780 (class 2606 OID 16560)
-- Name: application_stages application_stages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_stages
    ADD CONSTRAINT application_stages_pkey PRIMARY KEY (id);


--
-- TOC entry 4762 (class 2606 OID 16427)
-- Name: countries countries_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.countries
    ADD CONSTRAINT countries_pkey PRIMARY KEY (id);


--
-- TOC entry 4764 (class 2606 OID 16432)
-- Name: currencies currencies_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.currencies
    ADD CONSTRAINT currencies_pkey PRIMARY KEY (id);


--
-- TOC entry 4766 (class 2606 OID 16454)
-- Name: employers employers_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.employers
    ADD CONSTRAINT employers_pkey PRIMARY KEY (id);


--
-- TOC entry 4770 (class 2606 OID 16466)
-- Name: hrcontacts hrcontacts_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.hrcontacts
    ADD CONSTRAINT hrcontacts_pkey PRIMARY KEY (id);


--
-- TOC entry 4768 (class 2606 OID 16461)
-- Name: hrmanagers hrmanagers_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.hrmanagers
    ADD CONSTRAINT hrmanagers_pkey PRIMARY KEY (id);


--
-- TOC entry 4758 (class 2606 OID 16412)
-- Name: job_searching_processes job_searching_processes_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.job_searching_processes
    ADD CONSTRAINT job_searching_processes_pkey PRIMARY KEY (id);


--
-- TOC entry 4774 (class 2606 OID 16493)
-- Name: links_to_job_types links_to_job_types_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.links_to_job_types
    ADD CONSTRAINT links_to_job_types_pkey PRIMARY KEY (id);


--
-- TOC entry 4776 (class 2606 OID 16532)
-- Name: locations locations_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (id);


--
-- TOC entry 4760 (class 2606 OID 16417)
-- Name: messengers messengers_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.messengers
    ADD CONSTRAINT messengers_pkey PRIMARY KEY (id);


--
-- TOC entry 4778 (class 2606 OID 16545)
-- Name: offer offer_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.offer
    ADD CONSTRAINT offer_pkey PRIMARY KEY (id);


--
-- TOC entry 4756 (class 2606 OID 16405)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4772 (class 2606 OID 16473)
-- Name: vacancies vacancies_pkey; Type: CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.vacancies
    ADD CONSTRAINT vacancies_pkey PRIMARY KEY (id);


--
-- TOC entry 4782 (class 2606 OID 16595)
-- Name: vacancies_to_hrmanagers vacancies_to_hrmanagers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacancies_to_hrmanagers
    ADD CONSTRAINT vacancies_to_hrmanagers_pkey PRIMARY KEY (id);


--
-- TOC entry 4793 (class 2606 OID 16561)
-- Name: application_stages application_stages_vacancy_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_stages
    ADD CONSTRAINT application_stages_vacancy_fkey FOREIGN KEY (vacancy) REFERENCES public.vacancies(id);


--
-- TOC entry 4783 (class 2606 OID 16433)
-- Name: job_searching_processes currency; Type: FK CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.job_searching_processes
    ADD CONSTRAINT currency FOREIGN KEY (currency) REFERENCES public.currencies(id) NOT VALID;


--
-- TOC entry 4788 (class 2606 OID 16571)
-- Name: vacancies currency; Type: FK CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.vacancies
    ADD CONSTRAINT currency FOREIGN KEY (currency) REFERENCES public.currencies(id) NOT VALID;


--
-- TOC entry 4789 (class 2606 OID 16566)
-- Name: vacancies employer; Type: FK CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.vacancies
    ADD CONSTRAINT employer FOREIGN KEY (employer) REFERENCES public.employers(id) NOT VALID;


--
-- TOC entry 4785 (class 2606 OID 16576)
-- Name: hrmanagers employer; Type: FK CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.hrmanagers
    ADD CONSTRAINT employer FOREIGN KEY (employer) REFERENCES public.employers(id) NOT VALID;


--
-- TOC entry 4786 (class 2606 OID 16581)
-- Name: hrcontacts hrmanager; Type: FK CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.hrcontacts
    ADD CONSTRAINT hrmanager FOREIGN KEY (hr_manager) REFERENCES public.hrmanagers(id) NOT VALID;


--
-- TOC entry 4790 (class 2606 OID 16533)
-- Name: locations locations_country_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_country_fkey FOREIGN KEY (country) REFERENCES public.countries(id);


--
-- TOC entry 4787 (class 2606 OID 16586)
-- Name: hrcontacts messenger; Type: FK CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.hrcontacts
    ADD CONSTRAINT messenger FOREIGN KEY (messenger) REFERENCES public.messengers(id) NOT VALID;


--
-- TOC entry 4791 (class 2606 OID 16551)
-- Name: offer offer_currency_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.offer
    ADD CONSTRAINT offer_currency_fkey FOREIGN KEY (currency) REFERENCES public.currencies(id);


--
-- TOC entry 4792 (class 2606 OID 16546)
-- Name: offer offer_vacancy_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.offer
    ADD CONSTRAINT offer_vacancy_fkey FOREIGN KEY (vacancy) REFERENCES public.vacancies(id);


--
-- TOC entry 4784 (class 2606 OID 16438)
-- Name: job_searching_processes user; Type: FK CONSTRAINT; Schema: public; Owner: jobHunterBasic_su
--

ALTER TABLE ONLY public.job_searching_processes
    ADD CONSTRAINT "user" FOREIGN KEY ("user") REFERENCES public.users(id) ON DELETE CASCADE NOT VALID;


--
-- TOC entry 4794 (class 2606 OID 16596)
-- Name: vacancies_to_hrmanagers vacancies_to_hrmanagers_hrmanager_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacancies_to_hrmanagers
    ADD CONSTRAINT vacancies_to_hrmanagers_hrmanager_fkey FOREIGN KEY (hrmanager) REFERENCES public.hrmanagers(id);


--
-- TOC entry 4795 (class 2606 OID 16601)
-- Name: vacancies_to_hrmanagers vacancies_to_hrmanagers_vacancy_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacancies_to_hrmanagers
    ADD CONSTRAINT vacancies_to_hrmanagers_vacancy_fkey FOREIGN KEY (vacancy) REFERENCES public.vacancies(id);


-- Completed on 2024-10-23 23:48:56

--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

\connect postgres

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

-- Started on 2024-10-23 23:48:57

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 16384)
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- TOC entry 4832 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


-- Completed on 2024-10-23 23:48:57

--
-- PostgreSQL database dump complete
--

-- Completed on 2024-10-23 23:48:57

--
-- PostgreSQL database cluster dump complete
--

