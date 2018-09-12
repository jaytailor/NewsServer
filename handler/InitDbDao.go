package handler

var mdao DbDAO;

func init(){
	mdao = DbDAO{Server:"localhost?maxPoolSize=4096", Database:DATABASE}
	mdao.Connect()
}

