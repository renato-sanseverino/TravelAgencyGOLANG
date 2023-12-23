CREATE DATABASE travel_agency;
\c travel_agency

CREATE TABLE public.accommodations (
	id SERIAL,
	hotel varchar(120) NOT NULL,
	guests integer NOT NULL,
	checkin date NOT NULL,
	checkout date NULL,
	room integer NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.travelinsurance (
	id SERIAL,
	client_id integer NOT NULL,
	"purposeOfTrip" varchar(80) NOT NULL,
	luggage numeric(10, 2) NULL,
	medical_cover numeric(10, 2) NULL,
	price_total numeric(10, 2) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.itinerary (
	id SERIAL,
	destination varchar(120) NOT NULL,
	departure date NOT NULL,
	arrival date NULL,
	transport_kind varchar(65) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.client (
	id SERIAL,
	"name" varchar(120) NOT NULL,
	"address" varchar(200) NULL,
	occupation varchar(120) NULL,
	birth_date date NOT NULL,
	email varchar(80) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.travelpackages (
	id SERIAL,
	description varchar(120) NOT NULL,
	client_id integer NOT NULL,
	country varchar(80) NOT NULL,
	city varchar(80) NOT NULL,
	accommodation_id integer NULL,
	insurance_id integer NULL,
	price_total numeric(10, 2) NOT NULL,
	PRIMARY KEY (id)
);

ALTER TABLE public.travelpackages
    ADD CONSTRAINT fk_tour_accommodation
	    FOREIGN KEY (accommodation_id) REFERENCES public.accommodations(id);

ALTER TABLE public.travelpackages
    ADD CONSTRAINT fk_tour_insurance
		FOREIGN KEY (insurance_id) REFERENCES public.travelinsurance(id);

ALTER TABLE public.travelpackages
    ADD CONSTRAINT fk_tour_client
		FOREIGN KEY (client_id) REFERENCES public.client(id);


CREATE TABLE public.eventtickets (
	id SERIAL,
	client_id integer NOT NULL,
	description varchar(80) NOT NULL,
	"location" varchar(80) NULL,
	price numeric(10, 2) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.guidedtours (
	id SERIAL,
	description varchar(120) NOT NULL,
	"date" date NOT NULL,
	participants integer NOT NULL,
	PRIMARY KEY (id)
);

-- data for table client
INSERT INTO public.client ("name","address","occupation","birth_date","email") VALUES
	 ('Herbert Olga','Avenida 9 de Julho, 340','Engineer','1982-11-10','herbertolga@gmail.com.ca'),
	 ('Isabela Cristina','Rua Oscar Freire, 290','Teacher','1997-05-05','isabela112@gmail.com');

-- data for table itinerary
INSERT INTO public.itinerary (destination,departure,arrival,transport_kind) VALUES
	 ('São Paulo','2023-12-20',NULL,'bus'),
	 ('Belo Horizonte','2023-12-20',NULL,'bus'),
	 ('Recife','2023-12-20',NULL,'plane'),
	 ('Curitiba','2023-12-20',NULL,'bus');
