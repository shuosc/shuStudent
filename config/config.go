package config

import "shuStudent/tools"

var COURSE_PROXY_ENDPOINT = tools.EnvOrElse("COURSE_PROXY_ENDPOINT", "http://cloud.shu.xn--io0a7i:30000/api/shu-course-proxy/")
var GET_NAME_URL = tools.EnvOrElse("GET_NAME_URL", "http://xk.autoisp.shu.edu.cn")
