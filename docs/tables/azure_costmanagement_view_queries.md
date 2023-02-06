# Table: azure_costmanagement_view_queries

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| e_tag | string | X | √ |  | 
| location | string | X | √ |  | 
| name | string | X | √ |  | 
| azure_costmanagement_views_selefra_id | string | X | X | fk to azure_costmanagement_views.selefra_id | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 
| sku | string | X | √ |  | 
| tags | json | X | √ |  | 


