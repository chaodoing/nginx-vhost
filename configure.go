package main

type establish struct {
	// 站点名称
	Site string
	// 站点存储路径
	Path string
	// 站点日志路径
	Log string
	// 站点日志路径
	Error string
}

var Configure establish = establish{
	Site:  "",
	Path:  "",
	Log:   "",
	Error: "",
}

var Nginx string = `
server {
	# 监听端口义
	listen		80;
	
	# 主机名
	server_name				{{.Site}};	# 站点名称
	autoindex 				on;		# 目录索引
	autoindex_exact_size 	off;		# 显示文件大小
	autoindex_localtime 	on;			# 显示本地时间
	charset 				utf-8;		# 字符集
	gzip 					on;			# 开启gzip

	# 日志位置定义
	# access_log	off;
	access_log		{{.Log}}/access.log; # 访问日期
	error_log		{{.Log}}/error.log;  # 错误日志

	# 定义根目录和索引文件

	root 		{{.Path}}/public;
	index 		index.php index.html index.htm default.php default.html default.htm;

	error_page 403 	/error_page/403.html;
	error_page 404 	/error_page/404.html;
	error_page 500	/error_page/500.html;
	error_page 501	/error_page/501.html;
	error_page 502	/error_page/502.html;
	error_page 503	/error_page/503.html;
	error_page 503	/error_page/503.html;
	error_page 504	/error_page/504.html;

	location /error_page {
		root {{.Path}};
	}

	# 图片压缩配置
	location ~*\\.(jpg|jpeg|gif|png|ico|swf)$ {
		access_log	off; # 关闭图片访问日志
		gzip		off;
	}
	location ~*\\.(css|js|less)$ {
		access_log	off; # 关闭脚本访问日志
		gzip		on;
	}

	# 加载 rewrite 规则
	#location / {
	#	if (!-e \$request_filename) {
	#		rewrite ^(.*)\$ /index.php?s=\$1 last;
	#		rewrite ^.*$ /index.html last;
	#	}
	#}

	location ~ \\.php$ {
		fastcgi_pass	127.0.0.1:9000;
		fastcgi_index	index.php;
		fastcgi_param 	SCRIPT_FILENAME  \$document_root\$fastcgi_script_name;
		include			fastcgi_params;
	}
}
`
