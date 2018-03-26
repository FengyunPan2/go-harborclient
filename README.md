# this project provides rest api to ceate/list/delete image from harbor docker registry

hubimage User Case:
image:
curl -X POST  -H "Content-Type: application/json" -d '{"file": "/root/busybox_test", "projectName": "library","name": "busybox","tag": "test","enableWait": true }'  http://x.x.x.x:30909/api/v1/image
curl -X GET  -H "Content-Type: application/json" -d '{ "projectName": "library", "page": 1, "pageSize": 10}'  http://x.x.x.x:30909/api/v1/image
curl -X GET  -H "Content-Type: application/json" -d '{ "projectName":"library", "tag": ""}'  http://x.x.x.x:30909/api/v1/image/busybox
curl -X DELETE  -H "Content-Type: application/json" -d '{"projectName": "library", "tag": "test" }'  http://x.x.x.x:30909/api/v1/image/busybox

project:
curl -X GET  -H "Content-Type: application/json" -d '{"name": "", "public": true ,"owner": "", "page": 1, "pageSize": 10}'  http://x.x.x.x:30909/api/v1/project
curl -X GET  http://x.x.x.x:30909/api/v1/project/pp
curl -X DELETE  http://x.x.x.x:30909/api/v1/project/pp
curl -X POST  -H "Content-Type: application/json" -d '{"name": "", "public": true }'  http://x.x.x.x:30909/api/v1/project

user:
curl -X POST  -H "Content-Type: application/json" -d '{"username": "test3", "password": "Passw0rd","email": "","realname": "","comment": "","role_name": "","has_admin_role": 0 }'  http://x.x.x.x:30909/api/v1/user
curl -X GET  -H "Content-Type: application/json" -d '{"username": "", "email":"","page": 1, "pageSize": 10}'  http://x.x.x.x:30909/api/v1/user
curl -X GET http://x.x.x.x:30909/api/v1/user/test3
curl -X DELETE http://x.x.x.x:30909/api/v1/user/test3

Note: x.x.x.x is a ip of kubernetes cluster node
