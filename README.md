# wowdrel9000
A wordle clone

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
