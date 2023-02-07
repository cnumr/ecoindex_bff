# Ecoindex Back For Front

This project aims to provide a simple back for front for [ecoindex API](https://github.com/cnumr/ecoindex_api) project. It is mainly used by the [ecoindex browser plugin](https://github.com/vvatelot/ecoindex-browser-plugin).

It offers a way to retrieve easily the latest results for a given page, and also for the current website

It is built in Golang and Fiber to provide great performance and be as light as possible

## 🛠️ Tech Stack

- [Docker](https://www.docker.com/)
- [Docker Compose v2](https://docs.docker.com/compose/compose-v2/)
- [Golang](https://go.dev/)
- [Fiber](https://gofiber.io/)
- [Air](https://github.com/cosmtrek/air) (live relaod)

## 🛠️ Install Dependencies

```bash
go mod download
```

## 🧑🏻‍💻 Usage

To start the project, you first need configure your `.env` file and provide the url of the ecoindex API you want to reach by setting the environment variable `API_URL`. Default is set to `https://ecoindex.p.rapidapi.com`.

If you use production API url, you have to [request an API key](https://rapidapi.com/cnumr-cnumr-default/api/ecoindex/pricing) on RapidAPI platform. Once you get your API Key, you can set the env variable `API_KEY`

```bash
API_URL=https://ecoindex.p.rapidapi.com        # Or your own server url
API_KEY=your-generated-api-key                     # Optional if not production server

# You can also specify your application listening port (default is 3001)
APP_PORT=1337
```

Then you can launch your project simply using air command:

```bash
air
```

> You can now reach your Back For Front instance on [http://localhost:3001](http://localhost:3001) (regarding the `APP_PORT` you defined...)

## 🔧 Configuration

### Environment variables

| Name           | Description                                                                        | Default value                       |
|----------------|------------------------------------------------------------------------------------|-------------------------------------|
| `API_URL`      | The url of the ecoindex API you want to reach                                      | `"https://ecoindex.p.rapidapi.com"` |
| `API_KEY`      | The API key you want to use to reach the ecoindex API (if production server)       | `""`                                |
| `APP_PORT`     | The port on which the application will listen                                      | `3001`                              |
| `APP_URL`      | The url of the application                                                         | `"http://localhost:3001"`           |
| `ENV`          | The environment in which the application is running (in dev mode, enables logging) | `dev`                               |
| `CACHE_TTL`    | The time to live of the cache (in seconds)                                         | `604800` (1 week)                   |
| `ECOINDEX_URL` | The url of the ecoindex website                                                    | `"https://www.ecoindex.fr"`         |

## ➤ API Reference

### Get latest results info

```http
GET /api/results/?url=https://www.mywebsite.com/my-page/
```

#### Get latest results parameters

| Name  | Type     | Located in | Description                                                                                  |
|-------|----------|------------|----------------------------------------------------------------------------------------------|
| `url` | `string` | query      | **Required**. This is the url of the page from which you want to retrieve the latest results |

#### Get latest results responses

| Code | Description                      | Model                                         |
|------|----------------------------------|-----------------------------------------------|
| 200  | There are results in the DB      | [LatestResultResponse](#latestresultresponse) |
| 400  | The url is not valid             | String                                        |
| 404  | There is no result for this page | [LatestResultResponse](#latestresultresponse) |

### Add a new analysis to the tasks queue

This is an alias of the [ecoindex API](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/cnumr/ecoindex_api/main/docs/openapi.json#tag/Tasks/operation/Add_new_ecoindex_analysis_task_to_the_waiting_queue_v1_tasks_ecoindexes_post) Create a new task endpoint

```http
POST /api/tasks
{
    "url": "https://www.mywebsite.com/my-page/",
    "width": 1920,
    "height": 1080
}
```

### Get the result of a task

This is an alias of the [ecoindex API](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/cnumr/ecoindex_api/main/docs/openapi.json#tag/Tasks/operation/Get_ecoindex_analysis_task_by_id_v1_tasks_ecoindexes__id__get) Get the result of a task endpoint

```http
GET /api/tasks/a7c3d264-62c6-4f45-b1db-51d7db31d085
```

### Get the screenshot of a ecoindex result

This is an alias of the [ecoindex API](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/cnumr/ecoindex_api/main/docs/openapi.json#tag/Ecoindex/operation/Get_screenshot__version__ecoindexes__id__screenshot_get) Get screenshot of a ecoindex result endpoint

```http
GET /api/screenshot/a7c3d264-62c6-4f45-b1db-51d7db31d085
```

### Get Api Health

```http
GET /health
```

#### Get health response

| Code | Description         | Model  |
|------|---------------------|--------|
| 200  | `OK` API is healthy | String |

### Get Ecoindex badge

```http
GET /badge/?url=https://www.mywebsite.com/my-page/
```

#### Get badge parameters

| Name  | Type     | Located in | Description                                                                                  |
|-------|----------|------------|----------------------------------------------------------------------------------------------|
| `url` | `string` | query      | **Required**. This is the url of the page from which you want to retrieve the latest results |

#### Get badge responses

| Code | Description                            | Model  |
|------|----------------------------------------|--------|
| 200  | Badge of the result (format `svg/xml`) | String |
| 400  | The url is not valid                   | String |

### Redirect to ecoindex result page

```http
GET /redirect/?url=https://www.mywebsite.com/my-page/
```

#### Get redirect parameters

| Name  | Type     | Located in | Description                                                                                  |
|-------|----------|------------|----------------------------------------------------------------------------------------------|
| `url` | `string` | query      | **Required**. This is the url of the page from which you want to retrieve the latest results |

#### Get redirect responses

| Code | Description                 | Model  |
|------|-----------------------------|--------|
| 303  | Redirect to the result page | String |
| 400  | The url is not valid        | String |

### Models

#### Result

| Name       | Type     | Description                               |
|------------|----------|-------------------------------------------|
| `date`     | `string` | Date of the result                        |
| `grade`    | `string` | Ecoindex result grade                     |
| `id`       | `string` | Result UUID                               |
| `nodes`    | `int`    | Number of nodes in the DOM of the webpage |
| `requests` | `int`    | Number of requests made by the webpage    |
| `score`    | `int`    | Ecoindex result score                     |
| `size`     | `int`    | Size of the webpage                       |
| `url`      | `string` | Page URL                                  |

#### LatestResultResponse

| Name            | Type                | Description                                       |
|-----------------|---------------------|---------------------------------------------------|
| `count`         | `int`               | Number of total results existing for this website |
| `latest-result` | [Result](#result)   | Latest result for this exact webpage              |
| `older-results` | [Result](#result)[] | Older results for the same webpage                |
| `other-results` | [Result](#result)[] | Other results tor this website                    |

## [License](LICENSE)

## [Contributing](CONTRIBUTING.md)

## [Code of conduct](CODE_OF_CONDUCT.md)
