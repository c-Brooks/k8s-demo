#!/usr/bin/env node

import http from "k6/http";

const ep = "http://104.197.233.82/";

export default function() {
  http.get(ep);
};

