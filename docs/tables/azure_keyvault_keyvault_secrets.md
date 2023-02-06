# Table: azure_keyvault_keyvault_secrets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| location | string | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| type | string | X | √ |  | 
| azure_keyvault_keyvault_selefra_id | string | X | X | fk to azure_keyvault_keyvault.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 


