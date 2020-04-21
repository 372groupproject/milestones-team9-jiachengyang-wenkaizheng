This project is a simple version of shadowsocks software in Golang.
Our approach is to create two proxies for network transfer (Server and Local). 
Usually, when the user visits one website, the user computer will performance DNS lookup for server Ip, ARP protocol for getting mac address,
and then do the three times TCP handshake for establishing a network connection. Finally, send requests(https or http).
This whole process can be done by the user application and the server, so basically we have 2 nodes.
User -> Server
User <- Server
In our implementation, we add two more nodes which are Local Proxy and Server Proxy between them.
User -> Local Proxy -> Server Proxy -> Server
User <- Local Proxy <- Server Proxy <- Server
The reason is that we need to do encryption for all traffic data. With encryption, we proxy can go cross firewall.
We also do user sign-in process in the beginning. When Local proxy accepts a connection from the user, it will send username and password(in Configure file)first. And Server Proxy will check all user information(in CSV file). After users pass the verification,
Local Proxy will send encode and decode table (256-byte array) to Server Proxy for future encryption usage.
For users part, they need to set up their chrome with socks5 protocol.
Socks5 : https://tools.ietf.org/html/rfc1928
How to set up chrome(or other browsers) : will be given later because we want to find a very good source for users.
Socks5 will send some specific network packets to LocalProxy and then LocalPorxy transfers it to SeverProxy. Server Proxy will send a reply 
according to the packet it receives. After the sock5 protocol process is done, both proxies will continue to transfer the normal data packet.
We also include thread control in our project because each request(https,http) will generate threads for each proxy to deal with. 
The Socks5 protocol supports UDP as well, we focus on TCP in this project.
To support the UDP version is our future plan
