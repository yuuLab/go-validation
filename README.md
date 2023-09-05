# Go Validation example

Example application of Gin and ozzo-validation.

- https://github.com/gin-gonic/gin
- https://github.com/go-ozzo/ozzo-validation

# Trying it out

1. clone
   ```
   git clone https://github.com/yuuLab/go-validation.git
   ```
1. cd
   ```
   cd go-validation/app
   ```
1. run
   ```
   go run main.go
   ```
1. request

   ```
   curl -i -X POST \
    -H "Content-Type:application/json" \
    -d \
    '{
      "title": "",
      "author": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      "price": -1,
      "genre": "dummy"
    }' \
   'http://localhost:8080/books'
   ```
   -- output --
   ```json
   {
     "type": "VALIDATION_ERROR",
     "title": "Your request parameters didn't validate.",
     "invalid-params": [
       {
         "name": "title",
         "reason": "タイトルは必須項目です。"
       },
       {
         "name": "author",
         "reason": "著者名は 1文字 以上 50文字 以内です。"
       },
       {
         "name": "price",
         "reason": "価格は 1円 以上で指定してください。"
       },
       {
         "name": "genre",
         "reason": "値が正しくありません。"
       }
     ]
   }
   ```
