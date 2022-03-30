#!/bin/bash
yum -y install bash-completion  net-tools    2>>   error.log
yum update -y   2>>   error.log
rpm -Uvh http://nginx.org/packages/centos/7/noarch/RPMS/nginx-release-centos-7-0.el7.ngx.noarch.rpm  2>>   error.log
yum install -y nginx vim  2>>   error.log
systemctl start firewalld   2>>   error.log
systemctl enable firewalld   2>>   error.log
firewall-cmd --permanent --add-rich-rule='rule protocol value=icmp drop'
read -t 10 -p "input ssh port:"ssh_port
sed -i 's/#Port 22/Port '$ssh_port'/' /etc/ssh/sshd_config
sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/'  /etc/ssh/sshd_config  2>>   error.log

systemctl restart sshd
sed -i 's/22/'$ssh_port'/'  /usr/lib/firewalld/services/ssh.xml   2>>   error.log
firewall-cmd --reload   2>>   error.log
systemctl restart firewalld   2>>   error.log

echo  'export HISTTIMEFORMAT="%Y-%m-%d %H:%M:%S  "' >> /etc/profile  && source /etc/profile

echo "auth required pam_tally.so deny=3 unlock_time=5" >> /etc/pam.d/sshd