# Table: azure_appservice_web_app_vnet_connections

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| azure_appservice_web_apps_selefra_id | string | X | X | fk to azure_appservice_web_apps.selefra_id | 


