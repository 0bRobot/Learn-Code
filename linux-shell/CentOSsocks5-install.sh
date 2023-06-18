#!/bin/bash

function SS_status()
{
    mark=''
for (( ratio=0;${ratio}<=100;ratio+=5))
do
    sleep 0.2
    printf "处理中:[%-40s]%d%%\r" "${mark}" "${ratio}"
    mark="##${mark}"
done
echo
}


function SS5_uninstall()
{
Remove_Port=`ps -aux | grep 'sbin\/ss5' | awk -F ":" '{print $4}'`
firewall-cmd --remove-port="$Remove_Port"/tcp --permanent   > /dev/null 2>&1
firewall-cmd --reload > /dev/null 2>&1
kill -9 `netstat -altp | grep ss5 |tr -s " " |cut -d " " -f 7| awk -F '/' '{print $1}'` > /dev/null 2>&1
/etc/init.d/ss5 stop   > /dev/null 2>&1 
find / -name "ss5*" -exec rm -rf {} \;  > /dev/null 2>&1
SS_status
}


function SS5_install()
{
echo -e "\e[1;33m===========下载安装编译环境==========\e[0m"
yum install gcc gcc-c++ pam-devel openldap-devel openssl-devel wget -y  >  /dev/null 2>&1

if  [ "`echo $?`" == "0" ]
    then
echo "===========编译环境安装完成=========="
echo "====================================="
else
echo "软件包安装错误"
echo "请检查yum源及网络"
    exit 1
fi

echo -e "\e[1;33m=============下载软件包中============\e[0m"
echo "====================================="
wget --no-check-certificate https://jaist.dl.sourceforge.net/project/ss5/ss5/3.8.9-8/ss5-3.8.9-8.tar.gz >  /dev/null 2>&1

U_PATH=`pwd`
if [ -f $U_PATH/ss5-3.8.9-8.tar.gz ]
then
    tar -xzvf ss5-3.8.9-8.tar.gz &>/dev/null
echo  "====== SS5 软件包下载解压完成 ======="
else
echo "====================================="
echo "======== 文件未找到 请检查 =========="
echo "====================================="
    exit 2
fi

cd $U_PATH/ss5-3.8.9
if [ "`echo $?`" == "0" ]
then
echo -e "\e[1;33m============编译中 请稍等============\e[0m"
    ./configure --prefix=/usr/local/src/  >  /dev/null 2>&1
echo -e "\e[1;33m============== 软件安装中 ===========\e[0m"
    sleep 0.5
    make >  /dev/null 2>&1
    make install  >  /dev/null 2>&1
echo  "============= 安装完成 =============="
else 
echo "==============安装失败==============="
echo "==========请检查编译环境============="
    exit 3
fi

if [ -f /etc/opt/ss5/ss5.conf   ]
    then
    sed  -i '/^#auth/a auth    0.0.0.0\/0        -        u' /etc/opt/ss5/ss5.conf 
    sed  -i 's/^#permit -/permit u/g' /etc/opt/ss5/ss5.conf
echo  "========修改ss5.conf文件完成========="
    else
echo "==========修改配置文件失败==========="
    exit 4
fi

U_Passwd=`cat /dev/urandom | tr -dc [:alnum:] |head -c 20`
U_user=`cat /dev/urandom | tr -dc [:alpha:] |head -c 3`

if [ -f  /etc/opt/ss5/ss5.passwd ]
    then
    sed -i "1 c "$U_user"admin $U_Passwd" /etc/opt/ss5/ss5.passwd
else
echo "======找不到SS5 passoword文件========" 
    exit 5
fi

aa=`echo $RANDOM`
bb=`echo $RANDOM`
U_Port=`echo $(($aa%$bb+8699))`
if [ -f /etc/sysconfig/ss5  ] 
then
    sed -i "2 aSS5_OPTS=\" -u root -b 0.0.0.0\:$U_Port\"" /etc/sysconfig/ss5
else
echo "======找不到sysconfig/ss5文件========"
    exit 6
fi

if [ -f /etc/init.d/ss5 ]
then
    chmod 750  /etc/init.d/ss5
    /etc/init.d/ss5  restart >  /dev/null 2>&1 
else
echo "========= 找不到ss5启动脚本=========="
    exit 7
fi
sleep 0.5
Scheck=`ps -aux | grep 'sbin\/ss5' | awk -F ":" '{print $4}'`

if [ "$Scheck" == "$U_Port" ]
then
echo -e "======== SS5 服务启动成功 ==========="
else
echo "======== SS5 服务启动失败 ==========="
    exit 8
fi

StatUS_server=`systemctl status firewalld.service  | grep Active: | awk '{print $2}'`
if [ $StatUS_server == "active" ]
then
echo -e "\e[1;33m========== 检查防火墙状态 ===========\e[0m"
    firewall-cmd --add-port=$U_Port/tcp  --permanent >  /dev/null 2>&1
    firewall-cmd  --reload   >  /dev/null 2>&1
echo  "========= 防火墙策略已更新 =========="
fi
SS_status

ip=`ifconfig  |grep -E -A 1 'ens33|eth0' | grep broadcast |tr -s " " | cut -d " " -f 3`
echo ""
echo ""
echo -e "\e[9;34m============================Message===========================\e[0m"
echo "--------------------------------------------------------------"
echo -e "\e[1;32m+---------------+++++++++++++++++++++++++++------------------+\e[0m"
echo -e "\e[1;31m|  User: "$U_user"admin         Password: $U_Passwd     |\e[0m"
echo -e "\e[1;31m|  Ip: $ip               Port: $U_Port             |\e[0m"
echo -e "\e[1;32m+---------------+++++++++++++++++++++++++++------------------+\e[0m"
echo "--------------------------------------------------------------"
}


for (( a=1 ; a<4; a++ ))
do
echo "========================================================"
echo -e "\e[1;34m================== Socks5 一键脚本 =====================\e[0m"
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
    SS5_install
    echo 
    echo -e  "\e[1;31m+-----------------------------------------------------+\e[0m"
    echo -e  "\e[1;31m|\e[0m\e[5;34m+++++++++++++++++++++ 安装完成 +++++++++++++++++++++++\e[0m\e[1;31m|\e[0m "
    echo -e  "\e[1;31m+-----------------------------------------------------+\e[0m"
    echo
    break
elif [ "$u_FUNC"X == "1"X ]
then
    SS5_uninstall
    echo
    echo -e  "\e[1;34m+----------------------------------------------------+\e[0m"
    echo -e  "\e[1;34m|\e[0m\e[1;31m++++++++++++++++++++++卸载完成+++++++++++++++++++++++\e[0m\e[1;34m|\e[0m "
    echo -e  "\e[1;34m+----------------------------------------------------+\e[0m"
    echo 
    break
elif [ "$u_FUNC"X == "q"X ]
then
    echo -e "\e[4;33m #_# 拜拜~~  程序退出 #-#\e[0m"
    break
else
    echo 
    echo -e "\e[3;31m#########################################\e[0m"
    echo -e "\e[5;31m ------------输入错误,不要乱输------------ \e[0m"
    echo -e "\e[3;31m#########################################\e[0m"
    sleep 2.5
    clear
fi
done

