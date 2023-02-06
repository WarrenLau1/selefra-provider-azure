# Table: azure_sql_server_database_threat_protections

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| system_data | json | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| azure_sql_server_databases_selefra_id | string | X | X | fk to azure_sql_server_databases.selefra_id | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 


