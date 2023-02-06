# Table: azure_sql_server_blob_auditing_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 


