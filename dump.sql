CREATE TABLE IF NOT EXISTS feiraslivres (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  lng VARCHAR(10),
  lat VARCHAR(10),
  setcens VARCHAR(10), 
  areap VARCHAR(10),
  coddist VARCHAR(10),
  distrito VARCHAR(30),
  codsubpref VARCHAR(10),
  subprefe VARCHAR(30),
  regiao5 VARCHAR(10),
  regiao8 VARCHAR(30),
  nome VARCHAR(60),
  registro VARCHAR(10),
  logradouro VARCHAR(60),
  numero VARCHAR(10),
  bairro VARCHAR(60),
  referencia VARCHAR(255)
);

CREATE INDEX ix_search ON feiraslivres(distrito,regiao5,nome,bairro);

-- LOAD DATA LOCAL INFILE './database/DEINFO_AB_FEIRASLIVRES_2014.csv'
-- INTO TABLE feiraslivres
--   FIELDS TERMINATED BY ','
--   LINES TERMINATED BY '\n'
--   IGNORE 1 LINES
--   (id,lng,lat,setcens,areap,coddist,distrito,codsubpref,subprefe,regiao5,regiao8,nome,registro,logradouro,numero,bairro,referencia);

CREATE VIRTUAL TABLE feiraslivres USING csv('./DEINFO_AB_FEIRASLIVRES_2014.csv', USE_HEADER_ROW, NO_QUOTE);