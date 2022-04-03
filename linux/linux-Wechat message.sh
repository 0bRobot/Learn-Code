#!/bin/bash

URL="https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="   #WeChat  webhookurl
KEY="ef012324928934907482789279"   # webhookurl key
URLToken=$URL$KEY

#system info
SysTime=`date "+%Y-%m-%d %H:%M"`
UPusers=`uptime | cut -d " " -f 9`
UMem=$(free -m |sed -n '2p' |awk '{print $3}')
FMem=$(free -m |sed -n '2p' |awk '{print $4}')
Hostname=`hostname`
IP=`ifconfig | grep  eth0 -A 1| grep inet  | awk '{print $2}'`
Connects=`ss -at | grep ESTAB |wc -l`

#post date
curl  -s -o /dev/null --url "$URLToken" \
   -H 'Content-Type: application/json' \
   -d '{
       "msgtype":"markdown",
       "markdown": {
           "content":
        "#  ** <font color=\"warning\">System Status Message </font> ** 
    > 
    >**====================**
    >**====================**      
        >**   主机名     <font color=\"info\"> '$Hostname'</font>**
        >**   IP地址     <font color=\"info\">'$IP'</font>**  
        >**   连接数量   <font color=\"info\">  '$Connects'</font>**  
        >**   在线用户   <font color=\"info\">  '$UPusers'</font>**  
        >**   已用内存   <font color=\"info\">'$UMem' MB</font>** 
        >**   剩余内存   <font color=\"info\">'$FMem' MB</font>**
    >**====================**
    >**====================**
    "
           }
           ,"safe":1
    }'
    