# ------------------------ Admin auth ------------------------

curl -X 'POST' \
  'http://localhost:8080/api/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "login": "admin",
  "password": "presale"
}'

# {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQwNjU3MDksInN1YiI6IjNmMDAxODM4LWQ4NDYtNGFhZC05NjliLTJjNWExZmZlZDMxYyJ9.02ugGRdUVTyf0I93v4q_eROa3VOe9PnPpKQYtWFhVKo"}

# ------------------------ Not added user ------------------------

curl -X 'POST' \
  'http://localhost:8080/api/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "login": "123",
  "password": "123"
}'

# no user with login 123

# ------------------------ Wrong password ------------------------

curl -X 'POST' \
  'http://localhost:8080/api/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "login": "admin",
  "password": "123"
}'

# wrong password

