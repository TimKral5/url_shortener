# URL Shortener - Concept

The URL shortener should shorten any given URL into parts of its hash
value (about ten digits as string), which then can be used to access
the resources that the full URL refers to.

## Technologies

- **Programming Language:** Go
- **Database (persistent):** MariaDB
- **Database (cache):** Memcached

## Procedures

**Terminology:**

- **Client:** The origin of the request
- **Server:** The URL shortener
- **URL:** The full URL that has to be shortened
- **Hash:** The hash value of the full URL
- **Database:** The persistant storage for all data
- **Cache:** The temporary storage for efficiency
- **Short:** The shortened URL

### Create a new Entry

1. The **client** sends a request to the **server**, containing the
  full **URL** (e.g. `POST https://short.example.com/new`).
1. The **server** receives and parses the request data.
1. The **server** produces the **hash** value of the full **URL**.
1. The **hash** and the full **URL** are then stored to the
    **database**.
1. If the **database** finds that the entry already exists, simply
  return the **hash** along with a **status** message.
1. Lastly, the result of the request, along with the **hash** are
  returned.

### Fetch an Entry

1. The **client** sends a request using the **short** URL
  (e.g. `GET https://api.example.com/AABBCCDDEF`).
1. The **server** receives and parses the request data.
1. The **server** tries to query the data from **cache**.
    - If successful, the queried data is returned.
    - Then, the `usage_count` counter of the of the cache entry is
      increased.
    - The `last_used` field is updated as well.
    - *The request ends here.*
1. The **server** tries to query the data from **database**.
    - If the request fails and the entry is not found, return a
      status message.
    - *The request ends here.*
1. The full **URL** is returned or setup as a redirect.
1. The data is stored in the **cache** for quick access.

## Data Structures

### Database

```plain
User:
  id        INT
  username  STRING

Url:
  id        INT
  short     STRING
  full      STRING

UrlMetadata:
  id            INT
  fk_url        INT
  fk_author     INT
```

### Cache

```plain
Url:
  id            INT
  short         STRING
  full          STRING
  usage_count   INT
  last_used     DATETIME
```

