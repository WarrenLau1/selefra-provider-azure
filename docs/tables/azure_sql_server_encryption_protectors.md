# Table: azure_sql_server_encryption_protectors

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| location | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 


