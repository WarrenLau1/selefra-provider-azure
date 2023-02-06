# Table: azure_compute_virtual_machine_extensions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| azure_compute_virtual_machines_selefra_id | string | X | X | fk to azure_compute_virtual_machines.selefra_id | 
| location | string | X | √ |  | 
| properties | json | X | √ |  | 
| tags | json | X | √ |  | 
| id | string | X | √ |  | 


