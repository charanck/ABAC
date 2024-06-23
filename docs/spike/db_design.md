## Resource

| Field       | Type   | Constraint            | Description                                                                                           |
| ----------- | ------ | --------------------- | ----------------------------------------------------------------------------------------------------- |
| Id          | string | primary key           |                                                                                                       |
| Name        | string | required              | Name of the resource                                                                                  |
| Owner_id    | string | required              | It is the owner of the resource that is the service or app with which the resource is associated with |
| Policy_id   | string | foreign key, required | Policy for the resource                                                                               |
| Description | string |                       | Details about the resource                                                                            |
| Updated     | date   |                       | Last updated time                                                                                     |
| Deleted     | date   |                       | Deleted on time                                                                                       |
| Created     | date   |                       | Created on time                                                                                       |

## Policy

| Field       | Type   | Constraint  | Description                                                               |
| ----------- | ------ | ----------- | ------------------------------------------------------------------------- |
| Id          | string | primary key |                                                                           |
| Name        | string | required    | Name of the policy                                                        |
| Action      | string | required    | Type of action this policy if for(ALL, READ, WRITE, UPDATE, LIST, DELETE) |
| Policy      | string | required    | Rego policy in string format                                              |
| Description | string |             | Details about the policy                                                  |
| Updated     | date   |             | Last updated time                                                         |
| Deleted     | date   |             | Deleted on time                                                           |
| Created     | date   |             | Created on time                                                           |

## **Policy_input**

| Field     | Type   | Constraint                       | Description                      |
| --------- | ------ | -------------------------------- | -------------------------------- |
| Id        | string | primary key                      |                                  |
| Policy_id | string | foreign key, required            | Parent policy of this input      |
| Name      | string | required, unique for each policy | Name of the policy input         |
| Type      | string | required                         | Type of the value for this input |
| Requried  | bool   | requried                         | If the input is optional        |

## **User**

| Field       | Type   | Constraint | Description                            |
| ----------- | ------ | ---------- | -------------------------------------- |
| Id          | string | primay key |                                        |
| Name        | string | required   | Name of the user                       |
| External_id | string |            | Id from the external service or client |
| Updated     | date   |            | Last updated time                      |
| Deleted     | date   |            | Deleted on time                        |
| Created     | date   |            | Created on time                        |

## Role

| Field   | Type   | Constraint       | Description       |
| ------- | ------ | ---------------- | ----------------- |
| id      | string | primary key      |                   |
| Name    | string | requried, unique | Name of the role  |
| Updated | date   |                  | Last updated time |
| Deleted | date   |                  | Deleted on time   |
| Created | date   |                  | Created on time   |

## Role_user_mapping

| Field   | Type   | Constraint  | Description                     |
| ------- | ------ | ----------- | ------------------------------- |
| id      | string | primary key |                                 |
| Role_id | string | required    | Role to be mapped for the user  |
| User_id | string | required    | User to be mapped with the role |
| Updated | date   |             | Last updated time               |
| Deleted | date   |             | Deleted on time                 |
| Created | date   |             | Created on time                 |

## Attribute

| Field         | Type   | Constraint       | Description                |
| ------------- | ------ | ---------------- | -------------------------- |
| id            | string | primary key      |                            |
| Name          | string | required, Unique | Name of the attribute      |
| String_value  | string |                  | Value if type is string    |
| Integer_value | int    |                  | Value if type is integer   |
| Float_value   | float  |                  | Value if type is float     |
| Bool_value    | bool   |                  | Value if type is bool      |
| Date_value    | date   |                  | Value if type is date      |
| Type          | string | required         | data type of the attribute |
| Updated       | date   |                  | Last updated time          |
| Deleted       | date   |                  | Deleted on time            |
| Created       | date   |                  | Created on time            |

## Role_Attribute_mapping

| Field        | Type   | Constraint            | Description                          |
| ------------ | ------ | --------------------- | ------------------------------------ |
| id           | string | primary key           |                                      |
| Role_id      | string | foreign key, required | Role to be mapped with the attribute |
| Attribute_id | string | foreign key, requried | Attribute to be mapped with the role |
| Created      | date   |                       | Created on time                      |
| Updated      | date   |                       | Last updated time                    |
| Deleted      | date   |                       | Deleted on time                      |

## Resource_Attribute_mapping

| Field        | Type   | Constraint            | Description                              |
| ------------ | ------ | --------------------- | ---------------------------------------- |
| id           | string | primary key           |                                          |
| Resource_id  | string | foreign key, required | Resource to be mapped with the attribute |
| Attribute_id | string | foreign key, required | Attribute to be mapped with the resource |
| Created      | date   |                       | Created on time                          |
| Updated      | date   |                       | Last updated time                        |
| Deleted      | date   |                       | Deleted on time                          |

## User_Attribute_mapping

| Field        | Type   | Constraint            | Description                          |
| ------------ | ------ | --------------------- | ------------------------------------ |
| id           | string | primay key            |                                      |
| User_id      | string | foreign key, required | User to be mapped with the attribute |
| Attribute_id | string | foreign key, required | Attribute to be mapped with the user |
| Created      | date   |                       | Created on time                      |
| Updated      | date   |                       | Last updated time                    |
| Deleted      | date   |                       | Deleted on time                      |
