# web_access_local
 A local service program, providing websocket interface for web pages, accessing local resources and system calls.
 
 
1. Supported features

(1) Run the local application.

(2) Read the contents of the file.

(3) Read the file meta information.

(4) Read the registry.

(5) Read CPU information.

(6) Read memory information.

(7) Read the machine code.

(8) Read the U disk information.

(9) Customize the command.


2. Installation configuration

2.1. Install to local service program
  "web_access_local_setup/bin/debug/Web_access_local_setup.exe", configure the local websocket service port number and remote authentication URL.

2.2. Code implementation of web page
  According to personal needs, implement part of the content in "protocol.json" in the web page. The "auth" command is required. After establishing websocket connection, the first command must be "auth", refer to "test.html".
  
  
