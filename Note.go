##Latihan 1
berfungsi untuk belajar langkah awal cara membuat website menggunakan GO, mulai dari cara setup koneksi agar bisa diakses dibrowser, handler apa saja yg bisa dilakukan websitenya, dan cara kerja form


Notes
//sytax umum
/*Fungsi untuk open koneksi
	if err:= http.ListenAndServe(":8080",nil); err !=nil{
		log.Fatal(err)
		return
	}
*/

/*fungsi untuk apa yg terjadi ketika website diakses
	fileServer := http.FileServer (http.Dir("./source"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
*/