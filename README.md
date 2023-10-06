# Apply GORM and Gin in order to fulfil CURD

## Gorm (mysql: 192.168.1.72:3306)

```	
"gorm.io/driver/mysql"
"gorm.io/gorm"
```

```
go get -u github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql@v1.9.16
```

Gorm is a framework by which developers can
use go to get access to a variety of databases
including mysql. In this GIT, I fulfil the crud
through mysql.

There is a prerequisite that developers should be 
familiar with language of mysql.

## Gin (port: 9999)
```
	"github.com/gin-gonic/gin"
	"net/http"
```

```
 go get -u github.com/gin-gonic/gin
```

## Logrus

```
	"github.com/sirupsen/logrus"
```

```
 go get github.com/sirupsen/logrus
```