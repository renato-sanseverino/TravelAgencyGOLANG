CREATE DATABASE travel_agency;
\c travel_agency

CREATE TABLE public.accommodations (
	id uuid NOT NULL,
	hotel varchar(120) NOT NULL,
	guests integer NOT NULL,
	checkin date NOT NULL,
	checkout date NULL,
	room integer NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.travelinsurance (
	id uuid NOT NULL,
	client_id uuid NOT NULL,
	"purposeOfTrip" varchar(80) NOT NULL,
	luggage numeric(10, 2) NULL,
	medical_cover numeric(10, 2) NULL,
	price_total numeric(10, 2) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.itinerary (
	id uuid NOT NULL,
	destination varchar(120) NOT NULL,
	departure date NOT NULL,
	arrival date NULL,
	transport_kind varchar(65) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.client (
	id uuid NOT NULL,
	"name" varchar(120) NOT NULL,
	"address" varchar(200) NULL,
	occupation varchar(120) NULL,
	birth_date date NOT NULL,
	email varchar(80) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.travelpackages (
	id uuid NOT NULL,
	description varchar(120) NOT NULL,
	client_id uuid NOT NULL,
	country varchar(80) NOT NULL,
	city varchar(80) NOT NULL,
	accommodation_id uuid NULL,
	insurance_id uuid NULL,
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
	id uuid NOT NULL,
	client_id uuid NOT NULL,
	description varchar(80) NOT NULL,
	"location" varchar(80) NULL,
	price numeric(10, 2) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.guidedtours (
	id uuid NOT NULL,
	description varchar(120) NOT NULL,
	"date" date NOT NULL,
	participants integer NOT NULL,
	PRIMARY KEY (id)
);

-- data for table client
INSERT INTO public.client ("id","name","address","occupation","birth_date","email") VALUES
	 ('6b54dc91-52b6-46d1-ae82-d69390a9fe9e','Herbert Olga','Avenida 9 de Julho, 340','Engineer','1982-11-10','herbertolga@gmail.com.ca'),
	 ('519676ad-e360-4bdf-8132-1712d0c18bb9','Isabela Cristina','Rua Oscar Freire, 290','Teacher','1997-05-05','isabela112@gmail.com');

-- data for table itinerary
INSERT INTO public.itinerary ("id","destination","departure","arrival","transport_kind") VALUES
	 ('343d27ac-ad5e-400c-a735-9ebd97374e00','São Paulo','2023-12-20',NULL,'bus'),
	 ('82881cf8-2450-4bb1-88cb-e11f16f9b301','Belo Horizonte','2023-12-20',NULL,'bus'),
	 ('367aa41d-7e17-4f39-b661-0c48a8ceec59','Recife','2023-12-20',NULL,'plane'),
	 ('2adc5e80-8a9b-4f88-a05d-5c46e9391569','Curitiba','2023-12-20',NULL,'bus');
