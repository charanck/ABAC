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

## Policy

| Field       | Type   | Constraint  | Description                  |
| ----------- | ------ | ----------- | ---------------------------- |
| Id          | string | primary key |                              |
| Name        | string | required    | Name of the policy           |
| Policy      | string | required    | Rego policy in string format |
| Description | string |             | Details about the policy     |
| Updated     | date   |             | Last updated time            |
| Deleted     | date   |             | Deleted on time              |

## **Policy_input**

| Field     | Type   | Constraint                    | Description                      |
| --------- | ------ | ----------------------------- | -------------------------------- |
| Id        | string | primary key                   |                                  |
| Policy_id | string | foreign key, required         | Parent policy of this input      |
| Name      | string | foreign key, required, unique | Name of the policy               |
| Type      | string | required                      | Type of the value for this input |

## **User**

| Field       | Type   | Constraint | Description                            |
| ----------- | ------ | ---------- | -------------------------------------- |
| Id          | string | primay key |                                        |
| Name        | string | required   | Name of the user                       |
| External_id | string |            | Id from the external service or client |
| Updated     | date   |            | Last updated time                      |
| Deleted     | date   |            | Deleted on time                        |

## Role

| Field   | Type   | Constraint       | Description       |
| ------- | ------ | ---------------- | ----------------- |
| id      | string | primary key      |                   |
| Name    | string | requried, unique | Name of the role  |
| Updated | date   |                  | Last updated time |
| Deleted | date   |                  | Deleted on time   |

## Role_user_mapping

| Field   | Type   | Constraint  | Description                     |
| ------- | ------ | ----------- | ------------------------------- |
| id      | string | primary key |                                 |
| Role_id | string | required    | Role to be mapped for the user  |
| User_id | string | required    | User to be mapped with the role |

## Attribute

| Field   | Type   | Constraint       | Description                       |
| ------- | ------ | ---------------- | --------------------------------- |
| id      | string | primary key      |                                   |
| Name    | string | required, Unique | Name of the attribute             |
| Value   | string | required         | base64 enconded data of the value |
| type    | string | required         | data type of the attribute        |
| Updated | date   |                  | Last updated time                 |
| Deleted | date   |                  | Deleted on time                   |

## Role_Attribute_mapping

| Field        | Type   | Constraint            | Description                          |
| ------------ | ------ | --------------------- | ------------------------------------ |
| id           | string | primary key           |                                      |
| Role_id      | string | foreign key, required | Role to be mapped with the attribute |
| Attribute_id | string | foreign key, requried | Attribute to be mapped with the role |

## Resource_Attribute_mapping

| Field        | Type   | Constraint            | Description                              |
| ------------ | ------ | --------------------- | ---------------------------------------- |
| id           | string | primary key           |                                          |
| Resource_id  | string | foreign key, required | Resource to be mapped with the attribute |
| Attribute_id | string | foreign key, requried | Policy to be mapped with the resource    |

## User_Attribute_mapping

| Field        | Type   | Constraint            | Description                          |
| ------------ | ------ | --------------------- | ------------------------------------ |
| id           | string | primay key            |                                      |
| User_id      | string | foreign key, required | User to be mapped with the attribute |
| Attribute_id | string | foreign key, required | Attribute to be mapped with the user |
