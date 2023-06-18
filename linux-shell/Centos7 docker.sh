#!/bin/bash

function  RemoveOldDocker()
{
    echo "Uninstall old versions"
    yum remove docker docker-client docker-client-latest docker-common  docker-latest docker-latest-logrotate  docker-logrotate docker-engine &>  /dev/null 
    sleep 3
    clear
}

function Install_Docker()
{
    echo "Install the yum-utils package"
    yum install -y yum-utils &> /dev/null
    echo "Yum Add docker-repo "
    yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo &> /dev/null
    if [ `echo $?` == '0' ]
    then
        sleep 3
        clear
        echo "Install docker"
        yum install docker-ce docker-ce-cli containerd.io docker-compose-plugin -y &> /dev/null
        if [ `echo $?` == '0' ]
        then 
            echo "Docker install successfully"
            systemctl start docker &> /dev/null 
            systemctl enable docker &> /dev/null
            
        else
        echo "Check yum install docker-ce"
        fi    
    else
    echo "Check network and  yum-utils" 
    fi   
}


function compose_Install()
{
    # https://api.github.com/repos/{owner}/{repo}/releases/latest
    compose_latest=`curl -s https://api.github.com/repos/docker/compose/releases/latest | grep  tag_name |awk -F "\"" '{print $4}'`
    if [ `echo $?` == '0' ]
    then
        echo "docker latest version $compose_latest"
        curl -LJo  docker-compose-linux-x86_64  https://ghproxy.com/https://github.com/docker/compose/releases/download/$compose_latest/docker-compose-linux-x86_64         
        if [ `echo $?` == '0' ]
        then
            if [ -f `pwd`/docker-compose-linux-x86_64 ]
            then
            chmod 777  `pwd`/docker-compose-linux-x86_64 
            mv `pwd`/docker-compose-linux-x86_64 /usr/bin/docker-compose -f
            Docker_V=`docker -v | cut -d "," -f1`
            comeposeV=`docker-compose -v`
            echo "+-----------------------------------------------+"
            echo "|           Docker Install Successfully         |"
            echo "|      Docker-comepose Install Successfully     |"
            echo "|-----------------------------------------------|"
            echo "|             $Docker_V           |"
            echo "|         $comeposeV        |"
            echo "+-----------------------------------------------+"
            exit
            fi

        else
            echo "Docker Install Successfully"
            echo "Docker-commpose Not installed"
            echo "+---------------------------------------------------------------------------------------------+"
            echo "|           docker-compose download  timeout  Please manual operation download                |"
            echo "|  https://github.com/docker/compose/releases/download/$compose_V/docker-compose-linux-x86_64    |"
            echo "+---------------------------------------------------------------------------------------------+"
        fi
    else
    echo "network failed  !  get docker-compose version error"

    fi    

}

function  choose()
{
    clear
    sleep 2
    echo "Do you want install docker-compose"
    read -t 6 -p "input  Y/yse   N/no :   " Rstatus
    if  [ "$Rstatus"X == "Y"X ]||[ "$Rstatus"X == "y"X ]||[  "$Rstatus"X == "yes"X ]||[ "$Rstatus"X == "YES"X ] 
    then
        compose_Install
        exit
    elif [ "$Rstatus"X == "n"X ]||[ "$Rstatus"X == "N"X ]||[ "$Rstatus"X == "no"X ]||[ "$Rstatus"X == "NO"X ]
    then
        Docker_V=`docker -v | cut -d "," -f1`
        echo "+-----------------------------------------------+"
        echo "|           Docker Install Successfully         |"
        echo "|        Docker-commpose Not installed!!        |"
        echo "|             $Docker_V           |"
        echo "+-----------------------------------------------+"
        exit 
    else
        echo "+-----------------------------------------------+"
        echo "|           Docker Install Successfully         |"
        echo "|        Docker-commpose Not installed!!        |"
        echo "|             $Docker_V           |"
        echo "+-----------------------------------------------+"
        echo "----------- Input Erroe Program Exit ------------"
        exit
    fi


}

echo "Check Internet Connet"
ping www.baidu.com -c 1   &>  /dev/null 
echo "The network connection is normal"
sleep 2
clear
if [ `echo $?`X == '0'X ]
then
    yum list  &> /dev/null
    if [ `echo $?`X  = '0'X ]
    then
       RemoveOldDocker
       Install_Docker
       choose
    else
    echo "yum repo  Error!"
    fi
else 
echo "network connection Erroe"
fi
