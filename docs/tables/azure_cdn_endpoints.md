# Table: azure_cdn_endpoints

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| properties | json | X | √ |  | 
| tags | json | X | √ |  | 
| name | string | X | √ |  | 
| system_data | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| location | string | X | √ |  | 
| id | string | X | √ |  | 
| type | string | X | √ |  | 
| azure_cdn_profiles_selefra_id | string | X | X | fk to azure_cdn_profiles.selefra_id | 


