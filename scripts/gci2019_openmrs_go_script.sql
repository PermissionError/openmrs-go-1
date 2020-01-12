CREATE TABLE patient
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(100),
    illness        VARCHAR(200),
    birth_date     DATE,
    location_id    INT NOT NULL,
    FOREIGN KEY (location_id) REFERENCES location (id)
);

CREATE TABLE location
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100),
    hospital_id INT,
    FOREIGN KEY (hospital_id) REFERENCES hospital (id)
);

CREATE TABLE hospital
(
    id                 SERIAL PRIMARY KEY,
    name               VARCHAR(100),
    max_patient_amount INT
);

insert into hospital (name, max_patient_amount) values ('OpenHospital', 50);
insert into hospital (name, max_patient_amount) values ('OtherHospital', 76);

insert into location (name, hospital_id) values ('Registration Desk',1);
insert into location (name, hospital_id) values ('Inpatient',1);

insert into location (name, hospital_id) values ('Outpatient',2);
insert into location (name, hospital_id) values ('Emergency',2);

insert into patient (name, illness, birth_date, location_id) values ('John One Doe', 'Fever', '1978-11-01', 1);
insert into patient (name, illness, birth_date, location_id) values ('John Two Doe', 'Cold', '1977-10-08', 1);
insert into patient (name, illness, birth_date, location_id) values ('John Three Doe', 'Stomach Cancer', '1969-07-17', 2);
insert into patient (name, illness, birth_date, location_id) values ('John Four Doe', 'Type I Diabetes', '1966-06-12', 2);
insert into patient (name, illness, birth_date, location_id) values ('John Five Doe', 'Lead Poisoning', '1989-10-01', 3);
insert into patient (name, illness, birth_date, location_id) values ('John Six Doe', 'Heroin Overdose', '1974-11-03', 6);
insert into patient (name, illness, birth_date, location_id) values ('John Seven Doe', 'Severe Death', '1999-01-21', 5);
insert into patient (name, illness, birth_date, location_id) values ('John Eight Doe', 'Too Old', '1901-01-01', 6);
insert into patient (name, illness, birth_date, location_id) values ('John Nine Doe', 'Being Undead', '1600-01-01', 6);
