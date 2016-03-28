#!/bin/bash

docker run -d -p 4140:4140 -p 9990:9990 -v `pwd`:/myapp -w /myapp --net host buoyantio/linkerd linkerd.yaml
