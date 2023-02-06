# Table: azure_sql_managed_instance_encryption_protectors

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| azure_sql_managed_instances_selefra_id | string | X | X | fk to azure_sql_managed_instances.selefra_id | 
| selefra_id | string | √ | √ | random id | 


