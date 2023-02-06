# Table: azure_sql_server_admins

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 


