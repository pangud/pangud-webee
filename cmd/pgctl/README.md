# pd
pd is a client tool for managing you pangud applications.
```shell
pd get endpoints
pd config --apiserver=https://localhost:6443 #需要设置输入token、server名称
pd select endpointId/endpointName
pd get containers --endpoint endpointId/endpointName
pd logs -f containerId/containerName
pd get dbinsts --endpoint endpointId/endpointName
pd get databases --dbinst dbinstId/dbinstName
```