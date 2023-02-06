# Table: azure_compute_virtual_machine_scale_set_vms

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ |  | 
| zones | string_array | X | √ |  | 
| sku | json | X | √ |  | 
| properties | json | X | √ |  | 
| instance_id | string | X | √ |  | 
| name | string | X | √ |  | 
| resources | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| identity | json | X | √ |  | 
| plan | json | X | √ |  | 
| tags | json | X | √ |  | 
| id | string | X | √ |  | 
| location | string | X | √ |  | 
| azure_compute_virtual_machine_scale_sets_selefra_id | string | X | X | fk to azure_compute_virtual_machine_scale_sets.selefra_id | 


