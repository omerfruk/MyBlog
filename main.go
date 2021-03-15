package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/routers"
	"log"
)

func main() {
	engine := html.New("views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New())
	//db baglantıları ve migrate işlemi
	database.ConnectAndMigrate()
	routers.Router(app)
	// create islemleri
	/*service.CreateEntry("../img/bgCover.jpg","WELCOME TO MY PAGE","IT'S NICE TO MEET YOU","KNOW ME")

	service.CreatePortfolio("../img/computer.jpg","Computer","what is a computer and where are we in this")
	service.CreatePortfolio("../img/network.jpg","Network","How they get us all around with a click")
	service.CreatePortfolio("../img/user.jpg","User","Well, who are we users ")
	service.CreatePortfolio("../img/developer.jpg","Developer","what do developers do?")
	service.CreatePortfolio("../img/hardware.jpg","Hardware","what's the hardware? who are hardwareist")
	service.CreatePortfolio("../img/securty.jpg","Securty","who constitute our security?")

	service.CreateFooter("../img/fotom.jpg","ÖMER FARUK TASDEMIR","Developer","https://www.instagram.com/omer_fruk/?hl=tr","https://www.facebook.com/omerrf/","https://github.com/omerfruk","https://twitter.com/home?lang=tr")

	service.CreateInstructions("Let's learn something about technology","The most efficient thing which is invented by people is book","Computer","Web","Users","It is certain that the development progress of the computer without pausing, from the past to the present,\nwill continue exponentially. So where are we in this technology?","We can go from one side of the world to the other with one click. \nSo how do we make this journey?","We, I mean the users, can shape the progress of technology. It sounds amazing right?")

	service.CreateTopBar("ÖmFar.","home","future","researcher","contact","Login")


	*/
	/*service.CreateUser("","selamlar","Serhat","055451357851",false)
	service.CreateUser("","naber","ibrahim ","0551231",false)
	service.CreateUser("","iyilik","musa","055451123",false)
	service.CreateUser("","senden naber","emrullah","055412351",false)
	service.CreateUser("","iyilik","osman","055412351",false)*/
	//service.CreatePortfolio("../img/game.jpg","Game","Of course, the game everyone wants!")
	/*service.CreateFoto("../img/gallery/gallery1.jpg","Boş Sokaklar","Yanlızliğin en çok benzediği şeydir boş sokaklar")
	service.CreateFoto("../img/gallery/gallery2.jpg","Doğal mı Afet mi ?","Normal şartlarda bu fotoğrafın güzel gözükmesi gerekirken aslında odaklanmamız gerekn yer güzellik insan elinden mi yoksa doğal olarak olması gerektiği degilmi")
	service.CreateFoto("../img/gallery/gallery3.jpg","Doğ","Her kim olursa olsun insan Doğada kendini rahat hisseder çünki insanın DOĞAsında var ")
	service.CreateFoto("../img/gallery/gallery4.jpg","Yamaçlar","Kimi için ürkütücü kimi için ise bir heyecan kaynağı olan bu yamaçlar gerçekten de insana \"Vayyy\" dedirtiyor")
	service.CreateFoto("../img/gallery/gallery5.jpg","Soluk bir ilk bahar","insanın şöyle bir yerde yıllarını kaybetse acaba kim bulabilir")
	service.CreateFoto("../img/gallery/gallery6.jpg","Dağlar Dumanlı Dağlar","Kendimi şöyle güzel bir dağ yamacına bırakıpğ gidesim var hem kendimi kaybedeyim hem kendimi buliyim süper değilmi ?")
	service.CreateFoto("../img/gallery/gallery7.jpg","Yıldızların Altında","Düşünsen yıldızlı bir gece herşeyi geride bırakıyorsun ve bir an olsun sadece anın tadını çıkartıyorsun		\n++Yapılacaklar listesinde")
	service.CreateFoto("../img/gallery/gallery8.jpg","Kanyon","Sundan şüphem yok ki ben bir doğa hayranıyım")
	service.CreateFoto("../img/gallery/gallery9.jpg","Tarihi Yapıtlar","Aslında bakarsanız çagımız ilerledikçe pek çok şey geride kalıyor \n-Mimari yapılar\n-Hoşgörü\n-Aaa bir de insanlık")
	service.CreateFoto("../img/gallery/gallery10.jpg","Uretgenlik","insanlar öyele varlıklar ki istedikleri zaman sadece adi bir hareketleriyle dünya güzelleşirken bir hareketle daha da güzel olabiliyor")

*/
	log.Fatal(app.Listen(":4747"))

}
