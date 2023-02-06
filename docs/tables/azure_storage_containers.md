# Table: azure_storage_containers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| properties | json | X | √ |  | 
| etag | string | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| azure_storage_accounts_selefra_id | string | X | X | fk to azure_storage_accounts.selefra_id | 


