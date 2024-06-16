# wowdrel9000
A simple wordle clone with server-client architecture.

# Goal
- Learn how to use client-side javascript

## Screenshot
![Screenshot](https://github.com/aleparuokakauppa/wowdrel9000/blob/main/data/screenshot.jpg?raw=true)

## Server API docs
**Endpoint**: `/guess`
**Method**: `POST`
**Description**: Post a JSON request to the server. Expect a JSON response.
**Request example**:
```json
{
    "version": 1,
    "guess": "TREES"
}
```
**Response example**:
```json
{
    "version": 1,
    "win": false,
    "letters": [
        {
            "char": "T",
            "status": "miss"
        },
        {
            "char": "R",
            "status": "match"
        },
        {
            "char": "E",
            "status": "close"
        },
        {
            "char": "E",
            "status": "close"
        },
        {
            "char": "S",
            "status": "miss"
        },
    ]
}
```
