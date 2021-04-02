module github.com/suisrc/logger.zgo

go 1.16

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/suisrc/config.zgo v0.0.0
)

replace github.com/suisrc/config.zgo v0.0.0 => ../config
