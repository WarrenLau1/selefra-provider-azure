# Table: azure_sql_server_database_blob_auditing_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| azure_sql_server_databases_selefra_id | string | X | X | fk to azure_sql_server_databases.selefra_id | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 


