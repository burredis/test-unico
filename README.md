# test-unico

## Descrição
RESTful API em Golang e SQLite

## Utilização
Para utilizar o projeto é preciso realizar os passos a seguir:

#### Clone

    git clone git@github.com:burredis/test-unico.git

#### Start

    make run

### API
Exemplos de utilização e retorno:

##### GET /feiraslivres
Pesquisa feiras livres que contenham o texto em algum dessas propriedades: distrito|região|nome|bairro
```bash
curl --request GET http://localhost:8000/feiraslivres?q=VILA%20SONIA
```
```json
{
    "data": [
        {
            "id": 12,
            "nome": "VILA SONIA",
            "lat": "-23592640",
            "lng": "-46731043",
            "setcens": "355030894000165",
            "areap": "3550308005179",
            "coddist": "96",
            "distrito": "VILA SONIA",
            "codsubpref": "10",
            "subprefe": "BUTANTA",
            "regiao5": "Oeste",
            "regiao8": "Oeste",
            "registro": "6034-8",
            "logradouro": "AV GAL FRANCISCO MORAZAN",
            "numero": "",
            "bairro": "VL SONIA",
            "referencia": "PROXIMO DISTRITO POLICIAL"
        },
        {
            "id": 189,
            "nome": "VILA REGINA",
            "lat": "-23596387",
            "lng": "-46758441",
            "setcens": "355030894000067",
            "areap": "3550308005181",
            "coddist": "96",
            "distrito": "VILA SONIA",
            "codsubpref": "10",
            "subprefe": "BUTANTA",
            "regiao5": "Oeste",
            "regiao8": "Oeste",
            "registro": "1153-3",
            "logradouro": "RUA EDWARD CARMILO",
            "numero": "S/N",
            "bairro": "JD CELESTE",
            "referencia": ""
        },
        {
            "id": 445,
            "nome": "JARDIM CELESTE",
            "lat": "-23597146",
            "lng": "-46751961",
            "setcens": "355030894000034",
            "areap": "3550308005181",
            "coddist": "96",
            "distrito": "VILA SONIA",
            "codsubpref": "10",
            "subprefe": "BUTANTA",
            "regiao5": "Oeste",
            "regiao8": "Oeste",
            "registro": "4096-7",
            "logradouro": "RUA EDWARD CARMILLO",
            "numero": "",
            "bairro": "JD CELESTE",
            "referencia": "CLUBE SOLAR DOS AMIGOS"
        },
        {
            "id": 497,
            "nome": "MARIO MOURA ANDRADE",
            "lat": "-23600215",
            "lng": "-46741853",
            "setcens": "355030894000110",
            "areap": "3550308005179",
            "coddist": "96",
            "distrito": "VILA SONIA",
            "codsubpref": "10",
            "subprefe": "BUTANTA",
            "regiao5": "Oeste",
            "regiao8": "Oeste",
            "registro": "7071-8",
            "logradouro": "AV CLAUDIO FRANCHI",
            "numero": "49.000000",
            "bairro": "BUTANTA",
            "referencia": "AV PROF FRANCISCO MORATO"
        },
        {
            "id": 503,
            "nome": "JARDIM COLOMBO",
            "lat": "-23598735",
            "lng": "-46732209",
            "setcens": "355030894000043",
            "areap": "3550308005178",
            "coddist": "96",
            "distrito": "VILA SONIA",
            "codsubpref": "10",
            "subprefe": "BUTANTA",
            "regiao5": "Oeste",
            "regiao8": "Oeste",
            "registro": "4097-5",
            "logradouro": "RUA SALOMAO WAIMBERG",
            "numero": "145.000000",
            "bairro": "JD BONFIGLIOLI",
            "referencia": "RUA DR SILVIO DANTE BERTACHI"
        },
        {
            "id": 532,
            "nome": "VILA MORSE",
            "lat": "-23594908",
            "lng": "-46728718",
            "setcens": "355030894000084",
            "areap": "3550308005179",
            "coddist": "96",
            "distrito": "VILA SONIA",
            "codsubpref": "10",
            "subprefe": "BUTANTA",
            "regiao5": "Oeste",
            "regiao8": "Oeste",
            "registro": "1029-4",
            "logradouro": "RUA IBIAPABA",
            "numero": "",
            "bairro": "VL SONIA",
            "referencia": "RUA MANOEL JACINTO"
        },
        {
            "id": 789,
            "nome": "PORTAL DO MORUMBI",
            "lat": "-23614128",
            "lng": "-46742590",
            "setcens": "355030894000190",
            "areap": "3550308005178",
            "coddist": "96",
            "distrito": "VILA SONIA",
            "codsubpref": "10",
            "subprefe": "BUTANTA",
            "regiao5": "Oeste",
            "regiao8": "Oeste",
            "registro": "6104-2",
            "logradouro": "RUA ISAAC ALBENIZ C/ FRANCO ALFANO",
            "numero": "",
            "bairro": "MORUMBI",
            "referencia": "AV JUARES TAVORA"
        }
    ]
}
```


