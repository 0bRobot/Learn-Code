#!/bin/bash
yum -y install bash-completion  net-tools    
yum update -y   
rpm -Uvh http://nginx.org/packages/centos/7/noarch/RPMS/nginx-release-centos-7-0.el7.ngx.noarch.rpm  

systemctl start firewalld   
systemctl enable firewalld   
firewall-cmd --permanent --add-rich-rule='rule protocol value=icmp drop'




echo  'export HISTTIMEFORMAT="%Y-%m-%d %H:%M:%S  "' >> /etc/profile  && source /etc/profile

echo "auth required pam_tally.so deny=3 unlock_time=5" >> /etc/pam.d/sshd

function sshd_init()
{
if [ -f /etc/ssh/sshd_config ]
then
	sshStatUS=`systemctl status sshd | grep Active: | awk '{print $2}'`
	firewallStatUS=`systemctl status sshd | grep Active: | awk '{print $2}'`
	if [ "$sshStatUS" == "active" ]
	then
		read -t 10 -p "Input  Your SSH Port: " ssh_port
		sed -i 's/#Port 22/Port '$ssh_port'/' /etc/ssh/sshd_config
		sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/'  /etc/ssh/sshd_config  >  /dev/null 2>&1
		if [ "$firewallStatUS" == "active" ]
		then	
		sed -i 's/22/'$ssh_port'/'  /usr/lib/firewalld/services/ssh.xml  
		firewall-cmd --reload  >  /dev/null 2>&1
		systemctl restart sshd 
		echo "SSH Port Update Success"
		echo "Firewall Reload Success"
		fi
	fi		

fi
}


//SE_status关闭
function se_status()
{
if [ -f /etc/selinux/config ]
then
SE_status=`grep ^SELINUX= /etc/selinux/config | awk -F "=" '{print $2}'`
	if [ "$SE_status" == "disabled" ]
	then
		echo " SELINUX IS CLOSED "
		echo " NOTHING TO DO ! "
	else
		sed -i s/SELINUX=enforcing/SELINUX=disabled/ /etc/selinux/config 
	fi	
fi
}