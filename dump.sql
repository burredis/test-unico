DROP TABLE IF EXISTS feiraslivres;
CREATE TABLE IF NOT EXISTS feiraslivres(id INTEGER PRIMARY KEY AUTOINCREMENT,lng TEXT, lat TEXT, setcens TEXT, areap TEXT, coddist TEXT, distrito TEXT, codsubpref TEXT, subprefe TEXT, regiao5 TEXT, regiao8 TEXT, nome TEXT, registro TEXT, logradouro TEXT, numero TEXT, bairro TEXT, referencia TEXT);
CREATE INDEX ix_search ON feiraslivres(distrito, regiao5, nome, bairro); 