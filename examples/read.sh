# ------------------------ OK ------------------------

curl -X 'POST' \
  'http://localhost:8080/api/read' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQwNjY5NzksInN1YiI6IjNmMDAxODM4LWQ4NDYtNGFhZC05NjliLTJjNWExZmZlZDMxYyJ9.IaoC76J5LYi9Wgyjy6IDPGEl_hrrMacNrHrl9-fldqU' \
  -H 'Content-Type: application/json' \
  -d '{
  "keys": [
    "key1", "key2", "key5"
  ]
}'

# {"data":{"key1":"1","key2":2,"key5":5}}

# ------------------------ Not stored keys ------------------------

curl -X 'POST' \
  'http://localhost:8080/api/read' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQwNjY5NzksInN1YiI6IjNmMDAxODM4LWQ4NDYtNGFhZC05NjliLTJjNWExZmZlZDMxYyJ9.IaoC76J5LYi9Wgyjy6IDPGEl_hrrMacNrHrl9-fldqU' \
  -H 'Content-Type: application/json' \
  -d '{
  "keys": [
    "key8"
  ]
}'

# {"data":{}}

# ------------------------ Unauthorize ------------------------

curl -X 'POST' \
  'http://localhost:8080/api/read' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "keys": [
    "key1", "key2", "key5"
  ]
}'

