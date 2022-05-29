#!/bin/bash

echo "======================================="
echo "=======Installing  ENV  Loading========"
yum install gcc gcc-c++ pam-devel openldap-devel openssl-devel wget -y  >  /dev/null 2>&1

if  [ "`echo $?`" == "0" ]
    then
    echo "=======system ENV install Done=========="
    echo "========================================"
else
    echo "install Error!"
    echo "check yum.repo"
    exit 1
fi

    echo "==============Download soft============="
    echo "========================================"
wget --no-check-certificate https://jaist.dl.sourceforge.net/project/ss5/ss5/3.8.9-8/ss5-3.8.9-8.tar.gz >  /dev/null 2>&1

U_PATH=`pwd`
if [ -f $U_PATH/ss5-3.8.9-8.tar.gz ]
then
    tar -xzvf ss5-3.8.9-8.tar.gz &>/dev/null
    echo "====ss5 file Decomposition Finish!!======"
else
    echo "========================================="
    echo "=========File  Not Found !!!!============"
    echo "========================================="
    exit 2
fi

cd $U_PATH/ss5-3.8.9
if [ "`echo $?`" == "0" ]
then
    echo "===============Compiling ================"
    ./configure --prefix=/usr/local/src/  >  /dev/null 2>&1
    echo "===============Install =================="
    sleep 1
    make >  /dev/null 2>&1
    make install  >  /dev/null 2>&1
    echo "=============Install Done================="
else 
    echo "====!!!!======Compiling Error============="
    exit 3
fi

if [ -f /etc/opt/ss5/ss5.conf   ]
    then
    sed  -i '/^#auth/a auth    0.0.0.0\/0        -        u' /etc/opt/ss5/ss5.conf 
    sed  -i 's/^#permit -/permit u/g' /etc/opt/ss5/ss5.conf
    echo "======Modify ss5.conf successfully !======="
    else
    echo "=========Modify ss5.conf  Error============"
    exit 4
fi

U_Passwd=`cat /dev/urandom | tr -dc [:alnum:] |head -c 20`

if [ -f  /etc/opt/ss5/ss5.passwd ]
    then
    sed -i "1 c Madmin $U_Passwd" /etc/opt/ss5/ss5.passwd
else
    echo "Not Found  file ss5.passwd " 
    exit 5
fi

U_Port=`echo $RANDOM`
if [ -f /etc/sysconfig/ss5  ] 
then
    sed -i "2 aSS5_OPTS=\" -u root -b 0.0.0.0\:$U_Port\"" /etc/sysconfig/ss5
else
    echo "Not Found  file /etc/sysconfig/ss5"
    exit 6
fi

if [ -f /etc/init.d/ss5 ]
then
    chmod 750  /etc/init.d/ss5
    /etc/init.d/ss5  restart
else
    echo "Not Found file /etc/sysconfig/ss5"
    exit 7
fi
ss -alt | grep $U_Port  >  /dev/null 2>&1
if [ "`echo $?`" == "0" ]
then
    echo "======SS5 Server Start  Successfully!========="
else
    echo "==========SS5 Server Start  Fail!============="
    exit 8
fi

StatUS_server=`systemctl status firewalld.service  | grep Active: | awk '{print $2}'`
if [ $StatUS_server == "active" ]
then
    echo "=========Check Firewalld  Starting============"
    firewall-cmd --add-port=$U_Port/tcp  --permanent
    firewall-cmd  --reload
fi

ip=`ifconfig  |grep -E -A 1 'ens33|eth0' | grep broadcast |tr -s " " | cut -d " " -f 3`
echo ""
echo ""
echo "----------------------------------------------------------"
echo "------------+++++++++++++++++++++++++++-------------------"
echo "=========================Message=========================="
echo "       User: Madmin    Password:$U_Passwd                 "
echo "       Ip: $ip             Port:$U_Port                   "
echo "------------+++++++++++++++++++++++++++-------------------"
echo "----------------------------------------------------------"