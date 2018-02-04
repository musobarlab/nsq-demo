### Simple Demo, How to use NSQ as a Message Queue

### Running
  - Start NSQD, NSQLOOKUPD, and NSQADMIN

  ```shell
  docker-compose up
  ```

  - Start Backend API
    Go to publisher folder

    ```shell
    cd publisher/
    ```
    and

    ```shell
    go run main.go
    ```

    send message like this

    ```curl
      curl -X POST \
      http://localhost:3000/api/send \
      -H 'cache-control: no-cache' \
      -H 'content-type: application/json' \
      -H 'postman-token: d8e4a3cc-fcad-2d07-29f7-ad2f47ba3e66' \
      -d '{
      	"from": "Wuriyanto",
        	"content":{
        		"header": "Makan Makan",
        		"body": "Makan nasi"
        	}
        }'
    ```

    open consumer folder

    ```shell
    go run main.go
    ```

    you'll see the result like this

    ```shell
    2018/02/04 15:49:01 INF    1 [wurys/ch] (127.0.0.1:4150) connecting to nsqd
    2018/02/04 15:49:06 Got a message: {"from":"Wuriyanto","content":{"header":"Makan Makan","body":"Makan nasi"}}
    ```
