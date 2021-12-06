# meson.network-lts-terminal

```
./meson service install
./meson service start
./meson service status
./meson service stop
./meson service remove
```

```
-root [root path]
-space [size in GB]
-token [user token]
-server [server host]
-port [port number]
```

```
apis:

// health check 
GET /api/testapi/test
GET /api/testapi/health

// file operation cmd (only from server)
POST /api/v1/file/save
POST /api/v1/file/delete
POST /api/v1/file/pause

// terminal cmd (only from server)
POST /api/v1/node/restart
POST /api/v1/node/pause

// logs (only from server)
GET /api/v1/defaultlog
GET /api/v1/filerequestlog

// schedulejob status (only from server)
GET /api/v1/schedulejobstatus

// node status (only from server)
GET /api/v1/nodestatus

// file request 
GET /*

```
