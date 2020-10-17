Talker dökümentasyonu
=====================

**Talker dökümanlarına** hoş geldiniz. Burada bu paketin ne için
oluşturulduğunu, nasıl kullanabileceğinizi, paketin bağımlılıklarını ve
örnekleri bulabileceksiniz.

Girişten önce `Tolk’un yazarı Davy Kager ve diğer
geliştiricilere <https://github.com/dkager>`__ teşekkürler.

Eğer golang kullanmıyorsanız **C ve C++ için**
`Tolk’u <https://github.com/dkager/tolk>`__ deneyebilirsiniz. O gayet
basit bir kütüphane.

-  Bu paket şu anda **sadece Windows** içindir. İleriki süreçte Mac
   (VoiceOver) ve Linux (Orca) desteğinin eklenmesi planlanmaktadır.

İyi kodlamalar.

Hikaye
------

Görme engelli veya kör kullanıcılar bilgisayarı **ekran okuyucusu
yazılım paketleriyle** kullanırlar. Bu yazılımlar, işletim sisteminin
API’larını kullanarak kullanıcının oluşturduğu olaylara tepki verirler.

Açık kaynaklı ve özgür bir yazılım olan
`NVDA’yı <https://github.com/nvaccess/nvda>`__ inceleyebilirsiniz.
Repo’da ekran okuyucularını anlamanız için pekçok kod bulacaksınız.
Uygulamanızı erişilebilir yapabilmek için, NVDA başlangıç noktanız
olabilir.

Uygun tasarım oluşturulduğunda, genellikle ekran okuyucularına özel
olarak erişmeniz gerekmez. Örneğin ARIA’nın güzelce planlandığı web
sayfalarında ekstradan herhangi bir çabaya gerek yoktur.

Ama bazı oluşturacağınız yazılımlar için ekran okuyucularına ulaşmanız
mümkün olmayabilir. Örnek olarak OpenGL kullanan çeşitli QT arayüzleri,
UnrealEngine oyunları verilebilir. Veya bir satranç oyunu.

Görüntüyü render eden uygulamaları ekran okuyucuları tam olarak
algılayamaz. Bu senaryolarda, ekran okuyucularına API’leriyle erişmeniz
gerekir. Bu şekilde herhangi bir oyundaki olayları ekran okuyucusuna
gönderebilirsiniz.

Talker, Golang için Windows işletim sisteminde popüler olan ekran
okuyucularına erişmek için API’ler sağlar. Aslında Tolk için bir
sarmalayıcıdır.

Bağımlılıklar ve tavsiyeler
---------------------------

Projede **CGO** kullanılmıştır. Uygun bir GCC derleyici paketi bulmanız
gereklidir. Golang, sisteminizde bulunan GCC derleyicisi ile başarılı
olarak çalışıyorsa yeni bir derleyiciye ihtiyacınız yok.

Benim tavsiyem olan Clang, LLVM, GDB, GCC ve G++ gibi popüler derleyici
araçları sağlayan `şu
repoyu <https://github.com/brechtsanders/winlibs_mingw/releases>`__
kullanabilirsiniz.

Yalnızca **talker modülü** için değil daha sonra geliştireceğiniz ve
derleyeceğiniz her şey için güzel bir ortamınız oluşur. Daha az dosya,
daha çok iş 😁

**Note**: adımları takip ettiğinizde MinGW başarıyla sisteminize
kurulmuş olacaktır. Bunun için **CMD** kullanabiliyor olmanız
gereklidir.

Lütfen adımları sırasıyla takip edin.

-  Uygulamanız için gerekli olan mimariye karar vermeniz gerekli. DLL
   bağımlılıkları her iki mimari için farklıdır. Daha eski sistemlerde
   kullanılması gereken bir uygulamayı 32bit olarak derlemelisiniz.

   -  Örneğin 32bit derleme yapabilmeniz için, 32bit Go derleyici paketi
      ve 32bit gcc derleyici paketine ihtiyacınız var. Artık desteklenen
      ve tavsiye edilen 64bit’tir. Lütfen buna karar verin.

-  Bu sayfayı tarayıcınızda açın:
   `github.com/brechtsanders/winlibs_mingw/releases <https://github.com/brechtsanders/winlibs_mingw/releases>`__

   -  “…x86_64” dosyaları 64bit sistemler içindir. “…i686” kısaca 32bit
      olarak bilinir. Zip veya 7Z arşivi olarak indirmelisiniz. Örneğin
      **64bit GGo için** x86-64 dosyası indirilmelidir.

-  İndirdiğiniz dosyayı uygun bir konuma çıkartmalısınız. Konumda ASCII
   dışında karakterlerin olmamasına ve aralıkların olmamasına dikkat
   etmelisiniz (‘Ö’, ‘🍟’, ‘my user’). Bazı durumlarda, bu sorunların
   oluşmasını engelleyecektir. Derleyiciler genelde Windows diskinizin
   main konumuna kurulur (C:\MinGW32 gibi)
-  ``XX:MinGW\bin`` konumunu pathe eklemelisiniz.

   -  Örnek path: **“C:\MinGW32\bin”**. Alternatif olarak
      bir batchfile oluşturabilirsiniz. Bu fikir sizin için daha iyi
      görünüyorsa lütfen “batchfile örneği” başlığına bakın.
   -  ``XX:\xx\MinGW\bin`` konumunu kopyalayın. Örnek:
      ``C:\MinGW64\bin``
   -  Control Panel > Searchbox: “env” > “Edit environment variables for
      your account” or “Edit the system environment variables” > select
      path > Click edit button > click New button > Paste to clipboard >
      Click ok button > click ok button.
   -  Tebrikler! MinGW path’e eklendi.

