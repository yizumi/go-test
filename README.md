```
 _   _______  _____ ___________ ___________ 
| | | | ___ \/  ___|     |  _  \  ___| ___ \
| | | | |_/ /\ `--.|     | | | | |__ | |_/ /
| | | |  __/  `--. \     | | | |  __||    / 
| |_| | |    /\__/ /     | |/ /| |___| |\ \ 
 \___/\_|    \____/|_____|___/ \____/\_| \_|
```

Hi there! Welcome to UPSIDER's Coding Challenge!

# Coding Test Instructions

## Introduction
For this coding test, you will need to create a new endpoint called "PUT /users" to update an existing record. In addition, you will need to add a credential mechanism, authorization, and create a model of "UserHistory" which records the previous values if the user data is being modified.

## Requirements
Your solution should:

### Task 1: Create Update Endpoint
* Create a new endpoint called "PUT /users" that accepts PUT requests to update an existing record.
* Return a 200 response with a json of the updated object if the update is successful.
* Return a 404 response if the object is not found.
* Return a 400 response if there is a problem with the json format.

### Task 2: Authorization
* Add a credential/authorization mechanism. If the request has the "X-Update-Key" header with the value "{update_key}" that matches with the record, then it should pass, otherwise, it should return a 403 response.

### Task 3: Create History Record upon Update
* Create a model of "UserHistory" which records the previous values if the user data is being modified.
* Modify the behavior of the PUT request so it records the previous version of the record.

Happy coding!
