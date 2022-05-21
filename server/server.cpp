#include <functional>
#include <string>

#include <ladder/Buffer.h>
#include <ladder/Connection.h>
#include <ladder/Logging.h>

#include "server.h"

using namespace ladder;
using namespace std::placeholders;

void OnMessage(const ConnectionPtr& conn, Buffer* buffer, Handler* handler) {
  std::string message = buffer->ReadAll();
  LOG_INFO("Received: " + message);
  if (handler)
  conn->Send(handler->OnRequest(message));
}

Server::Server(unsigned short port)
  : server_(nullptr), handler_(nullptr)
{
  Logger::create("test_server.log");
  SocketAddr addr("0.0.0.0", static_cast<uint16_t>(port), false);
  server_ = new TcpServer(addr, false);
}

Server::~Server() {
  delete server_;
}

void Server::Start() {
  server_->Start();
}

void Server::SetRequestCallback(Handler* handler) {
  server_->SetReadCallback(std::bind(OnMessage, _1, _2, handler));
}

Handler::~Handler() {}
