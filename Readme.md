[![Docker Repository on Quay](https://quay.io/repository/elastifeed/es-rss/status "Docker Repository on Quay")](https://quay.io/repository/elastifeed/es-rss)
# es-rss
Simple RSS/Atom-Feed parser based on [this awesome go library](https://github.com/mmcdole/gofeed).

## Running
> The image is designed to scale. You can run on localy by any container runtime

e.g.: `podman run -p 8080:8080 quay.io/elastifeed/es-rss:latest`

## Rest API 
> Currently there is only the `/parse` endpoint present.
### Request
```json
{
    "url": "{{ rss_feed_url }}",
    "from_time": "{{ last_scrape_time }}"
}
```
### Response
```json
[
    {
        "title": "Some Title",
        "description": "This is some article about bla...",,
        "url": "Original article URL"
    },
    ...
]
```