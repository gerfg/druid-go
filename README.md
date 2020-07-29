# druid-go

# Native Queries

## Scan

Return uma lista de linhas (raw).

    Além do uso simples onde uma consulta Scan é emitida para o Broker, a consulta Scan também pode ser emitida diretamente para processos históricos ou tarefas de ingestão de streaming. Isso pode ser útil se você quiser recuperar grandes quantidades de dados em paralelo.

propriedade | descrição |   obrigatorio?  
------------|-----------|-----------------
queryType |	Este campo precisa ser sempre "scan" (primeiro parâmetro que o druid verifica para processar a query) | sim 
dataSource | fonte dos dados (string ou objeto), bastante similar  uma tabela | sim 
intervals | Objeto JSON representando intervalos no formato ISO-8601. Time range que delimita a query (baseando-se no __time) | sim
resultFormat | representação dos resultados, formatos: list (nomes das colunas e valores), compactedList (apenas valores), default: list | não
filter | | não
columns | Array de strings de dimensões e metricas para scanear.default: todas as colunas | não
batchSize | Numero máximo de linhas em buffer antes de retornar ao cliente (?) | não
limit | número de linhas a se retornar, default: todas as linhas | não
order | Ordenação das linhas baseandos-se pelo timestamp (__time ?). Asc e Desc apenas funcionam quando __time está incluida na lista de colunas | não
legacy | Retorna resultados consistentes com a extensão legacy "scan-query" (?) (legacy mode) | não
context | JSON Object adicional que pode ser utilizado para especificar algumas flags (query context properties)

#### dataSource

Podemos definir o tipo e o nome do dataSource:

```json
{
    "dataSource": {
        "type": "table",
        "name": "wikipedia"
    }
}
```

Ou apenas o nome do dataSource:

```json
{
    "dataSource": "wikipedia"
}
```

### Legacy mode
A scan query suporte o modo legacy projetado para o protocolo de compatibilidade com a antiga extensão "scan query contrib".

### Query context properties

| property |  description | values  | default  |
|---|---|---|---|
| maxRowsQueuedForOrdering  | Número máximo de linhas a ser retornado quando o "time ordering" está sendo utilizado. | inteiro entre: [1,2147483647]  | druid.query.scan.maxRowsQueuedForOrdering  |
| maxSegmentPartitionsOrderedInMemory | Número máximo de segmentos escaneados por histórico quando o "time ordering" está sendo utilizado. | inteiro entre: [1,2147483647] | druid.query.scan.maxSegmentPartitionsOrderedInMemory|

Exemplo:

```json
{
  "maxRowsQueuedForOrdering": 100001,
  "maxSegmentPartitionsOrderedInMemory": 100
}
```

## Timeseries

Este tipo de query utiliza de um "timeseries query object" e retorna um array de objetos JSON em que cada objeto representa um valor solicitado pela query timeseries.

|propriedade | descrição |   obrigatorio?  |
|------------|-----------|-----------------|
|queryType |	Este campo precisa ser sempre "timeseries" (primeiro parâmetro que o druid verifica para processar a query) | sim |
|dataSource | fonte dos dados (string ou objeto), bastante similar  uma tabela | sim |
|intervals | Objeto JSON representando intervalos no formato ISO-8601. Time range que delimita a query (baseando-se no __time) | sim| 
|descending | se deve ordenar o resultado de forma descendente, default: false -> ascendente | não |
| granularity | define a granularidade do "bucket query results", [granularity](https://druid.apache.org/docs/latest/querying/granularities.html) | sim|


## Query filters