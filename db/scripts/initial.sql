CREATE TABLE urls (
	hash varchar(100) NOT NULL,
	original_url varchar(2083) NOT NULL,
	CONSTRAINT urls_pkey PRIMARY KEY (hash)
);
