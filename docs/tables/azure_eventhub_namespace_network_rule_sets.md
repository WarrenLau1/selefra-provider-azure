# Table: azure_eventhub_namespace_network_rule_sets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| location | string | X | √ |  | 
| name | string | X | √ |  | 
| system_data | json | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| azure_eventhub_namespaces_selefra_id | string | X | X | fk to azure_eventhub_namespaces.selefra_id | 
| properties | json | X | √ |  | 
| id | string | X | √ |  | 


