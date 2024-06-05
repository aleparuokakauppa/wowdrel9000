# wowdrel9000
A wordle clone

## Server API documentation

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
            "match": false,
            "close": true,
            "miss": false
        },
        {
            "char": "R",
            "match": true,
            "close": true,
            "miss": false
        },
        {
            "char": "E",
            "match": false,
            "close": false,
            "miss": true
        },
        {
            "char": "E",
            "match": false,
            "close": true,
            "miss": true
        },
        {
            "char": "S",
            "match": false,
            "close": true,
            "miss": true
        },
    ]
}
```