-  Şimdi test zamanı. Herhangi bir CMD konsolu açmalısınız ve
   ``gcc --version`` yazmalısınız.

   -  **Uyarı**: açık olan konsolları kapatmak iyi bir fikir olabilir.
      -Şuna benzer bir çıktı görmelisiniz:
      ``gcc (MinGW-W64 x86_64-posix-seh, built by Brecht Sanders) 10.2.0 Copyright (C) 2020 Free Software Foundation, Inc. This is free software; see the source for copying conditions.  There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.``
   -  Eğer ``'gcc' is not recognized as an internal or external command``
      hatasını alıyorsanız, path’e ekleme işlemi başarısız olmuş
      demektir.

--------------

Ve evet, MinGW kurulmuş olmalı! Sırada bu modülün kullanımı var.

Kurulum, Fonksiyonlar, Kullanım ve derleme sonrası aşamalar
-----------------------------------------------------------

Kurulum
~~~~~~~

Projenin **GCCGo** bağımlılığı olduğu için Go’nun **gcc**’ye
ulaşabiliyor olması gerekir.

::

   go get github.com/SeanTolstoyevski/talker

Herhangi bir çıktı görmediğinizde modül kurulmuş demektir.

Fonksiyonlar
~~~~~~~~~~~~

-  ``talker.Load()``

Tolk’u yükler ve yapılandırır. Diğer fonksiyonları çalıştırmadan önce bu
fonksiyonu çalıştırdığınıza emin olmalısınız.

-  Dönüş: hiçbir şey.

C açıklaması: Initializes Tolk by loading and initializing the screen
reader drivers and setting the current screen reader driver, provided at
least one of the supported screen readers is active. Also initializes
COM if it has not already been initialized on the calling thread.
Calling this function more than once will only initialize COM. You
should call this function before using the functions below. Use
Tolk_IsLoaded to determine if Tolk has been initialized.

-  ``talker.Unload()``

Tolk’un kaynaklarını serbest bırakır. Uygulamalardan çıkma gibi
durumlarda çalıştırılması gerekir.

-  Dönüş: hiçbir şey

C açıklaması: Finalizes Tolk by finalizing and unloading the screen
reader drivers and clearing the current screen reader driver, provided
one was set. Also uninitializes COM on the calling thread. Calling this
function more than once will only uninitialize COM. You should not use
the functions below if this function has been called.

-  ``talker.IsLoaded() bool``

Tolk’un yapılandırılma durumunu kontrol eder. Eğer daha önce Tolk
yüklenmişse True, aksi türlü false döndürür

-  Dönüş: bool

C açıklaması: Tests if Tolk has been initialized.

-  ``talker.TrySAPI(yesno bool)``

SAPI’nın dahil edilip edilmeyeceiğini belirler. En iyi durumda Load’den
önce kullanmalısınız.

-  Dönüş: hiçbir şey

C açıklaması: Sets if Microsoft Speech API (SAPI) should be used in the
screen reader auto-detection process. The default is not to include
SAPI. The SAPI driver will use the system default synthesizer, voice and
soundcard. This function triggers the screen reader detection process if
needed. For best performance, you should call this function before
calling Tolk_Load.

-  ``talker.PreferSAPI(yesno bool)``

SAPI driver’inin ilk öğe veya son öğe olma durumunu ayarlar. Varsayılan
durumda SAPI son öğedir. En iyi durumda Load’den önce kullanmalısınız.

-  Dönüş: hiçbir şey

C açıklaması: If auto-detection for SAPI has been turned on through
Tolk_TrySAPI, sets if SAPI should be placed first (true) or last (false)
in the screen reader detection list. Putting it last is the default and
is good for using SAPI as a fallback option. Putting it first is good
for ensuring SAPI is used even when a screen reader is running, but keep
in mind screen readers will still be tried if SAPI is unavailable. This
function triggers the screen reader detection process if needed. For
best performance, you should call this function before calling
Tolk_Load.

-  ``talker.DetectScreenReader() string``

Şu an aktif olan ekran okuyucusunun adını döndürür (Örn. NVDA, JAWS…)

-  Dönüş: string

C açıklaması: Returns the common name for the currently active screen
reader driver, if one is set. If none is set, tries to detect the
currently active screen reader before looking up the name. If no screen
reader is active, NULL is returned. Note that the drivers hard-code the
common name, it is not requested from the screen reader itself. You
should call Tolk_Load once before using this function.

-  ``talker.Output(text string, interrupt bool) bool``

Metni aktif olan ekran okuyucusuna gönderir. Eğer “interrupt” true ise
mevcut konuşma duraklatılır. Çoğu zaman “interrupt”un true olmasına
gerek yoktur. Kullanıcılar mevcut konuşmanın kesilmesinden hoşlanmazlar.

-  Dönüş: bool

C açıklaması: Outputs text through the current screen reader driver, if
one is set. If none is set or if it encountered an error, tries to
detect the currently active screen reader before outputting the text.
This is the preferred function to use for sending text to a screen
reader, because it uses all of the supported output methods (speech
and/or braille depending on the current screen reader driver). You
should call Tolk_Load once before using this function. This function is
asynchronous.

Batcfile örneği
---------------


Kodu uygun şekilde düzenledikten sonra **.bat dosyası** olarak
kaydetmelisiniz.

Derleme yaparken bat dosyasını çalıştırmalı ve bu konsolu
kullanmalısınız:

::

   @echo off
   set PATH=C:MinGW64bin;%PATH%
   rem echo %PATH%
   "C:WINDOWSsystem32cmd.exe"
   "cd C:"
