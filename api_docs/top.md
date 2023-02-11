# TopPage APT

## TOPページ取得 GET /top

Retrieve the articles

```json
{
  "status": 200,
  "data": {
    "articles": [
      {
        "id": 1,
        "userId": 1,
        "thumbnailUrl": "https://example.com/thumbnail.jpg",
        "title": "example_title",
        "content": "example_content",
        "status": 2,
        "createdAt": "2023-01-28T21:30:17.957546+09:00",
        "updatedAt": "2023-01-28T21:30:17.957546+09:00",
        "categoryId": 1,
        "user": {
          "id": 1,
          "name": "example_name",
          "password": "$2a$10$s0IkhGD3R9qmwZ8/afJbP..uKGPNGl/ObUrVH8J2j181uk0KTfJ3q",
          "email": "example@mail.com",
          "createdAt": "2023-01-28T21:30:17.919039+09:00"
        },
        "category": {
          "id": 1,
          "name": "example_category",
          "description": "example_description",
          "parentId": 0,
          "createdAt": "2023-01-28T21:30:17.985704+09:00"
        },
        "tags": [
          {
            "id": 1,
            "name": "example_tag",
            "description": "example_description",
            "createdAt": "2023-01-28T21:30:18.005803+09:00"
          },
          {
            "id": 2,
            "name": "example_tag",
            "description": "example_description",
            "createdAt": "2023-01-28T21:30:18.005803+09:00"
          }
        ]
      }
    ],
    "pager": {
      "currentPage": 1,
      "lastPage": 3
    }
  }
}
```
