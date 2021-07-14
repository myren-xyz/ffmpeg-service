# ffmpeg converter service

## prerequisites
`.config.json` to load access tokens

## Endpoints

/api/v1/

- `convert`
- `subscribe` SSE endpoint


### convert
parameters:
- file_url
- - string | e.g: https://google.com/song/test.mp3
- upload_path
- - string | e.g: vdfa4234/audiofy/43290fdsa
- issuer
- - string | e.g: 60c9e35ab324586977a0e697

### subscribe
parameters:
- job_id
- - string | e.g: few46e
