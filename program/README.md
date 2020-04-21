This project is a simple version of shadowsocks software in Golang.  

Our approach is to create two proxies for network transfer (Server and Local).  
Usually, when users visit websites, their computers will performance DNS lookup for server IP, 
ARP request or ARP cache for getting mac address and then do the TCP handshake for establishing a network connection. 
Finally, they will send requests(https or http).  

This whole process can be done by user applications and servers, so basically we have 2 nodes.   
User -> Server   
User <- Server   
In our implementation, we add two more nodes, which are Local Proxy and Server Proxy sitting in-between users and real servers.    
User -> Local Proxy -> Server Proxy -> Server   
User <- Local Proxy <- Server Proxy <- Server   
The reason is that we need to encrypt all traffic data in order to bypass firewall.   
When a user initiates local proxy, it will send username and password, and then server proxy will compare them with pairs of username and corresponding password stored in data.csv.  
For matters of security, we only store sha512 values of salted usernames and salted passwords.
After authentication steps, local proxy will send encode and decode table (256-byte array) to server proxy for future encryption usage.   
When handling requests from user applications and responds from read servers, we use multiple go-routines so that we handel each request simultaneously.  


For users part, they need to set up their chrome with socks5 protocol.   
Socks5 : https://tools.ietf.org/html/rfc1928  
How to set up chrome :   
- install the Chrome extension SwitchyOmega  
- open options  
- added proxy:
  - Protocol: SOCKS5 
  - Server: 127.0.0.1
  - Port: 5209 (as local_port defined in config.json)  
- make sure you are using proxy rather than direct connection  
- do not forget to run local proxy

Browsers will send specific network packets to local proxy, and then local proxy transfers them to sever proxy.
 Server Proxy will respond them according to packets it receives. 
 After the sock5 protocol process is done, both proxies will continue to transfer the normal data packet.   
Although The Socks5 protocol supports UDP as well, due to the limit of time,  we focus on TCP in this project.   
