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

## ➤ API Reference

### Get latest results info

```http
GET /?url=https://www.mywebsite.com/my-page/
```

#### Get latest results parameters

| Name    | Type      | Located in | Description                                                                                  |
|:--------|:----------|:-----------|:---------------------------------------------------------------------------------------------|
| `url`   | `string`  | query      | **Required**. This is the url of the page from which you want to retrieve the latest results |
| `badge` | `boolean` | query      | If you want to get the badge in a SVG format (default is `false`)                            |

#### Get latest results responses

| Code | Description                 | Model                                                  |
|------|-----------------------------|--------------------------------------------------------|
| 200  | There are results in the DB | [LatestResultResponse](#latestresultresponse) / String |

### Add a new analysis to the tasks queue

This is an alias of the [ecoindex API](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/cnumr/ecoindex_api/main/docs/openapi.json#tag/Tasks/operation/Add_new_ecoindex_analysis_task_to_the_waiting_queue_v1_tasks_ecoindexes_post) Create a new task endpoint

### Get the result of a task

This is an alias of the [ecoindex API](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/cnumr/ecoindex_api/main/docs/openapi.json#tag/Tasks/operation/Get_ecoindex_analysis_task_by_id_v1_tasks_ecoindexes__id__get) Get the result of a task endpoint

### Get the screenshot of a ecoindex result

This is an alias of the [ecoindex API](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/cnumr/ecoindex_api/main/docs/openapi.json#tag/Ecoindex/operation/Get_screenshot__version__ecoindexes__id__screenshot_get) Get screenshot of a ecoindex result endpoint

### Models

#### Result

| Name       | Type     | Description                               |
|:-----------|:---------|:------------------------------------------|
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
|:----------------|:--------------------|:--------------------------------------------------|
| `count`         | `int`               | Number of total results existing for this website |
| `latest-result` | [Result](#result)   | Latest result for this exact webpage              |
| `older-results` | [Result](#result)[] | Older results for the same webpage                |
| `other-results` | [Result](#result)[] | Other results tor this website                    |

## [License](LICENSE)

## [Contributing](CONTRIBUTING.md)

## [Code of conduct](CODE_OF_CONDUCT.md)
