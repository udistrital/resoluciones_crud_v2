-- Database generated with pgModeler (desarrollooasQL Database Modeler).
-- pgModeler  version: 0.9.3-beta1
-- desarrollooasQL version: 9.6
-- Project Site: pgmodeler.io
-- Model Author: Juan Pablo Moreno Rico

-- object: resoluciones_new | type: SCHEMA --
-- DROP SCHEMA IF EXISTS resoluciones_new CASCADE;
CREATE SCHEMA IF NOT EXISTS resoluciones_new;
-- ddl-end --
ALTER SCHEMA resoluciones_new OWNER TO desarrollooas;
-- ddl-end --
COMMENT ON SCHEMA resoluciones_new IS E'Modelo de datos del esquema para el nuevo sistema de resoluciones';
-- ddl-end --

SET search_path TO pg_catalog,public,resoluciones_new;
-- ddl-end --

-- object: resoluciones_new.resolucion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS resoluciones_new.resolucion_id_seq CASCADE;
CREATE SEQUENCE resoluciones_new.resolucion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE resoluciones_new.resolucion_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.resolucion | type: TABLE --
-- DROP TABLE IF EXISTS resoluciones_new.resolucion CASCADE;
CREATE TABLE resoluciones_new.resolucion (
	id integer NOT NULL DEFAULT nextval('resoluciones_new.resolucion_id_seq'::regclass),
	numero_resolucion character varying NOT NULL DEFAULT 10,
	fecha_expedicion timestamp,
	vigencia integer NOT NULL,
	dependencia_id integer NOT NULL,
	tipo_resolucion_id integer NOT NULL,
	preambulo_resolucion text NOT NULL,
	consideracion_resolucion text NOT NULL,
	numero_semanas integer NOT NULL DEFAULT 0,
	periodo integer NOT NULL DEFAULT 0,
	titulo text,
	dependencia_firma_id integer,
	vigencia_carga integer,
	periodo_carga integer,
	cuadro_responsabilidades json,
	nuxeo_uid varchar(50),
	activo boolean DEFAULT true,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp,
	CONSTRAINT pk_resolucion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE resoluciones_new.resolucion IS E'Resolución de vinculación de docentes';
-- ddl-end --
ALTER TABLE resoluciones_new.resolucion OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.vinculacion_docente_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS resoluciones_new.vinculacion_docente_id_seq CASCADE;
CREATE SEQUENCE resoluciones_new.vinculacion_docente_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE resoluciones_new.vinculacion_docente_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.vinculacion_docente | type: TABLE --
-- DROP TABLE IF EXISTS resoluciones_new.vinculacion_docente CASCADE;
CREATE TABLE resoluciones_new.vinculacion_docente (
	id integer NOT NULL DEFAULT nextval('resoluciones_new.vinculacion_docente_id_seq'::regclass),
	numero_contrato character varying(10),
	vigencia integer,
	persona_id numeric(15,0) NOT NULL,
	numero_horas_semanales integer NOT NULL,
	numero_semanas integer NOT NULL,
	punto_salarial_id integer,
	salario_minimo_id integer,
	resolucion_vinculacion_docente_id integer NOT NULL,
	dedicacion_id integer NOT NULL,
	proyecto_curricular_id smallint NOT NULL,
	valor_contrato numeric(16,3),
	categoria character varying(15),
	emerito boolean NOT NULL DEFAULT false,
	disponibilidad integer,
	dependencia_academica integer,
	numero_rp numeric(6,0),
	vigencia_rp numeric(4,0),
	fecha_inicio timestamp,
	activo boolean DEFAULT true,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp,
	CONSTRAINT pk_contrato_docente PRIMARY KEY (id),
	CONSTRAINT uq_numero_contrato_vinculacion_docente UNIQUE (numero_contrato,vigencia)

);
-- ddl-end --
COMMENT ON TABLE resoluciones_new.vinculacion_docente IS E'Contiene información detallada de la vinculación del docente';
-- ddl-end --
ALTER TABLE resoluciones_new.vinculacion_docente OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.resolucion_vinculacion_docente_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS resoluciones_new.resolucion_vinculacion_docente_id_seq CASCADE;
CREATE SEQUENCE resoluciones_new.resolucion_vinculacion_docente_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE resoluciones_new.resolucion_vinculacion_docente_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.resolucion_vinculacion_docente | type: TABLE --
-- DROP TABLE IF EXISTS resoluciones_new.resolucion_vinculacion_docente CASCADE;
CREATE TABLE resoluciones_new.resolucion_vinculacion_docente (
	id integer NOT NULL DEFAULT nextval('resoluciones_new.resolucion_vinculacion_docente_id_seq'::regclass),
	facultad_id integer NOT NULL,
	dedicacion character varying(12) NOT NULL,
	nivel_academico character varying(15) NOT NULL,
	activo boolean DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_resolucion_vinculacion_docente PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE resoluciones_new.resolucion_vinculacion_docente IS E'Detalles de lavinculación del docente por medio de la resolución';
-- ddl-end --
ALTER TABLE resoluciones_new.resolucion_vinculacion_docente OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.componente_resolucion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS resoluciones_new.componente_resolucion_id_seq CASCADE;
CREATE SEQUENCE resoluciones_new.componente_resolucion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE resoluciones_new.componente_resolucion_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.componente_resolucion | type: TABLE --
-- DROP TABLE IF EXISTS resoluciones_new.componente_resolucion CASCADE;
CREATE TABLE resoluciones_new.componente_resolucion (
	id integer NOT NULL DEFAULT nextval('resoluciones_new.componente_resolucion_id_seq'::regclass),
	numero integer NOT NULL,
	resolucion_id smallint NOT NULL,
	texto text NOT NULL,
	tipo_componente character varying(15) NOT NULL,
	componente_resolucion_padre integer,
	activo boolean DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_componente_resolucion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE resoluciones_new.componente_resolucion IS E'Textos que componen una resolución, hace referencia a los diferentes artículos on sus respectivos parágrafos, si existen.';
-- ddl-end --
ALTER TABLE resoluciones_new.componente_resolucion OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.resolucion_estado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS resoluciones_new.resolucion_estado_id_seq CASCADE;
CREATE SEQUENCE resoluciones_new.resolucion_estado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE resoluciones_new.resolucion_estado_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.resolucion_estado | type: TABLE --
-- DROP TABLE IF EXISTS resoluciones_new.resolucion_estado CASCADE;
CREATE TABLE resoluciones_new.resolucion_estado (
	id integer NOT NULL DEFAULT nextval('resoluciones_new.resolucion_estado_id_seq'::regclass),
	usuario character varying(50),
	estado_resolucion_id integer NOT NULL,
	resolucion_id integer NOT NULL,
	activo boolean DEFAULT true,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp,
	CONSTRAINT pk_resolucion_estado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE resoluciones_new.resolucion_estado IS E'Estado de la resolución';
-- ddl-end --
ALTER TABLE resoluciones_new.resolucion_estado OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.modificacion_resolucion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS resoluciones_new.modificacion_resolucion_id_seq CASCADE;
CREATE SEQUENCE resoluciones_new.modificacion_resolucion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE resoluciones_new.modificacion_resolucion_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.modificacion_resolucion | type: TABLE --
-- DROP TABLE IF EXISTS resoluciones_new.modificacion_resolucion CASCADE;
CREATE TABLE resoluciones_new.modificacion_resolucion (
	id integer NOT NULL DEFAULT nextval('resoluciones_new.modificacion_resolucion_id_seq'::regclass),
	resolucion_nueva_id integer NOT NULL,
	resolucion_anterior_id integer NOT NULL,
	activo boolean DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_modificacion_resolucion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE resoluciones_new.modificacion_resolucion IS E'Modificaciones realizadas a una resolución por medio de otra';
-- ddl-end --
ALTER TABLE resoluciones_new.modificacion_resolucion OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.modificacion_vinculacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS resoluciones_new.modificacion_vinculacion_id_seq CASCADE;
CREATE SEQUENCE resoluciones_new.modificacion_vinculacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE resoluciones_new.modificacion_vinculacion_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: resoluciones_new.modificacion_vinculacion | type: TABLE --
-- DROP TABLE IF EXISTS resoluciones_new.modificacion_vinculacion CASCADE;
CREATE TABLE resoluciones_new.modificacion_vinculacion (
	id integer NOT NULL DEFAULT nextval('resoluciones_new.modificacion_vinculacion_id_seq'::regclass),
	modificacion_resolucion_id integer NOT NULL,
	vinculacion_docente_cancelada_id integer NOT NULL,
	vinculacion_docente_registrada_id integer,
	horas numeric(2,0),
	activo boolean DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_modificacion_vinculacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE resoluciones_new.modificacion_vinculacion IS E'Modificaciones realizadas a la vinculación de un docente a traves de la modificación de una resolución';
-- ddl-end --
ALTER TABLE resoluciones_new.modificacion_vinculacion OWNER TO desarrollooas;
-- ddl-end --

-- object: fk_vinculacion_docente_resolucion_vinculacion_docente | type: CONSTRAINT --
-- ALTER TABLE resoluciones_new.vinculacion_docente DROP CONSTRAINT IF EXISTS fk_vinculacion_docente_resolucion_vinculacion_docente CASCADE;
ALTER TABLE resoluciones_new.vinculacion_docente ADD CONSTRAINT fk_vinculacion_docente_resolucion_vinculacion_docente FOREIGN KEY (resolucion_vinculacion_docente_id)
REFERENCES resoluciones_new.resolucion_vinculacion_docente (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_resolucion_vinculacion_docente_resolucion | type: CONSTRAINT --
-- ALTER TABLE resoluciones_new.resolucion_vinculacion_docente DROP CONSTRAINT IF EXISTS fk_resolucion_vinculacion_docente_resolucion CASCADE;
ALTER TABLE resoluciones_new.resolucion_vinculacion_docente ADD CONSTRAINT fk_resolucion_vinculacion_docente_resolucion FOREIGN KEY (id)
REFERENCES resoluciones_new.resolucion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_componente_resolucion_resolucion | type: CONSTRAINT --
-- ALTER TABLE resoluciones_new.componente_resolucion DROP CONSTRAINT IF EXISTS fk_componente_resolucion_resolucion CASCADE;
ALTER TABLE resoluciones_new.componente_resolucion ADD CONSTRAINT fk_componente_resolucion_resolucion FOREIGN KEY (resolucion_id)
REFERENCES resoluciones_new.resolucion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_componente_resolucion_componente_resolucion_padre | type: CONSTRAINT --
-- ALTER TABLE resoluciones_new.componente_resolucion DROP CONSTRAINT IF EXISTS fk_componente_resolucion_componente_resolucion_padre CASCADE;
ALTER TABLE resoluciones_new.componente_resolucion ADD CONSTRAINT fk_componente_resolucion_componente_resolucion_padre FOREIGN KEY (componente_resolucion_padre)
REFERENCES resoluciones_new.componente_resolucion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_resolucion_estado_resolucion | type: CONSTRAINT --
-- ALTER TABLE resoluciones_new.resolucion_estado DROP CONSTRAINT IF EXISTS fk_resolucion_estado_resolucion CASCADE;
ALTER TABLE resoluciones_new.resolucion_estado ADD CONSTRAINT fk_resolucion_estado_resolucion FOREIGN KEY (resolucion_id)
REFERENCES resoluciones_new.resolucion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_modificacion_resolucion_resolucion_nueva | type: CONSTRAINT --
-- ALTER TABLE resoluciones_new.modificacion_resolucion DROP CONSTRAINT IF EXISTS fk_modificacion_resolucion_resolucion_nueva CASCADE;
ALTER TABLE resoluciones_new.modificacion_resolucion ADD CONSTRAINT fk_modificacion_resolucion_resolucion_nueva FOREIGN KEY (resolucion_nueva_id)
REFERENCES resoluciones_new.resolucion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_modificacion_resolucion_resolucion_anterior | type: CONSTRAINT --
-- ALTER TABLE resoluciones_new.modificacion_resolucion DROP CONSTRAINT IF EXISTS fk_modificacion_resolucion_resolucion_anterior CASCADE;
ALTER TABLE resoluciones_new.modificacion_resolucion ADD CONSTRAINT fk_modificacion_resolucion_resolucion_anterior FOREIGN KEY (resolucion_anterior_id)
REFERENCES resoluciones_new.resolucion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_modificacion_vinculacion_vinculacion_docente_cancelada | type: CONSTRAINT --
-- ALTER TABLE resoluciones_new.modificacion_vinculacion DROP CONSTRAINT IF EXISTS fk_modificacion_vinculacion_vinculacion_docente_cancelada CASCADE;
ALTER TABLE resoluciones_new.modificacion_vinculacion ADD CONSTRAINT fk_modificacion_vinculacion_vinculacion_docente_cancelada FOREIGN KEY (vinculacion_docente_cancelada_id)
REFERENCES resoluciones_new.vinculacion_docente (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_modificacion_vinculacion_vinculacion_docente_registrada | type: CONSTRAINT --
-- ALTER TABLE resoluciones_new.modificacion_vinculacion DROP CONSTRAINT IF EXISTS fk_modificacion_vinculacion_vinculacion_docente_registrada CASCADE;
ALTER TABLE resoluciones_new.modificacion_vinculacion ADD CONSTRAINT fk_modificacion_vinculacion_vinculacion_docente_registrada FOREIGN KEY (vinculacion_docente_registrada_id)
REFERENCES resoluciones_new.vinculacion_docente (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_modificacion_vinculacion_modificacion_resolucion | type: CONSTRAINT --
-- ALTER TABLE resoluciones_new.modificacion_vinculacion DROP CONSTRAINT IF EXISTS fk_modificacion_vinculacion_modificacion_resolucion CASCADE;
ALTER TABLE resoluciones_new.modificacion_vinculacion ADD CONSTRAINT fk_modificacion_vinculacion_modificacion_resolucion FOREIGN KEY (modificacion_resolucion_id)
REFERENCES resoluciones_new.modificacion_resolucion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
