package infrastructure

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Test struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
}

func NewConfig() *Config {

	c := new(Config)

	// 本番環境
	// c.DB.Production.Host = "us-cdbr-east-05.cleardb.net"
	// c.DB.Production.Username = "b797415be7f773"
	// c.DB.Production.Password = "2811cb95"
	// c.DB.Production.DBName = "heroku_acf2f27c78e8fd6"

	// 開発環境
	c.DB.Production.Host = "localhost"
	c.DB.Production.Username = "zoom_matching_user"
	c.DB.Production.Password = "password"
	c.DB.Production.DBName = "zoom_matching_db"

	// c.DB.Test.Host = "localhost"
	// c.DB.Test.Username = "zoom_matching_test_user"
	// c.DB.Test.Password = "password"
	// c.DB.Test.DBName = "zoom_matching_test_db"

	return c
}
