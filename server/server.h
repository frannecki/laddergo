#ifndef SERVER_H
#define SERVER_H
#include <ladder/TcpServer.h>

class Server {
public:
  Server(unsigned short);
  Server& operator=(const Server&) = delete;
  ~Server();
  void Start();
private:
  ladder::TcpServer* server_;
};

#endif
