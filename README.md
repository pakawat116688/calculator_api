# calculator_api
Go Redis Calculator

# lib
- github.com/gofiber/fiber/v2
- github.com/go-redis/redis/v9

# curl
- curl localhost:8000/plus -H 'content-type:application/json' -d '{"setter":5, "action":5}' -X GET
- curl localhost:8000/minus -H 'content-type:application/json' -d '{"setter":5, "action":5}' -X GET
- curl localhost:8000/multiply -H 'content-type:application/json' -d '{"setter":5, "action":5}' -X GET
- curl localhost:8000/divide -H 'content-type:application/json' -d '{"setter":5, "action":5}' -X GET