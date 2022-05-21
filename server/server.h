#ifndef SERVER_H
#define SERVER_H
#include <ladder/TcpServer.h>

class Handler;
class Server {
public:
  Server(unsigned short);
  Server& operator=(const Server&) = delete;
  ~Server();
  void Start();
  void SetRequestCallback(Handler*);
private:
  ladder::TcpServer* server_;
  Handler* handler_;
};

class Handler {
public:
  // should be customed from a derived class
  virtual std::string OnRequest(const std::string&) = 0;
  virtual ~Handler();
};

#endif
