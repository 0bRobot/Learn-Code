#!/bin/bash

function SS_status()
{
    mark=''
for (( ratio=0;${ratio}<=100;ratio+=5))
do
    sleep 0.1
    printf "处理中:[%-40s]%d%%\r" "${mark}" "${ratio}"
    mark="##${mark}"
done
echo
}
function CheckSYS()  ## No.1
{
    	if [ -f /etc/redhat-release ]
        then
        OS_ver=`cat /etc/redhat-release  | awk   '{ print  $4 }'| awk -F "." ' { print $1}'`
       
        if [ "$OS_ver"X == "6"X ]||[ "$OS_ver"X == "7"X ] 
            then
            echo -e "\e[1;33m====当前系统:`cat /etc/redhat-release`====\e[0m"
        else
            echo -e "\e[1;31m=========++++++ 此脚本不支持当前 CentOS版本 ++++++=========\e[0m"
           
        fi 
    else 
    echo -e "\e[1;31m===========++++++  抱歉--- 当前系统不支持 ++++++===========\e[0m"

    fi
    
}

CheckENV()   ##No.2
{
    yum install http://ngtech.co.il/repo/centos/7/squid-repo-1-1.el7.centos.noarch.rpm -y   >  /dev/null 2>&1
    yum -y install squid     >  /dev/null 2>&1
    #yum install httpd-tools -y  >  /dev/null 2>&1
    SquidProt=`echo $RANDOM`  #生成随机端口
    sed -i "s/http_port 3128/http_port "$SquidProt"/" squid.conf  /etc/squid/squid.conf  #修改配置文件
    systemctl start squid    >  /dev/null 2>&1
    systemctl enable squid   >  /dev/null 2>&1
    StatUS_firewalld=`systemctl status squid  | grep Active: | awk '{print $2}'`
    if [ "$StatUS_firewalld" == "active" ]
    then
        echo -e "\e[1;33m========== 检测到系统防火墙已开启 ===========\e[0m"
        firewall-cmd --add-port="$SquidProt"/tcp --permanent     >  /dev/null 2>&1
        firewall-cmd --reload    >  /dev/null 2>&1
        echo -e "\e[1;33m============ Squid 端口已放行 ==============\e[0m"
    else
        systemctl start firewalld    
        firewall-cmd --add-port="$SquidProt"/tcp --permanent  >  /dev/null 2>&1
        firewall-cmd --reload    >  /dev/null 2>&1
        echo -e "\e[1;33m============ Squid 端口已放行 ==============\e[0m"   
    fi

    SQUID_Status=`systemctl status squid  | grep Active: | awk '{print $2}'`
    if [ "$SQUID_Status" == "active" ]
    then 
        echo -e "\e[1;33m=========== Squid 服务启动成功 =============\e[0m"
    else
        echo -e "\e[1;31m======== 未检测到服务，请手动检查 ===========\e[0m"
    fi
    SS_status
    clear 
Sqip=`ifconfig  |grep -E -A 1 'ens33|eth0' | grep broadcast |tr -s " " | cut -d " " -f 3`
    echo ""
    echo ""
    echo -e "\e[9;34m============================Message===========================\e[0m"
    echo "--------------------------------------------------------------"
    echo -e "\e[1;32m----------------+++++++++++++++++++++++++++-------------------\e[0m"
    #echo -e "\e[1;31m  User: "$U_user"admin   Password: $U_Passwd                 \e[0m"
    echo -e "\e[1;31m  Ip: $Sqip                      Port: $SquidProt                       \e[0m"
    echo -e "\e[1;32m----------------+++++++++++++++++++++++++++-------------------\e[0m"
    echo "--------------------------------------------------------------"
}


function Uninstall()
{
    systemctl stop squid  >  /dev/null 2>&1
    SWstat=`echo $?`
    if [ "$SWstat" == "0" ]
    then
    yum remove -y squid >  /dev/null 2>&1
    echo -e "\e[1;32m----------------+++++++++++++++++++++++++++-------------------\e[0m"
    echo -e "\e[1;33m=====================  Squid 卸载完成 ！ ======================\e[0m"
    echo "--------------------------------------------------------------"
    rm -rf $0
    else
    echo -e "\e[1;33m=====================  Squid 没有安装 ！ ======================\e[0m"
    rm -rf $0
    fi
}



function CheckInit()  #环境检查
{
    CheckSYS  #检查系统版本
    CheckENV #环境检查
   # User_passwd
}



for (( i = 1; i < 4; i++ ))
do
echo "========================================================"
echo -e "\e[1;34m==================  Http代理 一键脚本 ==================\e[0m"
echo "--------------------------------------------------------"
echo -e "\e[1;32m================== 输入 0 进行安装 =====================\e[0m"
echo -e "\e[1;32m================== 输入 1 进行卸载 =====================\e[0m"
echo -e "\e[1;32m================== 输入 q 退出程序 =====================\e[0m"
echo "========================================================"
echo "--------------------------------------------------------"
echo -e "\e[2;31m----------------- 目前仅适用于Centos  ------------------\e[0m"
echo "--------------------------------------------------------"

read -p "请按键选择:  " u_FUNC

if [ "$u_FUNC"X == "0"X ]
then
    CheckInit

    break
elif [ "$u_FUNC"X == "1"X ]
then
    Uninstall
    break

elif [ "$u_FUNC"X == "q"X ]
then
    echo -e "\e[4;33m #_# 拜拜~~  程序退出 #-#\e[0m"
    break

else
    echo 
    echo -e "\e[3;31m#########################################\e[0m"
    echo -e "\e[5;31m------------输入错误,不要乱输------------\e[0m"
    echo -e "\e[3;31m#########################################\e[0m"
    sleep 2.5
    clear
sleep 3 
clear
fi
done


function User_passwd() ##No.3  未用
{

    if [ -d  /etc/squid ]
    then
    s_user=`cat /dev/urandom | tr -dc [:alpha:] |head -c 6`
    s_Passwd=`cat /dev/urandom | tr -dc [:alnum:] |head -c 20`
    htpasswd -cb /etc/squid/.us_passwd   $s_user $s_Passwd
    echo  $s_user  
    echo  $s_Passwd
    chown squid:squid  /etc/squid/.us_passwd 
    sed -i "1aauth_param basic program /usr/lib64/squid/basic_ncsa_auth /etc/squid/.us_passwd" /etc/squid/squid.conf
    sed -i "2aauth_param basic children  3"  /etc/squid/squid.conf
    sed -i "3aacl auth_users proxy_auth $s_user"  /etc/squid/squid.conf
    sed -i "4ahttp_access allow auth_users"  /etc/squid/squid.conf
    systemctl restart squid.service
    fi
}