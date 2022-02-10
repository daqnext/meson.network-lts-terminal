# meson.network-lts-terminal

### Usage
service command:
```
//run meson as system service

./meson service install 
./meson service start
./meson service status
./meson service stop
./meson service restart
./meson service remove
```

log command:
```
//print logs

./meson logs [command options]  //print logs

OPTIONS:
--num   define print lines, default is 20
--onlyerr   if true only print error logs, default is false

example: ./meson logs -num=30 -onlyerr=true
```

config command:
```
//print current config

./meson config show //print current confing
```
```
//set config

./meson config set [command options] //set config

OPTIONS:
--dest          //set server host     
--log_level     //set log level(TRAC,DEBU,INFO,WARN,ERRO)
--token         //set token   
--port          //set https service port
--addpath       //add provide folder    
--removepath    //remove provide folder 

example:
./meson config set --dest=target.serverhost.com
./meson config set --log_level=INFO
./meson config set --token=zrTusiei77sdfieiwcx==
./meson config set --port=443
./meson config set --addpath=/root/path/mesonfolder
./meson config set --removepath=/root/path/mesonfolder
```
```
//check config

./meson config check //check current config is correct or not
```