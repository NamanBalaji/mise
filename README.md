# Mise

##### Mise is an in-memory key-value caching server that runs on top of HTTP

## Features
- Current supports the following:
   - All basic data types
   - Arrays
   - Linked Lists

- All  basic data types support SET GET and DELETE
- Arrays support `SET` `GET` `GET_RANGE` `DELETE_INDEX` `ADD`
- Linked Lists supports `SET_LIST` `GET_FIRST` `GET_LAST` `SET_FIRST` `SET_LAST` `DELETE_FIRST` `DELETE_LAST`. Linked lists are intended to be used as stacks and Queues
- Sorted Sets supports `Set_SORTED_SET`, `GET_MIN`, `GET_MAX`, `ADD`, `DELETE_MIN` `DELETE_MAX`. Sorted sets can currently hold only integer types

## Installation
```
Download the app from the release page
```
While running users can set custom port numbers by using the `-p` flag 

## Usage
  - Ping
    ```
    METHOD: GET 
    URL:  http://localhost:<PORT_NUM>/ping
    RESPONSE: Pong
    ```
  - SET: Used to set basic data types and arrays
    ```
    METHOD: POST
    URL:  http://localhost:<PORT_NUM>/set
    BODY: {
        "key": <string>,
        "value": <any value or an array>
    }
    RESPONSE: {
    "message": "OK",
    "status": 0
    }
    ```
  - GET: Used for retrieving basic data typea and arrays
    ```
    METHOD: POST
    URL:  http://localhost:<PORT_NUM>/get
    BODY: {
        "key": <string>,
    }
    RESPONSE: {
    "value": <some value>,
    "status": 0,
    "message": "OK"
    }
    ```
  - DELETE: Used to delete a key-value pair
     ```
     METHOD: POST
    URL:  http://localhost:<PORT_NUM>/delete
    BODY: {
        "key": <string>,
    }
    RESPONSE: {
    "value": <some value>,
    "status": 0,
    "message": "OK"
    }
    ```
 - ADD: Used to append to an array
     ```
     METHOD: POST
    URL:  http://localhost:<PORT_NUM>/add
    BODY: {
        "key": <string>,
        "value": <some value>,
        "index": <int> //put -1 if want to append to the end
    }
    RESPONSE: {
        "message": "OK",
        "status": 0
    }
    ```
- GET_RANGE : Used to get a range from an array
    ```
    METHOD: POST
    URL:  http://localhost:<PORT_NUM>/get-range
    BODY: {
        "key": <string>,
        "start": <int>,
        "stop": <int> //put -1 if want to range till the  end
    }
    RESPONSE: {
    "value": [
        <some value>,
        <some value>, 
        ...
    ],
    "status": 0,
    "message": "OK"
    }
    ```
- DELETE_INDEX: Used to delete value at a particular index from an array
     ```
    METHOD: DELETE
    URL:  http://localhost:<PORT_NUM>/delete-element
    BODY: {
        "key": <string>,
        "index": <int> //put -1 if want to delete the last element
    }
    RESPONSE: {
    "value": <some value,
    "status": 0,
    "message": "OK"
    }
    ```
- SET_LIST: Used to set a key wth a linked list. 
     ```
    METHOD: POST
    URL:  http://localhost:<PORT_NUM>/set-list
    BODY: {
        "key": <string>,
        "value": <some value> // can be a basic data type or an array
    }
    RESPONSE: {
        "message": "OK",
        "status": 0
    }
    ```
- GET_FIRST/GET_LAST: Works similar to a peek function
    ```
    METHOD: POST
    URL:  http://localhost:<PORT_NUM>/get-list
    BODY: {
        "key": <string>,
        "get_first": <bool> // set `true` to peek at the head
    }
    RESPONSE: {
        "value": <some value>
        "message": "OK",
        "status": 0
    }
    ```
- SET_FIRST/SET_LAST: Works similar to a push/offer function
    ```
    METHOD: POST
    URL:  http://localhost:<PORT_NUM>/add-list
    BODY: {
        "key": <string>,
        "value": <some value>,
        "add_first": <bool> //set true if you want to add to the head
    }
    RESPONSE: {
        "message": "OK",
        "status": 0
    }
    ```
- DELETE_FIRST/DELETE_LAST: Works similar to a pop/remove function
     ```
    METHOD: DELETE
    URL:  http://localhost:<PORT_NUM>/delete-list
    BODY: {
        "key": <string>,
        "delete_first": <bool> //set true if you want to delete the head
    }
    RESPONSE: {
        "value": <some value>,
        "message": "OK",
        "status": 0
    }
    ```
- SET SORTED SET: Initializes and appends values to a sorted set
    ```
    METHOD: POST
    URL:  http://localhost:<PORT_NUM>/set-sortedSet
    BODY: {
        "key": <string>,
        "value": <int> // accepts integers and integer arrays
    }
    RESPONSE: {
        "message": "OK",
        "status": 0
    }
    ```
- ADD SORTED SET: Appends a value to a sorted set
    ```
    METHOD: POST
    URL:  http://localhost:<PORT_NUM>/add-sortedSet
    BODY: {
        "key": <string>,
        "value": <int>
    }
    RESPONSE: {
        "message": "OK",
        "status": 0
    }
    ```
- GET SORTED SET: Retrieves the minimum or maximum value from the sorted set
    ```
    METHOD: POST
    URL:  http://localhost:<PORT_NUM>/get-sortedSet
    BODY: {
        "key": <string>,
        "max": <bool> //if true returns thr maximum value
    }
    RESPONSE: {
        "value": <some int>
        "message": "OK",
        "status": 0
    }
    ```

- DELETE SORTED SET: Deletes the minimum or maximum value from the sorted set
    ```
    METHOD: POST
    URL:  http://localhost:<PORT_NUM>/delete-sortedSet
    BODY: {
        "key": <string>,
        "max": <bool> //if true returns thr maximum value
    }
    RESPONSE: {
        "value": <some int>
        "message": "OK",
        "status": 0
    }
    ```

- SIZE: Fetches the size of a key's associated value. Returns the number of elements for arrays, Sorted sets and linked lists.
```
    METHOD: GET
    URL:  http://localhost:<PORT_NUM>/size/<key>
    RESPONSE: {
        "value": <some int>
    }
    ```
    
### Run locally
- Download and install `go`
- Clone this repo and make changes.
- Then run `run.sh`

### Future work
- Add support for Trees
- Persist data after server dies
- Create a CLI tool
- Develop client SDKs 