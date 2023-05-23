# Rest API Go

## Required `go` and `docker`

### Endpoints

- GET `http://localhost:3000/facts`

- FindFacts
```bash
[
	{
		"id": 1,
		"created_at": "2023-05-23T04:56:09.414173Z",
		"updated_at": "2023-05-23T04:56:09.414173Z",
		"question": "What language do you speak in brazil?",
		"answer": "Language portuguese"
	},
	{
		"id": 2,
		"created_at": "2023-05-23T05:15:34.510283Z",
		"updated_at": "2023-05-23T05:15:34.510283Z",
		"question": "What language do you speak in USA?",
		"answer": "Language english"
	}
]
```

- POST `http://localhost:3000/facts`
- CreateFact
```bash
{
	"question": "What language do you speak in USA?",
	"answer": "Language english"
}
```