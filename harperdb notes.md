## Get Schema
- Provides information about all the tables from the database - *dev*. If ```schema``` is not provided in the body the default is considered as ```data```
```json
{
    "operation": "describe_schema",
    "schema": "dev"   
}
```

## Describe All
- Provides information about all the database and tables in it.
```json
{
    "operation": "describe_all"
}
```

## Create schema
- Creates empty database
```json
{
    "operation": "create_schema",
    "schema_name": "my_app"
}
```

## Create table
- Creates table in the database. Default database name is ```data```
```json
{
    "operation": "create_table",
    "table": "table_name",
    "primary_key": "id_name",
    "database": "data"
}
```

## Insert Data
- Insert data into table
```json
{
    "operation": "insert",
    "database": "data",
    "table": "call-forwards",
    "records": [
        {
            "id": 1,
            "phoneNumber": "+123456",
            "status": true
        },
        {
            "id": 2,
            "phoneNumber": "+1234567",
            "status": true
        }
    ]
}
```

## Get records
