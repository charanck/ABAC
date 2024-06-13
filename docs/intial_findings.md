# Rough Draft:

Attribute based access control determine the access of the user based on the attributes of the **user, resource and the environment**

## Types of attributes

* User attributes - Eg: user's role, user's age, user's department etc...
* Resource attributes - Eg: visibility, owner etc..
* Environment attributes - Eg: time, physical location, device, network etc...

## Policy

    Policies are used to evalute the attributes and determine the access. Policies are used to mix multiple attribute checks to enforce security

### Policy language

    In this project the policies are expressed in a high-level declarative language called Rego, it is built for expressing policies over complex hierarchical data structure.

## Further analysis required for?

1. [X] Database schema
2. [ ] How to store the attributes for resource, user and environment
3. [ ] How generate rego policies dynamically
