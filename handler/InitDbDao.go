package handler

var mdao DbDAO;

func init(){
	mdao = DbDAO{Server:"127.0.0.1?maxPoolSize=4096", Database:DATABASE}
	mdao.Connect()
}

