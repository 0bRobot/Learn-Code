#!/bin/bash
ip="192.168.110.201"
user="root"
pass="123456"

/usr/bin/expect <<-EOF

set timeout 20

spawn ssh $user@$ip
expect {
 "yes/no" { send "yes\r"; exp_continue }
 "password"  { send "$pass\r" }
}

expect "]#"
send "touch e{1..4}\r"
expect eof
EOF

