# ------------------------ Write data ------------------------

curl -X 'POST' \
  'http://localhost:8080/api/write' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQwNjY0MTAsInN1YiI6IjNmMDAxODM4LWQ4NDYtNGFhZC05NjliLTJjNWExZmZlZDMxYyJ9.rq37KNAmz2wYJgNxFYgCV--W_m0_HuPTv-JfzZ7z6G4' \
  -H 'Content-Type: application/json' \
  -d '{
  "data": {"key4":"4", "key5":5, "key6":6.0}
}'

# {"status":"success"}

# ------------------------ Write not unique key ------------------------

curl -X 'POST' \
  'http://localhost:8080/api/write' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQwNjY0MTAsInN1YiI6IjNmMDAxODM4LWQ4NDYtNGFhZC05NjliLTJjNWExZmZlZDMxYyJ9.rq37KNAmz2wYJgNxFYgCV--W_m0_HuPTv-JfzZ7z6G4' \
  -H 'Content-Type: application/json' \
  -d '{
  "data": {"key4":"4", "key5":5, "key6":6.0}
}'

# Duplicate key exists in unique index "primary" in space "data" with old tuple - ["key4", "4"] and new tuple - ["key4", "4"] (ClientError, code 0x3), see ./src/box/memtx_tree.cc line 1176

# ------------------------ Unauthorize ------------------------

curl -X 'POST' \
  'http://localhost:8080/api/write' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "data": {"key4":"4", "key5":5, "key6":6.0}
}'


