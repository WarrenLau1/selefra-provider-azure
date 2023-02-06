# Table: azure_postgresql_server_firewall_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| azure_postgresql_servers_selefra_id | string | X | X | fk to azure_postgresql_servers.selefra_id | 


