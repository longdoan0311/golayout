app:
  mod: "development"
server:
  port: 3011

store:
  use: "imdb" # one of: imdb, postgres
  postgres:
    host: "<<YOUR_HOST>>"
    port: <<YOUR_POST>>
    db: ""
    user: ""
    pass: ""
    sslmode: false
    maxopencnn: 30
    maxcnnlifetime: 30 # minutes