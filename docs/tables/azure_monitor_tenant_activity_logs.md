# Table: azure_monitor_tenant_activity_logs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| caller | string | X | √ |  | 
| category | json | X | √ |  | 
| http_request | json | X | √ |  | 
| resource_provider_name | json | X | √ |  | 
| sub_status | json | X | √ |  | 
| properties | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| claims | json | X | √ |  | 
| id | string | X | √ |  | 
| level | string | X | √ |  | 
| operation_id | string | X | √ |  | 
| operation_name | json | X | √ |  | 
| authorization | json | X | √ |  | 
| event_data_id | string | X | √ |  | 
| event_timestamp | timestamp | X | √ |  | 
| resource_group_name | string | X | √ |  | 
| resource_id | string | X | √ |  | 
| submission_timestamp | timestamp | X | √ |  | 
| tenant_id | string | X | √ |  | 
| correlation_id | string | X | √ |  | 
| description | string | X | √ |  | 
| event_name | json | X | √ |  | 
| resource_type | json | X | √ |  | 
| status | json | X | √ |  | 


