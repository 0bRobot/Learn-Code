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

function Install_soft()
{
wget --spider www.baidu.com  &> /dev/null
if [ `echo $?`X == '0'X ]
then
    echo "Wget install  Successfully"
else
    yum install wget -y &> /dev/null   
fi
yum remove mariadb* -y  &> /dev/null 

if [ ! -f  `pwd`/mysql80-community-release-el7-6.noarch.rpm ]
then
    wget https://repo.mysql.com//mysql80-community-release-el7-6.noarch.rpm  &> /dev/null
    if [ -f `pwd`/mysql80-community-release-el7-6.noarch.rpm ]
    then
        rpm -ivh mysql80-community-release-el7-6.noarch.rpm &> /dev/null
        rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022  &> /dev/null
    fi 
fi
}

function check_mysql()
{
SS_status
yum install  mysql mysql-server -y &> /dev/null
if [ `echo $?`X == '0'X ]
then 
    echo "+------------------------------------+"
    echo "|    Mysql Install Successfully!!    |"
    echo "+------------------------------------+" 
fi
systemctl start mysqld.service 
systemctl enable mysqld.service 

mysqlpath=`systemctl status mysqld.service |grep Active`

if [[ $mysqlpath =~ running ]]
then
     if [ -f  /var/log/mysqld.log ]
     then
     echo "+----------------------------------------------------------------+"
     echo "|================================================================|"
     echo "|    Mysql Default Password  : `grep 'temporary password'   /var/log/mysqld.log | awk -F "for" '{print $NF}'`     |"
     echo "+================================================================+"
     else
     echo "Check  Mysql Services"
     fi

fi
}

SS_status
ping www.baidu.com -c 1   &>  /dev/null 

if [ `echo $?`X == '0'X ]
then
    yum list  &> /dev/null
    if [ `echo $?`X  = '0'X ]
    then
       Install_soft
       check_mysql
    else
    echo "yum repo  Error!"
    fi
else 
echo "network connection Erroe"
fi