package main

type Configure struct {
	NGINX      string
	VHOST      string
	ServerName string
	RootPath   string
	LogPath    string
	ErrorPath  string
}

type StdIn struct {
	ServerName string
	RootPath string
	LogPath string
	ErrPath string
}