##### GET /feiraslivres/:id
Recupera a feira livre de acordo com o id informado
```bash
curl --request GET http://localhost:8000/feiraslivres/100
```
```json
{
    "data": {
        "id": 100,
        "nome": "CONGONHAS",
        "lat": "-23623517",
        "lng": "-46664328",
        "setcens": "355030815000060",
        "areap": "3550308005100",
        "coddist": "15",
        "distrito": "CAMPO BELO",
        "codsubpref": "14",
        "subprefe": "SANTO AMARO",
        "regiao5": "Sul",
        "regiao8": "Sul 2",
        "registro": "5071-7",
        "logradouro": "AV INVERNADA",
        "numero": "351.000000",
        "bairro": "CONGONHAS",
        "referencia": "ENTRE BARAO REGO BARROS"
    }
}
```

##### POST /feiraslivres
Salva a feira livre na base de dados
```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '"{"lat":"-46550164","lng":"-23558733","setcens":"355030885000091","areap":"3550308005040","coddist":"87","distrito":"VILA FORMOSA","codsubpref":"26","subprefe":"ARICANDUVA-FORMOSA-CARRAO","regiao5":"Leste","regiao8":"Leste 1","nome":"VILA FORMOSA","registro":"4041-0","logradouro":"RUA MARAGOJIPE","numero":"S/N","bairro":"VL FORMOSA","referencia":"TV RUA PRETORIA"}"' \
  http://localhost:8000/feiraslivres
```
```json
{}
```

##### PATCH /feiraslivres/:id
Altera os dados da feira livre atráves do id informado
```bash
curl --header "Content-Type: application/json" \
  --request PATCH \
  --data '"{"lat":"-46550164","lng":"-23558733","setcens":"355030885000091","areap":"3550308005040","coddist":"87","distrito":"VILA FORMOSA","codsubpref":"26","subprefe":"ARICANDUVA-FORMOSA-CARRAO","regiao5":"Leste","regiao8":"Leste 1","nome":"VILA FORMOSA","registro":"4041-0","logradouro":"RUA MARAGOJIPE","numero":"S/N","bairro":"VL FORMOSA","referencia":"TV RUA PRETORIA"}"' \
  http://localhost:8000/feiraslivres/100
```
```json
{}
```

##### DELETE /feiraslivres/:id
Remove a feira livre da base de dados atráves do id informado
```bash
curl --header "Content-Type: application/json" \
  --request PATCH \
  --data '"{"lat":"-46550164","lng":"-23558733","setcens":"355030885000091","areap":"3550308005040","coddist":"87","distrito":"VILA FORMOSA","codsubpref":"26","subprefe":"ARICANDUVA-FORMOSA-CARRAO","regiao5":"Leste","regiao8":"Leste 1","nome":"VILA FORMOSA","registro":"4041-0","logradouro":"RUA MARAGOJIPE","numero":"S/N","bairro":"VL FORMOSA","referencia":"TV RUA PRETORIA"}"' \
  http://localhost:8000/feiraslivres/100
```
```json
{}
```

### Documentação
http://localhost:8000/swagger/index.html

## Teste
Para executar os testes unitários utilize os comandos a seguir:

##### io

    go test -v unico/app/helper/io

##### feiralivre

    go test -v unico/feiralivre

### REF.:

https://www.thepolyglotdeveloper.com/2017/04/using-sqlite-database-golang-application/

https://dev.to/fevziomurtekin/using-sqlite-in-go-programming-3g2c

https://eltonminetto.dev/post/2020-06-29-clean-architecture-2anos-depois/

https://www.restapiexample.com/golang-tutorial/write-log-files-in-golang/

https://viblo.asia/p/document-golang-restful-api-with-swagger-OeVKB17Q5kW