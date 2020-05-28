# nginx站点配置工具

## nginx.conf

```
worker_processes  1;
events {
    worker_connections  1024;
}
http {
        include                     mime.types;
        default_type                application/octet-stream;
        server_tokens               off;
        sendfile                    off;
        keepalive_timeout           65;
        gzip                        on;
        client_max_body_size        4096m;
        expires                     off;

	dav_methods PUT DELETE MKCOL COPY MOVE;

        # include
        include  [vhost_path]/*.conf;
}
```

```shell
vhost_linux -v [vhost_path] -nginx `which nginx`
# 请输入站点名称: 开始配置站点
```