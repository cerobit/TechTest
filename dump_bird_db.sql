--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Drop databases (except postgres and template1)
--

DROP DATABASE birddemo;




--
-- Drop roles
--

DROP ROLE postgres;


--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'md5cf0e7d1525b2a85d3dcf7e5b68359b47';






--
-- Databases
--

--
-- Database "template1" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1 (Debian 13.1-1.pgdg100+1)
-- Dumped by pg_dump version 13.1 (Debian 13.1-1.pgdg100+1)

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

UPDATE pg_catalog.pg_database SET datistemplate = false WHERE datname = 'template1';
DROP DATABASE template1;
--
-- Name: template1; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE template1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE template1 OWNER TO postgres;

\connect template1

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
-- Name: DATABASE template1; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE template1 IS 'default template for new databases';


--
-- Name: template1; Type: DATABASE PROPERTIES; Schema: -; Owner: postgres
--

ALTER DATABASE template1 IS_TEMPLATE = true;


\connect template1

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
-- Name: DATABASE template1; Type: ACL; Schema: -; Owner: postgres
--

REVOKE CONNECT,TEMPORARY ON DATABASE template1 FROM PUBLIC;
GRANT CONNECT ON DATABASE template1 TO PUBLIC;


--
-- PostgreSQL database dump complete
--

--
-- Database "birddemo" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1 (Debian 13.1-1.pgdg100+1)
-- Dumped by pg_dump version 13.1 (Debian 13.1-1.pgdg100+1)

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
-- Name: birddemo; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE birddemo WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE birddemo OWNER TO postgres;

\connect birddemo

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
-- Name: birdinfo; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA birdinfo;


ALTER SCHEMA birdinfo OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: birds; Type: TABLE; Schema: birdinfo; Owner: postgres
--

CREATE TABLE birdinfo.birds (
    id integer NOT NULL,
    specie character varying NOT NULL,
    name character varying NOT NULL,
    characteristics character varying NOT NULL
);


ALTER TABLE birdinfo.birds OWNER TO postgres;

--
-- Name: birds_id_seq; Type: SEQUENCE; Schema: birdinfo; Owner: postgres
--

CREATE SEQUENCE birdinfo.birds_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE birdinfo.birds_id_seq OWNER TO postgres;

--
-- Name: birds_id_seq; Type: SEQUENCE OWNED BY; Schema: birdinfo; Owner: postgres
--

ALTER SEQUENCE birdinfo.birds_id_seq OWNED BY birdinfo.birds.id;


--
-- Name: birds id; Type: DEFAULT; Schema: birdinfo; Owner: postgres
--

ALTER TABLE ONLY birdinfo.birds ALTER COLUMN id SET DEFAULT nextval('birdinfo.birds_id_seq'::regclass);


--
-- Data for Name: birds; Type: TABLE DATA; Schema: birdinfo; Owner: postgres
--

COPY birdinfo.birds (id, specie, name, characteristics) FROM stdin;
1	Galliforme	Faisan	Ave de vistoso plumaje
2	Gruidae	Grulla comun	Ave migratoria 
3	Petirrojo		Beaituful bird from north america
6	Petirrojo		Beaituful bird from north america
7	Petirrojo		Beaituful bird from north america
10	Petirrojo		Beaituful bird from north america
11		Petirrojo	Beatiful Bird from north america
12	Petirrojo	Petirrojo	Beatiful Bird from north america
13	Petirrojo	Petirrojo	Beatiful Bird from north america
14	Petirrojo	Petirrojo	Beatiful Bird from north america
5	PetiAvi- 2  	PetiAzul	Beaituful bird from north america
16	Petirrojo	Petirrojo	Beatiful Bird from north america
18	Petirrojo Test 2	Petirrojo	Beatiful Bird from north america
20	Petirrojo Test 2	Petirrojo	Beatiful Bird from north america
21	Petirrojo Test 2	Petirrojo	Beatiful Bird from north america
22	Petirrojo Test 2	Petirrojo	Beatiful Bird from north america
\.


--
-- Name: birds_id_seq; Type: SEQUENCE SET; Schema: birdinfo; Owner: postgres
--

SELECT pg_catalog.setval('birdinfo.birds_id_seq', 23, true);


--
-- Name: birds birds_pk; Type: CONSTRAINT; Schema: birdinfo; Owner: postgres
--

ALTER TABLE ONLY birdinfo.birds
    ADD CONSTRAINT birds_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1 (Debian 13.1-1.pgdg100+1)
-- Dumped by pg_dump version 13.1 (Debian 13.1-1.pgdg100+1)

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

DROP DATABASE postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO postgres;

\connect postgres

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
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

