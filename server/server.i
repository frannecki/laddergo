%module(directors="1") server

%{
#include "server.h"
%}

%include <typemaps.i>
%include <cpointer.i>
%include <std_string.i>

%feature("director") Handler;

%include "server.h"
