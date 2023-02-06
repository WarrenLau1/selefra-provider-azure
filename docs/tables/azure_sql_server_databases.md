# Table: azure_sql_server_databases

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| managed_by | string | X | √ |  | 
| type | string | X | √ |  | 
| properties | json | X | √ |  | 
| sku | json | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 
| location | string | X | √ |  | 
| identity | json | X | √ |  | 
| tags | json | X | √ |  | 


